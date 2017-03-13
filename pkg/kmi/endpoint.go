package kmi

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a struct which collects all endpoints for the kmi service
type Endpoints struct {
	AddKMIEndpoint    endpoint.Endpoint
	RemoveKMIEndpoint endpoint.Endpoint
	GetKMIEndpoint    endpoint.Endpoint
	KMIEndpoint       endpoint.Endpoint
}

// AddKMIRequest is the request struct for the AddKMIEndpoint
type AddKMIRequest struct {
	Path string
}

// AddKMIResponse is the response struct for the AddKMIEndpoint
type AddKMIResponse struct {
	ID    uint
	Error error
}

// MakeAddKMIEndpoint creates a gokit endpoint which invokes AddKMI
func MakeAddKMIEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddKMIRequest)
		id, err := s.AddKMI(req.Path)
		return AddKMIResponse{
			ID:    id,
			Error: err,
		}, nil
	}
}

// RemoveKMIRequest is the request struct for the RemoveKMIEndpoint
type RemoveKMIRequest struct {
	ID uint
}

// RemoveKMIResponse is the response struct for the RemoveKMIEndpoint
type RemoveKMIResponse struct {
	Error error
}

// MakeRemoveKMIEndpoint creates a gokit endpoint which invokes RemoveKMI
func MakeRemoveKMIEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveKMIRequest)
		err := s.RemoveKMI(req.ID)
		return RemoveKMIResponse{
			Error: err,
		}, nil
	}
}

// GetKMIRequest is the request struct for the GetKMIEndpoint
type GetKMIRequest struct {
	ID uint
}

// GetKMIResponse is the response struct for the GetKMIEndpoint
type GetKMIResponse struct {
	KMI   *KMI
	Error error
}

// MakeGetKMIEndpoint creates a gokit endpoint which invokes GetKMI
func MakeGetKMIEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		k := &KMI{}
		req := request.(GetKMIRequest)
		err := s.GetKMI(req.ID, k)
		return GetKMIResponse{
			KMI:   k,
			Error: err,
		}, nil
	}
}

// KMIRequest is the request struct for the KMIEndpoint
type KMIRequest struct{}

// KMIResponse is the response struct for the KMIEndpoint
type KMIResponse struct {
	KMDI  *[]KMDI
	Error error
}

// MakeKMIEndpoint creates a gokit endpoint which invokes KMI
func MakeKMIEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		k := &[]KMDI{}
		err := s.KMI(k)
		return KMIResponse{
			KMDI:  k,
			Error: err,
		}, nil
	}
}
