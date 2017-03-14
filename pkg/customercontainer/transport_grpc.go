package customercontainer

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/ttdennis/kontainer.io/pkg/pb"
	oldcontext "golang.org/x/net/context"
)

// MakeGRPCServer makes a set of Endpoints available as a gRPC customercontainerServer
func MakeGRPCServer(ctx context.Context, endpoints Endpoints, logger log.Logger) pb.CustomerContainerServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcServer{
		createcontainer: grpctransport.NewServer(
			endpoints.CreateContainerEndpoint,
			DecodeGRPCCreateContainerRequest,
			EncodeGRPCCreateContainerResponse,
			options...,
		),
		editcontainer: grpctransport.NewServer(
			endpoints.EditContainerEndpoint,
			DecodeGRPCEditContainerRequest,
			EncodeGRPCEditContainerResponse,
			options...,
		),
		removecontainer: grpctransport.NewServer(
			endpoints.RemoveContainerEndpoint,
			DecodeGRPCRemoveContainerRequest,
			EncodeGRPCRemoveContainerResponse,
			options...,
		),
		instances: grpctransport.NewServer(
			endpoints.InstancesEndpoint,
			DecodeGRPCInstancesRequest,
			EncodeGRPCInstancesResponse,
			options...,
		),
	}
}

type grpcServer struct {
	createcontainer grpctransport.Handler
	editcontainer   grpctransport.Handler
	removecontainer grpctransport.Handler
	instances       grpctransport.Handler
}

func (s *grpcServer) CreateContainer(ctx oldcontext.Context, req *pb.CreateContainerRequest) (*pb.CreateContainerResponse, error) {
	_, res, err := s.createcontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.CreateContainerResponse), nil
}

func (s *grpcServer) EditContainer(ctx oldcontext.Context, req *pb.EditContainerRequest) (*pb.EditContainerResponse, error) {
	_, res, err := s.editcontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.EditContainerResponse), nil
}

func (s *grpcServer) RemoveContainer(ctx oldcontext.Context, req *pb.RemoveContainerRequest) (*pb.RemoveContainerResponse, error) {
	_, res, err := s.removecontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.RemoveContainerResponse), nil
}

func (s *grpcServer) Instances(ctx oldcontext.Context, req *pb.InstancesRequest) (*pb.InstancesResponse, error) {
	_, res, err := s.instances.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.InstancesResponse), nil
}

func convertContainerConfig(cfg *pb.ContainerConfig) *ContainerConfig {
	return &ContainerConfig{
		ImageName: cfg.ImageName,
	}
}

// DecodeGRPCCreateContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateContainer request to a messages/customercontainer.proto-domain createcontainer request.
func DecodeGRPCCreateContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateContainerRequest)
	return CreateContainerRequest{
		Refid: int(req.Refid),
		Cfg:   convertContainerConfig(req.Cfg),
	}, nil
}

// DecodeGRPCEditContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC EditContainer request to a messages/customercontainer.proto-domain editcontainer request.
func DecodeGRPCEditContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.EditContainerRequest)
	return EditContainerRequest{
		ID:  req.ID,
		Cfg: convertContainerConfig(req.Cfg),
	}, nil
}

// DecodeGRPCRemoveContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemoveContainer request to a messages/customercontainer.proto-domain removecontainer request.
func DecodeGRPCRemoveContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RemoveContainerRequest)
	return RemoveContainerRequest{
		ID: req.ID,
	}, nil
}

// DecodeGRPCInstancesRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC Instances request to a messages/customercontainer.proto-domain instances request.
func DecodeGRPCInstancesRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.InstancesRequest)
	return InstancesRequest{
		Refid: int(req.Refid),
	}, nil
}

// EncodeGRPCCreateContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain createcontainer response to a gRPC CreateContainer response.
func EncodeGRPCCreateContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(CreateContainerResponse)
	return &pb.CreateContainerResponse{
		Error: res.Error.Error(),
		Name:  res.Name,
		ID:    res.ID,
	}, nil
}

// EncodeGRPCEditContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain editcontainer response to a gRPC EditContainer response.
func EncodeGRPCEditContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(EditContainerResponse)
	return &pb.EditContainerResponse{
		Error: res.Error.Error(),
	}, nil
}

// EncodeGRPCRemoveContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain removecontainer response to a gRPC RemoveContainer response.
func EncodeGRPCRemoveContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveContainerResponse)
	return &pb.RemoveContainerResponse{
		Error: res.Error.Error(),
	}, nil
}

// EncodeGRPCInstancesResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain instances response to a gRPC Instances response.
func EncodeGRPCInstancesResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(InstancesResponse)
	return &pb.InstancesResponse{
		Instances: res.Instances,
	}, nil
}