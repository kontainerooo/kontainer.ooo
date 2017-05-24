package container

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
	oldcontext "golang.org/x/net/context"
)

// MakeGRPCServer makes a set of Endpoints available as a gRPC containerServer
func MakeGRPCServer(ctx context.Context, endpoints Endpoints, logger log.Logger) pb.ContainerServiceServer {
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
		stopcontainer: grpctransport.NewServer(
			endpoints.StopContainerEndpoint,
			DecodeGRPCStopContainerRequest,
			EncodeGRPCStopContainerResponse,
			options...,
		),
		isrunning: grpctransport.NewServer(
			endpoints.IsRunningEndpoint,
			DecodeGRPCIsRunningRequest,
			EncodeGRPCIsRunningResponse,
			options...,
		),
		execute: grpctransport.NewServer(
			endpoints.ExecuteEndpoint,
			DecodeGRPCExecuteRequest,
			EncodeGRPCExecuteResponse,
			options...,
		),
	}
}

type grpcServer struct {
	createcontainer grpctransport.Handler
	removecontainer grpctransport.Handler
	instances       grpctransport.Handler
	startcontainer  grpctransport.Handler
	stopcontainer   grpctransport.Handler
	isrunning       grpctransport.Handler
	execute         grpctransport.Handler
}

func (s *grpcServer) CreateContainer(ctx oldcontext.Context, req *pb.CreateContainerRequest) (*pb.CreateContainerResponse, error) {
	_, res, err := s.createcontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.CreateContainerResponse), nil
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

func (s *grpcServer) StartContainer(ctx oldcontext.Context, req *pb.StartContainerRequest) (*pb.StartContainerResponse, error) {
	_, res, err := s.startcontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StartContainerResponse), nil
}

func (s *grpcServer) StopContainer(ctx oldcontext.Context, req *pb.StopContainerRequest) (*pb.StopContainerResponse, error) {
	_, res, err := s.stopcontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StopContainerResponse), nil
}

func (s *grpcServer) IsRunning(ctx oldcontext.Context, req *pb.IsRunningRequest) (*pb.IsRunningResponse, error) {
	_, res, err := s.isrunning.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.IsRunningResponse), nil
}

func (s *grpcServer) Execute(ctx oldcontext.Context, req *pb.ExecuteRequest) (*pb.ExecuteResponse, error) {
	_, res, err := s.execute.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ExecuteResponse), nil
}

// DecodeGRPCCreateContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateContainer request to a messages/container.proto-domain createcontainer request.
func DecodeGRPCCreateContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateContainerRequest)
	return CreateContainerRequest{
		RefID: uint(req.RefID),
		KmiID: uint(req.KmiID),
		Name:  req.Name,
	}, nil
}

// DecodeGRPCRemoveContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemoveContainer request to a messages/container.proto-domain removecontainer request.
func DecodeGRPCRemoveContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RemoveContainerRequest)
	return RemoveContainerRequest{
		RefID: uint(req.RefID),
		ID:    req.ID,
	}, nil
}

// DecodeGRPCInstancesRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC Instances request to a messages/container.proto-domain instances request.
func DecodeGRPCInstancesRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.InstancesRequest)
	return InstancesRequest{
		RefID: uint(req.RefID),
	}, nil
}

// DecodeGRPCStopContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC StopContainer request to a messages/container.proto-domain stopcontainer request.
func DecodeGRPCStopContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.StopContainerRequest)
	return StopContainerRequest{
		RefID: uint(req.RefID),
		ID:    req.ID,
	}, nil
}

// DecodeGRPCIsRunningRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC IsRunning request to a messages/container.proto-domain isrunning request.
func DecodeGRPCIsRunningRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.IsRunningRequest)
	return IsRunningRequest{
		RefID: uint(req.RefID),
	}, nil
}

// DecodeGRPCExecuteRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC Execute request to a messages/container.proto-domain execute request.
func DecodeGRPCExecuteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ExecuteRequest)
	return ExecuteRequest{
		RefID: uint(req.RefID),
		ID:    req.ID,
		CMD:   req.Cmd,
	}, nil
}

// EncodeGRPCCreateContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain createcontainer response to a gRPC CreateContainer response.
func EncodeGRPCCreateContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(CreateContainerResponse)
	gRPCRes := &pb.CreateContainerResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemoveContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain removecontainer response to a gRPC RemoveContainer response.
func EncodeGRPCRemoveContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveContainerResponse)
	gRPCRes := &pb.RemoveContainerResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCInstancesResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain instances response to a gRPC Instances response.
func EncodeGRPCInstancesResponse(_ context.Context, response interface{}) (interface{}, error) {
	gRPCRes := &pb.InstancesResponse{}
	return gRPCRes, nil
}

// EncodeGRPCStopContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain stopcontainer response to a gRPC StopContainer response.
func EncodeGRPCStopContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(StopContainerResponse)
	gRPCRes := &pb.StopContainerResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCIsRunningResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain isrunning response to a gRPC IsRunning response.
func EncodeGRPCIsRunningResponse(_ context.Context, response interface{}) (interface{}, error) {
	gRPCRes := &pb.IsRunningResponse{}
	return gRPCRes, nil
}

// EncodeGRPCExecuteResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain execute response to a gRPC Execute response.
func EncodeGRPCExecuteResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(ExecuteResponse)
	gRPCRes := &pb.ExecuteResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}
