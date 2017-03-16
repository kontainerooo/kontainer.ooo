package user

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/ttdennis/kontainer.io/pkg/pb"
	ws "github.com/ttdennis/kontainer.io/pkg/websocket"
)

// MakeWebsocketService makes a set of user Endpoints available as a websocket Service
func MakeWebsocketService(endpoints Endpoints) *ws.ServiceDescription {
	service := ws.NewServiceDescription("userService", ws.ProtoIDFromString("USR"))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"CreateUser",
		ws.ProtoIDFromString("CRT"),
		endpoints.CreateUserEndpoint,
		DecodeWSCreateUserRequest,
		EncodeGRPCCreateUserResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"EditUser",
		ws.ProtoIDFromString("EDT"),
		endpoints.EditUserEndpoint,
		DecodeWSEditUserRequest,
		EncodeGRPCEditUserResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"ChangeUsername",
		ws.ProtoIDFromString("CHU"),
		endpoints.ChangeUsernameEndpoint,
		DecodeWSChangeUsernameRequest,
		EncodeGRPCChangeUsernameResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"DeleteUser",
		ws.ProtoIDFromString("DLT"),
		endpoints.DeleteUserEndpoint,
		DecodeWSDeleteUserRequest,
		EncodeGRPCDeleteUserResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"ResetPassword",
		ws.ProtoIDFromString("RST"),
		endpoints.ResetPasswordEndpoint,
		DecodeWSResetPasswordRequest,
		EncodeGRPCResetPasswordResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"GetUser",
		ws.ProtoIDFromString("GET"),
		endpoints.GetUserEndpoint,
		DecodeWSGetUserRequest,
		EncodeGRPCGetUserResponse,
	))

	return service
}

// DecodeWSCreateUserRequest is a websocket.DecodeRequestFunc that converts a
// WS CreateUser request to a messages/user.proto-domain createuser request.
func DecodeWSCreateUserRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.CreateUserRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCCreateUserRequest(ctx, req)
}

// DecodeWSEditUserRequest is a websocket.DecodeRequestFunc that converts a
// WS EditUser request to a messages/user.proto-domain edituser request.
func DecodeWSEditUserRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.EditUserRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCEditUserRequest(ctx, req)
}

// DecodeWSChangeUsernameRequest is a websocket.DecodeRequestFunc that converts a
// WS ChangeUsername request to a messages/user.proto-domain changeusername request.
func DecodeWSChangeUsernameRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.ChangeUsernameRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCChangeUsernameRequest(ctx, req)
}

// DecodeWSDeleteUserRequest is a websocket.DecodeRequestFunc that converts a
// WS DeleteUser request to a messages/user.proto-domain deleteuser request.
func DecodeWSDeleteUserRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.DeleteUserRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCDeleteUserRequest(ctx, req)
}

// DecodeWSResetPasswordRequest is a websocket.DecodeRequestFunc that converts a
// WS ResetPassword request to a messages/user.proto-domain resetpassword request.
func DecodeWSResetPasswordRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.ResetPasswordRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCResetPasswordRequest(ctx, req)
}

// DecodeWSGetUserRequest is a websocket.DecodeRequestFunc that converts a
// WS GetUser request to a messages/user.proto-domain getuser request.
func DecodeWSGetUserRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.GetUserRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCGetUserRequest(ctx, req)
}
