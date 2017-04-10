package kmi

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
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

func convertPBFrontendModule(f *FrontendModule) *pb.FrontendModule {
	return &pb.FrontendModule{
		Template:   f.Template,
		Parameters: f.Parameters.ToStringMap(),
	}
}

func convertPBFrontendModuleArray(f FrontendArray) []*pb.FrontendModule {
	a := make([]*pb.FrontendModule, len(f))
	for i, m := range f {
		a[i] = convertPBFrontendModule(m)
	}
	return a
}

func convertPBKMDI(k KMDI) *pb.KMDI {
	return &pb.KMDI{
		ID:          uint32(k.ID),
		Name:        k.Name,
		Version:     k.Version,
		Description: k.Description,
	}
}

func convertPBKMI(k *KMI) *pb.KMI {
	return &pb.KMI{
		KMDI:        convertPBKMDI(k.KMDI),
		Dockerfile:  k.Dockerfile,
		Context:     k.Context,
		Commands:    k.Commands.ToStringMap(),
		Environment: k.Environment.ToStringMap(),
		Frontend:    convertPBFrontendModuleArray(k.Frontend),
		Imports:     k.Imports,
		Interfaces:  k.Interfaces.ToStringMap(),
		Mounts:      k.Mounts,
		Variables:   k.Variables,
		Resources:   k.Resources.ToStringMap(),
	}
}

func convertPBKMDIArray(k *[]KMDI) []*pb.KMDI {
	a := make([]*pb.KMDI, len(*k))
	for i, d := range *k {
		a[i] = convertPBKMDI(d)
	}
	return a
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
	gRPCRes := &pb.AddKMIResponse{
		ID: uint32(res.ID),
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemoveKMIResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/KMI.proto-domain removeKMI response to a gRPC RemoveKMI response.
func EncodeGRPCRemoveKMIResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveKMIResponse)
	gRPCRes := &pb.RemoveKMIResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCGetKMIResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/KMI.proto-domain getKMI response to a gRPC GetKMI response.
func EncodeGRPCGetKMIResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetKMIResponse)
	gRPCRes := &pb.GetKMIResponse{}
	if res.KMI != nil {
		gRPCRes.Kmi = convertPBKMI(res.KMI)
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCKMIResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/KMI.proto-domain KMI response to a gRPC KMI response.
func EncodeGRPCKMIResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(KMIResponse)
	gRPCRes := &pb.KMIResponse{}
	if res.KMDI != nil {
		gRPCRes.Kmdi = convertPBKMDIArray(res.KMDI)
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}
