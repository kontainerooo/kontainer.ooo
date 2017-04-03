package routing

import (
	"context"
	"strings"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
	oldcontext "golang.org/x/net/context"
)

// MakeGRPCServer makes a set of Endpoints available as a gRPC routingServer
func MakeGRPCServer(ctx context.Context, endpoints Endpoints, logger log.Logger) pb.RoutingServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{

		createConfig: grpctransport.NewServer(
			endpoints.CreateConfigEndpoint,
			DecodeGRPCCreateConfigRequest,
			EncodeGRPCCreateConfigResponse,
			options...,
		),

		editConfig: grpctransport.NewServer(
			endpoints.EditConfigEndpoint,
			DecodeGRPCEditConfigRequest,
			EncodeGRPCEditConfigResponse,
			options...,
		),

		getConfig: grpctransport.NewServer(
			endpoints.GetConfigEndpoint,
			DecodeGRPCGetConfigRequest,
			EncodeGRPCGetConfigResponse,
			options...,
		),

		removeConfig: grpctransport.NewServer(
			endpoints.RemoveConfigEndpoint,
			DecodeGRPCRemoveConfigRequest,
			EncodeGRPCRemoveConfigResponse,
			options...,
		),

		addLocation: grpctransport.NewServer(
			endpoints.AddLocationEndpoint,
			DecodeGRPCAddLocationRequest,
			EncodeGRPCAddLocationResponse,
			options...,
		),

		removeLocation: grpctransport.NewServer(
			endpoints.RemoveLocationEndpoint,
			DecodeGRPCRemoveLocationRequest,
			EncodeGRPCRemoveLocationResponse,
			options...,
		),

		changeListenStatement: grpctransport.NewServer(
			endpoints.ChangeListenStatementEndpoint,
			DecodeGRPCChangeListenStatementRequest,
			EncodeGRPCChangeListenStatementResponse,
			options...,
		),

		addServerName: grpctransport.NewServer(
			endpoints.AddServerNameEndpoint,
			DecodeGRPCAddServerNameRequest,
			EncodeGRPCAddServerNameResponse,
			options...,
		),

		removeServerName: grpctransport.NewServer(
			endpoints.RemoveServerNameEndpoint,
			DecodeGRPCRemoveServerNameRequest,
			EncodeGRPCRemoveServerNameResponse,
			options...,
		),

		configurations: grpctransport.NewServer(
			endpoints.ConfigurationsEndpoint,
			DecodeGRPCConfigurationsRequest,
			EncodeGRPCConfigurationsResponse,
			options...,
		),
	}
}

type grpcServer struct {
	createConfig          grpctransport.Handler
	editConfig            grpctransport.Handler
	getConfig             grpctransport.Handler
	removeConfig          grpctransport.Handler
	addLocation           grpctransport.Handler
	removeLocation        grpctransport.Handler
	changeListenStatement grpctransport.Handler
	addServerName         grpctransport.Handler
	removeServerName      grpctransport.Handler
	configurations        grpctransport.Handler
}

func (s *grpcServer) CreateConfig(ctx oldcontext.Context, req *pb.CreateConfigRequest) (*pb.CreateConfigResponse, error) {
	_, res, err := s.createConfig.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.CreateConfigResponse), nil
}

func (s *grpcServer) EditConfig(ctx oldcontext.Context, req *pb.EditConfigRequest) (*pb.EditConfigResponse, error) {
	_, res, err := s.editConfig.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.EditConfigResponse), nil
}

func (s *grpcServer) GetConfig(ctx oldcontext.Context, req *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	_, res, err := s.getConfig.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetConfigResponse), nil
}

func (s *grpcServer) RemoveConfig(ctx oldcontext.Context, req *pb.RemoveConfigRequest) (*pb.RemoveConfigResponse, error) {
	_, res, err := s.removeConfig.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.RemoveConfigResponse), nil
}

func (s *grpcServer) AddLocation(ctx oldcontext.Context, req *pb.AddLocationRequest) (*pb.AddLocationResponse, error) {
	_, res, err := s.addLocation.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.AddLocationResponse), nil
}

func (s *grpcServer) RemoveLocation(ctx oldcontext.Context, req *pb.RemoveLocationRequest) (*pb.RemoveLocationResponse, error) {
	_, res, err := s.removeLocation.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.RemoveLocationResponse), nil
}

func (s *grpcServer) ChangeListenStatement(ctx oldcontext.Context, req *pb.ChangeListenStatementRequest) (*pb.ChangeListenStatementResponse, error) {
	_, res, err := s.changeListenStatement.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ChangeListenStatementResponse), nil
}

func (s *grpcServer) AddServerName(ctx oldcontext.Context, req *pb.AddServerNameRequest) (*pb.AddServerNameResponse, error) {
	_, res, err := s.addServerName.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.AddServerNameResponse), nil
}

func (s *grpcServer) RemoveServerName(ctx oldcontext.Context, req *pb.RemoveServerNameRequest) (*pb.RemoveServerNameResponse, error) {
	_, res, err := s.removeServerName.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.RemoveServerNameResponse), nil
}

func (s *grpcServer) Configurations(ctx oldcontext.Context, req *pb.ConfigurationsRequest) (*pb.ConfigurationsResponse, error) {
	_, res, err := s.configurations.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ConfigurationsResponse), nil
}

func convertPBRules(r map[string]string) map[string][]string {
	m := make(map[string][]string)
	for k, v := range r {
		m[k] = strings.Split(v, " ")
	}
	return m
}

func convertPBLocation(l *pb.Location) *LocationRule {
	return &LocationRule{
		Location: l.Location,
		Rules:    convertPBRules(l.Rules),
	}
}

func convertPBLocations(l []*pb.Location) LocationRules {
	lr := make(LocationRules, len(l))
	for i, r := range l {
		lr[i] = convertPBLocation(r)
	}
	return lr
}

func convertPBListenStatement(l *pb.ListenStatement) *ListenStatement {
	ip, _ := abstraction.NewInet(l.IPAddress)
	return &ListenStatement{
		IPAddress: ip,
		Keyword:   l.Keyword,
		Port:      uint16(l.Port),
	}
}

func convertPBLog(l *pb.Log) Log {
	return Log{
		Path:    l.Path,
		Keyword: l.Keyword,
	}
}

func convertPBSSLSettings(s *pb.SSLSettings) SSLSettings {
	return SSLSettings{
		Protocols:           s.Protocols,
		Ciphers:             s.Ciphers,
		PreferServerCiphers: s.PreferServerCiphers,
		Certificate:         s.Certificate,
		CertificateKey:      s.CertificateKey,
		Curve:               s.Curve,
	}
}

// ConvertPBConfig convert *pb.RouterConfig to *RouterConfig
func ConvertPBConfig(c *pb.RouterConfig) *RouterConfig {
	return &RouterConfig{
		RefID:           uint(c.RefID),
		Name:            c.Name,
		ListenStatement: convertPBListenStatement(c.ListenStatement),
		ServerName:      c.ServerName,
		AccessLog:       convertPBLog(c.AccessLog),
		ErrorLog:        convertPBLog(c.ErrorLog),
		RootPath:        c.RootPath,
		SSLSettings:     convertPBSSLSettings(c.SSLSettings),
		LocationRules:   convertPBLocations(c.LocationRules),
	}
}

func convertRules(r map[string][]string) map[string]string {
	m := make(map[string]string)
	for k, v := range r {
		m[k] = strings.Join(v, " ")
	}
	return m
}

// ConvertLocation convert *Location to *pb.Location
func ConvertLocation(l *LocationRule) *pb.Location {
	return &pb.Location{
		Location: l.Location,
		Rules:    convertRules(l.Rules),
	}
}

func convertLocations(l LocationRules) []*pb.Location {
	lr := make([]*pb.Location, len(l))
	for i, r := range l {
		lr[i] = ConvertLocation(r)
	}
	return lr
}

// ConvertListenStatement convert *ListenStatement to *pb.ListenStatement
func ConvertListenStatement(l *ListenStatement) *pb.ListenStatement {
	return &pb.ListenStatement{
		IPAddress: string(l.IPAddress),
		Keyword:   l.Keyword,
		Port:      uint32(l.Port),
	}
}

func convertLog(l Log) *pb.Log {
	return &pb.Log{
		Path:    l.Path,
		Keyword: l.Keyword,
	}
}

func convertSSLSettings(s SSLSettings) *pb.SSLSettings {
	return &pb.SSLSettings{
		Protocols:           s.Protocols,
		Ciphers:             s.Ciphers,
		PreferServerCiphers: s.PreferServerCiphers,
		Certificate:         s.Certificate,
		CertificateKey:      s.CertificateKey,
		Curve:               s.Curve,
	}
}

// ConvertConfiguration convert routing domain RouterConfig to *pb.RouterConfig
func ConvertConfiguration(c RouterConfig) *pb.RouterConfig {
	return &pb.RouterConfig{
		RefID:           uint32(c.RefID),
		Name:            c.Name,
		ListenStatement: ConvertListenStatement(c.ListenStatement),
		ServerName:      c.ServerName,
		AccessLog:       convertLog(c.AccessLog),
		ErrorLog:        convertLog(c.ErrorLog),
		RootPath:        c.RootPath,
		SSLSettings:     convertSSLSettings(c.SSLSettings),
		LocationRules:   convertLocations(c.LocationRules),
	}
}

// ConvertConfigurations convert routing domain *[]RouterConfig to []*pb.RouterConfig
func ConvertConfigurations(c *[]RouterConfig) []*pb.RouterConfig {
	p := make([]*pb.RouterConfig, len(*c))
	for i, r := range *c {
		p[i] = ConvertConfiguration(r)
	}
	return p
}

// DecodeGRPCCreateConfigRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateConfig request to a messages/routing.proto-domain createconfig request.
func DecodeGRPCCreateConfigRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateConfigRequest)
	return CreateConfigRequest{
		Config: ConvertPBConfig(req.Config),
	}, nil
}

// DecodeGRPCEditConfigRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC EditConfig request to a messages/routing.proto-domain editconfig request.
func DecodeGRPCEditConfigRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.EditConfigRequest)
	return EditConfigRequest{
		IDRequest: IDRequest{
			RefID: uint(req.RefID),
			Name:  req.Name,
		},
		Config: ConvertPBConfig(req.Config),
	}, nil
}

// DecodeGRPCGetConfigRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetConfig request to a messages/routing.proto-domain getconfig request.
func DecodeGRPCGetConfigRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetConfigRequest)
	return GetConfigRequest{
		IDRequest: IDRequest{
			RefID: uint(req.RefID),
			Name:  req.Name,
		},
	}, nil
}

// DecodeGRPCRemoveConfigRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemoveConfig request to a messages/routing.proto-domain removeconfig request.
func DecodeGRPCRemoveConfigRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RemoveConfigRequest)
	return RemoveConfigRequest{
		IDRequest: IDRequest{
			RefID: uint(req.RefID),
			Name:  req.Name,
		},
	}, nil
}

// DecodeGRPCAddLocationRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC AddLocation request to a messages/routing.proto-domain addlocation request.
func DecodeGRPCAddLocationRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddLocationRequest)
	return AddLocationRequest{
		IDRequest: IDRequest{
			RefID: uint(req.RefID),
			Name:  req.Name,
		},
		Location: convertPBLocation(req.Location),
	}, nil
}

// DecodeGRPCRemoveLocationRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemoveLocation request to a messages/routing.proto-domain removelocation request.
func DecodeGRPCRemoveLocationRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RemoveLocationRequest)
	return RemoveLocationRequest{
		IDRequest: IDRequest{
			RefID: uint(req.RefID),
			Name:  req.Name,
		},
		LID: int(req.Id),
	}, nil
}

// DecodeGRPCChangeListenStatementRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC ChangeListenStatement request to a messages/routing.proto-domain changelistenstatement request.
func DecodeGRPCChangeListenStatementRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ChangeListenStatementRequest)
	return ChangeListenStatementRequest{
		IDRequest: IDRequest{
			RefID: uint(req.RefID),
			Name:  req.Name,
		},
		ListenStatement: convertPBListenStatement(req.ListenStatement),
	}, nil
}

// DecodeGRPCAddServerNameRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC AddServerName request to a messages/routing.proto-domain addservername request.
func DecodeGRPCAddServerNameRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddServerNameRequest)
	return AddServerNameRequest{
		IDRequest: IDRequest{
			RefID: uint(req.RefID),
			Name:  req.Name,
		},
		ServerName: req.ServerName,
	}, nil
}

// DecodeGRPCRemoveServerNameRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemoveServerName request to a messages/routing.proto-domain removeservername request.
func DecodeGRPCRemoveServerNameRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RemoveServerNameRequest)
	return RemoveServerNameRequest{
		IDRequest: IDRequest{
			RefID: uint(req.RefID),
			Name:  req.Name,
		},
		ID: int(req.Id),
	}, nil
}

// DecodeGRPCConfigurationsRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC Configurations request to a messages/routing.proto-domain configurations request.
func DecodeGRPCConfigurationsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return ConfigurationsRequest{}, nil
}

// EncodeGRPCCreateConfigResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain createconfig response to a gRPC CreateConfig response.
func EncodeGRPCCreateConfigResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(CreateConfigResponse)
	gRPCRes := &pb.CreateConfigResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCEditConfigResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain editconfig response to a gRPC EditConfig response.
func EncodeGRPCEditConfigResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(EditConfigResponse)
	gRPCRes := &pb.EditConfigResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCGetConfigResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain getconfig response to a gRPC GetConfig response.
func EncodeGRPCGetConfigResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetConfigResponse)
	gRPCRes := &pb.GetConfigResponse{
		Config: ConvertConfiguration(res.Config),
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemoveConfigResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain removeconfig response to a gRPC RemoveConfig response.
func EncodeGRPCRemoveConfigResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveConfigResponse)
	gRPCRes := &pb.RemoveConfigResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCAddLocationResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain addlocation response to a gRPC AddLocation response.
func EncodeGRPCAddLocationResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(AddLocationResponse)
	gRPCRes := &pb.AddLocationResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemoveLocationResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain removelocation response to a gRPC RemoveLocation response.
func EncodeGRPCRemoveLocationResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveLocationResponse)
	gRPCRes := &pb.RemoveLocationResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCChangeListenStatementResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain changelistenstatement response to a gRPC ChangeListenStatement response.
func EncodeGRPCChangeListenStatementResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(ChangeListenStatementResponse)
	gRPCRes := &pb.ChangeListenStatementResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCAddServerNameResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain addservername response to a gRPC AddServerName response.
func EncodeGRPCAddServerNameResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(AddServerNameResponse)
	gRPCRes := &pb.AddServerNameResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemoveServerNameResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain removeservername response to a gRPC RemoveServerName response.
func EncodeGRPCRemoveServerNameResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveServerNameResponse)
	gRPCRes := &pb.RemoveServerNameResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCConfigurationsResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/routing.proto-domain configurations response to a gRPC Configurations response.
func EncodeGRPCConfigurationsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(ConfigurationsResponse)
	gRPCRes := &pb.ConfigurationsResponse{
		Configurations: ConvertConfigurations(res.Configurations),
	}
	return gRPCRes, nil
}
