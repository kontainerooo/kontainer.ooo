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

	containerPB "github.com/kontainerooo/kontainer.ooo/pkg/container/pb"
	kmiPB "github.com/kontainerooo/kontainer.ooo/pkg/kmi/pb"
)

// New creates a set of endpoints based on a gRPC connection
func New(conn *grpc.ClientConn, logger log.Logger) *container.Endpoints {

	var CreateContainerEndpoint endpoint.Endpoint
	{
		CreateContainerEndpoint = grpctransport.NewClient(
			conn,
			"container.ContainerService",
			"CreateContainer",
			EncodeGRPCCreateContainerRequest,
			DecodeGRPCCreateContainerResponse,
			containerPB.CreateContainerResponse{},
		).Endpoint()
	}

	var RemoveContainerEndpoint endpoint.Endpoint
	{
		RemoveContainerEndpoint = grpctransport.NewClient(
			conn,
			"container.ContainerService",
			"RemoveContainer",
			EncodeGRPCRemoveContainerRequest,
			DecodeGRPCRemoveContainerResponse,
			containerPB.RemoveContainerResponse{},
		).Endpoint()
	}

	var InstancesEndpoint endpoint.Endpoint
	{
		InstancesEndpoint = grpctransport.NewClient(
			conn,
			"container.ContainerService",
			"Instances",
			EncodeGRPCInstancesRequest,
			DecodeGRPCInstancesResponse,
			containerPB.InstancesResponse{},
		).Endpoint()
	}

	var StopContainerEndpoint endpoint.Endpoint
	{
		StopContainerEndpoint = grpctransport.NewClient(
			conn,
			"container.ContainerService",
			"StopContainer",
			EncodeGRPCStopContainerRequest,
			DecodeGRPCStopContainerResponse,
			containerPB.StopContainerResponse{},
		).Endpoint()
	}

	var ExecuteEndpoint endpoint.Endpoint
	{
		ExecuteEndpoint = grpctransport.NewClient(
			conn,
			"container.ContainerService",
			"Execute",
			EncodeGRPCExecuteRequest,
			DecodeGRPCExecuteResponse,
			containerPB.ExecuteResponse{},
		).Endpoint()
	}

	var GetEnvEndpoint endpoint.Endpoint
	{
		GetEnvEndpoint = grpctransport.NewClient(
			conn,
			"container.ContainerService",
			"GetEnv",
			EncodeGRPCGetEnvRequest,
			DecodeGRPCGetEnvResponse,
			containerPB.GetEnvResponse{},
		).Endpoint()
	}

	var SetEnvEndpoint endpoint.Endpoint
	{
		SetEnvEndpoint = grpctransport.NewClient(
			conn,
			"container.ContainerService",
			"SetEnv",
			EncodeGRPCSetEnvRequest,
			DecodeGRPCSetEnvResponse,
			containerPB.SetEnvResponse{},
		).Endpoint()
	}

	var IDForNameEndpoint endpoint.Endpoint
	{
		IDForNameEndpoint = grpctransport.NewClient(
			conn,
			"container.ContainerService",
			"IDForName",
			EncodeGRPCIDForNameRequest,
			DecodeGRPCIDForNameResponse,
			containerPB.IDForNameResponse{},
		).Endpoint()
	}

	var GetContainerKMIEndpoint endpoint.Endpoint
	{
		GetContainerKMIEndpoint = grpctransport.NewClient(
			conn,
			"container.ContainerService",
			"GetContainerKMI",
			EncodeGRPCGetContainerKMIRequest,
			DecodeGRPCGetContainerKMIResponse,
			containerPB.GetContainerKMIResponse{},
		).Endpoint()
	}

	return &container.Endpoints{
		CreateContainerEndpoint: CreateContainerEndpoint,
		RemoveContainerEndpoint: RemoveContainerEndpoint,
		InstancesEndpoint:       InstancesEndpoint,
		StopContainerEndpoint:   StopContainerEndpoint,
		ExecuteEndpoint:         ExecuteEndpoint,
		GetEnvEndpoint:          GetEnvEndpoint,
		SetEnvEndpoint:          SetEnvEndpoint,
		IDForNameEndpoint:       IDForNameEndpoint,
		GetContainerKMIEndpoint: GetContainerKMIEndpoint,
	}
}

func getError(e string) error {
	if e != "" {
		return errors.New(e)
	}
	return nil
}

func pbContainersToContainers(pbc []*containerPB.Container) []container.Container {
	cts := []container.Container{}
	for _, c := range pbc {
		cts = append(cts, container.Container{
			RefID:         uint(c.RefID),
			ContainerID:   c.ContainerID,
			ContainerName: c.ContainerName,
			KMI:           container.CKMI(*convertKMI(c.Kmi)),
		})
	}

	return cts
}

func convertPBFrontendModule(f *kmi.FrontendModule) *kmiPB.FrontendModule {
	return &kmiPB.FrontendModule{
		Template:   f.Template,
		Parameters: f.Parameters.ToStringMap(),
	}
}

func convertPBFrontendModuleArray(f kmi.FrontendArray) []*kmiPB.FrontendModule {
	a := make([]*kmiPB.FrontendModule, len(f))
	for i, m := range f {
		a[i] = convertPBFrontendModule(m)
	}
	return a
}

func convertPBKMDI(k kmi.KMDI) *kmiPB.KMDI {
	return &kmiPB.KMDI{
		ID:          uint32(k.ID),
		Name:        k.Name,
		Version:     k.Version,
		Description: k.Description,
	}
}

func convertPBKMI(k *kmi.KMI) *kmiPB.KMI {
	return &kmiPB.KMI{
		KMDI:            convertPBKMDI(k.KMDI),
		ProvisionScript: k.ProvisionScript,
		Commands:        k.Commands.ToStringMap(),
		Environment:     k.Environment.ToStringMap(),
		Frontend:        convertPBFrontendModuleArray(k.Frontend),
		Imports:         k.Imports,
		Interfaces:      k.Interfaces.ToStringMap(),
		Resources:       k.Resources.ToStringMap(),
	}
}

func convertPBKMDIArray(k *[]kmi.KMDI) []*kmiPB.KMDI {
	a := make([]*kmiPB.KMDI, len(*k))
	for i, d := range *k {
		a[i] = convertPBKMDI(d)
	}
	return a
}

func convertFrontendModule(f *kmiPB.FrontendModule) *kmi.FrontendModule {
	return &kmi.FrontendModule{
		Template:   f.Template,
		Parameters: abstraction.NewJSONFromMap(f.Parameters),
	}
}

func convertFrontendModuleArray(f []*kmiPB.FrontendModule) kmi.FrontendArray {
	a := make(kmi.FrontendArray, len(f))
	for i, m := range f {
		a[i] = convertFrontendModule(m)
	}
	return a
}

func convertKMDI(k *kmiPB.KMDI) kmi.KMDI {
	return kmi.KMDI{
		ID:          uint(k.ID),
		Name:        k.Name,
		Version:     k.Version,
		Description: k.Description,
	}
}

func convertKMI(k *kmiPB.KMI) *kmi.KMI {
	return &kmi.KMI{
		KMDI:            convertKMDI(k.KMDI),
		ProvisionScript: k.ProvisionScript,
		Commands:        abstraction.NewJSONFromMap(k.Commands),
		Environment:     abstraction.NewJSONFromMap(k.Environment),
		Frontend:        convertFrontendModuleArray(k.Frontend),
		Imports:         pq.StringArray(k.Imports),
		Interfaces:      abstraction.NewJSONFromMap(k.Interfaces),
		Resources:       abstraction.NewJSONFromMap(k.Resources),
	}
}

func convertKMDIArray(k []*kmiPB.KMDI) *[]kmi.KMDI {
	a := make([]kmi.KMDI, len(k))
	for i, d := range k {
		a[i] = convertKMDI(d)
	}
	return &a
}

// EncodeGRPCCreateContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain createcontainer request to a gRPC CreateContainer request.
func EncodeGRPCCreateContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.CreateContainerRequest)
	return &containerPB.CreateContainerRequest{
		RefID: uint32(req.RefID),
		KmiID: uint32(req.KmiID),
		Name:  req.Name,
	}, nil
}

// DecodeGRPCCreateContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC CreateContainer response to a messages/container.proto-domain createcontainer response.
func DecodeGRPCCreateContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*containerPB.CreateContainerResponse)
	return &container.CreateContainerResponse{
		ID:    response.ID,
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain removecontainer request to a gRPC RemoveContainer request.
func EncodeGRPCRemoveContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.RemoveContainerRequest)
	return &containerPB.RemoveContainerRequest{
		RefID: uint32(req.RefID),
		ID:    req.ID,
	}, nil
}

// DecodeGRPCRemoveContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveContainer response to a messages/container.proto-domain removecontainer response.
func DecodeGRPCRemoveContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*containerPB.RemoveContainerResponse)
	return &container.RemoveContainerResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCInstancesRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain instances request to a gRPC Instances request.
func EncodeGRPCInstancesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.InstancesRequest)
	return &containerPB.InstancesRequest{
		RefID: uint32(req.RefID),
	}, nil
}

// DecodeGRPCInstancesResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC Instances response to a messages/container.proto-domain instances response.
func DecodeGRPCInstancesResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*containerPB.InstancesResponse)
	return &container.InstancesResponse{
		Containers: pbContainersToContainers(response.Instances),
	}, nil
}

// EncodeGRPCStopContainerRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain stopcontainer request to a gRPC StopContainer request.
func EncodeGRPCStopContainerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.StopContainerRequest)
	return &containerPB.StopContainerRequest{
		RefID: uint32(req.RefID),
		ID:    req.ID,
	}, nil
}

// DecodeGRPCStopContainerResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC StopContainer response to a messages/container.proto-domain stopcontainer response.
func DecodeGRPCStopContainerResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*containerPB.StopContainerResponse)
	return &container.StopContainerResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCExecuteRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain execute request to a gRPC Execute request.
func EncodeGRPCExecuteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.ExecuteRequest)
	return &containerPB.ExecuteRequest{
		RefID: uint32(req.RefID),
		ID:    req.ID,
		Cmd:   req.CMD,
	}, nil
}

// DecodeGRPCExecuteResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC Execute response to a messages/container.proto-domain execute response.
func DecodeGRPCExecuteResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*containerPB.ExecuteResponse)
	return &container.ExecuteResponse{
		Error:    getError(response.Error),
		Response: response.Response,
	}, nil
}

// EncodeGRPCGetEnvRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain getenv request to a gRPC GetEnv request.
func EncodeGRPCGetEnvRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.GetEnvRequest)
	return &containerPB.GetEnvRequest{
		RefID: uint32(req.RefID),
		ID:    req.ID,
		Key:   req.Key,
	}, nil
}

// DecodeGRPCGetEnvResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetEnv response to a messages/container.proto-domain getenv response.
func DecodeGRPCGetEnvResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*containerPB.GetEnvResponse)
	return &container.GetEnvResponse{
		Value: response.Value,
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCSetEnvRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain setenv request to a gRPC SetEnv request.
func EncodeGRPCSetEnvRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.SetEnvRequest)
	return &containerPB.SetEnvRequest{
		RefID: uint32(req.RefID),
		ID:    req.ID,
		Key:   req.Key,
		Value: req.Value,
	}, nil
}

// DecodeGRPCSetEnvResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC SetEnv response to a messages/container.proto-domain setenv response.
func DecodeGRPCSetEnvResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*containerPB.SetEnvResponse)
	return &container.SetEnvResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCIDForNameRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain idforname request to a gRPC IDForName request.
func EncodeGRPCIDForNameRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.IDForNameRequest)
	return &containerPB.IDForNameRequest{
		RefID: uint32(req.RefID),
		Name:  req.Name,
	}, nil
}

// DecodeGRPCIDForNameResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC IDForName response to a messages/container.proto-domain idforname response.
func DecodeGRPCIDForNameResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*containerPB.IDForNameResponse)
	return &container.IDForNameResponse{
		ID:    response.ID,
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCGetContainerKMIRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain getcontainerkmi request to a gRPC GetContainerKMI request.
func EncodeGRPCGetContainerKMIRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.GetContainerKMIRequest)
	return &containerPB.GetContainerKMIRequest{
		ContainerID: req.ContainerID,
	}, nil
}

// DecodeGRPCGetContainerKMIResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetContainerKMI response to a messages/container.proto-domain getcontainerkmi response.
func DecodeGRPCGetContainerKMIResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*containerPB.GetContainerKMIResponse)
	kmi := convertKMI(response.ContainerKMI)
	return &container.GetContainerKMIResponse{
		Error:        getError(response.Error),
		ContainerKMI: *kmi,
	}, nil
}
