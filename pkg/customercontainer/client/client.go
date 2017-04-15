package client

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"github.com/kontainerooo/kontainer.ooo/pkg/customercontainer"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
)

// New creates a set of endpoints based on a gRPC connection
func New(conn *grpc.ClientConn, logger log.Logger) *customercontainer.Endpoints {
	var CreateContainerEndpoint endpoint.Endpoint
	{
		CreateContainerEndpoint = grpctransport.NewClient(
			conn,
			"customercontainerService",
			"CreateContainer",
			EncodeGRPCCreateContainerRequest,
			DecodeGRPCCreateContainerResponse,
			pb.CreateContainerResponse{},
		).Endpoint()
	}
	var EditContainerEndpoint endpoint.Endpoint
	{
		EditContainerEndpoint = grpctransport.NewClient(
			conn,
			"customercontainerService",
			"EditContainer",
			EncodeGRPCEditContainerRequest,
			DecodeGRPCEditContainerResponse,
			pb.EditContainerResponse{},
		).Endpoint()
	}
	var RemoveContainerEndpoint endpoint.Endpoint
	{
		RemoveContainerEndpoint = grpctransport.NewClient(
			conn,
			"customercontainerService",
			"RemoveContainer",
			EncodeGRPCRemoveContainerRequest,
			DecodeGRPCRemoveContainerResponse,
			pb.RemoveContainerResponse{},
		).Endpoint()
	}
	var InstancesEndpoint endpoint.Endpoint
	{
		InstancesEndpoint = grpctransport.NewClient(
			conn,
			"customercontainerService",
			"Instances",
			EncodeGRPCInstancesRequest,
			DecodeGRPCInstancesResponse,
			pb.InstancesResponse{},
		).Endpoint()
	}
	var GetContainerIDEndpoint endpoint.Endpoint
	{
		GetContainerIDEndpoint = grpctransport.NewClient(
			conn,
			"customercontainerService",
			"GetContainerID",
			EncodeGRPCGetContainerIDRequest,
			DecodeGRPCGetContainerIDResponse,
			pb.GetContainerIDResponse{},
		).Endpoint()
	}

	return &customercontainer.Endpoints{
		CreateContainerEndpoint: CreateContainerEndpoint,
		EditContainerEndpoint:   EditContainerEndpoint,
		RemoveContainerEndpoint: RemoveContainerEndpoint,
		InstancesEndpoint:       InstancesEndpoint,
		GetContainerIDEndpoint:  GetContainerIDEndpoint,
	}
}

func getError(e string) error {
	if e != "" {
		return errors.New(e)
	}
	return nil
}

func configToPbConfig(cfg *customercontainer.ContainerConfig) *pb.ContainerConfig {
	return &pb.ContainerConfig{
		ImageName: cfg.ImageName,
	}
}

// EncodeGRPCCreateContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain createcontainer request to a gRPC CreateContainer request.
func EncodeGRPCCreateContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*customercontainer.CreateContainerRequest)
	return &pb.CreateContainerRequest{
		RefID: uint32(req.RefID),
		KmiID: uint32(req.KMIID),
	}, nil
}

// DecodeGRPCCreateContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC CreateContainer response to a messages/customercontainer.proto-domain createcontainer response.
func DecodeGRPCCreateContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.CreateContainerResponse)
	return &customercontainer.CreateContainerResponse{
		Error: getError(response.Error),
		Name:  response.Name,
		ID:    response.ID,
	}, nil
}

// EncodeGRPCEditContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain editcontainer request to a gRPC EditContainer request.
func EncodeGRPCEditContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*customercontainer.EditContainerRequest)
	return &pb.EditContainerRequest{
		ID:  req.ID,
		Cfg: configToPbConfig(req.Cfg),
	}, nil
}

// DecodeGRPCEditContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC EditContainer response to a messages/customercontainer.proto-domain editcontainer response.
func DecodeGRPCEditContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.EditContainerResponse)
	return &customercontainer.EditContainerResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain removecontainer request to a gRPC RemoveContainer request.
func EncodeGRPCRemoveContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*customercontainer.RemoveContainerRequest)
	return &pb.RemoveContainerRequest{
		ID: req.ID,
	}, nil
}

// DecodeGRPCRemoveContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveContainer response to a messages/customercontainer.proto-domain removecontainer response.
func DecodeGRPCRemoveContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.RemoveContainerResponse)
	return &customercontainer.RemoveContainerResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCInstancesRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain instances request to a gRPC Instances request.
func EncodeGRPCInstancesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*customercontainer.InstancesRequest)
	return &pb.InstancesRequest{
		RefID: uint32(req.RefID),
	}, nil
}

// DecodeGRPCInstancesResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC Instances response to a messages/customercontainer.proto-domain instances response.
func DecodeGRPCInstancesResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.InstancesResponse)
	return &customercontainer.InstancesResponse{
		Instances: response.Instances,
	}, nil
}

// EncodeGRPCGetContainerIDRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain getcontainerid request to a gRPC GetContainerID request.
func EncodeGRPCGetContainerIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*customercontainer.GetContainerIDRequest)
	return &pb.GetContainerIDRequest{
		RefID:         uint32(req.RefID),
		ContainerName: req.ContainerName,
	}, nil
}

// DecodeGRPCGetContainerIDResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetContainerID response to a messages/customercontainer.proto-domain getcontainerid response.
func DecodeGRPCGetContainerIDResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.GetContainerIDResponse)
	return &customercontainer.GetContainerIDResponse{
		ContainerID: response.ContainerID,
		Error:       getError(response.Error),
	}, nil
}
