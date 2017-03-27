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
		getrulesbyref: grpctransport.NewServer(
			endpoints.GetRulesByRefEndpoint,
			DecodeGRPCGetRulesByRefRequest,
			EncodeGRPCGetRulesByRefResponse,
			options...,
		),
	}
}

type grpcServer struct {
	addrule       grpctransport.Handler
	removerule    grpctransport.Handler
	getrulesbyref grpctransport.Handler
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

func (s *grpcServer) GetRulesByRef(ctx oldcontext.Context, req *pb.GetRulesByRefRequest) (*pb.GetRulesByRefResponse, error) {
	_, res, err := s.getrulesbyref.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetRulesByRefResponse), nil
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

// DecodeGRPCAddRuleRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC AddRule request to a messages/iptables.proto-domain addrule request.
func DecodeGRPCAddRuleRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddRuleRequest)
	rule, err := convertToNativeRule(*req.Rule)
	if err != nil {
		return AddRuleRequest{}, err
	}

	return AddRuleRequest{
		Refid: req.Refid,
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

// DecodeGRPCGetRulesByRefRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetRulesByRef request to a messages/iptables.proto-domain GetRulesByRef request.
func DecodeGRPCGetRulesByRefRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetRulesByRefRequest)
	return GetRulesByRefRequest{
		Refid: req.Refid,
	}, nil
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

// EncodeGRPCGetRulesByRefResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/iptables.proto-domain GetRulesByRef response to a gRPC GetRulesByRef response.
func EncodeGRPCGetRulesByRefResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetRulesByRefResponse)
	rules := []*pb.IPTRule{}

	gRPCRes := &pb.GetRulesByRefResponse{}
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
