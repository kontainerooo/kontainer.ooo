package routing

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
)

// MakeWebsocketService makes a set of routing Endpoints available as a websocket Service
func MakeWebsocketService(endpoints Endpoints) *ws.ServiceDescription {
	service := ws.NewServiceDescription("routingService", ws.ProtoIDFromString("RTG"))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"CreateConfig",
		ws.ProtoIDFromString("CRT"),
		endpoints.CreateConfigEndpoint,
		DecodeWSCreateConfigRequest,
		EncodeGRPCCreateConfigResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"EditConfig",
		ws.ProtoIDFromString("EDT"),
		endpoints.EditConfigEndpoint,
		DecodeWSEditConfigRequest,
		EncodeGRPCEditConfigResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"GetConfig",
		ws.ProtoIDFromString("GET"),
		endpoints.GetConfigEndpoint,
		DecodeWSGetConfigRequest,
		EncodeGRPCGetConfigResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"RemoveConfig",
		ws.ProtoIDFromString("REM"),
		endpoints.RemoveConfigEndpoint,
		DecodeWSRemoveConfigRequest,
		EncodeGRPCRemoveConfigResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"AddLocation",
		ws.ProtoIDFromString("ALO"),
		endpoints.AddLocationEndpoint,
		DecodeWSAddLocationRequest,
		EncodeGRPCAddLocationResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"RemoveLocation",
		ws.ProtoIDFromString("RLO"),
		endpoints.RemoveLocationEndpoint,
		DecodeWSRemoveLocationRequest,
		EncodeGRPCRemoveLocationResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"ChangeListenStatement",
		ws.ProtoIDFromString("CLS"),
		endpoints.ChangeListenStatementEndpoint,
		DecodeWSChangeListenStatementRequest,
		EncodeGRPCChangeListenStatementResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"AddServerName",
		ws.ProtoIDFromString("ASM"),
		endpoints.AddServerNameEndpoint,
		DecodeWSAddServerNameRequest,
		EncodeGRPCAddServerNameResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"RemoveServerName",
		ws.ProtoIDFromString("RSM"),
		endpoints.RemoveServerNameEndpoint,
		DecodeWSRemoveServerNameRequest,
		EncodeGRPCRemoveServerNameResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"Configurations",
		ws.ProtoIDFromString("CON"),
		endpoints.ConfigurationsEndpoint,
		DecodeWSConfigurationsRequest,
		EncodeGRPCConfigurationsResponse,
	))

	return service
}

// DecodeWSCreateConfigRequest is a websocket.DecodeRequestFunc that converts a
// WS CreateConfig request to a messages/routing.proto-domain createconfig request.
func DecodeWSCreateConfigRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.CreateConfigRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCCreateConfigRequest(ctx, req)
}

// DecodeWSEditConfigRequest is a websocket.DecodeRequestFunc that converts a
// WS EditConfig request to a messages/routing.proto-domain editconfig request.
func DecodeWSEditConfigRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.EditConfigRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCEditConfigRequest(ctx, req)
}

// DecodeWSGetConfigRequest is a websocket.DecodeRequestFunc that converts a
// WS GetConfig request to a messages/routing.proto-domain getconfig request.
func DecodeWSGetConfigRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.GetConfigRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCGetConfigRequest(ctx, req)
}

// DecodeWSRemoveConfigRequest is a websocket.DecodeRequestFunc that converts a
// WS RemoveConfig request to a messages/routing.proto-domain removeconfig request.
func DecodeWSRemoveConfigRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.RemoveConfigRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCRemoveConfigRequest(ctx, req)
}

// DecodeWSAddLocationRequest is a websocket.DecodeRequestFunc that converts a
// WS AddLocation request to a messages/routing.proto-domain addlocation request.
func DecodeWSAddLocationRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.AddLocationRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCAddLocationRequest(ctx, req)
}

// DecodeWSRemoveLocationRequest is a websocket.DecodeRequestFunc that converts a
// WS RemoveLocation request to a messages/routing.proto-domain removelocation request.
func DecodeWSRemoveLocationRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.RemoveLocationRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCRemoveLocationRequest(ctx, req)
}

// DecodeWSChangeListenStatementRequest is a websocket.DecodeRequestFunc that converts a
// WS ChangeListenStatement request to a messages/routing.proto-domain changelistenstatement request.
func DecodeWSChangeListenStatementRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.ChangeListenStatementRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCChangeListenStatementRequest(ctx, req)
}

// DecodeWSAddServerNameRequest is a websocket.DecodeRequestFunc that converts a
// WS AddServerName request to a messages/routing.proto-domain addservername request.
func DecodeWSAddServerNameRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.AddServerNameRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCAddServerNameRequest(ctx, req)
}

// DecodeWSRemoveServerNameRequest is a websocket.DecodeRequestFunc that converts a
// WS RemoveServerName request to a messages/routing.proto-domain removeservername request.
func DecodeWSRemoveServerNameRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.RemoveServerNameRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCRemoveServerNameRequest(ctx, req)
}

// DecodeWSConfigurationsRequest is a websocket.DecodeRequestFunc that converts a
// WS Configurations request to a messages/routing.proto-domain configurations request.
func DecodeWSConfigurationsRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.ConfigurationsRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCConfigurationsRequest(ctx, req)
}
