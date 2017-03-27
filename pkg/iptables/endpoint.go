package iptables

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a struct which collects all endpoints for the iptables service
type Endpoints struct {
	AddRuleEndpoint         endpoint.Endpoint
	RemoveRuleEndpoint      endpoint.Endpoint
	GetRulesForUserEndpoint endpoint.Endpoint
}

// AddRuleRequest is the request struct for the AddRuleEndpoint
type AddRuleRequest struct {
	Refid uint
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

// GetRulesForUserRequest is the request struct for the GetRulesForUserEndpoint
type GetRulesForUserRequest struct {
	Refid uint
}

// GetRulesForUserResponse is the response struct for the GetRulesForUserEndpoint
type GetRulesForUserResponse struct {
	Rules []Rule
	Error error
}

// MakeGetRulesForUserEndpoint creates a gokit endpoint which invokes GetRulesForUser
func MakeGetRulesForUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRulesForUserRequest)
		rules, err := s.GetRulesForUser(req.Refid)
		return GetRulesForUserResponse{
			Rules: rules,
			Error: err,
		}, nil
	}
}
