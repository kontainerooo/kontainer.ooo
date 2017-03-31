package firewall

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
)

// Endpoints is a struct which collects all endpoints for the firewall service
type Endpoints struct {
	InitBridgeEndpoint      endpoint.Endpoint
	AllowConnectionEndpoint endpoint.Endpoint
	BlockConnectionEndpoint endpoint.Endpoint
	AllowPortEndpoint       endpoint.Endpoint
	BlockPortEndpoint       endpoint.Endpoint
}

// InitBridgeRequest is the request struct for the InitBridgeEndpoint
type InitBridgeRequest struct {
	IP    abstraction.Inet
	NetIf string
}

// InitBridgeResponse is the response struct for the InitBridgeEndpoint
type InitBridgeResponse struct {
	Error error
}

// MakeInitBridgeEndpoint creates a gokit endpoint which invokes InitBridge
func MakeInitBridgeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InitBridgeRequest)
		err := s.InitBridge(req.IP, req.NetIf)
		return InitBridgeResponse{
			Error: err,
		}, nil
	}
}

// AllowConnectionRequest is the request struct for the AllowConnectionEndpoint
type AllowConnectionRequest struct {
	SrcIP      abstraction.Inet
	SrcNetwork string
	DstIP      abstraction.Inet
	DstNetwork string
}

// AllowConnectionResponse is the response struct for the AllowConnectionEndpoint
type AllowConnectionResponse struct {
	Error error
}

// MakeAllowConnectionEndpoint creates a gokit endpoint which invokes AllowConnection
func MakeAllowConnectionEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AllowConnectionRequest)
		err := s.AllowConnection(req.SrcIP, req.SrcNetwork, req.DstIP, req.DstNetwork)
		return AllowConnectionResponse{
			Error: err,
		}, nil
	}
}

// BlockConnectionRequest is the request struct for the BlockConnectionEndpoint
type BlockConnectionRequest struct {
	SrcIP      abstraction.Inet
	SrcNetwork string
	DstIP      abstraction.Inet
	DstNetwork string
}

// BlockConnectionResponse is the response struct for the BlockConnectionEndpoint
type BlockConnectionResponse struct {
	Error error
}

// MakeBlockConnectionEndpoint creates a gokit endpoint which invokes BlockConnection
func MakeBlockConnectionEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BlockConnectionRequest)
		err := s.BlockConnection(req.SrcIP, req.SrcNetwork, req.DstIP, req.DstNetwork)
		return BlockConnectionResponse{
			Error: err,
		}, nil
	}
}

// AllowPortRequest is the request struct for the AllowPortEndpoint
type AllowPortRequest struct {
	SrcIP    abstraction.Inet
	SrcNw    string
	DstIP    abstraction.Inet
	DstNw    string
	Port     uint16
	Protocol string
}

// AllowPortResponse is the response struct for the AllowPortEndpoint
type AllowPortResponse struct {
	Error error
}

// MakeAllowPortEndpoint creates a gokit endpoint which invokes AllowPort
func MakeAllowPortEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AllowPortRequest)
		err := s.AllowPort(req.SrcIP, req.SrcNetwork, req.DstIP, req.DstNetwork, req.Port, req.Protocol)
		return AllowPortResponse{
			Error: err,
		}, nil
	}
}

// BlockPortRequest is the request struct for the BlockPortEndpoint
type BlockPortRequest struct {
	SrcIP    abstraction.Inet
	SrcNw    string
	DstIP    abstraction.Inet
	DstNw    string
	Port     uint16
	Protocol string
}

// BlockPortResponse is the response struct for the BlockPortEndpoint
type BlockPortResponse struct {
	Error error
}

// MakeBlockPortEndpoint creates a gokit endpoint which invokes BlockPort
func MakeBlockPortEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BlockPortRequest)
		err := s.BlockPort(req.SrcIP, req.SrcNetwork, req.DstIP, req.DstNetwork, req.Port, req.Protocol)
		return BlockPortResponse{
			Error: err,
		}, nil
	}
}
