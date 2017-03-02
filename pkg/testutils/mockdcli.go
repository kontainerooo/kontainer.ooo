package testutils

import (
	"context"
	"errors"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

var (
	// ErrClientError is returned, when the client should return an error
	ErrClientError = errors.New("client failure")

	// ErrAlreadyRunning is returned, if a container which should be started is already running
	ErrAlreadyRunning = errors.New("container already running")
)

// MockDCli simulates a docker client for testing purposes
type MockDCli struct {
	containers map[string]bool
	err        bool
}

// SetError sets the err property of MockDCli to be true, causing the next instruction to return an error
func (d *MockDCli) SetError() {
	d.err = true
}

func (d *MockDCli) produceError() bool {
	if d.err {
		d.err = false
		return true
	}
	return false
}

// IsRunning returns
func (d *MockDCli) IsRunning(container string) bool {
	return d.containers[container]
}

// ContainerStart is
func (d *MockDCli) ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error {
	if d.produceError() {
		return ErrClientError
	}
	if !d.IsRunning(container) {
		d.containers[container] = true
		return nil
	}
	return ErrAlreadyRunning
}

// ContainerKill is
func (d *MockDCli) ContainerKill(ctx context.Context, container, signal string) error {
	if d.produceError() || !d.IsRunning(container) {
		return ErrClientError
	}
	d.containers[container] = false
	return nil
}

// ContainerExecCreate is
func (d *MockDCli) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (string, error) {
	if d.produceError() || !d.IsRunning(container) {
		return "", ErrClientError
	}
	return strings.Join(config.Cmd, " "), nil
}

// ContainerCreate is
func (d *MockDCli) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	if d.produceError() {
		return container.ContainerCreateCreatedBody{}, ErrClientError
	}

	return container.ContainerCreateCreatedBody{}, nil
}

// ImageInspectWithRaw is
func (d *MockDCli) ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error) {
	if d.produceError() {
		return types.ImageInspect{}, nil, ErrClientError
	}

	return types.ImageInspect{}, nil, nil
}

// NewMockDCli returns a new instance of MockDCli
func NewMockDCli() *MockDCli {
	return &MockDCli{
		containers: make(map[string]bool),
	}
}
