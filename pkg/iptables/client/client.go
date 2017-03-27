package client

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/iptables"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
)

// New creates a set of endpoints based on a gRPC connection
func New(conn *grpc.ClientConn, logger log.Logger) *iptables.Endpoints {

	var CreateChainEndpoint endpoint.Endpoint
	{
		CreateChainEndpoint = grpctransport.NewClient(
			conn,
			"IPTablesService",
			"CreateChain",
			EncodeGRPCCreateChainRequest,
			DecodeGRPCCreateChainResponse,
			pb.CreateChainResponse{},
		).Endpoint()
	}

	var AddRuleEndpoint endpoint.Endpoint
	{
		AddRuleEndpoint = grpctransport.NewClient(
			conn,
			"IPTablesService",
			"AddRule",
			EncodeGRPCAddRuleRequest,
			DecodeGRPCAddRuleResponse,
			pb.AddRuleResponse{},
		).Endpoint()
	}

	var RemoveRuleEndpoint endpoint.Endpoint
	{
		RemoveRuleEndpoint = grpctransport.NewClient(
			conn,
			"IPTablesService",
			"RemoveRule",
			EncodeGRPCRemoveRuleRequest,
			DecodeGRPCRemoveRuleResponse,
			pb.RemoveRuleResponse{},
		).Endpoint()
	}

	var GetRulesByRefEndpoint endpoint.Endpoint
	{
		GetRulesByRefEndpoint = grpctransport.NewClient(
			conn,
			"IPTablesService",
			"GetRulesByRef",
			EncodeGRPCGetRulesByRefRequest,
			DecodeGRPCGetRulesByRefResponse,
			pb.GetRulesByRefResponse{},
		).Endpoint()
	}

	return &iptables.Endpoints{
		CreateChainEndpoint:   CreateChainEndpoint,
		AddRuleEndpoint:       AddRuleEndpoint,
		RemoveRuleEndpoint:    RemoveRuleEndpoint,
		GetRulesByRefEndpoint: GetRulesByRefEndpoint,
	}
}

func getError(e string) error {
	if e != "" {
		return errors.New(e)
	}
	return nil
}

func convertToIPTRule(r iptables.Rule) *pb.IPTRule {
	return &pb.IPTRule{
		Operation:       r.Operation,
		Target:          r.Target,
		Chain:           r.Chain,
		Protocol:        r.Protocol,
		In:              r.In,
		Out:             r.Out,
		Source:          string(r.Source),
		Destination:     string(r.Destination),
		SourcePort:      uint32(r.SourcePort),
		DestinationPort: uint32(r.DestinationPort),
		State:           r.State,
	}
}

func convertToNativeRule(r *pb.IPTRule) (iptables.Rule, error) {
	sourceIP, err := abstraction.NewInet(r.Source)
	if err != nil {
		return iptables.Rule{}, err
	}

	destIP, err := abstraction.NewInet(r.Destination)
	if err != nil {
		return iptables.Rule{}, err
	}

	return iptables.Rule{
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

// EncodeGRPCCreateChainRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/iptables.proto-domain createchain request to a gRPC CreateChain request.
func EncodeGRPCCreateChainRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*iptables.CreateChainRequest)
	return &pb.CreateChainRequest{
		Name: req.Name,
	}, nil
}

// DecodeGRPCCreateChainResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC CreateChain response to a messages/iptables.proto-domain createchain response.
func DecodeGRPCCreateChainResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.CreateChainResponse)
	return &iptables.CreateChainResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCAddRuleRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/iptables.proto-domain addrule request to a gRPC AddRule request.
func EncodeGRPCAddRuleRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*iptables.AddRuleRequest)
	return &pb.AddRuleRequest{
		Refid: req.Refid,
		Rule:  convertToIPTRule(req.Rule),
	}, nil
}

// DecodeGRPCAddRuleResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC AddRule response to a messages/iptables.proto-domain addrule response.
func DecodeGRPCAddRuleResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.AddRuleResponse)
	return &iptables.AddRuleResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveRuleRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/iptables.proto-domain removerule request to a gRPC RemoveRule request.
func EncodeGRPCRemoveRuleRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*iptables.RemoveRuleRequest)
	return &pb.RemoveRuleRequest{
		ID: req.ID,
	}, nil
}

// DecodeGRPCRemoveRuleResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveRule response to a messages/iptables.proto-domain removerule response.
func DecodeGRPCRemoveRuleResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.RemoveRuleResponse)
	return &iptables.RemoveRuleResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCGetRulesByRefRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/iptables.proto-domain getrulesbyid request to a gRPC GetRulesByRef request.
func EncodeGRPCGetRulesByRefRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*iptables.GetRulesByRefRequest)
	return &pb.GetRulesByRefRequest{
		Refid: req.Refid,
	}, nil
}

// DecodeGRPCGetRulesByRefResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetRulesByRef response to a messages/iptables.proto-domain getrulesbyref response.
func DecodeGRPCGetRulesByRefResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.GetRulesByRefResponse)

	rules := []iptables.Rule{}
	for _, v := range response.Rules {
		rule, err := convertToNativeRule(v)
		if err != nil {
			return &iptables.GetRulesByRefResponse{
				Error: getError(response.Error),
			}, err
		}
		rules = append(rules, rule)
	}

	return &iptables.GetRulesByRefResponse{
		Error: getError(response.Error),
		Rules: rules,
	}, nil
}
