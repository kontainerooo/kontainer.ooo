package user

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/ttdennis/kontainer.io/pkg/pb"
	oldcontext "golang.org/x/net/context"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC UserServiceServer
func MakeGRPCServer(ctx context.Context, endpoints Endpoints, logger log.Logger) pb.UserServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		createUser: grpctransport.NewServer(
			ctx,
			endpoints.CreateUserEndpoint,
			DecodeGRPCCreateUserRequest,
			EncodeGRPCCreateUserResponse,
			options...,
		),
		editUser: grpctransport.NewServer(
			ctx,
			endpoints.EditUserEndpoint,
			DecodeGRPCEditUserRequest,
			EncodeGRPCEditUserResponse,
			options...,
		),
		changeUsername: grpctransport.NewServer(
			ctx,
			endpoints.ChangeUsernameEndpoint,
			DecodeGRPCChangeUsernameRequest,
			EncodeGRPCChangeUsernameResponse,
			options...,
		),
		deleteUser: grpctransport.NewServer(
			ctx,
			endpoints.DeleteUserEndpoint,
			DecodeGRPCDeleteUserRequest,
			EncodeGRPCDeleteUserResponse,
			options...,
		),
		resetPassword: grpctransport.NewServer(
			ctx,
			endpoints.ResetPasswordEndpoint,
			DecodeGRPCResetPasswordRequest,
			EncodeGRPCResetPasswordResponse,
			options...,
		),
		getUser: grpctransport.NewServer(
			ctx,
			endpoints.GetUserEndpoint,
			DecodeGRPCGetUserRequest,
			EncodeGRPCGetUserResponse,
			options...,
		),
	}
}

type grpcServer struct {
	createUser     grpctransport.Handler
	editUser       grpctransport.Handler
	changeUsername grpctransport.Handler
	deleteUser     grpctransport.Handler
	resetPassword  grpctransport.Handler
	getUser        grpctransport.Handler
}

func (s *grpcServer) CreateUser(ctx oldcontext.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, res, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.CreateUserResponse), nil
}

func (s *grpcServer) EditUser(ctx oldcontext.Context, req *pb.EditUserRequest) (*pb.EditUserResponse, error) {
	_, res, err := s.editUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.EditUserResponse), nil
}

func (s *grpcServer) ChangeUsername(ctx oldcontext.Context, req *pb.ChangeUsernameRequest) (*pb.ChangeUsernameResponse, error) {
	_, res, err := s.changeUsername.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ChangeUsernameResponse), nil
}

func (s *grpcServer) DeleteUser(ctx oldcontext.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, res, err := s.deleteUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.DeleteUserResponse), nil
}

func (s *grpcServer) ResetPassword(ctx oldcontext.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	_, res, err := s.resetPassword.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ResetPasswordResponse), nil
}

func (s *grpcServer) GetUser(ctx oldcontext.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, res, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetUserResponse), nil
}

func convertPbAddress(pb *pb.Address) *Address {
	return &Address{
		ID:         uint(pb.ID),
		Postcode:   pb.Postcode,
		City:       pb.City,
		Country:    pb.Country,
		Street:     pb.Street,
		Houseno:    int(pb.Houseno),
		Additional: pb.Additional,
	}
}

func convertPbConfig(pb *pb.Config) (*Config, *Address) {
	cfg := &Config{
		Admin:     pb.Admin,
		Email:     pb.Email,
		Password:  pb.Password,
		Salt:      pb.Salt,
		Image:     pb.Image,
		Phone:     pb.Phone,
		AddressID: uint(pb.AddressID),
	}

	adr := convertPbAddress(pb.Address)
	cfg.Address = *adr

	return cfg, adr
}

func convertUser(usr *User) *pb.User {
	return &pb.User{
		ID:       uint32(usr.ID),
		Username: usr.Username,
		Config: &pb.Config{
			Admin:     usr.Admin,
			Email:     usr.Email,
			Phone:     usr.Phone,
			Image:     usr.Image,
			AddressID: uint32(usr.AddressID),
			Address: &pb.Address{
				ID:         uint32(usr.Address.ID),
				Postcode:   usr.Address.Postcode,
				City:       usr.Address.City,
				Country:    usr.Address.Country,
				Street:     usr.Address.Street,
				Houseno:    int32(usr.Address.Houseno),
				Additional: usr.Address.Additional,
			},
		},
	}
}

// DecodeGRPCCreateUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateUser request to a user-domain createUser request.
func DecodeGRPCCreateUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateUserRequest)
	cfg, adr := convertPbConfig(req.Config)
	return createUserRequest{
		Username: req.Username,
		Cfg:      cfg,
		Adr:      adr,
	}, nil
}

// DecodeGRPCEditUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC EditUser request to a user-domain editUser request.
func DecodeGRPCEditUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.EditUserRequest)
	cfg, _ := convertPbConfig(req.Config)
	return editUserRequest{
		ID:  uint(req.ID),
		Cfg: cfg,
	}, nil
}

// DecodeGRPCChangeUsernameRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC ChangeUsername request to a user-domain changeUsername request.
func DecodeGRPCChangeUsernameRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ChangeUsernameRequest)
	return changeUsernameRequest{
		ID:       uint(req.ID),
		Username: req.Username,
	}, nil
}

// DecodeGRPCDeleteUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC DeleteUser request to a user-domain deleteUser request.
func DecodeGRPCDeleteUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DeleteUserRequest)
	return deleteUserRequest{
		ID: uint(req.ID),
	}, nil
}

// DecodeGRPCResetPasswordRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC ResetPassword request to a user-domain resetPassword request.
func DecodeGRPCResetPasswordRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ResetPasswordRequest)
	return resetPasswordRequest{
		Email: req.Email,
	}, nil
}

// DecodeGRPCGetUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetUser request to a user-domain getUser request.
func DecodeGRPCGetUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetUserRequest)
	return getUserRequest{
		ID: uint(req.ID),
	}, nil
}

// EncodeGRPCCreateUserResponse is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createUser response to a gRPC CreateUser response.
func EncodeGRPCCreateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(createUserResponse)
	return &pb.CreateUserResponse{
		ID:    uint32(res.ID),
		Error: res.Error.Error(),
	}, nil
}

// EncodeGRPCEditUserResponse is a transport/grpc.EncodeRequestFunc that converts a
// user-domain editUser response to a gRPC EditUser response.
func EncodeGRPCEditUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(editUserResponse)
	return &pb.EditUserResponse{
		Error: res.Error.Error(),
	}, nil
}

// EncodeGRPCChangeUsernameResponse is a transport/grpc.EncodeRequestFunc that converts a
// user-domain changeUsername response to a gRPC ChangeUsername response.
func EncodeGRPCChangeUsernameResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(changeUsernameResponse)
	return &pb.ChangeUsernameResponse{
		Error: res.Error.Error(),
	}, nil
}

// EncodeGRPCDeleteUserResponse is a transport/grpc.EncodeRequestFunc that converts a
// user-domain deleteUser response to a gRPC DeleteUser response.
func EncodeGRPCDeleteUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(deleteUserResponse)
	return &pb.DeleteUserResponse{
		Error: res.Error.Error(),
	}, nil
}

// EncodeGRPCResetPasswordResponse is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createUser response to a gRPC CreateUser response.
func EncodeGRPCResetPasswordResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(resetPasswordResponse)
	return &pb.ResetPasswordResponse{
		Error: res.Error.Error(),
	}, nil
}

// EncodeGRPCGetUserResponse is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getUser response to a gRPC GetUser response.
func EncodeGRPCGetUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(getUserResponse)
	user := convertUser(res.User)
	return &pb.GetUserResponse{
		User:  user,
		Error: res.Error.Error(),
	}, nil
}
