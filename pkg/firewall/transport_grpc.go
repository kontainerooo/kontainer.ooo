package firewall

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/firewall/pb"
	oldcontext "golang.org/x/net/context"
)

// MakeGRPCServer makes a set of Endpoints available as a gRPC firewallServer
func MakeGRPCServer(ctx context.Context, endpoints Endpoints, logger log.Logger) pb.FirewallServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		initbridge: grpctransport.NewServer(
			endpoints.InitBridgeEndpoint,
			DecodeGRPCInitBridgeRequest,
			EncodeGRPCInitBridgeResponse,
			options...,
		),
		allowconnection: grpctransport.NewServer(
			endpoints.AllowConnectionEndpoint,
			DecodeGRPCAllowConnectionRequest,
			EncodeGRPCAllowConnectionResponse,
			options...,
		),
		blockconnection: grpctransport.NewServer(
			endpoints.BlockConnectionEndpoint,
			DecodeGRPCBlockConnectionRequest,
			EncodeGRPCBlockConnectionResponse,
			options...,
		),
		allowport: grpctransport.NewServer(
			endpoints.AllowPortEndpoint,
			DecodeGRPCAllowPortRequest,
			EncodeGRPCAllowPortResponse,
			options...,
		),
		blockport: grpctransport.NewServer(
			endpoints.BlockPortEndpoint,
			DecodeGRPCBlockPortRequest,
			EncodeGRPCBlockPortResponse,
			options...,
		),
	}
}

type grpcServer struct {
	initbridge      grpctransport.Handler
	allowconnection grpctransport.Handler
	blockconnection grpctransport.Handler
	allowport       grpctransport.Handler
	blockport       grpctransport.Handler
}

func (s *grpcServer) InitBridge(ctx oldcontext.Context, req *pb.InitBridgeRequest) (*pb.InitBridgeResponse, error) {
	_, res, err := s.initbridge.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.InitBridgeResponse), nil
}

func (s *grpcServer) AllowConnection(ctx oldcontext.Context, req *pb.AllowConnectionRequest) (*pb.AllowConnectionResponse, error) {
	_, res, err := s.allowconnection.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.AllowConnectionResponse), nil
}

func (s *grpcServer) BlockConnection(ctx oldcontext.Context, req *pb.BlockConnectionRequest) (*pb.BlockConnectionResponse, error) {
	_, res, err := s.blockconnection.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.BlockConnectionResponse), nil
}

func (s *grpcServer) AllowPort(ctx oldcontext.Context, req *pb.AllowPortRequest) (*pb.AllowPortResponse, error) {
	_, res, err := s.allowport.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.AllowPortResponse), nil
}

func (s *grpcServer) BlockPort(ctx oldcontext.Context, req *pb.BlockPortRequest) (*pb.BlockPortResponse, error) {
	_, res, err := s.blockport.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.BlockPortResponse), nil
}

// DecodeGRPCInitBridgeRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC InitBridge request to a messages/firewall.proto-domain initbridge request.
func DecodeGRPCInitBridgeRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.InitBridgeRequest)
	ip, err := abstraction.NewInet(req.IP)
	if err != nil {
		return InitBridgeRequest{}, err
	}
	return InitBridgeRequest{
		IP:    ip,
		NetIf: req.NetworkName,
	}, nil
}

// DecodeGRPCAllowConnectionRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC AllowConnection request to a messages/firewall.proto-domain allowconnection request.
func DecodeGRPCAllowConnectionRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AllowConnectionRequest)
	srcIP, err := abstraction.NewInet(req.SrcIP)
	if err != nil {
		return AllowConnectionRequest{}, err
	}
	dstIP, err := abstraction.NewInet(req.DstIP)
	if err != nil {
		return AllowConnectionRequest{}, err
	}
	return AllowConnectionRequest{
		SrcIP:      srcIP,
		DstIP:      dstIP,
		SrcNetwork: req.SrcNetwork,
		DstNetwork: req.DstNetwork,
	}, nil
}

// DecodeGRPCBlockConnectionRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC BlockConnection request to a messages/firewall.proto-domain blockconnection request.
func DecodeGRPCBlockConnectionRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.BlockConnectionRequest)
	srcIP, err := abstraction.NewInet(req.SrcIP)
	if err != nil {
		return BlockConnectionRequest{}, err
	}
	dstIP, err := abstraction.NewInet(req.DstIP)
	if err != nil {
		return BlockConnectionRequest{}, err
	}
	return BlockConnectionRequest{
		SrcIP:      srcIP,
		DstIP:      dstIP,
		SrcNetwork: req.SrcNetwork,
		DstNetwork: req.DstNetwork,
	}, nil
}

// DecodeGRPCAllowPortRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC AllowPort request to a messages/firewall.proto-domain allowport request.
func DecodeGRPCAllowPortRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AllowPortRequest)
	srcIP, err := abstraction.NewInet(req.SrcIP)
	if err != nil {
		return AllowPortRequest{}, err
	}
	dstIP, err := abstraction.NewInet(req.DstIP)
	if err != nil {
		return AllowPortRequest{}, err
	}
	return AllowPortRequest{
		SrcIP:      srcIP,
		DstIP:      dstIP,
		SrcNetwork: req.SrcNetwork,
		DstNetwork: req.DstNetwork,
		Protocol:   req.Protocol,
		Port:       uint16(req.Port),
	}, nil
}

// DecodeGRPCBlockPortRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC BlockPort request to a messages/firewall.proto-domain blockport request.
func DecodeGRPCBlockPortRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.BlockPortRequest)
	srcIP, err := abstraction.NewInet(req.SrcIP)
	if err != nil {
		return BlockPortRequest{}, err
	}
	dstIP, err := abstraction.NewInet(req.DstIP)
	if err != nil {
		return BlockPortRequest{}, err
	}
	return BlockPortRequest{
		SrcIP:      srcIP,
		DstIP:      dstIP,
		SrcNetwork: req.SrcNetwork,
		DstNetwork: req.DstNetwork,
		Protocol:   req.Protocol,
		Port:       uint16(req.Port),
	}, nil
}

// EncodeGRPCInitBridgeResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/firewall.proto-domain initbridge response to a gRPC InitBridge response.
func EncodeGRPCInitBridgeResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(InitBridgeResponse)
	gRPCRes := &pb.InitBridgeResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCAllowConnectionResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/firewall.proto-domain allowconnection response to a gRPC AllowConnection response.
func EncodeGRPCAllowConnectionResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(AllowConnectionResponse)
	gRPCRes := &pb.AllowConnectionResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCBlockConnectionResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/firewall.proto-domain blockconnection response to a gRPC BlockConnection response.
func EncodeGRPCBlockConnectionResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(BlockConnectionResponse)
	gRPCRes := &pb.BlockConnectionResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCAllowPortResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/firewall.proto-domain allowport response to a gRPC AllowPort response.
func EncodeGRPCAllowPortResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(AllowPortResponse)
	gRPCRes := &pb.AllowPortResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCBlockPortResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/firewall.proto-domain blockport response to a gRPC BlockPort response.
func EncodeGRPCBlockPortResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(BlockPortResponse)
	gRPCRes := &pb.BlockPortResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}
