package customercontainer

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
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
		getcontainerid: grpctransport.NewServer(
			endpoints.GetContainerIDEndpoint,
			DecodeGRPCGetContainerIDRequest,
			EncodeGRPCGetContainerIDResponse,
			options...,
		),
	}
}

type grpcServer struct {
	createcontainer grpctransport.Handler
	editcontainer   grpctransport.Handler
	removecontainer grpctransport.Handler
	instances       grpctransport.Handler
	getcontainerid  grpctransport.Handler
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

func (s *grpcServer) GetContainerID(ctx oldcontext.Context, req *pb.GetContainerIDRequest) (*pb.GetContainerIDResponse, error) {
	_, res, err := s.getcontainerid.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetContainerIDResponse), nil
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
		RefID: uint(req.RefID),
		KMIID: uint(req.KmiID),
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
		RefID: uint(req.RefID),
	}, nil
}

// DecodeGRPCGetContainerIDRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetContainerID request to a messages/customercontainer.proto-domain getcontainerid request.
func DecodeGRPCGetContainerIDRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetContainerIDRequest)
	return GetContainerIDRequest{
		RefID:         uint(req.RefID),
		ContainerName: req.ContainerName,
	}, nil
}

// EncodeGRPCCreateContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain createcontainer response to a gRPC CreateContainer response.
func EncodeGRPCCreateContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(CreateContainerResponse)
	gRPCRes := &pb.CreateContainerResponse{
		ID:   res.ID,
		Name: res.Name,
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCEditContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain editcontainer response to a gRPC EditContainer response.
func EncodeGRPCEditContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(EditContainerResponse)
	gRPCRes := &pb.EditContainerResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemoveContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain removecontainer response to a gRPC RemoveContainer response.
func EncodeGRPCRemoveContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveContainerResponse)
	gRPCRes := &pb.RemoveContainerResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCInstancesResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain instances response to a gRPC Instances response.
func EncodeGRPCInstancesResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(InstancesResponse)
	gRPCRes := &pb.InstancesResponse{
		Instances: res.Instances,
	}
	return gRPCRes, nil
}

// EncodeGRPCGetContainerIDResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/customercontainer.proto-domain getcontainerid response to a gRPC GetContainerID response.
func EncodeGRPCGetContainerIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetContainerIDResponse)
	gRPCRes := &pb.GetContainerIDResponse{
		ContainerID: res.ContainerID,
		Error:       res.Error.Error(),
	}
	return gRPCRes, nil
}
