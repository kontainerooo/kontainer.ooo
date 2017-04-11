package testutils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
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

	// ErrNetNotFound is returned when a network that is operated on does not exist
	ErrNetNotFound = errors.New("network not found")
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

// MockDCli simulates a docker client for testing purposes
type MockDCli struct {
	running         map[string]bool
	containers      map[string]types.Container
	images          map[string]bool
	err             bool
	idNotExist      bool
	dockerIsOffline bool
	networks        map[string]mockNetwork
}

type mockNetwork struct {
	ID         string
	containers []string
}

// SetError sets the err property of MockDCli to be true, causing the next instruction to return an error
func (d *MockDCli) SetError() {
	d.err = true
}

// SetDockerOffline simulates a non responding docker daemon
func (d *MockDCli) SetDockerOffline() {
	d.dockerIsOffline = true
}

// SetIDNotExisting simulates a non existing docker container id
func (d *MockDCli) SetIDNotExisting() {
	d.idNotExist = true
}

func (d *MockDCli) produceError() bool {
	if d.err {
		d.err = false
		return true
	}
	return false
}

// GetNetworks returns available networks
func (d *MockDCli) GetNetworks() map[string]string {
	nwMap := make(map[string]string)
	for k, v := range d.networks {
		nwMap[k] = v.ID
	}
	return nwMap
}

// IsRunning checks if a mocked container is running
func (d *MockDCli) IsRunning(container string) bool {
	return d.running[container]
}

// CreateMockImage creates a mock image
func (d *MockDCli) CreateMockImage(image string) {
	d.images[image] = true
}

// ContainerStart sets a given container to the status running
func (d *MockDCli) ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error {
	if d.produceError() {
		return ErrClientError
	}
	if !d.IsRunning(container) {
		d.running[container] = true
		return nil
	}
	return ErrAlreadyRunning
}

// ContainerKill sets the given container to the status stopped
func (d *MockDCli) ContainerKill(ctx context.Context, container string, signal string) error {
	if d.produceError() || !d.IsRunning(container) {
		return ErrClientError
	}
	d.running[container] = false
	return nil
}

// ContainerExecCreate creates a new execution on a given container
func (d *MockDCli) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (string, error) {
	if d.produceError() || !d.IsRunning(container) {
		return "", ErrClientError
	}
	return strings.Join(config.Cmd, " "), nil
}

// ContainerCreate creates a mock container
func (d *MockDCli) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	if d.produceError() || d.dockerIsOffline {
		return container.ContainerCreateCreatedBody{}, ErrClientError
	}

	id := fmt.Sprintf("%d", rand.Int())
	d.containers[id] = types.Container{
		ID: id,
	}

	return container.ContainerCreateCreatedBody{
		ID: id,
	}, nil
}

// ContainerRename renames a container with a given ID
func (d *MockDCli) ContainerRename(ctx context.Context, containerID, newContainerName string) error {
	_, ok := d.containers[containerID]
	if d.produceError() || !ok || d.idNotExist {
		return ErrClientError
	}

	d.containers[containerID] = types.Container{
		ID:    d.containers[containerID].ID,
		Names: []string{newContainerName},
	}

	return nil
}

// ContainerRemove removes a container with a given ID
func (d *MockDCli) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error {
	_, ok := d.containers[containerID]
	if d.produceError() || !ok {
		return ErrClientError
	}

	delete(d.containers, containerID)

	return nil
}

// ContainerList lists all containers on the host
func (d *MockDCli) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	if d.produceError() {
		return nil, ErrClientError
	}

	var tContainers []types.Container
	for cid := range d.containers {
		tmp := types.Container{
			ID:    cid,
			Names: []string{fmt.Sprintf("123-%s", cid)},
		}
		tContainers = append(tContainers, tmp)
	}

	return tContainers, nil
}

// ImageInspectWithRaw gets information about a mock image
func (d *MockDCli) ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error) {
	if d.images[imageID] {
		return types.ImageInspect{}, nil, nil
	}

	return types.ImageInspect{}, nil, fmt.Errorf("Image does not exist")
}

// ImageBuild builds a mock image
func (d *MockDCli) ImageBuild(ctx context.Context, buildContext io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error) {
	d.images[options.Tags[0]] = true
	return types.ImageBuildResponse{
		Body: nopCloser{bytes.NewBufferString(fmt.Sprintf("{\"stream\":\"sha1:%s\"}", options.Tags[0]))},
	}, nil
}

// IsErrImageNotFound returns true if the error means the image was not found
func (d *MockDCli) IsErrImageNotFound(err error) bool {
	if err != nil {
		return true
	}

	return false
}

// NetworkCreate creates a new mock network
func (d *MockDCli) NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error) {
	if d.produceError() {
		return types.NetworkCreateResponse{}, ErrClientError
	}
	_, ok := d.networks[name]
	if ok {
		return types.NetworkCreateResponse{}, errors.New("Network already exists")
	}

	id := fmt.Sprintf("%d", rand.Int())
	d.networks[name] = mockNetwork{
		ID: id,
	}

	return types.NetworkCreateResponse{
		ID: id,
	}, nil
}

// NetworkRemove removes a mock network
func (d *MockDCli) NetworkRemove(ctx context.Context, networkID string) error {
	if d.produceError() {
		return ErrClientError
	}

	for k, v := range d.networks {
		if v.ID == networkID {
			delete(d.networks, k)
			return nil
		}
	}

	return ErrNetNotFound
}

// NetworkConnect connects a container to a mock network
func (d *MockDCli) NetworkConnect(ctx context.Context, networkID, containerID string, config *network.EndpointSettings) error {
	if d.produceError() {
		return ErrClientError
	}

	for k, v := range d.networks {
		if v.ID == networkID {
			_, ok := d.networks[k]
			if ok {
				m := d.networks[k]
				m.containers = append(d.networks[k].containers, containerID)
				d.networks[k] = m
				return nil
			}
			d.networks[k] = mockNetwork{
				ID:         networkID,
				containers: []string{containerID},
			}
			return nil
		}
	}

	return ErrNetNotFound
}

// NetworkDisconnect disconnects a container from a mock network
func (d *MockDCli) NetworkDisconnect(ctx context.Context, networkID, containerID string, force bool) error {
	if d.produceError() {
		return ErrClientError
	}

	for _, v := range d.networks {
		if v.ID == networkID {
			// Remove containerID fom array
			for idx, cont := range v.containers {
				if cont == containerID {
					v.containers[idx] = v.containers[len(v.containers)-1]
					v.containers = v.containers[:len(v.containers)-1]

					return nil
				}

				return errors.New("container not connected to network")
			}
		}
	}

	return ErrNetNotFound
}

// NetworkInspect returns network information
func (d *MockDCli) NetworkInspect(ctx context.Context, networkID string, verbose bool) (types.NetworkResource, error) {
	for _, v := range d.networks {
		if v.ID == networkID {
			cts := make(map[string]types.EndpointResource)

			for _, containers := range v.containers {
				cts[containers] = types.EndpointResource{
					IPv4Address: "127.0.0.2",
				}
			}
			return types.NetworkResource{
				Containers: cts,
			}, nil

		}
	}
	return types.NetworkResource{}, errors.New("No containers in network")

}

// NewMockDCli returns a new instance of MockDCli
func NewMockDCli() *MockDCli {
	return &MockDCli{
		running:    make(map[string]bool),
		containers: make(map[string]types.Container),
		images:     make(map[string]bool),
		networks:   make(map[string]mockNetwork),
	}
}
