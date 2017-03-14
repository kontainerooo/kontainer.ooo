package containerlifecycle

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a struct which collects all endpoints for the containerlifecycle service
type Endpoints struct {
	StartContainerEndpoint endpoint.Endpoint
	StartCommandEndpoint   endpoint.Endpoint
	StopContainerEndpoint  endpoint.Endpoint
}

// StartContainerRequest is the request struct for the StartContainerEndpoint
type StartContainerRequest struct {
	ID string
}

// StartContainerResponse is the response struct for the StartContainerEndpoint
type StartContainerResponse struct {
	Error error
}

// MakeStartContainerEndpoint creates a gokit endpoint which invokes StartContainer
func MakeStartContainerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StartContainerRequest)
		err := s.StartContainer(req.ID)
		return StartContainerResponse{
			Error: err,
		}, nil
	}
}

// StartCommandRequest is the request struct for the StartCommandEndpoint
type StartCommandRequest struct {
	ID  string
	Cmd string
}

// StartCommandResponse is the response struct for the StartCommandEndpoint
type StartCommandResponse struct {
	ID    string
	Error error
}

// MakeStartCommandEndpoint creates a gokit endpoint which invokes StartCommand
func MakeStartCommandEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StartCommandRequest)
		id, err := s.StartCommand(req.ID, req.Cmd)
		return StartCommandResponse{
			ID:    id,
			Error: err,
		}, nil
	}
}

// StopContainerRequest is the request struct for the StopContainerEndpoint
type StopContainerRequest struct {
	ID string
}

// StopContainerResponse is the response struct for the StopContainerEndpoint
type StopContainerResponse struct {
	Error error
}

// MakeStopContainerEndpoint creates a gokit endpoint which invokes StopContainer
func MakeStopContainerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StopContainerRequest)
		err := s.StopContainer(req.ID)
		return StopContainerResponse{
			Error: err,
		}, nil
	}
}
