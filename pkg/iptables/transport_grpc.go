package iptables

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
	oldcontext "golang.org/x/net/context"
)

// MakeGRPCServer makes a set of Endpoints available as a gRPC iptablesServer
func MakeGRPCServer(ctx context.Context, endpoints Endpoints, logger log.Logger) pb.IPTablesServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		createchain: grpctransport.NewServer(
			endpoints.CreateChainEndpoint,
			DecodeGRPCCreateChainRequest,
			EncodeGRPCCreateChainResponse,
			options...,
		),
		addrule: grpctransport.NewServer(
			endpoints.AddRuleEndpoint,
			DecodeGRPCAddRuleRequest,
			EncodeGRPCAddRuleResponse,
			options...,
		),
		removerule: grpctransport.NewServer(
			endpoints.RemoveRuleEndpoint,
			DecodeGRPCRemoveRuleRequest,
			EncodeGRPCRemoveRuleResponse,
			options...,
		),
		getrulesforuser: grpctransport.NewServer(
			endpoints.GetRulesForUserEndpoint,
			DecodeGRPCGetRulesForUserRequest,
			EncodeGRPCGetRulesForUserResponse,
			options...,
		),
	}
}

type grpcServer struct {
	createchain     grpctransport.Handler
	addrule         grpctransport.Handler
	removerule      grpctransport.Handler
	getrulesforuser grpctransport.Handler
}

func (s *grpcServer) CreateChain(ctx oldcontext.Context, req *pb.CreateChainRequest) (*pb.CreateChainResponse, error) {
	_, res, err := s.createchain.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.CreateChainResponse), nil
}

func (s *grpcServer) AddRule(ctx oldcontext.Context, req *pb.AddRuleRequest) (*pb.AddRuleResponse, error) {
	_, res, err := s.addrule.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.AddRuleResponse), nil
}

func (s *grpcServer) RemoveRule(ctx oldcontext.Context, req *pb.RemoveRuleRequest) (*pb.RemoveRuleResponse, error) {
	_, res, err := s.removerule.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.RemoveRuleResponse), nil
}

func (s *grpcServer) GetRulesForUser(ctx oldcontext.Context, req *pb.GetRulesForUserRequest) (*pb.GetRulesForUserResponse, error) {
	_, res, err := s.getrulesforuser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetRulesForUserResponse), nil
}

func convertToNativeRule(r pb.IPTRule) (Rule, error) {
	sourceIP, err := abstraction.NewInet(r.Source)
	if err != nil {
		return Rule{}, err
	}

	destIP, err := abstraction.NewInet(r.Destination)
	if err != nil {
		return Rule{}, err
	}

	return Rule{
		Operation:       r.Operation,
		Target:          r.Target,
		Chain:           r.Chain,
		Protocol:        r.Protocol,
		In:              r.In,
		Out:             r.Out,
		Source:          sourceIP,
		Destination:     destIP,
		SourcePort:      uint16(r.SourcePort),
		DestinationPort: uint16(r.DestinationPort),
		State:           r.State,
	}, nil
}

// DecodeGRPCCreateChainRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateChain request to a messages/iptables.proto-domain createchain request.
func DecodeGRPCCreateChainRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateChainRequest)
	return CreateChainRequest{
		Name: req.Name,
	}, nil
}

// DecodeGRPCAddRuleRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC AddRule request to a messages/iptables.proto-domain addrule request.
func DecodeGRPCAddRuleRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddRuleRequest)
	rule, err := convertToNativeRule(*req.Rule)
	if err != nil {
		return AddRuleRequest{}, err
	}

	return AddRuleRequest{
		Refid: uint(req.Refid),
		Rule:  rule,
	}, nil
}

// DecodeGRPCRemoveRuleRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemoveRule request to a messages/iptables.proto-domain removerule request.
func DecodeGRPCRemoveRuleRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RemoveRuleRequest)
	return RemoveRuleRequest{
		ID: req.ID,
	}, nil
}

// DecodeGRPCGetRulesForUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetRulesForUser request to a messages/iptables.proto-domain getrulesforuser request.
func DecodeGRPCGetRulesForUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetRulesForUserRequest)
	return GetRulesForUserRequest{
		Refid: uint(req.Refid),
	}, nil
}

// EncodeGRPCCreateChainResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/iptables.proto-domain createchain response to a gRPC CreateChain response.
func EncodeGRPCCreateChainResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(CreateChainResponse)
	gRPCRes := &pb.CreateChainResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCAddRuleResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/iptables.proto-domain addrule response to a gRPC AddRule response.
func EncodeGRPCAddRuleResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(AddRuleResponse)
	gRPCRes := &pb.AddRuleResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemoveRuleResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/iptables.proto-domain removerule response to a gRPC RemoveRule response.
func EncodeGRPCRemoveRuleResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveRuleResponse)
	gRPCRes := &pb.RemoveRuleResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCGetRulesForUserResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/iptables.proto-domain getrulesforuser response to a gRPC GetRulesForUser response.
func EncodeGRPCGetRulesForUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetRulesForUserResponse)
	rules := []*pb.IPTRule{}

	gRPCRes := &pb.GetRulesForUserResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}

	for _, v := range res.Rules {
		rule := pb.IPTRule{
			Operation:       v.Operation,
			Target:          v.Target,
			Chain:           v.Chain,
			Protocol:        v.Protocol,
			In:              v.In,
			Out:             v.Out,
			Source:          string(v.Source),
			Destination:     string(v.Destination),
			SourcePort:      uint32(v.SourcePort),
			DestinationPort: uint32(v.DestinationPort),
			State:           v.State,
		}
		rules = append(rules, &rule)
	}
	gRPCRes.Rules = rules

	return gRPCRes, nil
}
