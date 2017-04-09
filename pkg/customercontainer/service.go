// Package customercontainer provides basic container functions to create and remove containers
package customercontainer

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/builder/dockerignore"
	"github.com/docker/docker/cli/command/image/build"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
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
	CreateDockerImage(refid int, kmiID uint) (string, error)

	// AddKMIClient adds the kmi Endpoints to the service
	AddKMIClient(ke *kmi.Endpoints)

	// AddLogger provides a logger to the service
	AddLogger(l log.Logger)
}

type service struct {
	dcli      abstraction.DCli
	kmiClient *kmi.Endpoints
	logger    log.Logger
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

func (s *service) CreateDockerImage(refid int, kmiID uint) (string, error) {

	buildBuf := bytes.NewBuffer(nil)

	if s.kmiClient == nil {
		return "", errors.New("No KMI client")
	}

	kmiResponse, err := s.kmiClient.GetKMIEndpoint(context.Background(), &kmi.GetKMIRequest{
		ID: kmiID,
	})
	if err != nil {
		return "", err
	}

	kmi := kmiResponse.(*kmi.GetKMIResponse).KMI

	imageTag := fmt.Sprintf("kro/%s:%d", strings.ToLower(kmi.Name), refid)
	_, _, err = s.dcli.ImageInspectWithRaw(context.Background(), imageTag)

	if !s.dcli.IsErrImageNotFound(err) {
		return "", err
	}

	dockerfile, err := addEnvToDockerfile(kmi.Dockerfile, kmi.Environment)
	if err != nil {
		return "", err
	}

	buildContext, err := createBuildContext(kmi.Context, dockerfile)
	if err != nil {
		return "", err
	}

	buildOptions := generateBuildOptions(kmi, refid, imageTag)

	res, err := s.dcli.ImageBuild(context.Background(), buildContext, buildOptions)
	if err != nil {
		return "", err
	}

	err = jsonmessage.DisplayJSONMessagesStream(res.Body, buildBuf, 1, false, nil)

	if err != nil {
		if jerr, ok := err.(*jsonmessage.JSONError); ok {
			if jerr.Code == 0 {
				jerr.Code = 1
			}
			return "", fmt.Errorf("%s", jerr.Message)
		}
	}

	// Response is sha1:IMAGE_ID
	answer := strings.Split(string(buildBuf.Bytes()), ":")
	imageID := strings.Replace(answer[1], "\n", "", -1)

	s.logger.Log("msg", "Created image", "ImageID", imageID)

	return imageID, nil
}

func (s *service) AddKMIClient(ke *kmi.Endpoints) {
	s.kmiClient = ke
}

func (s *service) AddLogger(l log.Logger) {
	s.logger = l
}

func generateBuildOptions(kmi *kmi.KMI, userID int, imageTag string) types.ImageBuildOptions {
	var (
		cpus string
		mem  int64
		swap int64
	)
	labels := make(map[string]string)
	labels["user"] = string(userID)

	tags := []string{
		imageTag,
	}

	intCpus, ok := kmi.Resources["cpus"]
	if ok {
		cpus = strconv.FormatInt(int64(intCpus.(int)), 10)
	}

	intMem, ok := kmi.Resources["mem"]
	if ok {
		mem = int64(intMem.(int))
	}

	intSwap, ok := kmi.Resources["swap"]
	if ok {
		swap = int64(intSwap.(int))
	}

	return types.ImageBuildOptions{
		Tags:           tags,
		SuppressOutput: true,
		NoCache:        true,
		Remove:         false,
		ForceRemove:    false,
		PullParent:     false,
		CPUSetCPUs:     cpus,
		Memory:         mem,
		MemorySwap:     swap,
		Dockerfile:     "./Dockerfile",
		Labels:         labels,
	}
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
	if err == nil {
		excludes, err = dockerignore.ReadAll(f)
		if err != nil {
			return nil, err
		}
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

	if len(env) == 0 {
		return dockerfile, nil
	}

	envString := "ENV"
	for k, v := range env {
		if envKeyValid(k) {
			// To optimize performance all ENVs are put into one line
			// so that docker only needs to create one container layer for
			// all environment variables
			envString = fmt.Sprintf("%s %s=\"%s\"", envString, k, envValueEscape(string(v.(int))))
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
