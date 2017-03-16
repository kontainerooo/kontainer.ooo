package containerlifecycle

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/ttdennis/kontainer.io/pkg/pb"
	ws "github.com/ttdennis/kontainer.io/pkg/websocket"
)

// MakeWebsocketService makes a set of containerlifecycle Endpoints available as a websocket Service
func MakeWebsocketService(endpoints Endpoints) *ws.ServiceDescription {
	service := ws.NewServiceDescription("containerlifecycleService", ws.ProtoIDFromString("CLS"))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"StartContainer",
		ws.ProtoIDFromString("SCN"),
		endpoints.StartContainerEndpoint,
		DecodeWSStartContainerRequest,
		EncodeGRPCStartContainerResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"StartCommand",
		ws.ProtoIDFromString("SCM"),
		endpoints.StartCommandEndpoint,
		DecodeWSStartCommandRequest,
		EncodeGRPCStartCommandResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"StopContainer",
		ws.ProtoIDFromString("STP"),
		endpoints.StopContainerEndpoint,
		DecodeWSStopContainerRequest,
		EncodeGRPCStopContainerResponse,
	))

	return service
}

// DecodeWSStartContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS StartContainer request to a messages/containerlifecycle.proto-domain startcontainer request.
func DecodeWSStartContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.StartContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCStartContainerRequest(ctx, req)
}

// DecodeWSStartCommandRequest is a websocket.DecodeRequestFunc that converts a
// WS StartCommand request to a messages/containerlifecycle.proto-domain startcommand request.
func DecodeWSStartCommandRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.StartCommandRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCStartCommandRequest(ctx, req)
}

// DecodeWSStopContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS StopContainer request to a messages/containerlifecycle.proto-domain stopcontainer request.
func DecodeWSStopContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.StopContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCStopContainerRequest(ctx, req)
}
