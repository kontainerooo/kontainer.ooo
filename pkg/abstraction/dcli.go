package abstraction

import (
	"context"

	engine "github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
)

// DCli is an interface to abstract dockers engine api client
type DCli interface {
	// Every function below invokes engine api client's function and returns its error property
	ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error
	ContainerKill(ctx context.Context, container, signal string) error
	ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (string, error)
}

type dcliAbstract struct {
	cli engine.Client
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
