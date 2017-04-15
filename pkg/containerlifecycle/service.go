package containerlifecycle

import (
	"context"
	"strings"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
)

// Service is awesome
type Service interface {
	// StartContainer starts a container
	StartContainer(id string) error

	// StartCommand starts a command in a given container
	StartCommand(id string, cmd string) (string, error)

	// StopContainer stops a container
	StopContainer(id string) error

	// IsRunning checks if a container is running
	IsRunning(id string) bool
}

type service struct {
	dcli abstraction.DCli
}

func (s *service) StartContainer(id string) error {
	return s.dcli.ContainerStart(context.Background(), id, dockerTypes.ContainerStartOptions{})
}

func (s *service) StartCommand(id string, cmd string) (string, error) {
	idStr, err := s.dcli.ContainerExecCreate(context.Background(), id, dockerTypes.ExecConfig{
		Cmd: strings.Split(cmd, " "),
	})
	return idStr, err
}

func (s *service) StopContainer(id string) error {
	return s.dcli.ContainerKill(context.Background(), id, "SIGKILL")
}

func (s *service) IsRunning(id string) bool {
	info, err := s.dcli.ContainerInspect(context.Background(), id)
	if err != nil {
		return false
	}

	return info.State.Running
}

// NewService creates a containerlifecycle service with necessary dependencies.
func NewService(dcli abstraction.DCli) Service {
	return &service{
		dcli: dcli,
	}
}
