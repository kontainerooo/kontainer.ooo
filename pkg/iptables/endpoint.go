package iptables

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a struct which collects all endpoints for the iptables service
type Endpoints struct {
	CreateChainEndpoint   endpoint.Endpoint
	AddRuleEndpoint       endpoint.Endpoint
	RemoveRuleEndpoint    endpoint.Endpoint
	GetRulesByRefEndpoint endpoint.Endpoint
}

// CreateChainRequest is the request struct for the CreateChainEndpoint
type CreateChainRequest struct {
	Name string
}

// CreateChainResponse is the response struct for the CreateChainEndpoint
type CreateChainResponse struct {
	Error error
}

// MakeCreateChainEndpoint creates a gokit endpoint which invokes CreateChain
func MakeCreateChainEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateChainRequest)
		err := s.CreateChain(req.Name)
		return CreateChainResponse{
			Error: err,
		}, nil
	}
}

// AddRuleRequest is the request struct for the AddRuleEndpoint
type AddRuleRequest struct {
	Refid string
	Rule  Rule
}

// AddRuleResponse is the response struct for the AddRuleEndpoint
type AddRuleResponse struct {
	Error error
}

// MakeAddRuleEndpoint creates a gokit endpoint which invokes AddRule
func MakeAddRuleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRuleRequest)
		err := s.AddRule(req.Refid, req.Rule)
		return AddRuleResponse{
			Error: err,
		}, nil
	}
}

// RemoveRuleRequest is the request struct for the RemoveRuleEndpoint
type RemoveRuleRequest struct {
	ID string
}

// RemoveRuleResponse is the response struct for the RemoveRuleEndpoint
type RemoveRuleResponse struct {
	Error error
}

// MakeRemoveRuleEndpoint creates a gokit endpoint which invokes RemoveRule
func MakeRemoveRuleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveRuleRequest)
		err := s.RemoveRule(req.ID)
		return RemoveRuleResponse{
			Error: err,
		}, nil
	}
}

// GetRulesByRefRequest is the request struct for the GetRulesByRefEndpoint
type GetRulesByRefRequest struct {
	Refid string
}

// GetRulesByRefResponse is the response struct for the GetRulesByRefEndpoint
type GetRulesByRefResponse struct {
	Rules []Rule
	Error error
}

// MakeGetRulesByRefEndpoint creates a gokit endpoint which invokes GetRulesByRef
func MakeGetRulesByRefEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRulesByRefRequest)
		rules, err := s.GetRulesByRef(req.Refid)
		return GetRulesByRefResponse{
			Rules: rules,
			Error: err,
		}, nil
	}
}
