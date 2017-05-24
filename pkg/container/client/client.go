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
	"github.com/kontainerooo/kontainer.ooo/pkg/container"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
)

// New creates a set of endpoints based on a gRPC connection
func New(conn *grpc.ClientConn, logger log.Logger) *container.Endpoints {

	var CreateContainerEndpoint endpoint.Endpoint
	{
		CreateContainerEndpoint = grpctransport.NewClient(
			conn,
			"containerService",
			"CreateContainer",
			EncodeGRPCCreateContainerRequest,
			DecodeGRPCCreateContainerResponse,
			pb.CreateContainerResponse{},
		).Endpoint()
	}

	var RemoveContainerEndpoint endpoint.Endpoint
	{
		RemoveContainerEndpoint = grpctransport.NewClient(
			conn,
			"containerService",
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
			"containerService",
			"Instances",
			EncodeGRPCInstancesRequest,
			DecodeGRPCInstancesResponse,
			pb.InstancesResponse{},
		).Endpoint()
	}

	var StopContainerEndpoint endpoint.Endpoint
	{
		StopContainerEndpoint = grpctransport.NewClient(
			conn,
			"containerService",
			"StopContainer",
			EncodeGRPCStopContainerRequest,
			DecodeGRPCStopContainerResponse,
			pb.StopContainerResponse{},
		).Endpoint()
	}

	var IsRunningEndpoint endpoint.Endpoint
	{
		IsRunningEndpoint = grpctransport.NewClient(
			conn,
			"containerService",
			"IsRunning",
			EncodeGRPCIsRunningRequest,
			DecodeGRPCIsRunningResponse,
			pb.IsRunningResponse{},
		).Endpoint()
	}

	var ExecuteEndpoint endpoint.Endpoint
	{
		ExecuteEndpoint = grpctransport.NewClient(
			conn,
			"containerService",
			"Execute",
			EncodeGRPCExecuteRequest,
			DecodeGRPCExecuteResponse,
			pb.ExecuteResponse{},
		).Endpoint()
	}

	return &container.Endpoints{
		CreateContainerEndpoint: CreateContainerEndpoint,
		RemoveContainerEndpoint: RemoveContainerEndpoint,
		InstancesEndpoint:       InstancesEndpoint,
		StopContainerEndpoint:   StopContainerEndpoint,
		IsRunningEndpoint:       IsRunningEndpoint,
		ExecuteEndpoint:         ExecuteEndpoint,
	}
}

func getError(e string) error {
	if e != "" {
		return errors.New(e)
	}
	return nil
}

func pbContainersToContainers(pbc []*pb.Container) []container.Container {
	cts := []container.Container{}
	for _, c := range pbc {
		cts = append(cts, container.Container{
			RefID:         uint(c.RefID),
			ContainerID:   c.ContainerID,
			ContainerName: c.ContainerName,
			Running:       c.Running,
			KMI:           *convertCKMI(c.Kmi),
		})
	}

	return cts
}

func convertPBFrontendModule(f *kmi.FrontendModule) *pb.FrontendModule {
	return &pb.FrontendModule{
		Template:   f.Template,
		Parameters: f.Parameters.ToStringMap(),
	}
}

func convertPBFrontendModuleArray(f kmi.FrontendArray) []*pb.FrontendModule {
	a := make([]*pb.FrontendModule, len(f))
	for i, m := range f {
		a[i] = convertPBFrontendModule(m)
	}
	return a
}

func convertPBKMDI(k kmi.KMDI) *pb.KMDI {
	return &pb.KMDI{
		ID:          uint32(k.ID),
		Name:        k.Name,
		Version:     k.Version,
		Description: k.Description,
	}
}

func convertPBKMI(k *kmi.KMI) *pb.KMI {
	return &pb.KMI{
		KMDI:            convertPBKMDI(k.KMDI),
		ProvisionScript: k.ProvisionScript,
		Commands:        k.Commands.ToStringMap(),
		Environment:     k.Environment.ToStringMap(),
		Frontend:        convertPBFrontendModuleArray(k.Frontend),
		Imports:         k.Imports,
		Interfaces:      k.Interfaces.ToStringMap(),
		Variables:       k.Variables,
		Resources:       k.Resources.ToStringMap(),
	}
}

func convertPBKMDIArray(k *[]kmi.KMDI) []*pb.KMDI {
	a := make([]*pb.KMDI, len(*k))
	for i, d := range *k {
		a[i] = convertPBKMDI(d)
	}
	return a
}

func convertCFrontendModule(f *pb.CFrontendModule) *kmi.FrontendModule {
	return &kmi.FrontendModule{
		Template:   f.Template,
		Parameters: abstraction.NewJSONFromMap(f.Parameters),
	}
}

func convertCFrontendModuleArray(f []*pb.CFrontendModule) kmi.FrontendArray {
	a := make(kmi.FrontendArray, len(f))
	for i, m := range f {
		a[i] = convertCFrontendModule(m)
	}
	return a
}

func convertCKMDI(k *pb.CKMDI) kmi.KMDI {
	return kmi.KMDI{
		ID:          uint(k.ID),
		Name:        k.Name,
		Version:     k.Version,
		Description: k.Description,
	}
}

func convertCKMI(k *pb.CKMI) *kmi.KMI {
	return &kmi.KMI{
		KMDI:            convertCKMDI(k.KMDI),
		ProvisionScript: k.ProvisionScript,
		Commands:        abstraction.NewJSONFromMap(k.Commands),
		Environment:     abstraction.NewJSONFromMap(k.Environment),
		Frontend:        convertCFrontendModuleArray(k.Frontend),
		Imports:         pq.StringArray(k.Imports),
		Interfaces:      abstraction.NewJSONFromMap(k.Interfaces),
		Variables:       pq.StringArray(k.Variables),
		Resources:       abstraction.NewJSONFromMap(k.Resources),
	}
}

func convertKMDIArray(k []*pb.CKMDI) *[]kmi.KMDI {
	a := make([]kmi.KMDI, len(k))
	for i, d := range k {
		a[i] = convertCKMDI(d)
	}
	return &a
}

// EncodeGRPCCreateContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain createcontainer request to a gRPC CreateContainer request.
func EncodeGRPCCreateContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.CreateContainerRequest)
	return &pb.CreateContainerRequest{
		RefID: uint32(req.RefID),
		KmiID: uint32(req.KmiID),
		Name:  req.Name,
	}, nil
}

// DecodeGRPCCreateContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC CreateContainer response to a messages/container.proto-domain createcontainer response.
func DecodeGRPCCreateContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.CreateContainerResponse)
	return &container.CreateContainerResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain removecontainer request to a gRPC RemoveContainer request.
func EncodeGRPCRemoveContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.RemoveContainerRequest)
	return &pb.RemoveContainerRequest{
		RefID: uint32(req.RefID),
		ID:    req.ID,
	}, nil
}

// DecodeGRPCRemoveContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveContainer response to a messages/container.proto-domain removecontainer response.
func DecodeGRPCRemoveContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.RemoveContainerResponse)
	return &container.RemoveContainerResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCInstancesRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain instances request to a gRPC Instances request.
func EncodeGRPCInstancesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.InstancesRequest)
	return &pb.InstancesRequest{
		RefID: uint32(req.RefID),
	}, nil
}

// DecodeGRPCInstancesResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC Instances response to a messages/container.proto-domain instances response.
func DecodeGRPCInstancesResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.InstancesResponse)
	return &container.InstancesResponse{
		Containers: pbContainersToContainers(response.Instances),
	}, nil
}

// EncodeGRPCStopContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain stopcontainer request to a gRPC StopContainer request.
func EncodeGRPCStopContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.StopContainerRequest)
	return &pb.StopContainerRequest{
		RefID: uint32(req.RefID),
		ID:    req.ID,
	}, nil
}

// DecodeGRPCStopContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC StopContainer response to a messages/container.proto-domain stopcontainer response.
func DecodeGRPCStopContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.StopContainerResponse)
	return &container.StopContainerResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCIsRunningRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain isrunning request to a gRPC IsRunning request.
func EncodeGRPCIsRunningRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.IsRunningRequest)
	return &pb.IsRunningRequest{
		RefID: uint32(req.RefID),
		ID:    req.ID,
	}, nil
}

// DecodeGRPCIsRunningResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC IsRunning response to a messages/container.proto-domain isrunning response.
func DecodeGRPCIsRunningResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.IsRunningResponse)
	return &container.IsRunningResponse{
		IsRunning: response.IsRunning,
	}, nil
}

// EncodeGRPCExecuteRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain execute request to a gRPC Execute request.
func EncodeGRPCExecuteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.ExecuteRequest)
	return &pb.ExecuteRequest{
		RefID: uint32(req.RefID),
		ID:    req.ID,
		Cmd:   req.CMD,
	}, nil
}

// DecodeGRPCExecuteResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC Execute response to a messages/container.proto-domain execute response.
func DecodeGRPCExecuteResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.ExecuteResponse)
	return &container.ExecuteResponse{
		Error: getError(response.Error),
	}, nil
}
