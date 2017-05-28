package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"github.com/kontainerooo/kontainer.ooo/pkg/user"
	"github.com/kontainerooo/kontainer.ooo/pkg/user/pb"
)

// New creates a set of endpoints based on a gRPC connection
func New(conn *grpc.ClientConn, logger log.Logger) *user.Endpoints {

	var CreateUserEndpoint endpoint.Endpoint
	{
		CreateUserEndpoint = grpctransport.NewClient(
			conn,
			"user.UserService",
			"CreateUser",
			EncodeGRPCCreateUserRequest,
			DecodeGRPCCreateUserResponse,
			pb.CreateUserResponse{},
		).Endpoint()
	}

	var EditUserEndpoint endpoint.Endpoint
	{
		EditUserEndpoint = grpctransport.NewClient(
			conn,
			"user.UserService",
			"EditUser",
			EncodeGRPCEditUserRequest,
			DecodeGRPCEditUserResponse,
			pb.EditUserResponse{},
		).Endpoint()
	}

	var ChangeUsernameEndpoint endpoint.Endpoint
	{
		ChangeUsernameEndpoint = grpctransport.NewClient(
			conn,
			"user.UserService",
			"ChangeUsername",
			EncodeGRPCChangeUsernameRequest,
			DecodeGRPCChangeUsernameResponse,
			pb.ChangeUsernameResponse{},
		).Endpoint()
	}

	var DeleteUserEndpoint endpoint.Endpoint
	{
		DeleteUserEndpoint = grpctransport.NewClient(
			conn,
			"user.UserService",
			"DeleteUser",
			EncodeGRPCDeleteUserRequest,
			DecodeGRPCDeleteUserResponse,
			pb.DeleteUserResponse{},
		).Endpoint()
	}

	var ResetPasswordEndpoint endpoint.Endpoint
	{
		ResetPasswordEndpoint = grpctransport.NewClient(
			conn,
			"user.UserService",
			"ResetPassword",
			EncodeGRPCResetPasswordRequest,
			DecodeGRPCResetPasswordResponse,
			pb.ResetPasswordResponse{},
		).Endpoint()
	}

	var GetUserEndpoint endpoint.Endpoint
	{
		GetUserEndpoint = grpctransport.NewClient(
			conn,
			"user.UserService",
			"GetUser",
			EncodeGRPCGetUserRequest,
			DecodeGRPCGetUserResponse,
			pb.GetUserResponse{},
		).Endpoint()
	}

	var CheckLoginCredentialsEndpoint endpoint.Endpoint
	{
		CheckLoginCredentialsEndpoint = grpctransport.NewClient(
			conn,
			"user.UserService",
			"CheckLoginCredentials",
			EncodeGRPCCheckLoginCredentialsRequest,
			DecodeGRPCCheckLoginCredentialsResponse,
			pb.CheckLoginCredentialsResponse{},
		).Endpoint()
	}

	return &user.Endpoints{
		CreateUserEndpoint:            CreateUserEndpoint,
		EditUserEndpoint:              EditUserEndpoint,
		ChangeUsernameEndpoint:        ChangeUsernameEndpoint,
		DeleteUserEndpoint:            DeleteUserEndpoint,
		ResetPasswordEndpoint:         ResetPasswordEndpoint,
		GetUserEndpoint:               GetUserEndpoint,
		CheckLoginCredentialsEndpoint: CheckLoginCredentialsEndpoint,
	}
}

func getError(e string) error {
	if e != "" {
		return errors.New(e)
	}
	return nil
}

func convertPBUser(usr *pb.User) *user.User {
	cfg, _ := user.ConvertPbConfig(usr.Config)
	return &user.User{
		ID:       uint(usr.ID),
		Username: usr.Username,
		Config:   *cfg,
	}
}

// EncodeGRPCCreateUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/user.proto-domain createuser request to a gRPC CreateUser request.
func EncodeGRPCCreateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	fmt.Println(request)
	req := request.(*user.CreateUserRequest)
	return &pb.CreateUserRequest{
		Username: req.Username,
		Config:   user.ConvertConfig(req.Cfg, true),
	}, nil
}

// DecodeGRPCCreateUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC CreateUser response to a messages/user.proto-domain createuser response.
func DecodeGRPCCreateUserResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.CreateUserResponse)
	return &user.CreateUserResponse{
		ID:    uint(response.ID),
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCEditUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/user.proto-domain edituser request to a gRPC EditUser request.
func EncodeGRPCEditUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*user.EditUserRequest)
	return &pb.EditUserRequest{
		ID:     uint32(req.ID),
		Config: user.ConvertConfig(req.Cfg, true),
	}, nil
}

// DecodeGRPCEditUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC EditUser response to a messages/user.proto-domain edituser response.
func DecodeGRPCEditUserResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.EditUserResponse)
	return &user.EditUserResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCChangeUsernameRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/user.proto-domain changeusername request to a gRPC ChangeUsername request.
func EncodeGRPCChangeUsernameRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*user.ChangeUsernameRequest)
	return &pb.ChangeUsernameRequest{
		ID:       uint32(req.ID),
		Username: req.Username,
	}, nil
}

// DecodeGRPCChangeUsernameResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC ChangeUsername response to a messages/user.proto-domain changeusername response.
func DecodeGRPCChangeUsernameResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.ChangeUsernameResponse)
	return &user.ChangeUsernameResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCDeleteUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/user.proto-domain deleteuser request to a gRPC DeleteUser request.
func EncodeGRPCDeleteUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*user.DeleteUserRequest)
	return &pb.DeleteUserRequest{
		ID: uint32(req.ID),
	}, nil
}

// DecodeGRPCDeleteUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC DeleteUser response to a messages/user.proto-domain deleteuser response.
func DecodeGRPCDeleteUserResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.DeleteUserResponse)
	return &user.DeleteUserResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCResetPasswordRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/user.proto-domain resetpassword request to a gRPC ResetPassword request.
func EncodeGRPCResetPasswordRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*user.ResetPasswordRequest)
	return &pb.ResetPasswordRequest{
		Email: req.Email,
	}, nil
}

// DecodeGRPCResetPasswordResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC ResetPassword response to a messages/user.proto-domain resetpassword response.
func DecodeGRPCResetPasswordResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.ResetPasswordResponse)
	return &user.ResetPasswordResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCGetUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/user.proto-domain getuser request to a gRPC GetUser request.
func EncodeGRPCGetUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*user.GetUserRequest)
	return &pb.GetUserRequest{
		ID: uint32(req.ID),
	}, nil
}

// DecodeGRPCGetUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetUser response to a messages/user.proto-domain getuser response.
func DecodeGRPCGetUserResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.GetUserResponse)
	return &user.GetUserResponse{
		User:  convertPBUser(response.User),
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCCheckLoginCredentialsRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/user.proto-domain CheckLoginCredentials request to a gRPC CheckLoginCredentials request.
func EncodeGRPCCheckLoginCredentialsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*user.CheckLoginCredentialsRequest)
	return &pb.CheckLoginCredentialsRequest{
		Username: req.Username,
		Password: req.Password,
	}, nil
}

// DecodeGRPCCheckLoginCredentialsResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC CheckLoginCredentials response to a messages/user.proto-domain getuser response.
func DecodeGRPCCheckLoginCredentialsResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.CheckLoginCredentialsResponse)
	return &user.CheckLoginCredentialsResponse{
		ID: uint(response.ID),
	}, nil
}
