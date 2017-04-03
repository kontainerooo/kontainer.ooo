package client

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"github.com/kontainerooo/kontainer.ooo/pkg/network"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
)

// New creates a set of endpoints based on a gRPC connection
func New(conn *grpc.ClientConn, logger log.Logger) *network.Endpoints {
	var CreatePrimaryNetworkForContainerEndpoint endpoint.Endpoint
	{
		CreatePrimaryNetworkForContainerEndpoint = grpctransport.NewClient(
			conn,
			"networkService",
			"CreatePrimaryNetworkForContainer",
			EncodeGRPCCreatePrimaryNetworkForContainerRequest,
			DecodeGRPCCreatePrimaryNetworkForContainerResponse,
			pb.CreatePrimaryNetworkForContainerResponse{},
		).Endpoint()
	}
	var CreateNetworkEndpoint endpoint.Endpoint
	{
		CreateNetworkEndpoint = grpctransport.NewClient(
			conn,
			"networkService",
			"CreateNetwork",
			EncodeGRPCCreateNetworkRequest,
			DecodeGRPCCreateNetworkResponse,
			pb.CreateNetworkResponse{},
		).Endpoint()
	}
	var RemoveNetworkByNameEndpoint endpoint.Endpoint
	{
		RemoveNetworkByNameEndpoint = grpctransport.NewClient(
			conn,
			"networkService",
			"RemoveNetworkByName",
			EncodeGRPCRemoveNetworkByNameRequest,
			DecodeGRPCRemoveNetworkByNameResponse,
			pb.RemoveNetworkByNameResponse{},
		).Endpoint()
	}
	var AddContainerToNetworkEndpoint endpoint.Endpoint
	{
		AddContainerToNetworkEndpoint = grpctransport.NewClient(
			conn,
			"networkService",
			"AddContainerToNetwork",
			EncodeGRPCAddContainerToNetworkRequest,
			DecodeGRPCAddContainerToNetworkResponse,
			pb.AddContainerToNetworkResponse{},
		).Endpoint()
	}
	var RemoveContainerFromNetworkEndpoint endpoint.Endpoint
	{
		RemoveContainerFromNetworkEndpoint = grpctransport.NewClient(
			conn,
			"networkService",
			"RemoveContainerFromNetwork",
			EncodeGRPCRemoveContainerFromNetworkRequest,
			DecodeGRPCRemoveContainerFromNetworkResponse,
			pb.RemoveContainerFromNetworkResponse{},
		).Endpoint()
	}
	var ExposePortToContainerEndpoint endpoint.Endpoint
	{
		ExposePortToContainerEndpoint = grpctransport.NewClient(
			conn,
			"networkService",
			"ExposePortToContainer",
			EncodeGRPCExposePortToContainerRequest,
			DecodeGRPCExposePortToContainerResponse,
			pb.ExposePortToContainerResponse{},
		).Endpoint()
	}
	var RemovePortFromContainerEndpoint endpoint.Endpoint
	{
		RemovePortFromContainerEndpoint = grpctransport.NewClient(
			conn,
			"networkService",
			"RemovePortFromContainer",
			EncodeGRPCRemovePortFromContainerRequest,
			DecodeGRPCRemovePortFromContainerResponse,
			pb.RemovePortFromContainerResponse{},
		).Endpoint()
	}

	return &network.Endpoints{
		CreatePrimaryNetworkForContainerEndpoint: CreatePrimaryNetworkForContainerEndpoint,
		CreateNetworkEndpoint:                    CreateNetworkEndpoint,
		RemoveNetworkByNameEndpoint:              RemoveNetworkByNameEndpoint,
		AddContainerToNetworkEndpoint:            AddContainerToNetworkEndpoint,
		RemoveContainerFromNetworkEndpoint:       RemoveContainerFromNetworkEndpoint,
		ExposePortToContainerEndpoint:            ExposePortToContainerEndpoint,
		RemovePortFromContainerEndpoint:          RemovePortFromContainerEndpoint,
	}
}

func getError(e string) error {
	if e != "" {
		return errors.New(e)
	}
	return nil
}

func nwConfigToPBConfig(c network.Config) *pb.NetworkConfig {
	return &pb.NetworkConfig{
		Driver: c.Driver,
		Name:   c.Name,
	}
}

// EncodeGRPCCreatePrimaryNetworkForContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain createprimarynetworkforcontainer request to a gRPC CreatePrimaryNetworkForContainer request.
func EncodeGRPCCreatePrimaryNetworkForContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*network.CreatePrimaryNetworkForContainerRequest)
	return &pb.CreatePrimaryNetworkForContainerRequest{
		Refid:       uint32(req.Refid),
		Config:      nwConfigToPBConfig(*req.Config),
		ContainerID: req.ContainerID,
	}, nil
}

// DecodeGRPCCreatePrimaryNetworkForContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC CreatePrimaryNetworkForContainer response to a messages/network.proto-domain createprimarynetworkforcontainer response.
func DecodeGRPCCreatePrimaryNetworkForContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.CreatePrimaryNetworkForContainerResponse)
	return &network.CreatePrimaryNetworkForContainerResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCCreateNetworkRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain createnetwork request to a gRPC CreateNetwork request.
func EncodeGRPCCreateNetworkRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*network.CreateNetworkRequest)
	return &pb.CreateNetworkRequest{
		Refid:  uint32(req.Refid),
		Config: nwConfigToPBConfig(*req.Config),
	}, nil
}

// DecodeGRPCCreateNetworkResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC CreateNetwork response to a messages/network.proto-domain createnetwork response.
func DecodeGRPCCreateNetworkResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.CreateNetworkResponse)
	return &network.CreateNetworkResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveNetworkByNameRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain removenetworkbyname request to a gRPC RemoveNetworkByName request.
func EncodeGRPCRemoveNetworkByNameRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*network.RemoveNetworkByNameRequest)
	return &pb.RemoveNetworkByNameRequest{
		Refid: uint32(req.Refid),
		Name:  req.Name,
	}, nil
}

// DecodeGRPCRemoveNetworkByNameResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveNetworkByName response to a messages/network.proto-domain removenetworkbyname response.
func DecodeGRPCRemoveNetworkByNameResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.RemoveNetworkByNameResponse)
	return &network.RemoveNetworkByNameResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCAddContainerToNetworkRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain addcontainertonetwork request to a gRPC AddContainerToNetwork request.
func EncodeGRPCAddContainerToNetworkRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*network.AddContainerToNetworkRequest)
	return &pb.AddContainerToNetworkRequest{
		Refid:       uint32(req.Refid),
		Name:        req.Name,
		ContainerID: req.ContainerID,
	}, nil
}

// DecodeGRPCAddContainerToNetworkResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC AddContainerToNetwork response to a messages/network.proto-domain addcontainertonetwork response.
func DecodeGRPCAddContainerToNetworkResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.AddContainerToNetworkResponse)
	return &network.AddContainerToNetworkResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveContainerFromNetworkRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain removecontainerfromnetwork request to a gRPC RemoveContainerFromNetwork request.
func EncodeGRPCRemoveContainerFromNetworkRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*network.RemoveContainerFromNetworkRequest)
	return &pb.RemoveContainerFromNetworkRequest{
		Refid:       uint32(req.Refid),
		Name:        req.Name,
		ContainerID: req.ContainerID,
	}, nil
}

// DecodeGRPCRemoveContainerFromNetworkResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveContainerFromNetwork response to a messages/network.proto-domain removecontainerfromnetwork response.
func DecodeGRPCRemoveContainerFromNetworkResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.RemoveContainerFromNetworkResponse)
	return &network.RemoveContainerFromNetworkResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCExposePortToContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain exposeporttocontainer request to a gRPC ExposePortToContainer request.
func EncodeGRPCExposePortToContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*network.ExposePortToContainerRequest)
	return &pb.ExposePortToContainerRequest{
		Refid:          uint32(req.Refid),
		SrcContainerID: req.SrcContainerID,
		DstContainerID: req.DstContainerID,
		Port:           uint32(req.Port),
		Protocol:       req.Protocol,
	}, nil
}

// DecodeGRPCExposePortToContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC ExposePortToContainer response to a messages/network.proto-domain exposeporttocontainer response.
func DecodeGRPCExposePortToContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.ExposePortToContainerResponse)
	return &network.ExposePortToContainerResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemovePortFromContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain removeportfromcontainer request to a gRPC RemovePortFromContainer request.
func EncodeGRPCRemovePortFromContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*network.RemovePortFromContainerRequest)
	return &pb.RemovePortFromContainerRequest{
		Refid:          uint32(req.Refid),
		SrcContainerID: req.SrcContainerID,
		DstContainerID: req.DstContainerID,
		Port:           uint32(req.Port),
		Protocol:       req.Protocol,
	}, nil
}

// DecodeGRPCRemovePortFromContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemovePortFromContainer response to a messages/network.proto-domain removeportfromcontainer response.
func DecodeGRPCRemovePortFromContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.RemovePortFromContainerResponse)
	return &network.RemovePortFromContainerResponse{
		Error: getError(response.Error),
	}, nil
}
