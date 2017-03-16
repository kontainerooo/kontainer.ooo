// Package customercontainer provides basic container functions to create and remove containers
package customercontainer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
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
}

type service struct {
	dcli abstraction.DCli
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

	s.dcli.ImageBuild(context.Background(), nil, types.ImageBuildOptions{
		Tags: []string{
			fmt.Sprintf("%d-%s", refid, kmi.Name),
		},
		SuppressOutput: false,
		RemoteContext: kmi.Context
		NoCache:     true,
		Remove:      false,
		ForceRemove: false,
		PullParent:  false,
		CPUSetCPUs:  "1",
		// CPUSetMems:  "1",
		Memory:     0x1fffffff, // 500 MB
		MemorySwap: 0x3fffffff, // 1 GB
		// CgroupParent   string
		// NetworkMode    string
		// ShmSize        int64
		Dockerfile: kmi.Dockerfile + kmi.Env,
		// Ulimits        []*units.Ulimit
		// BuildArgs needs to be a *string instead of just a string so that
		// we can tell the difference between "" (empty string) and no value
		// at all (nil). See the parsing of buildArgs in
		// api/server/router/build/build_routes.go for even more info.
		// BuildArgs   map[string]*string
		// AuthConfigs map[string]AuthConfig
		// Context     io.Reader
		Labels: labels,
		// squash the resulting image's layers to the parent
		// preserves the original image and creates a new one from the parent with all
		// the changes applied to a single layer
		Squash: false, // TODO: Test disk impact
		// CacheFrom specifies images that are used for matching cache. Images
		// specified here do not need to have a valid parent chain to match cache.
		// CacheFrom   []string
		// ExtraHosts  []string // List of extra hosts
	})
	return nil
}

// NewService creates a customercontainer with necessary dependencies.
func NewService(dcli abstraction.DCli) Service {
	return &service{
		dcli: dcli,
	}
}
