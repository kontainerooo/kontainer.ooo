// Package customercontainer provides basic container functions to create and remove containers
package customercontainer

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/builder/dockerignore"
	"github.com/docker/docker/cli/command/image/build"
	"github.com/docker/docker/pkg/archive"
	"github.com/ttdennis/kontainer.io/pkg/abstraction"
	"github.com/ttdennis/kontainer.io/pkg/kmi"
)

// Service Customer Container service
type Service interface {
	// CreateContainer instanciates a container for a User with the id refid and returns its id
	CreateContainer(refid int, cfg *ContainerConfig) (name string, id string, err error)

	// EditContainer is used to edit a container instances configuration by id
	EditContainer(id string, cfg *ContainerConfig) error

	// RemoveContainer is used to remove a container instance by id
	RemoveContainer(id string) error

	// Instances returns a list of instances of an user by id
	Instances(refid int) []string

	// CreateDockerImage creates a Docker image from a given KMI
	CreateDockerImage(refid int, kmi kmi.KMI) error

	// AddKMIClient adds the kmi Endpoints to the service
	AddKMIClient(ke *kmi.Endpoints)
}

type service struct {
	dcli      abstraction.DCli
	kmiClient *kmi.Endpoints
}

// imageExists checks if a docker image exists.
func (s *service) imageExists(image string) bool {
	_, _, err := s.dcli.ImageInspectWithRaw(context.Background(), image)

	if err == nil {
		return true
	}

	return !s.dcli.IsErrImageNotFound(err)
}

func (s *service) CreateContainer(refid int, cfg *ContainerConfig) (name string, id string, err error) {
	securityOpts := []string{
		"no-new-privileges",
	}

	b := bytes.NewBuffer(nil)
	if err := json.Compact(b, []byte(SeccompProfile)); err != nil {
		return "", "", fmt.Errorf("compacting json for seccomp profile failed: %v", err)
	}

	// Use docker standard
	securityOpts = append(securityOpts, fmt.Sprintf("seccomp=%s", b.Bytes()))

	// TODO: which CAPS are dropped?
	dropCaps := &strslice.StrSlice{"NET_RAW"}

	// Check if the image exists
	if exists := s.imageExists(cfg.ImageName); !exists {
		return "", "", fmt.Errorf("Image does not exist")
	}

	// Create the container
	r, err := s.dcli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image:        cfg.ImageName,
			Cmd:          []string{"sh"},
			Tty:          true,
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
			OpenStdin:    true,
			StdinOnce:    true,
		},
		&container.HostConfig{
			SecurityOpt: securityOpts,
			CapDrop:     *dropCaps,
			NetworkMode: "none",
			LogConfig: container.LogConfig{
				Type: "none",
			},
		},
		nil, "")

	if err != nil {
		return "", "", err
	}

	// Name the container $userID-$contaierID
	containerName := fmt.Sprintf("%d-%s", refid, r.ID)
	if err := s.dcli.ContainerRename(context.Background(), r.ID, containerName); err != nil {
		return "", "", err
	}

	return containerName, r.ID, nil
}

func (s *service) EditContainer(id string, cfg *ContainerConfig) error {
	// TODO: implement
	return nil
}

func (s *service) RemoveContainer(id string) error {
	return s.dcli.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{})
}

func (s *service) Instances(refid int) []string {

	containers, _ := s.dcli.ContainerList(context.Background(), types.ContainerListOptions{})

	var containerList []string
	userid := fmt.Sprintf("%d", refid)
	for _, v := range containers {
		if len(v.Names) > 0 && strings.HasPrefix(v.Names[0], userid) {
			entry := v.ID
			containerList = append(containerList, entry)
		}
	}

	return containerList
}

func (s *service) CreateDockerImage(refid int, kmi kmi.KMI) error {
	labels := make(map[string]string)
	labels["user"] = string(refid)

	dockerfile, err := addEnvToDockerfile(kmi.Dockerfile, kmi.Environment)
	if err != nil {
		return err
	}

	buildContext, err := createBuildContext(kmi.Container, dockerfile)
	if err != nil {
		return err
	}

	s.dcli.ImageBuild(context.Background(), buildContext, types.ImageBuildOptions{
		Tags: []string{
			fmt.Sprintf("%d-%s", refid, kmi.Name),
		},
		SuppressOutput: false,
		NoCache:        true,
		Remove:         false,
		ForceRemove:    false,
		PullParent:     false,
		CPUSetCPUs:     kmi.Resources["cpus"].(string),
		Memory:         int64(kmi.Resources["mem"].(int)),
		MemorySwap:     int64(kmi.Resources["swap"].(int)),
		Dockerfile:     dockerfile,
		Labels:         labels,
	})
	return nil
}

func (s *service) AddKMIClient(ke *kmi.Endpoints) {
	s.kmiClient = ke
}

func createBuildContext(path string, dockerfileContent string) (io.Reader, error) {
	var buildCtx io.ReadCloser

	// Read dockerignore file from container path
	f, err := os.Open(filepath.Join(path, ".dockerignore"))
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	defer f.Close()

	var excludes []string
	excludes, err = dockerignore.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if err := build.ValidateContextDirectory(path, excludes); err != nil {
		return nil, fmt.Errorf("Cannot create context from container path: %s", err)
	}

	compression := archive.Uncompressed

	// Create tar from build context
	buildCtx, err = archive.TarWithOptions(path, &archive.TarOptions{
		ExcludePatterns: excludes,
		Compression:     compression,
	})
	if err != nil {
		return nil, err
	}

	// Add custom dockerfile to tar archive
	buildCtxReader, err := addDockerfileToTar(buildCtx, dockerfileContent)
	if err != nil {
		return nil, err
	}

	return buildCtxReader, nil
}

func addDockerfileToTar(inputTar io.ReadCloser, dockerfileContent string) (io.Reader, error) {
	// pipeReader, pipeWriter := io.Pipe()
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	// Open current archive for reading
	tr := tar.NewReader(inputTar)

	defer inputTar.Close()

	for {

		hdr, err := tr.Next()

		// End of archive
		if err == io.EOF {
			// Write dockerfile to tar
			dockerfileBytes := []byte(dockerfileContent)
			hdr = &tar.Header{
				Name: "Dockerfile",
				Mode: 0600,
				Size: int64(len(dockerfileBytes)),
			}

			if err := tw.WriteHeader(hdr); err != nil {
				return nil, err
			}

			if _, err := tw.Write(dockerfileBytes); err != nil {
				return nil, err
			}

			tw.Close()
			return buf, nil
		}

		if err != nil {
			return nil, err
		}

		// Copy header from file
		err = tw.WriteHeader(hdr)
		if err != nil {
			return nil, err
		}

		// Copy contents of file
		content := io.Reader(tr)
		_, err = io.Copy(tw, content)
		if err != nil {
			return nil, err
		}

	}
}

func addEnvToDockerfile(dockerfile string, env map[string]interface{}) (string, error) {
	envString := "ENV"
	for k, v := range env {
		if envKeyValid(k) {
			// To optimize performance all ENVs are put into one line
			// so that docker only needs to create one container layer for
			// all environment variables
			envString = fmt.Sprintf("%s %s=\"%s\"", envString, k, envValueEscape(v.(string)))
		} else {
			return "", fmt.Errorf("Invalid ENV key (%s)", k)
		}
	}

	return fmt.Sprintf("%s\n%s", dockerfile, envString), nil
}

func envKeyValid(envKey string) bool {
	r, _ := regexp.Compile("^[a-zA-Z_][a-zA-Z0-9_]*$")
	return r.MatchString(envKey)
}

func envValueEscape(envValue string) string {
	return strings.Replace(envValue, "\"", "\\\"", -1)
}

// NewService creates a customercontainer with necessary dependencies.
func NewService(dcli abstraction.DCli) Service {
	return &service{
		dcli: dcli,
	}
}
