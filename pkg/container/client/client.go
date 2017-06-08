package client

import (
	"context"
	"errors"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"github.com/kontainerooo/kontainer.ooo/pkg/container"

	containerPB "github.com/kontainerooo/kontainer.ooo/pkg/container/pb"
	kmiClient "github.com/kontainerooo/kontainer.ooo/pkg/kmi/client"
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
	var SetLinkEndpoint endpoint.Endpoint
	{
		SetLinkEndpoint = grpctransport.NewClient(
			conn,
			"container.ContainerService",
			"SetLink",
			EncodeGRPCSetLinkRequest,
			DecodeGRPCSetLinkResponse,
			containerPB.SetLinkResponse{},
		).Endpoint()
	}

	var RemoveLinkEndpoint endpoint.Endpoint
	{
		RemoveLinkEndpoint = grpctransport.NewClient(
			conn,
			"container.ContainerService",
			"RemoveLink",
			EncodeGRPCRemoveLinkRequest,
			DecodeGRPCRemoveLinkResponse,
			containerPB.RemoveLinkResponse{},
		).Endpoint()
	}

	var GetLinksEndpoint endpoint.Endpoint
	{
		GetLinksEndpoint = grpctransport.NewClient(
			conn,
			"container.ContainerService",
			"GetLinks",
			EncodeGRPCGetLinksRequest,
			DecodeGRPCGetLinksResponse,
			containerPB.GetLinksResponse{},
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
		SetLinkEndpoint:         SetLinkEndpoint,
		RemoveLinkEndpoint:      RemoveLinkEndpoint,
		GetLinksEndpoint:        GetLinksEndpoint,
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
			KMI: container.CKMI{
				KMI: *kmiClient.ConvertKMI(c.Kmi),
			},
		})
	}

	return cts
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
	kmi := kmiClient.ConvertKMI(response.ContainerKMI)
	return &container.GetContainerKMIResponse{
		Error:        getError(response.Error),
		ContainerKMI: *kmi,
	}, nil
}

// EncodeGRPCSetLinkRequest is a transport/grpc.EncodeRequestFunc that converts a
// container.proto-domain setlink request to a gRPC SetLink request.
func EncodeGRPCSetLinkRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.SetLinkRequest)
	return &containerPB.SetLinkRequest{
		RefID:         uint32(req.RefID),
		ContainerID:   req.ContainerID,
		LinkID:        req.LinkID,
		LinkName:      req.LinkName,
		LinkInterface: req.LinkInterface,
	}, nil
}

// DecodeGRPCSetLinkResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC SetLink response to a container.proto-domain setlink response.
func DecodeGRPCSetLinkResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*containerPB.SetLinkResponse)
	return &container.SetLinkResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveLinkRequest is a transport/grpc.EncodeRequestFunc that converts a
// container.proto-domain removelink request to a gRPC RemoveLink request.
func EncodeGRPCRemoveLinkRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.RemoveLinkRequest)
	return &containerPB.RemoveLinkRequest{
		RefID:         uint32(req.RefID),
		ContainerID:   req.ContainerID,
		LinkID:        req.LinkID,
		LinkName:      req.LinkName,
		LinkInterface: req.LinkInterface,
	}, nil
}

// DecodeGRPCRemoveLinkResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveLink response to a container.proto-domain removelink response.
func DecodeGRPCRemoveLinkResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*containerPB.RemoveLinkResponse)
	return &container.RemoveLinkResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCGetLinksRequest is a transport/grpc.EncodeRequestFunc that converts a
// container.proto-domain getlinks request to a gRPC GetLinks request.
func EncodeGRPCGetLinksRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*container.GetLinksRequest)
	return &containerPB.GetLinksRequest{
		RefID:       uint32(req.RefID),
		ContainerID: req.ContainerID,
	}, nil
}

// DecodeGRPCGetLinksResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetLinks response to a container.proto-domain getlinks response.
func DecodeGRPCGetLinksResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*containerPB.GetLinksResponse)

	arrayMap := make(map[string][]string)
	for k, v := range response.Links {
		arrayMap[k] = strings.Split(v, ",")
	}

	return &container.GetLinksResponse{
		Error: getError(response.Error),
		Links: arrayMap,
	}, nil
}
