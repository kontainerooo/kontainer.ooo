package customercontainer

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a struct which collects all endpoints for the customercontainer service
type Endpoints struct {
	CreateContainerEndpoint endpoint.Endpoint
	EditContainerEndpoint   endpoint.Endpoint
	RemoveContainerEndpoint endpoint.Endpoint
	InstancesEndpoint       endpoint.Endpoint
}

// CreateContainerRequest is the request struct for the CreateContainerEndpoint
type CreateContainerRequest struct {
	Refid int
	Cfg   *ContainerConfig
}

// CreateContainerResponse is the response struct for the CreateContainerEndpoint
type CreateContainerResponse struct {
	Name  string
	ID    string
	Error error
}

// MakeCreateContainerEndpoint creates a gokit endpoint which invokes CreateContainer
func MakeCreateContainerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateContainerRequest)
		name, id, err := s.CreateContainer(req.Refid, req.Cfg)
		return CreateContainerResponse{
			Error: err,
			ID:    id,
			Name:  name,
		}, nil
	}
}

// EditContainerRequest is the request struct for the EditContainerEndpoint
type EditContainerRequest struct {
	ID  string
	Cfg *ContainerConfig
}

// EditContainerResponse is the response struct for the EditContainerEndpoint
type EditContainerResponse struct {
	Error error
}

// MakeEditContainerEndpoint creates a gokit endpoint which invokes EditContainer
func MakeEditContainerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EditContainerRequest)
		err := s.EditContainer(req.ID, req.Cfg)
		return EditContainerResponse{
			Error: err,
		}, nil
	}
}

// RemoveContainerRequest is the request struct for the RemoveContainerEndpoint
type RemoveContainerRequest struct {
	ID string
}

// RemoveContainerResponse is the response struct for the RemoveContainerEndpoint
type RemoveContainerResponse struct {
	Error error
}

// MakeRemoveContainerEndpoint creates a gokit endpoint which invokes RemoveContainer
func MakeRemoveContainerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveContainerRequest)
		err := s.RemoveContainer(req.ID)
		return RemoveContainerResponse{
			Error: err,
		}, nil
	}
}

// InstancesRequest is the request struct for the InstancesEndpoint
type InstancesRequest struct {
	Refid int
}

// InstancesResponse is the response struct for the InstancesEndpoint
type InstancesResponse struct {
	Instances []string
}

// MakeInstancesEndpoint creates a gokit endpoint which invokes Instances
func MakeInstancesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InstancesRequest)
		inst := s.Instances(req.Refid)
		return InstancesResponse{
			Instances: inst,
		}, nil
	}
}
