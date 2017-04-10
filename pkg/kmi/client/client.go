package client

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
)

// New creates a set of endpoints based on a gRPC connection
func New(conn *grpc.ClientConn, logger log.Logger) *kmi.Endpoints {

	var AddKMIEndpoint endpoint.Endpoint
	{
		AddKMIEndpoint = grpctransport.NewClient(
			conn,
			"KMIService",
			"AddKMI",
			EncodeGRPCAddKMIRequest,
			DecodeGRPCAddKMIResponse,
			pb.AddKMIResponse{},
		).Endpoint()
	}

	var RemoveKMIEndpoint endpoint.Endpoint
	{
		RemoveKMIEndpoint = grpctransport.NewClient(
			conn,
			"KMIService",
			"RemoveKMI",
			EncodeGRPCRemoveKMIRequest,
			DecodeGRPCRemoveKMIResponse,
			pb.RemoveKMIResponse{},
		).Endpoint()
	}

	var GetKMIEndpoint endpoint.Endpoint
	{
		GetKMIEndpoint = grpctransport.NewClient(
			conn,
			"KMIService",
			"GetKMI",
			EncodeGRPCGetKMIRequest,
			DecodeGRPCGetKMIResponse,
			pb.GetKMIResponse{},
		).Endpoint()
	}

	var KMIEndpoint endpoint.Endpoint
	{
		KMIEndpoint = grpctransport.NewClient(
			conn,
			"KMIService",
			"KMI",
			EncodeGRPCKMIRequest,
			DecodeGRPCKMIResponse,
			pb.KMIResponse{},
		).Endpoint()
	}

	return &kmi.Endpoints{
		AddKMIEndpoint:    AddKMIEndpoint,
		RemoveKMIEndpoint: RemoveKMIEndpoint,
		GetKMIEndpoint:    GetKMIEndpoint,
		KMIEndpoint:       KMIEndpoint,
	}
}

func getError(e string) error {
	if e != "" {
		return errors.New(e)
	}
	return nil
}

func convertFrontendModule(f *pb.FrontendModule) *kmi.FrontendModule {
	return &kmi.FrontendModule{
		Template:   f.Template,
		Parameters: abstraction.NewJSONFromMap(f.Parameters),
	}
}

func convertFrontendModuleArray(f []*pb.FrontendModule) kmi.FrontendArray {
	a := make(kmi.FrontendArray, len(f))
	for i, m := range f {
		a[i] = convertFrontendModule(m)
	}
	return a
}

func convertKMDI(k *pb.KMDI) kmi.KMDI {
	return kmi.KMDI{
		ID:          uint(k.ID),
		Name:        k.Name,
		Version:     k.Version,
		Description: k.Description,
	}
}

func convertKMI(k *pb.KMI) *kmi.KMI {
	return &kmi.KMI{
		KMDI:        convertKMDI(k.KMDI),
		Dockerfile:  k.Dockerfile,
		Context:     k.Context,
		Commands:    abstraction.NewJSONFromMap(k.Commands),
		Environment: abstraction.NewJSONFromMap(k.Environment),
		Frontend:    convertFrontendModuleArray(k.Frontend),
		Imports:     pq.StringArray(k.Imports),
		Interfaces:  abstraction.NewJSONFromMap(k.Interfaces),
		Mounts:      pq.StringArray(k.Mounts),
		Variables:   pq.StringArray(k.Variables),
		Resources:   abstraction.NewJSONFromMap(k.Resources),
	}
}

func convertKMDIArray(k []*pb.KMDI) *[]kmi.KMDI {
	a := make([]kmi.KMDI, len(k))
	for i, d := range k {
		a[i] = convertKMDI(d)
	}
	return &a
}

// EncodeGRPCAddKMIRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/kmi.proto-domain addkmi request to a gRPC AddKMI request.
func EncodeGRPCAddKMIRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*kmi.AddKMIRequest)
	return &pb.AddKMIRequest{
		Path: req.Path,
	}, nil
}

// DecodeGRPCAddKMIResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC AddKMI response to a messages/kmi.proto-domain addkmi response.
func DecodeGRPCAddKMIResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.AddKMIResponse)
	return &kmi.AddKMIResponse{
		ID:    uint(response.ID),
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveKMIRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/kmi.proto-domain removekmi request to a gRPC RemoveKMI request.
func EncodeGRPCRemoveKMIRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*kmi.RemoveKMIRequest)
	return &pb.RemoveKMIRequest{
		ID: uint32(req.ID),
	}, nil
}

// DecodeGRPCRemoveKMIResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveKMI response to a messages/kmi.proto-domain removekmi response.
func DecodeGRPCRemoveKMIResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.RemoveKMIResponse)
	return &kmi.RemoveKMIResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCGetKMIRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/kmi.proto-domain getkmi request to a gRPC GetKMI request.
func EncodeGRPCGetKMIRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*kmi.GetKMIRequest)
	return &pb.GetKMIRequest{
		ID: uint32(req.ID),
	}, nil
}

// DecodeGRPCGetKMIResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetKMI response to a messages/kmi.proto-domain getkmi response.
func DecodeGRPCGetKMIResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.GetKMIResponse)
	return &kmi.GetKMIResponse{
		Error: getError(response.Error),
		KMI:   convertKMI(response.Kmi),
	}, nil
}

// EncodeGRPCKMIRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/kmi.proto-domain kmi request to a gRPC KMI request.
func EncodeGRPCKMIRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return &pb.KMIRequest{}, nil
}

// DecodeGRPCKMIResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC KMI response to a messages/kmi.proto-domain kmi response.
func DecodeGRPCKMIResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.KMIResponse)
	return &kmi.KMIResponse{
		KMDI:  convertKMDIArray(response.Kmdi),
		Error: getError(response.Error),
	}, nil
}
