package network

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a struct which collects all endpoints for the network service
type Endpoints struct {
	CreatePrimaryNetworkForContainerEndpoint endpoint.Endpoint
	CreateNetworkEndpoint                    endpoint.Endpoint
	RemoveNetworkByNameEndpoint              endpoint.Endpoint
	AddContainerToNetworkEndpoint            endpoint.Endpoint
	RemoveContainerFromNetworkEndpoint       endpoint.Endpoint
	ExposePortToContainerEndpoint            endpoint.Endpoint
	RemovePortFromContainerEndpoint          endpoint.Endpoint
}

// CreatePrimaryNetworkForContainerRequest is the request struct for the CreatePrimaryNetworkForContainerEndpoint
type CreatePrimaryNetworkForContainerRequest struct {
	RefID       uint
	Config      *Config
	ContainerID string
}

// CreatePrimaryNetworkForContainerResponse is the response struct for the CreatePrimaryNetworkForContainerEndpoint
type CreatePrimaryNetworkForContainerResponse struct {
	Error error
}

// MakeCreatePrimaryNetworkForContainerEndpoint creates a gokit endpoint which invokes CreatePrimaryNetworkForContainer
func MakeCreatePrimaryNetworkForContainerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePrimaryNetworkForContainerRequest)
		err := s.CreatePrimaryNetworkForContainer(req.RefID, req.Config, req.ContainerID)
		return CreatePrimaryNetworkForContainerResponse{
			Error: err,
		}, nil
	}
}

// CreateNetworkRequest is the request struct for the CreateNetworkEndpoint
type CreateNetworkRequest struct {
	RefID  uint
	Config *Config
}

// CreateNetworkResponse is the response struct for the CreateNetworkEndpoint
type CreateNetworkResponse struct {
	Error error
}

// MakeCreateNetworkEndpoint creates a gokit endpoint which invokes CreateNetwork
func MakeCreateNetworkEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateNetworkRequest)
		err := s.CreateNetwork(req.RefID, req.Config)
		return CreateNetworkResponse{
			Error: err,
		}, nil
	}
}

// RemoveNetworkByNameRequest is the request struct for the RemoveNetworkByNameEndpoint
type RemoveNetworkByNameRequest struct {
	RefID uint
	Name  string
}

// RemoveNetworkByNameResponse is the response struct for the RemoveNetworkByNameEndpoint
type RemoveNetworkByNameResponse struct {
	Error error
}

// MakeRemoveNetworkByNameEndpoint creates a gokit endpoint which invokes RemoveNetworkByName
func MakeRemoveNetworkByNameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveNetworkByNameRequest)
		err := s.RemoveNetworkByName(req.RefID, req.Name)
		return RemoveNetworkByNameResponse{
			Error: err,
		}, nil
	}
}

// AddContainerToNetworkRequest is the request struct for the AddContainerToNetworkEndpoint
type AddContainerToNetworkRequest struct {
	RefID       uint
	Name        string
	ContainerID string
}

// AddContainerToNetworkResponse is the response struct for the AddContainerToNetworkEndpoint
type AddContainerToNetworkResponse struct {
	Error error
}

// MakeAddContainerToNetworkEndpoint creates a gokit endpoint which invokes AddContainerToNetwork
func MakeAddContainerToNetworkEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddContainerToNetworkRequest)
		err := s.AddContainerToNetwork(req.RefID, req.Name, req.ContainerID)
		return AddContainerToNetworkResponse{
			Error: err,
		}, nil
	}
}

// RemoveContainerFromNetworkRequest is the request struct for the RemoveContainerFromNetworkEndpoint
type RemoveContainerFromNetworkRequest struct {
	RefID       uint
	Name        string
	ContainerID string
}

// RemoveContainerFromNetworkResponse is the response struct for the RemoveContainerFromNetworkEndpoint
type RemoveContainerFromNetworkResponse struct {
	Error error
}

// MakeRemoveContainerFromNetworkEndpoint creates a gokit endpoint which invokes RemoveContainerFromNetwork
func MakeRemoveContainerFromNetworkEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveContainerFromNetworkRequest)
		err := s.RemoveContainerFromNetwork(req.RefID, req.Name, req.ContainerID)
		return RemoveContainerFromNetworkResponse{
			Error: err,
		}, nil
	}
}

// ExposePortToContainerRequest is the request struct for the ExposePortToContainerEndpoint
type ExposePortToContainerRequest struct {
	RefID          uint
	SrcContainerID string
	Port           uint16
	Protocol       string
	DstContainerID string
}

// ExposePortToContainerResponse is the response struct for the ExposePortToContainerEndpoint
type ExposePortToContainerResponse struct {
	Error error
}

// MakeExposePortToContainerEndpoint creates a gokit endpoint which invokes ExposePortToContainer
func MakeExposePortToContainerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ExposePortToContainerRequest)
		err := s.ExposePortToContainer(req.RefID, req.SrcContainerID, req.Port, req.Protocol, req.DstContainerID)
		return ExposePortToContainerResponse{
			Error: err,
		}, nil
	}
}

// RemovePortFromContainerRequest is the request struct for the RemovePortFromContainerEndpoint
type RemovePortFromContainerRequest struct {
	RefID          uint
	SrcContainerID string
	Port           uint16
	Protocol       string
	DstContainerID string
}

// RemovePortFromContainerResponse is the response struct for the RemovePortFromContainerEndpoint
type RemovePortFromContainerResponse struct {
	Error error
}

// MakeRemovePortFromContainerEndpoint creates a gokit endpoint which invokes RemovePortFromContainer
func MakeRemovePortFromContainerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemovePortFromContainerRequest)
		err := s.RemovePortFromContainer(req.RefID, req.SrcContainerID, req.Port, req.Protocol, req.DstContainerID)
		return RemovePortFromContainerResponse{
			Error: err,
		}, nil
	}
}
