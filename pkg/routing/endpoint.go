package routing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a struct which collects all endpoints for the routing service
type Endpoints struct {
	CreateConfigEndpoint          endpoint.Endpoint
	EditConfigEndpoint            endpoint.Endpoint
	GetConfigEndpoint             endpoint.Endpoint
	RemoveConfigEndpoint          endpoint.Endpoint
	AddLocationEndpoint           endpoint.Endpoint
	RemoveLocationEndpoint        endpoint.Endpoint
	ChangeListenStatementEndpoint endpoint.Endpoint
	AddServerNameEndpoint         endpoint.Endpoint
	RemoveServerNameEndpoint      endpoint.Endpoint
	ConfigurationsEndpoint        endpoint.Endpoint
}

// IDRequest combines a RefID and a name
type IDRequest struct {
	RefID uint
	Name  string
}

// CreateConfigRequest is the request struct for the CreateConfigEndpoint
type CreateConfigRequest struct {
	Config *RouterConfig
}

// CreateConfigResponse is the response struct for the CreateConfigEndpoint
type CreateConfigResponse struct {
	Error error
}

// MakeCreateConfigEndpoint creates a gokit endpoint which invokes CreateConfig
func MakeCreateConfigEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateConfigRequest)
		err := s.CreateRouterConfig(req.Config)
		return CreateConfigResponse{err}, nil
	}
}

// EditConfigRequest is the request struct for the EditConfigEndpoint
type EditConfigRequest struct {
	IDRequest
	Config *RouterConfig
}

// EditConfigResponse is the response struct for the EditConfigEndpoint
type EditConfigResponse struct {
	Error error
}

// MakeEditConfigEndpoint creates a gokit endpoint which invokes EditConfig
func MakeEditConfigEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EditConfigRequest)
		err := s.EditRouterConfig(req.RefID, req.Name, req.Config)
		return EditConfigResponse{err}, nil
	}
}

// GetConfigRequest is the request struct for the GetConfigEndpoint
type GetConfigRequest struct {
	IDRequest
}

// GetConfigResponse is the response struct for the GetConfigEndpoint
type GetConfigResponse struct {
	Config RouterConfig
	Error  error
}

// MakeGetConfigEndpoint creates a gokit endpoint which invokes GetConfig
func MakeGetConfigEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetConfigRequest)
		conf := RouterConfig{}
		err := s.GetRouterConfig(req.RefID, req.Name, &conf)
		return GetConfigResponse{
			Config: conf,
			Error:  err,
		}, nil
	}
}

// RemoveConfigRequest is the request struct for the RemoveConfigEndpoint
type RemoveConfigRequest struct {
	IDRequest
}

// RemoveConfigResponse is the response struct for the RemoveConfigEndpoint
type RemoveConfigResponse struct {
	Error error
}

// MakeRemoveConfigEndpoint creates a gokit endpoint which invokes RemoveConfig
func MakeRemoveConfigEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveConfigRequest)
		err := s.RemoveRouterConfig(req.RefID, req.Name)
		return RemoveConfigResponse{err}, nil
	}
}

// AddLocationRequest is the request struct for the AddLocationEndpoint
type AddLocationRequest struct {
	IDRequest
	Location *LocationRule
}

// AddLocationResponse is the response struct for the AddLocationEndpoint
type AddLocationResponse struct {
	Error error
}

// MakeAddLocationEndpoint creates a gokit endpoint which invokes AddLocation
func MakeAddLocationEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddLocationRequest)
		err := s.AddLocationRule(req.RefID, req.Name, req.Location)
		return AddLocationResponse{err}, nil
	}
}

// RemoveLocationRequest is the request struct for the RemoveLocationEndpoint
type RemoveLocationRequest struct {
	IDRequest
	LID int
}

// RemoveLocationResponse is the response struct for the RemoveLocationEndpoint
type RemoveLocationResponse struct {
	Error error
}

// MakeRemoveLocationEndpoint creates a gokit endpoint which invokes RemoveLocation
func MakeRemoveLocationEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveLocationRequest)
		err := s.RemoveLocationRule(req.RefID, req.Name, req.LID)
		return RemoveLocationResponse{err}, nil
	}
}

// ChangeListenStatementRequest is the request struct for the ChangeListenStatementEndpoint
type ChangeListenStatementRequest struct {
	IDRequest
	ListenStatement *ListenStatement
}

// ChangeListenStatementResponse is the response struct for the ChangeListenStatementEndpoint
type ChangeListenStatementResponse struct {
	Error error
}

// MakeChangeListenStatementEndpoint creates a gokit endpoint which invokes ChangeListenStatement
func MakeChangeListenStatementEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeListenStatementRequest)
		err := s.ChangeListenStatement(req.RefID, req.Name, req.ListenStatement)
		return ChangeListenStatementResponse{err}, nil
	}
}

// AddServerNameRequest is the request struct for the AddServerNameEndpoint
type AddServerNameRequest struct {
	IDRequest
	ServerName string
}

// AddServerNameResponse is the response struct for the AddServerNameEndpoint
type AddServerNameResponse struct {
	Error error
}

// MakeAddServerNameEndpoint creates a gokit endpoint which invokes AddServerName
func MakeAddServerNameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddServerNameRequest)
		err := s.AddServerName(req.RefID, req.Name, req.ServerName)
		return AddServerNameResponse{err}, nil
	}
}

// RemoveServerNameRequest is the request struct for the RemoveServerNameEndpoint
type RemoveServerNameRequest struct {
	IDRequest
	ID int
}

// RemoveServerNameResponse is the response struct for the RemoveServerNameEndpoint
type RemoveServerNameResponse struct {
	Error error
}

// MakeRemoveServerNameEndpoint creates a gokit endpoint which invokes RemoveServerName
func MakeRemoveServerNameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveServerNameRequest)
		err := s.RemoveServerName(req.RefID, req.Name, req.ID)
		return RemoveServerNameResponse{err}, nil
	}
}

// ConfigurationsRequest is the request struct for the ConfigurationsEndpoint
type ConfigurationsRequest struct{}

// ConfigurationsResponse is the response struct for the ConfigurationsEndpoint
type ConfigurationsResponse struct {
	Configurations *[]RouterConfig
}

// MakeConfigurationsEndpoint creates a gokit endpoint which invokes Configurations
func MakeConfigurationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		configs := make([]RouterConfig, 0)
		s.Configurations(&configs)
		return ConfigurationsResponse{
			Configurations: &configs,
		}, nil
	}
}
