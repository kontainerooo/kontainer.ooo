package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a struct which collects all endpoints for the user service
type Endpoints struct {
	CreateUserEndpoint     endpoint.Endpoint
	EditUserEndpoint       endpoint.Endpoint
	ChangeUsernameEndpoint endpoint.Endpoint
	DeleteUserEndpoint     endpoint.Endpoint
	ResetPasswordEndpoint  endpoint.Endpoint
	GetUserEndpoint        endpoint.Endpoint
}

// CreateUserRequest is the request struct for the CreateUserEndpoint
type CreateUserRequest struct {
	Username string
	Cfg      *Config
	Adr      *Address
}

// CreateUserResponse is the response struct for the CreateUserEndpoint
type CreateUserResponse struct {
	ID    uint
	Error error
}

// MakeCreateUserEndpoint creates a gokit endpoint which invokes CreateUser
func MakeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		createReq := request.(CreateUserRequest)
		id, err := s.CreateUser(createReq.Username, createReq.Cfg, createReq.Adr)
		return CreateUserResponse{
			ID:    id,
			Error: err,
		}, nil
	}
}

// EditUserRequest is the request struct for the EditUserEndpoint
type EditUserRequest struct {
	ID  uint
	Cfg *Config
}

// EditUserResponse is the response struct for the EditUserEndpoint
type EditUserResponse struct {
	Error error
}

// MakeEditUserEndpoint creates a gokit endpoint which invokes EditUser
func MakeEditUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		editReq := request.(EditUserRequest)
		err := s.EditUser(editReq.ID, editReq.Cfg)
		return EditUserResponse{
			Error: err,
		}, nil
	}
}

// ChangeUsernameRequest is the request struct for the ChangeUsernameEndpoint
type ChangeUsernameRequest struct {
	ID       uint
	Username string
}

// ChangeUsernameResponse is the response struct for the ChangeUsernameResponse
type ChangeUsernameResponse struct {
	Error error
}

// MakeChangeUsernameEndpoint creates a gokit endpoint which invokes ChangeUsername
func MakeChangeUsernameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		changeReq := request.(ChangeUsernameRequest)
		err := s.ChangeUsername(changeReq.ID, changeReq.Username)
		return ChangeUsernameResponse{
			Error: err,
		}, nil
	}
}

// DeleteUserRequest is the request struct for the DeleteUserEndpoint
type DeleteUserRequest struct {
	ID uint
}

// DeleteUserResponse is the response struct for the DeleteUserEndpoint
type DeleteUserResponse struct {
	Error error
}

// MakeDeleteUserEndpoint creates a gokit endpoint which invokes DeleteUser
func MakeDeleteUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		deleteReq := request.(DeleteUserRequest)
		err := s.DeleteUser(deleteReq.ID)
		return DeleteUserResponse{
			Error: err,
		}, nil
	}
}

// ResetPasswordRequest is the request struct for the ResetPasswordEndpoint
type ResetPasswordRequest struct {
	Email string
}

// ResetPasswordResponse is the response struct for the ResetPasswordEndpoint
type ResetPasswordResponse struct {
	Error error
}

// MakeResetPasswordEndpoint creates a gokit endpoint which invokes ResetPassword
func MakeResetPasswordEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ResetPasswordReq := request.(ResetPasswordRequest)
		err := s.ResetPassword(ResetPasswordReq.Email)
		return ResetPasswordResponse{
			Error: err,
		}, nil
	}
}

// GetUserRequest is the request struct for the GetUserEndpoint
type GetUserRequest struct {
	ID uint
}

// GetUserResponse is the response struct for the GetUserEndpoint
type GetUserResponse struct {
	User  *User
	Error error
}

// MakeGetUserEndpoint creates a gokit endpoiint which invokes GetUser
func MakeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		getRequest := request.(GetUserRequest)
		user := &User{}
		err := s.GetUser(getRequest.ID, user)
		return GetUserResponse{
			User:  user,
			Error: err,
		}, nil
	}
}
