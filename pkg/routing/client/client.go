package client

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
)

// New creates a set of endpoints based on a gRPC connection
func New(conn *grpc.ClientConn, logger log.Logger) *routing.Endpoints {

	var CreateConfigEndpoint endpoint.Endpoint
	{
		CreateConfigEndpoint = grpctransport.NewClient(
			conn,
			"routingService",
			"CreateConfig",
			EncodeGRPCCreateConfigRequest,
			DecodeGRPCCreateConfigResponse,
			pb.CreateConfigResponse{},
		).Endpoint()
	}

	var EditConfigEndpoint endpoint.Endpoint
	{
		EditConfigEndpoint = grpctransport.NewClient(
			conn,
			"routingService",
			"EditConfig",
			EncodeGRPCEditConfigRequest,
			DecodeGRPCEditConfigResponse,
			pb.EditConfigResponse{},
		).Endpoint()
	}

	var GetConfigEndpoint endpoint.Endpoint
	{
		GetConfigEndpoint = grpctransport.NewClient(
			conn,
			"routingService",
			"GetConfig",
			EncodeGRPCGetConfigRequest,
			DecodeGRPCGetConfigResponse,
			pb.GetConfigResponse{},
		).Endpoint()
	}

	var RemoveConfigEndpoint endpoint.Endpoint
	{
		RemoveConfigEndpoint = grpctransport.NewClient(
			conn,
			"routingService",
			"RemoveConfig",
			EncodeGRPCRemoveConfigRequest,
			DecodeGRPCRemoveConfigResponse,
			pb.RemoveConfigResponse{},
		).Endpoint()
	}

	var AddLocationEndpoint endpoint.Endpoint
	{
		AddLocationEndpoint = grpctransport.NewClient(
			conn,
			"routingService",
			"AddLocation",
			EncodeGRPCAddLocationRequest,
			DecodeGRPCAddLocationResponse,
			pb.AddLocationResponse{},
		).Endpoint()
	}

	var RemoveLocationEndpoint endpoint.Endpoint
	{
		RemoveLocationEndpoint = grpctransport.NewClient(
			conn,
			"routingService",
			"RemoveLocation",
			EncodeGRPCRemoveLocationRequest,
			DecodeGRPCRemoveLocationResponse,
			pb.RemoveLocationResponse{},
		).Endpoint()
	}

	var ChangeListenStatementEndpoint endpoint.Endpoint
	{
		ChangeListenStatementEndpoint = grpctransport.NewClient(
			conn,
			"routingService",
			"ChangeListenStatement",
			EncodeGRPCChangeListenStatementRequest,
			DecodeGRPCChangeListenStatementResponse,
			pb.ChangeListenStatementResponse{},
		).Endpoint()
	}

	var AddServerNameEndpoint endpoint.Endpoint
	{
		AddServerNameEndpoint = grpctransport.NewClient(
			conn,
			"routingService",
			"AddServerName",
			EncodeGRPCAddServerNameRequest,
			DecodeGRPCAddServerNameResponse,
			pb.AddServerNameResponse{},
		).Endpoint()
	}

	var RemoveServerNameEndpoint endpoint.Endpoint
	{
		RemoveServerNameEndpoint = grpctransport.NewClient(
			conn,
			"routingService",
			"RemoveServerName",
			EncodeGRPCRemoveServerNameRequest,
			DecodeGRPCRemoveServerNameResponse,
			pb.RemoveServerNameResponse{},
		).Endpoint()
	}

	var ConfigurationsEndpoint endpoint.Endpoint
	{
		ConfigurationsEndpoint = grpctransport.NewClient(
			conn,
			"routingService",
			"Configurations",
			EncodeGRPCConfigurationsRequest,
			DecodeGRPCConfigurationsResponse,
			pb.ConfigurationsResponse{},
		).Endpoint()
	}

	return &routing.Endpoints{
		CreateConfigEndpoint:          CreateConfigEndpoint,
		EditConfigEndpoint:            EditConfigEndpoint,
		GetConfigEndpoint:             GetConfigEndpoint,
		RemoveConfigEndpoint:          RemoveConfigEndpoint,
		AddLocationEndpoint:           AddLocationEndpoint,
		RemoveLocationEndpoint:        RemoveLocationEndpoint,
		ChangeListenStatementEndpoint: ChangeListenStatementEndpoint,
		AddServerNameEndpoint:         AddServerNameEndpoint,
		RemoveServerNameEndpoint:      RemoveServerNameEndpoint,
		ConfigurationsEndpoint:        ConfigurationsEndpoint,
	}
}

func getError(e string) error {
	if e != "" {
		return errors.New(e)
	}
	return nil
}

func convertPBConfigs(r []*pb.RouterConfig) *[]routing.RouterConfig {
	p := make([]routing.RouterConfig, len(r))
	for i, c := range r {
		p[i] = *routing.ConvertPBConfig(c)
	}
	return &p
}

// EncodeGRPCCreateConfigRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain createconfig request to a gRPC CreateConfig request.
func EncodeGRPCCreateConfigRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*routing.CreateConfigRequest)
	return &pb.CreateConfigRequest{
		Config: routing.ConvertConfiguration(*req.Config),
	}, nil
}

// DecodeGRPCCreateConfigResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC CreateConfig response to a messages/routing.proto-domain createconfig response.
func DecodeGRPCCreateConfigResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.CreateConfigResponse)
	return &routing.CreateConfigResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCEditConfigRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain editconfig request to a gRPC EditConfig request.
func EncodeGRPCEditConfigRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*routing.EditConfigRequest)
	return &pb.EditConfigRequest{
		RefID:  uint32(req.RefID),
		Name:   req.Name,
		Config: routing.ConvertConfiguration(*req.Config),
	}, nil
}

// DecodeGRPCEditConfigResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC EditConfig response to a messages/routing.proto-domain editconfig response.
func DecodeGRPCEditConfigResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.EditConfigResponse)
	return &routing.EditConfigResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCGetConfigRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain getconfig request to a gRPC GetConfig request.
func EncodeGRPCGetConfigRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*routing.GetConfigRequest)
	return &pb.GetConfigRequest{
		RefID: uint32(req.RefID),
		Name:  req.Name,
	}, nil
}

// DecodeGRPCGetConfigResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetConfig response to a messages/routing.proto-domain getconfig response.
func DecodeGRPCGetConfigResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.GetConfigResponse)
	return &routing.GetConfigResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveConfigRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain removeconfig request to a gRPC RemoveConfig request.
func EncodeGRPCRemoveConfigRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*routing.RemoveConfigRequest)
	return &pb.RemoveConfigRequest{
		RefID: uint32(req.RefID),
		Name:  req.Name,
	}, nil
}

// DecodeGRPCRemoveConfigResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveConfig response to a messages/routing.proto-domain removeconfig response.
func DecodeGRPCRemoveConfigResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.RemoveConfigResponse)
	return &routing.RemoveConfigResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCAddLocationRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain addlocation request to a gRPC AddLocation request.
func EncodeGRPCAddLocationRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*routing.AddLocationRequest)
	return &pb.AddLocationRequest{
		RefID:    uint32(req.RefID),
		Name:     req.Name,
		Location: routing.ConvertLocation(req.Location),
	}, nil
}

// DecodeGRPCAddLocationResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC AddLocation response to a messages/routing.proto-domain addlocation response.
func DecodeGRPCAddLocationResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.AddLocationResponse)
	return &routing.AddLocationResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveLocationRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain removelocation request to a gRPC RemoveLocation request.
func EncodeGRPCRemoveLocationRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*routing.RemoveLocationRequest)
	return &pb.RemoveLocationRequest{
		RefID: uint32(req.RefID),
		Name:  req.Name,
		Id:    int32(req.LID),
	}, nil
}

// DecodeGRPCRemoveLocationResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveLocation response to a messages/routing.proto-domain removelocation response.
func DecodeGRPCRemoveLocationResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.RemoveLocationResponse)
	return &routing.RemoveLocationResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCChangeListenStatementRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain changelistenstatement request to a gRPC ChangeListenStatement request.
func EncodeGRPCChangeListenStatementRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*routing.ChangeListenStatementRequest)
	return &pb.ChangeListenStatementRequest{
		RefID:           uint32(req.RefID),
		Name:            req.Name,
		ListenStatement: routing.ConvertListenStatement(req.ListenStatement),
	}, nil
}

// DecodeGRPCChangeListenStatementResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC ChangeListenStatement response to a messages/routing.proto-domain changelistenstatement response.
func DecodeGRPCChangeListenStatementResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.ChangeListenStatementResponse)
	return &routing.ChangeListenStatementResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCAddServerNameRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain addservername request to a gRPC AddServerName request.
func EncodeGRPCAddServerNameRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*routing.AddServerNameRequest)
	return &pb.AddServerNameRequest{
		RefID:      uint32(req.RefID),
		Name:       req.Name,
		ServerName: req.ServerName,
	}, nil
}

// DecodeGRPCAddServerNameResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC AddServerName response to a messages/routing.proto-domain addservername response.
func DecodeGRPCAddServerNameResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.AddServerNameResponse)
	return &routing.AddServerNameResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveServerNameRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain removeservername request to a gRPC RemoveServerName request.
func EncodeGRPCRemoveServerNameRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*routing.RemoveServerNameRequest)
	return &pb.RemoveServerNameRequest{
		RefID: uint32(req.RefID),
		Name:  req.Name,
		Id:    int32(req.ID),
	}, nil
}

// DecodeGRPCRemoveServerNameResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveServerName response to a messages/routing.proto-domain removeservername response.
func DecodeGRPCRemoveServerNameResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.RemoveServerNameResponse)
	return &routing.RemoveServerNameResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCConfigurationsRequest is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain configurations request to a gRPC Configurations request.
func EncodeGRPCConfigurationsRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &pb.ConfigurationsRequest{}, nil
}

// DecodeGRPCConfigurationsResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC Configurations response to a messages/routing.proto-domain configurations response.
func DecodeGRPCConfigurationsResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.ConfigurationsResponse)
	return &routing.ConfigurationsResponse{
		Configurations: convertPBConfigs(response.Configurations),
	}, nil
}
