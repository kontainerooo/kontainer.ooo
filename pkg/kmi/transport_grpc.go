package kmi

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/ttdennis/kontainer.io/pkg/pb"
	oldcontext "golang.org/x/net/context"
)

// MakeGRPCServer makes a set of Endpoints available as a gRPC KMIServer
func MakeGRPCServer(ctx context.Context, endpoints Endpoints, logger log.Logger) pb.KMIServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		addKMI: grpctransport.NewServer(
			endpoints.AddKMIEndpoint,
			DecodeGRPCAddKMIRequest,
			EncodeGRPCAddKMIResponse,
			options...,
		),
		removeKMI: grpctransport.NewServer(
			endpoints.RemoveKMIEndpoint,
			DecodeGRPCRemoveKMIRequest,
			EncodeGRPCRemoveKMIResponse,
			options...,
		),
		getKMI: grpctransport.NewServer(
			endpoints.GetKMIEndpoint,
			DecodeGRPCGetKMIRequest,
			EncodeGRPCGetKMIResponse,
			options...,
		),
		kmi: grpctransport.NewServer(
			endpoints.KMIEndpoint,
			DecodeGRPCKMIRequest,
			EncodeGRPCKMIResponse,
			options...,
		),
	}
}

type grpcServer struct {
	addKMI    grpctransport.Handler
	removeKMI grpctransport.Handler
	getKMI    grpctransport.Handler
	kmi       grpctransport.Handler
}

func (s *grpcServer) AddKMI(ctx oldcontext.Context, req *pb.AddKMIRequest) (*pb.AddKMIResponse, error) {
	_, res, err := s.addKMI.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.AddKMIResponse), nil
}

func (s *grpcServer) RemoveKMI(ctx oldcontext.Context, req *pb.RemoveKMIRequest) (*pb.RemoveKMIResponse, error) {
	_, res, err := s.removeKMI.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.RemoveKMIResponse), nil
}

func (s *grpcServer) GetKMI(ctx oldcontext.Context, req *pb.GetKMIRequest) (*pb.GetKMIResponse, error) {
	_, res, err := s.getKMI.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetKMIResponse), nil
}

func (s *grpcServer) KMI(ctx oldcontext.Context, req *pb.KMIRequest) (*pb.KMIResponse, error) {
	_, res, err := s.kmi.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.KMIResponse), nil
}

func transformIntoPBKMI(kmi interface{}) *pb.KMI {
	return kmi.(*pb.KMI)
}

func transformIntoPBKMDI(kmdi interface{}) []*pb.KMDI {
	return kmdi.([]*pb.KMDI)
}

// DecodeGRPCAddKMIRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC AddKMI request to a messages/KMI.proto-domain addKMI request.
func DecodeGRPCAddKMIRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddKMIRequest)
	return AddKMIRequest{
		Path: req.Path,
	}, nil
}

// DecodeGRPCRemoveKMIRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemoveKMI request to a messages/KMI.proto-domain removeKMI request.
func DecodeGRPCRemoveKMIRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RemoveKMIRequest)
	return RemoveKMIRequest{
		ID: uint(req.ID),
	}, nil
}

// DecodeGRPCGetKMIRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetKMI request to a messages/KMI.proto-domain getKMI request.
func DecodeGRPCGetKMIRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetKMIRequest)
	return GetKMIRequest{
		ID: uint(req.ID),
	}, nil
}

// DecodeGRPCKMIRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC KMI request to a messages/KMI.proto-domain KMI request.
func DecodeGRPCKMIRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return KMIRequest{}, nil
}

// EncodeGRPCAddKMIResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/KMI.proto-domain addKMI response to a gRPC AddKMI response.
func EncodeGRPCAddKMIResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(AddKMIResponse)
	return &pb.AddKMIResponse{
		ID:    uint32(res.ID),
		Error: res.Error.Error(),
	}, nil
}

// EncodeGRPCRemoveKMIResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/KMI.proto-domain removeKMI response to a gRPC RemoveKMI response.
func EncodeGRPCRemoveKMIResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveKMIResponse)
	return &pb.RemoveKMIResponse{
		Error: res.Error.Error(),
	}, nil
}

// EncodeGRPCGetKMIResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/KMI.proto-domain getKMI response to a gRPC GetKMI response.
func EncodeGRPCGetKMIResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetKMIResponse)
	return &pb.GetKMIResponse{
		Kmi:   transformIntoPBKMI(res.KMI),
		Error: res.Error.Error(),
	}, nil
}

// EncodeGRPCKMIResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/KMI.proto-domain KMI response to a gRPC KMI response.
func EncodeGRPCKMIResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(KMIResponse)
	return &pb.KMIResponse{
		Kmdi:  transformIntoPBKMDI(res.KMDI),
		Error: res.Error.Error(),
	}, nil
}
