package kmi

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi/pb"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
)

// MakeWebsocketService makes a set of kmi Endpoints available as a websocket Service
func MakeWebsocketService(endpoints Endpoints) *ws.ServiceDescription {
	service := ws.NewServiceDescription("kmiService", ws.ProtoIDFromString("KMI"))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"AddKMI",
		ws.ProtoIDFromString("ADD"),
		endpoints.AddKMIEndpoint,
		DecodeWSAddKMIRequest,
		EncodeGRPCAddKMIResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"RemoveKMI",
		ws.ProtoIDFromString("REM"),
		endpoints.RemoveKMIEndpoint,
		DecodeWSRemoveKMIRequest,
		EncodeGRPCRemoveKMIResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"GetKMI",
		ws.ProtoIDFromString("GET"),
		endpoints.GetKMIEndpoint,
		DecodeWSGetKMIRequest,
		EncodeGRPCGetKMIResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"KMI",
		ws.ProtoIDFromString("ALL"),
		endpoints.KMIEndpoint,
		DecodeWSKMIRequest,
		EncodeGRPCKMIResponse,
	))

	return service
}

// DecodeWSAddKMIRequest is a websocket.DecodeRequestFunc that converts a
// WS AddKMI request to a messages/kmi.proto-domain addkmi request.
func DecodeWSAddKMIRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.AddKMIRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCAddKMIRequest(ctx, req)
}

// DecodeWSRemoveKMIRequest is a websocket.DecodeRequestFunc that converts a
// WS RemoveKMI request to a messages/kmi.proto-domain removekmi request.
func DecodeWSRemoveKMIRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.RemoveKMIRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCRemoveKMIRequest(ctx, req)
}

// DecodeWSGetKMIRequest is a websocket.DecodeRequestFunc that converts a
// WS GetKMI request to a messages/kmi.proto-domain getkmi request.
func DecodeWSGetKMIRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.GetKMIRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCGetKMIRequest(ctx, req)
}

// DecodeWSKMIRequest is a websocket.DecodeRequestFunc that converts a
// WS KMI request to a messages/kmi.proto-domain kmi request.
func DecodeWSKMIRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.KMIRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCKMIRequest(ctx, req)
}
