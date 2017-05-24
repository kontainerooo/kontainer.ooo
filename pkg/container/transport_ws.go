package container

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
)

// MakeWebsocketService makes a set of container Endpoints available as a websocket Service
func MakeWebsocketService(endpoints Endpoints) *ws.ServiceDescription {
	service := ws.NewServiceDescription("containerService", ws.ProtoIDFromString("CNT"))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"CreateContainer",
		ws.ProtoIDFromString("CRT"),
		endpoints.CreateContainerEndpoint,
		DecodeWSCreateContainerRequest,
		EncodeGRPCCreateContainerResponse,
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
		"StopContainer",
		ws.ProtoIDFromString("STO"),
		endpoints.StopContainerEndpoint,
		DecodeWSStopContainerRequest,
		EncodeGRPCStopContainerResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"Execute",
		ws.ProtoIDFromString("EXE"),
		endpoints.ExecuteEndpoint,
		DecodeWSExecuteRequest,
		EncodeGRPCExecuteResponse,
	))

	return service
}

// DecodeWSCreateContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS CreateContainer request to a messages/container.proto-domain createcontainer request.
func DecodeWSCreateContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.CreateContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCCreateContainerRequest(ctx, req)
}

// DecodeWSRemoveContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS RemoveContainer request to a messages/container.proto-domain removecontainer request.
func DecodeWSRemoveContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.RemoveContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCRemoveContainerRequest(ctx, req)
}

// DecodeWSInstancesRequest is a websocket.DecodeRequestFunc that converts a
// WS Instances request to a messages/container.proto-domain instances request.
func DecodeWSInstancesRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.InstancesRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCInstancesRequest(ctx, req)
}

// DecodeWSStopContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS StopContainer request to a messages/container.proto-domain stopcontainer request.
func DecodeWSStopContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.StopContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCStopContainerRequest(ctx, req)
}

// DecodeWSExecuteRequest is a websocket.DecodeRequestFunc that converts a
// WS Execute request to a messages/container.proto-domain execute request.
func DecodeWSExecuteRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.ExecuteRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCExecuteRequest(ctx, req)
}
