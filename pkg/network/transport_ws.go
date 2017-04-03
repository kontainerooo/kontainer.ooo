package network

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
)

// MakeWebsocketService makes a set of network Endpoints available as a websocket Service
func MakeWebsocketService(endpoints Endpoints) *ws.ServiceDescription {
	service := ws.NewServiceDescription("networkService", ws.ProtoIDFromString("network"))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"CreatePrimaryNetworkForContainer",
		ws.ProtoIDFromString(""),
		endpoints.CreatePrimaryNetworkForContainerEndpoint,
		DecodeWSCreatePrimaryNetworkForContainerRequest,
		EncodeGRPCCreatePrimaryNetworkForContainerResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"CreateNetwork",
		ws.ProtoIDFromString(""),
		endpoints.CreateNetworkEndpoint,
		DecodeWSCreateNetworkRequest,
		EncodeGRPCCreateNetworkResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"RemoveNetworkByName",
		ws.ProtoIDFromString(""),
		endpoints.RemoveNetworkByNameEndpoint,
		DecodeWSRemoveNetworkByNameRequest,
		EncodeGRPCRemoveNetworkByNameResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"AddContainerToNetwork",
		ws.ProtoIDFromString(""),
		endpoints.AddContainerToNetworkEndpoint,
		DecodeWSAddContainerToNetworkRequest,
		EncodeGRPCAddContainerToNetworkResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"RemoveContainerFromNetwork",
		ws.ProtoIDFromString(""),
		endpoints.RemoveContainerFromNetworkEndpoint,
		DecodeWSRemoveContainerFromNetworkRequest,
		EncodeGRPCRemoveContainerFromNetworkResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"ExposePortToContainer",
		ws.ProtoIDFromString(""),
		endpoints.ExposePortToContainerEndpoint,
		DecodeWSExposePortToContainerRequest,
		EncodeGRPCExposePortToContainerResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"RemovePortFromContainer",
		ws.ProtoIDFromString(""),
		endpoints.RemovePortFromContainerEndpoint,
		DecodeWSRemovePortFromContainerRequest,
		EncodeGRPCRemovePortFromContainerResponse,
	))

	return service
}

// DecodeWSCreatePrimaryNetworkForContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS CreatePrimaryNetworkForContainer request to a messages/network.proto-domain createprimarynetworkforcontainer request.
func DecodeWSCreatePrimaryNetworkForContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.CreatePrimaryNetworkForContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCCreatePrimaryNetworkForContainerRequest(ctx, req)
}

// DecodeWSCreateNetworkRequest is a websocket.DecodeRequestFunc that converts a
// WS CreateNetwork request to a messages/network.proto-domain createnetwork request.
func DecodeWSCreateNetworkRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.CreateNetworkRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCCreateNetworkRequest(ctx, req)
}

// DecodeWSRemoveNetworkByNameRequest is a websocket.DecodeRequestFunc that converts a
// WS RemoveNetworkByName request to a messages/network.proto-domain removenetworkbyname request.
func DecodeWSRemoveNetworkByNameRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.RemoveNetworkByNameRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCRemoveNetworkByNameRequest(ctx, req)
}

// DecodeWSAddContainerToNetworkRequest is a websocket.DecodeRequestFunc that converts a
// WS AddContainerToNetwork request to a messages/network.proto-domain addcontainertonetwork request.
func DecodeWSAddContainerToNetworkRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.AddContainerToNetworkRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCAddContainerToNetworkRequest(ctx, req)
}

// DecodeWSRemoveContainerFromNetworkRequest is a websocket.DecodeRequestFunc that converts a
// WS RemoveContainerFromNetwork request to a messages/network.proto-domain removecontainerfromnetwork request.
func DecodeWSRemoveContainerFromNetworkRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.RemoveContainerFromNetworkRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCRemoveContainerFromNetworkRequest(ctx, req)
}

// DecodeWSExposePortToContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS ExposePortToContainer request to a messages/network.proto-domain exposeporttocontainer request.
func DecodeWSExposePortToContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.ExposePortToContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCExposePortToContainerRequest(ctx, req)
}

// DecodeWSRemovePortFromContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS RemovePortFromContainer request to a messages/network.proto-domain removeportfromcontainer request.
func DecodeWSRemovePortFromContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.RemovePortFromContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCRemovePortFromContainerRequest(ctx, req)
}
