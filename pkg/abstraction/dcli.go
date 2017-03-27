package abstraction

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	engine "github.com/docker/docker/client"
)

// DCli is an interface to abstract dockers engine api client
type DCli interface {
	// Every function below invokes engine api client's function and returns its error property
	ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error
	ContainerKill(ctx context.Context, container, signal string) error
	ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (string, error)
	ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error)
	ContainerRename(ctx context.Context, containerID, newContainerName string) error
	ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error
	ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
	ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error)
	ImageBuild(ctx context.Context, buildContext io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error)
	IsErrImageNotFound(err error) bool
	NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error)
	NetworkRemove(ctx context.Context, networkID string) error
	NetworkConnect(ctx context.Context, networkID, containerID string, config *network.EndpointSettings) error
	NetworkDisconnect(ctx context.Context, networkID, containerID string, force bool) error
}

type dcliAbstract struct {
	cli *engine.Client
}

func (d dcliAbstract) ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error {
	return d.cli.ContainerStart(ctx, container, options)
}

func (d dcliAbstract) ContainerKill(ctx context.Context, container, signal string) error {
	return d.cli.ContainerKill(ctx, container, signal)
}

func (d dcliAbstract) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (string, error) {
	idStruct, err := d.cli.ContainerExecCreate(ctx, container, config)
	return idStruct.ID, err
}

func (d dcliAbstract) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	return d.cli.ContainerCreate(ctx, config, hostConfig, networkingConfig, containerName)
}

func (d dcliAbstract) ContainerRename(ctx context.Context, containerID, newContainerName string) error {
	return d.cli.ContainerRename(ctx, containerID, newContainerName)
}

func (d dcliAbstract) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error {
	return d.cli.ContainerRemove(ctx, containerID, options)
}

func (d dcliAbstract) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	return d.cli.ContainerList(ctx, options)
}

func (d dcliAbstract) ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error) {
	return d.cli.ImageInspectWithRaw(ctx, imageID)
}

func (d dcliAbstract) ImageBuild(ctx context.Context, buildContext io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error) {
	return d.cli.ImageBuild(ctx, buildContext, options)
}

func (d dcliAbstract) IsErrImageNotFound(err error) bool {
	return engine.IsErrImageNotFound(err)
}

func (d dcliAbstract) NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error) {
	return d.cli.NetworkCreate(ctx, name, options)
}

func (d dcliAbstract) NetworkRemove(ctx context.Context, networkID string) error {
	return d.cli.NetworkRemove(ctx, networkID)
}

func (d dcliAbstract) NetworkConnect(ctx context.Context, networkID, containerID string, config *network.EndpointSettings) error {
	return d.cli.NetworkConnect(ctx, networkID, containerID, config)
}

func (d dcliAbstract) NetworkDisconnect(ctx context.Context, networkID, containerID string, force bool) error {
	return d.cli.NetworkDisconnect(ctx, networkID, containerID, force)
}

// NewDCLI returns an new Wrapper instance
func NewDCLI(dcli *engine.Client) DCli {
	return &dcliAbstract{
		cli: dcli,
	}
}
