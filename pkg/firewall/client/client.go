package client

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"github.com/kontainerooo/kontainer.ooo/pkg/firewall"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
)

// New creates a set of endpoints based on a gRPC connection
func New(conn *grpc.ClientConn, logger log.Logger) *firewall.Endpoints {
	var InitBridgeEndpoint endpoint.Endpoint
	{
		InitBridgeEndpoint = grpctransport.NewClient(
			conn,
			"firewallService",
			"InitBridge",
			EncodeGRPCInitBridgeRequest,
			DecodeGRPCInitBridgeResponse,
			pb.InitBridgeResponse{},
		).Endpoint()
	}
	var AllowConnectionEndpoint endpoint.Endpoint
	{
		AllowConnectionEndpoint = grpctransport.NewClient(
			conn,
			"firewallService",
			"AllowConnection",
			EncodeGRPCAllowConnectionRequest,
			DecodeGRPCAllowConnectionResponse,
			pb.AllowConnectionResponse{},
		).Endpoint()
	}
	var BlockConnectionEndpoint endpoint.Endpoint
	{
		BlockConnectionEndpoint = grpctransport.NewClient(
			conn,
			"firewallService",
			"BlockConnection",
			EncodeGRPCBlockConnectionRequest,
			DecodeGRPCBlockConnectionResponse,
			pb.BlockConnectionResponse{},
		).Endpoint()
	}
	var AllowPortEndpoint endpoint.Endpoint
	{
		AllowPortEndpoint = grpctransport.NewClient(
			conn,
			"firewallService",
			"AllowPort",
			EncodeGRPCAllowPortRequest,
			DecodeGRPCAllowPortResponse,
			pb.AllowPortResponse{},
		).Endpoint()
	}
	var BlockPortEndpoint endpoint.Endpoint
	{
		BlockPortEndpoint = grpctransport.NewClient(
			conn,
			"firewallService",
			"BlockPort",
			EncodeGRPCBlockPortRequest,
			DecodeGRPCBlockPortResponse,
			pb.BlockPortResponse{},
		).Endpoint()
	}

	return &firewall.Endpoints{
		InitBridgeEndpoint:      InitBridgeEndpoint,
		AllowConnectionEndpoint: AllowConnectionEndpoint,
		BlockConnectionEndpoint: BlockConnectionEndpoint,
		AllowPortEndpoint:       AllowPortEndpoint,
		BlockPortEndpoint:       BlockPortEndpoint,
	}
}

func getError(e string) error {
	if e != "" {
		return errors.New(e)
	}
	return nil
}

// EncodeGRPCInitBridgeRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/firewall.proto-domain initbridge request to a gRPC InitBridge request.
func EncodeGRPCInitBridgeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*firewall.InitBridgeRequest)
	return &pb.InitBridgeRequest{
		IP:          string(req.IP),
		NetworkName: req.NetIf,
	}, nil
}

// DecodeGRPCInitBridgeResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC InitBridge response to a messages/firewall.proto-domain initbridge response.
func DecodeGRPCInitBridgeResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.InitBridgeResponse)
	return &firewall.InitBridgeResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCAllowConnectionRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/firewall.proto-domain allowconnection request to a gRPC AllowConnection request.
func EncodeGRPCAllowConnectionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*firewall.AllowConnectionRequest)
	return &pb.AllowConnectionRequest{
		SrcIP:      string(req.SrcIP),
		DstIP:      string(req.DstIP),
		SrcNetwork: req.SrcNetwork,
		DstNetwork: req.DstNetwork,
	}, nil
}

// DecodeGRPCAllowConnectionResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC AllowConnection response to a messages/firewall.proto-domain allowconnection response.
func DecodeGRPCAllowConnectionResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.AllowConnectionResponse)
	return &firewall.AllowConnectionResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCBlockConnectionRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/firewall.proto-domain blockconnection request to a gRPC BlockConnection request.
func EncodeGRPCBlockConnectionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*firewall.BlockConnectionRequest)
	return &pb.BlockConnectionRequest{
		SrcIP:      string(req.SrcIP),
		DstIP:      string(req.DstIP),
		SrcNetwork: req.SrcNetwork,
		DstNetwork: req.DstNetwork,
	}, nil
}

// DecodeGRPCBlockConnectionResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC BlockConnection response to a messages/firewall.proto-domain blockconnection response.
func DecodeGRPCBlockConnectionResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.BlockConnectionResponse)
	return &firewall.BlockConnectionResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCAllowPortRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/firewall.proto-domain allowport request to a gRPC AllowPort request.
func EncodeGRPCAllowPortRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*firewall.AllowPortRequest)
	return &pb.AllowPortRequest{
		SrcIP:      string(req.SrcIP),
		DstIP:      string(req.DstIP),
		SrcNetwork: req.SrcNetwork,
		DstNetwork: req.DstNetwork,
		Protocol:   req.Protocol,
		Port:       uint32(req.Port),
	}, nil
}

// DecodeGRPCAllowPortResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC AllowPort response to a messages/firewall.proto-domain allowport response.
func DecodeGRPCAllowPortResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.AllowPortResponse)
	return &firewall.AllowPortResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCBlockPortRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/firewall.proto-domain blockport request to a gRPC BlockPort request.
func EncodeGRPCBlockPortRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*firewall.BlockPortRequest)
	return &pb.BlockPortRequest{
		SrcIP:      string(req.SrcIP),
		DstIP:      string(req.DstIP),
		SrcNetwork: req.SrcNetwork,
		DstNetwork: req.DstNetwork,
		Protocol:   req.Protocol,
		Port:       uint32(req.Port),
	}, nil
}

// DecodeGRPCBlockPortResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC BlockPort response to a messages/firewall.proto-domain blockport response.
func DecodeGRPCBlockPortResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.BlockPortResponse)
	return &firewall.BlockPortResponse{
		Error: getError(response.Error),
	}, nil
}
