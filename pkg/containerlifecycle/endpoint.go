package containerlifecycle

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a struct which collects all endpoints for the containerlifecycle service
type Endpoints struct {
	
			StartContainerEndpoint endpoint.Endpoint
	
			StartCommandEndpoint endpoint.Endpoint
	
			StopContainerEndpoint endpoint.Endpoint
	
}



// StartContainerRequest is the request struct for the StartContainerEndpoint
type StartContainerRequest struct {

}

// StartContainerResponse is the response struct for the StartContainerEndpoint
type StartContainerResponse struct {

}

// MakeStartContainerEndpoint creates a gokit endpoint which invokes StartContainer
func MakeStartContainerEndpoint(s Service) endpoint.Endpoint {
   return func(ctx context.Context, request interface{}) (interface{}, error) {
	     req := request.(StartContainerRequest)
	     err := s.StartContainer()
	     return StartContainerResponse{

       }, nil
   }
}

// StartCommandRequest is the request struct for the StartCommandEndpoint
type StartCommandRequest struct {

}

// StartCommandResponse is the response struct for the StartCommandEndpoint
type StartCommandResponse struct {

}

// MakeStartCommandEndpoint creates a gokit endpoint which invokes StartCommand
func MakeStartCommandEndpoint(s Service) endpoint.Endpoint {
   return func(ctx context.Context, request interface{}) (interface{}, error) {
	     req := request.(StartCommandRequest)
	     err := s.StartCommand()
	     return StartCommandResponse{

       }, nil
   }
}

// StopContainerRequest is the request struct for the StopContainerEndpoint
type StopContainerRequest struct {

}

// StopContainerResponse is the response struct for the StopContainerEndpoint
type StopContainerResponse struct {

}

// MakeStopContainerEndpoint creates a gokit endpoint which invokes StopContainer
func MakeStopContainerEndpoint(s Service) endpoint.Endpoint {
   return func(ctx context.Context, request interface{}) (interface{}, error) {
	     req := request.(StopContainerRequest)
	     err := s.StopContainer()
	     return StopContainerResponse{

       }, nil
   }
}

