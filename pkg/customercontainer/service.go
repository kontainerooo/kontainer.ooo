package customercontainer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/ttdennis/kontainer.io/pkg/abstraction"
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
	// TODO: implement
	return nil
}

// NewService creates a customercontainer with necessary dependencies.
func NewService(dcli abstraction.DCli) Service {
	return &service{
		dcli: dcli,
	}
}
