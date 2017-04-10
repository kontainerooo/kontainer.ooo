package customercontainer

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
)

// MakeWebsocketService makes a set of customercontainer Endpoints available as a websocket Service
func MakeWebsocketService(endpoints Endpoints) *ws.ServiceDescription {
	service, _ := ws.NewServiceDescription("customercontainerService", ws.ProtoIDFromString("CCS"))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"CreateContainer",
		ws.ProtoIDFromString("CRT"),
		endpoints.CreateContainerEndpoint,
		DecodeWSCreateContainerRequest,
		EncodeGRPCCreateContainerResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"EditContainer",
		ws.ProtoIDFromString("EDT"),
		endpoints.EditContainerEndpoint,
		DecodeWSEditContainerRequest,
		EncodeGRPCEditContainerResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"RemoveContainer",
		ws.ProtoIDFromString("REM"),
		endpoints.RemoveContainerEndpoint,
		DecodeWSRemoveContainerRequest,
		EncodeGRPCRemoveContainerResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"Instances",
		ws.ProtoIDFromString("ALL"),
		endpoints.InstancesEndpoint,
		DecodeWSInstancesRequest,
		EncodeGRPCInstancesResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"CreateDockerImage",
		ws.ProtoIDFromString("CDI"),
		endpoints.CreateDockerImageEndpoint,
		DecodeWSCreateDockerImageRequest,
		EncodeGRPCCreateDockerImageResponse,
	))

	return service
}

// DecodeWSCreateContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS CreateContainer request to a messages/customercontainer.proto-domain createcontainer request.
func DecodeWSCreateContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.CreateContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCCreateContainerRequest(ctx, req)
}

// DecodeWSEditContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS EditContainer request to a messages/customercontainer.proto-domain editcontainer request.
func DecodeWSEditContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.EditContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCEditContainerRequest(ctx, req)
}

// DecodeWSRemoveContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS RemoveContainer request to a messages/customercontainer.proto-domain removecontainer request.
func DecodeWSRemoveContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.RemoveContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCRemoveContainerRequest(ctx, req)
}

// DecodeWSInstancesRequest is a websocket.DecodeRequestFunc that converts a
// WS Instances request to a messages/customercontainer.proto-domain instances request.
func DecodeWSInstancesRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.InstancesRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCInstancesRequest(ctx, req)
}

// DecodeWSCreateDockerImageRequest is a websocket.DecodeRequestFunc that converts a
// WS CreateDockerImage request to a messages/customercontainer.proto-domain CreateDockerImage request.
func DecodeWSCreateDockerImageRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.CreateDockerImageRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCCreateDockerImageRequest(ctx, req)
}
