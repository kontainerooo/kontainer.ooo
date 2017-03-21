package containerlifecycle

import (
	"context"
	"strings"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
)

// Service is awesome
type Service interface {
	// StartContainer
	StartContainer(id string) error

	// StartCommand
	StartCommand(id string, cmd string) (string, error)

	// StopContainer
	StopContainer(id string) error
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

// NewService creates a UserService with necessary dependencies.
func NewService(dcli abstraction.DCli) Service {
	return &service{
		dcli: dcli,
	}
}
