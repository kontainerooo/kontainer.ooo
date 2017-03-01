package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a struct which collects all endpoints for the user service
type Endpoints struct {
	CreateUserEndpoint     endpoint.Endpoint
	EditUserEndpoint       endpoint.Endpoint
	ChangeUserNameEndpoint endpoint.Endpoint
	DeleteUserEndpoint     endpoint.Endpoint
	ResetPasswordEndpoint  endpoint.Endpoint
	GetUserEndpoint        endpoint.Endpoint
}

type createUserRequest struct {
	Username string
	Cfg      *Config
	Adr      *Address
}

type createUserResponse struct {
	ID    uint
	Error error
}

func MakeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		createReq := request.(createUserRequest)
		id, err := s.CreateUser(createReq.Username, createReq.Cfg, createReq.Adr)
		return createUserResponse{
			ID:    id,
			Error: err,
		}, nil
	}
}

type editUserRequest struct {
	ID  uint
	Cfg *Config
}

type editUserResponse struct {
	Error error
}

func MakeEditUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		editReq := request.(editUserRequest)
		err := s.EditUser(editReq.ID, editReq.Cfg)
		return editUserResponse{
			Error: err,
		}, nil
	}
}

type changeUsernameRequest struct {
	ID       uint
	Username string
}

type changeUsernameResponse struct {
	Error error
}

func MakeChangeUsernameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		changeReq := request.(changeUsernameRequest)
		err := s.ChangeUsername(changeReq.ID, changeReq.Username)
		return changeUsernameResponse{
			Error: err,
		}, nil
	}
}

type deleteUserRequest struct {
	ID uint
}

type deleteUserResponse struct {
	Error error
}

func MakeDeleteUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		deleteReq := request.(deleteUserRequest)
		err := s.DeleteUser(deleteReq.ID)
		return deleteUserResponse{
			Error: err,
		}, nil
	}
}

type resetPasswordRequest struct {
	Email string
}

type resetPasswordResponse struct {
	Error error
}

func MakeResetPasswordEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		resetPasswordReq := request.(resetPasswordRequest)
		err := s.ResetPassword(resetPasswordReq.Email)
		return resetPasswordResponse{
			Error: err,
		}, nil
	}
}

type getUserRequest struct {
	ID uint
}
type getUserResponse struct {
	User  *User
	Error error
}

func MakeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		getRequest := request.(getUserRequest)
		user := &User{}
		err := s.GetUser(getRequest.ID, user)
		return getUserResponse{
			User:  user,
			Error: err,
		}, nil
	}
}
