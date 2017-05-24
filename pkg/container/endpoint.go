package container

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a struct which collects all endpoints for the container service
type Endpoints struct {
	CreateContainerEndpoint endpoint.Endpoint
	RemoveContainerEndpoint endpoint.Endpoint
	InstancesEndpoint       endpoint.Endpoint
	StartContainerEndpoint  endpoint.Endpoint
	StopContainerEndpoint   endpoint.Endpoint
	ExecuteEndpoint         endpoint.Endpoint
}

// CreateContainerRequest is the request struct for the CreateContainerEndpoint
type CreateContainerRequest struct {
	RefID uint `bart:"ref"`
	KmiID uint
	Name  string
}

// CreateContainerResponse is the response struct for the CreateContainerEndpoint
type CreateContainerResponse struct {
	ID    string
	Error error
}

// MakeCreateContainerEndpoint creates a gokit endpoint which invokes CreateContainer
func MakeCreateContainerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateContainerRequest)
		id, err := s.CreateContainer(req.RefID, req.KmiID, req.Name)
		return CreateContainerResponse{
			ID:    id,
			Error: err,
		}, nil
	}
}

// RemoveContainerRequest is the request struct for the RemoveContainerEndpoint
type RemoveContainerRequest struct {
	RefID uint `bart:"ref"`
	ID    string
}

// RemoveContainerResponse is the response struct for the RemoveContainerEndpoint
type RemoveContainerResponse struct {
	Error error
}

// MakeRemoveContainerEndpoint creates a gokit endpoint which invokes RemoveContainer
func MakeRemoveContainerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveContainerRequest)
		err := s.RemoveContainer(req.RefID, req.ID)
		return RemoveContainerResponse{
			Error: err,
		}, nil
	}
}

// InstancesRequest is the request struct for the InstancesEndpoint
type InstancesRequest struct {
	RefID uint `bart:"ref"`
}

// InstancesResponse is the response struct for the InstancesEndpoint
type InstancesResponse struct {
	Containers []Container
}

// MakeInstancesEndpoint creates a gokit endpoint which invokes Instances
func MakeInstancesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InstancesRequest)
		cnt := s.Instances(req.RefID)
		return InstancesResponse{
			Containers: cnt,
		}, nil
	}
}

// StopContainerRequest is the request struct for the StopContainerEndpoint
type StopContainerRequest struct {
	RefID uint `bart:"ref"`
	ID    string
}

// StopContainerResponse is the response struct for the StopContainerEndpoint
type StopContainerResponse struct {
	Error error
}

// MakeStopContainerEndpoint creates a gokit endpoint which invokes StopContainer
func MakeStopContainerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StopContainerRequest)
		err := s.StopContainer(req.RefID, req.ID)
		return StopContainerResponse{
			Error: err,
		}, nil
	}
}

// ExecuteRequest is the request struct for the ExecuteEndpoint
type ExecuteRequest struct {
	RefID uint `bart:"ref"`
	ID    string
	CMD   string
}

// ExecuteResponse is the response struct for the ExecuteEndpoint
type ExecuteResponse struct {
	Response string
	Error    error
}

// MakeExecuteEndpoint creates a gokit endpoint which invokes Execute
func MakeExecuteEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ExecuteRequest)
		res, err := s.Execute(req.RefID, req.ID, req.CMD)
		return ExecuteResponse{
			Response: res,
			Error:    err,
		}, nil
	}
}
