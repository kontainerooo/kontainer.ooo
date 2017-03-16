package containerlifecycle

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/ttdennis/kontainer.io/pkg/pb"
	oldcontext "golang.org/x/net/context"
)

// MakeGRPCServer makes a set of Endpoints available as a gRPC containerlifecycleServer
func MakeGRPCServer(ctx context.Context, endpoints Endpoints, logger log.Logger) pb.ContainerLifecycleServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		startcontainer: grpctransport.NewServer(
			endpoints.StartContainerEndpoint,
			DecodeGRPCStartContainerRequest,
			EncodeGRPCStartContainerResponse,
			options...,
		),
		startcommand: grpctransport.NewServer(
			endpoints.StartCommandEndpoint,
			DecodeGRPCStartCommandRequest,
			EncodeGRPCStartCommandResponse,
			options...,
		),
		stopcontainer: grpctransport.NewServer(
			endpoints.StopContainerEndpoint,
			DecodeGRPCStopContainerRequest,
			EncodeGRPCStopContainerResponse,
			options...,
		),
	}
}

type grpcServer struct {
	startcontainer grpctransport.Handler
	startcommand   grpctransport.Handler
	stopcontainer  grpctransport.Handler
}

func (s *grpcServer) StartContainer(ctx oldcontext.Context, req *pb.StartContainerRequest) (*pb.StartContainerResponse, error) {
	_, res, err := s.startcontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StartContainerResponse), nil
}

func (s *grpcServer) StartCommand(ctx oldcontext.Context, req *pb.StartCommandRequest) (*pb.StartCommandResponse, error) {
	_, res, err := s.startcommand.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StartCommandResponse), nil
}

func (s *grpcServer) StopContainer(ctx oldcontext.Context, req *pb.StopContainerRequest) (*pb.StopContainerResponse, error) {
	_, res, err := s.stopcontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StopContainerResponse), nil
}

// DecodeGRPCStartContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC StartContainer request to a messages/containerlifecycle.proto-domain startcontainer request.
func DecodeGRPCStartContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.StartContainerRequest)
	return StartContainerRequest{
		ID: req.ID,
	}, nil
}

// DecodeGRPCStartCommandRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC StartCommand request to a messages/containerlifecycle.proto-domain startcommand request.
func DecodeGRPCStartCommandRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.StartCommandRequest)
	return StartCommandRequest{
		ID:  req.ID,
		Cmd: req.Cmd,
	}, nil
}

// DecodeGRPCStopContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC StopContainer request to a messages/containerlifecycle.proto-domain stopcontainer request.
func DecodeGRPCStopContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.StopContainerRequest)
	return StopContainerRequest{
		ID: req.ID,
	}, nil
}

// EncodeGRPCStartContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/containerlifecycle.proto-domain startcontainer response to a gRPC StartContainer response.
func EncodeGRPCStartContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(StartContainerResponse)
	gRPCRes := &pb.StartContainerResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCStartCommandResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/containerlifecycle.proto-domain startcommand response to a gRPC StartCommand response.
func EncodeGRPCStartCommandResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(StartCommandResponse)
	gRPCRes := &pb.StartCommandResponse{
		ID: res.ID,
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCStopContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/containerlifecycle.proto-domain stopcontainer response to a gRPC StopContainer response.
func EncodeGRPCStopContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(StopContainerResponse)
	gRPCRes := &pb.StopContainerResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}
