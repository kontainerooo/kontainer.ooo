package network

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/kontainerooo/kontainer.ooo/pkg/network/pb"
	oldcontext "golang.org/x/net/context"
)

// MakeGRPCServer makes a set of Endpoints available as a gRPC networkServer
func MakeGRPCServer(ctx context.Context, endpoints Endpoints, logger log.Logger) pb.NetworkServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcServer{
		createprimarynetworkforcontainer: grpctransport.NewServer(
			endpoints.CreatePrimaryNetworkForContainerEndpoint,
			DecodeGRPCCreatePrimaryNetworkForContainerRequest,
			EncodeGRPCCreatePrimaryNetworkForContainerResponse,
			options...,
		),
		createnetwork: grpctransport.NewServer(
			endpoints.CreateNetworkEndpoint,
			DecodeGRPCCreateNetworkRequest,
			EncodeGRPCCreateNetworkResponse,
			options...,
		),
		removenetworkbyname: grpctransport.NewServer(
			endpoints.RemoveNetworkByNameEndpoint,
			DecodeGRPCRemoveNetworkByNameRequest,
			EncodeGRPCRemoveNetworkByNameResponse,
			options...,
		),
		addcontainertonetwork: grpctransport.NewServer(
			endpoints.AddContainerToNetworkEndpoint,
			DecodeGRPCAddContainerToNetworkRequest,
			EncodeGRPCAddContainerToNetworkResponse,
			options...,
		),
		removecontainerfromnetwork: grpctransport.NewServer(
			endpoints.RemoveContainerFromNetworkEndpoint,
			DecodeGRPCRemoveContainerFromNetworkRequest,
			EncodeGRPCRemoveContainerFromNetworkResponse,
			options...,
		),
		exposeporttocontainer: grpctransport.NewServer(
			endpoints.ExposePortToContainerEndpoint,
			DecodeGRPCExposePortToContainerRequest,
			EncodeGRPCExposePortToContainerResponse,
			options...,
		),
		removeportfromcontainer: grpctransport.NewServer(
			endpoints.RemovePortFromContainerEndpoint,
			DecodeGRPCRemovePortFromContainerRequest,
			EncodeGRPCRemovePortFromContainerResponse,
			options...,
		),
	}
}

type grpcServer struct {
	createprimarynetworkforcontainer grpctransport.Handler
	createnetwork                    grpctransport.Handler
	removenetworkbyname              grpctransport.Handler
	addcontainertonetwork            grpctransport.Handler
	removecontainerfromnetwork       grpctransport.Handler
	exposeporttocontainer            grpctransport.Handler
	removeportfromcontainer          grpctransport.Handler
}

func (s *grpcServer) CreatePrimaryNetworkForContainer(ctx oldcontext.Context, req *pb.CreatePrimaryNetworkForContainerRequest) (*pb.CreatePrimaryNetworkForContainerResponse, error) {
	_, res, err := s.createprimarynetworkforcontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.CreatePrimaryNetworkForContainerResponse), nil
}

func (s *grpcServer) CreateNetwork(ctx oldcontext.Context, req *pb.CreateNetworkRequest) (*pb.CreateNetworkResponse, error) {
	_, res, err := s.createnetwork.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.CreateNetworkResponse), nil
}

func (s *grpcServer) RemoveNetworkByName(ctx oldcontext.Context, req *pb.RemoveNetworkByNameRequest) (*pb.RemoveNetworkByNameResponse, error) {
	_, res, err := s.removenetworkbyname.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.RemoveNetworkByNameResponse), nil
}

func (s *grpcServer) AddContainerToNetwork(ctx oldcontext.Context, req *pb.AddContainerToNetworkRequest) (*pb.AddContainerToNetworkResponse, error) {
	_, res, err := s.addcontainertonetwork.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.AddContainerToNetworkResponse), nil
}

func (s *grpcServer) RemoveContainerFromNetwork(ctx oldcontext.Context, req *pb.RemoveContainerFromNetworkRequest) (*pb.RemoveContainerFromNetworkResponse, error) {
	_, res, err := s.removecontainerfromnetwork.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.RemoveContainerFromNetworkResponse), nil
}

func (s *grpcServer) ExposePortToContainer(ctx oldcontext.Context, req *pb.ExposePortToContainerRequest) (*pb.ExposePortToContainerResponse, error) {
	_, res, err := s.exposeporttocontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ExposePortToContainerResponse), nil
}

func (s *grpcServer) RemovePortFromContainer(ctx oldcontext.Context, req *pb.RemovePortFromContainerRequest) (*pb.RemovePortFromContainerResponse, error) {
	_, res, err := s.removeportfromcontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.RemovePortFromContainerResponse), nil
}

func nwConfigToPBConfig(c Config) *pb.NetworkConfig {
	return &pb.NetworkConfig{
		Driver: c.Driver,
		Name:   c.Name,
	}
}

func pbConfigToNWConfig(c pb.NetworkConfig) *Config {
	return &Config{
		Driver: c.Driver,
		Name:   c.Name,
	}
}

// DecodeGRPCCreatePrimaryNetworkForContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreatePrimaryNetworkForContainer request to a messages/network.proto-domain createprimarynetworkforcontainer request.
func DecodeGRPCCreatePrimaryNetworkForContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreatePrimaryNetworkForContainerRequest)
	return CreatePrimaryNetworkForContainerRequest{
		RefID:       uint(req.RefID),
		Config:      pbConfigToNWConfig(*req.Config),
		ContainerID: req.ContainerID,
	}, nil
}

// DecodeGRPCCreateNetworkRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateNetwork request to a messages/network.proto-domain createnetwork request.
func DecodeGRPCCreateNetworkRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateNetworkRequest)
	return CreateNetworkRequest{
		RefID:  uint(req.RefID),
		Config: pbConfigToNWConfig(*req.Config),
	}, nil
}

// DecodeGRPCRemoveNetworkByNameRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemoveNetworkByName request to a messages/network.proto-domain removenetworkbyname request.
func DecodeGRPCRemoveNetworkByNameRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RemoveNetworkByNameRequest)
	return RemoveNetworkByNameRequest{
		RefID: uint(req.RefID),
		Name:  req.Name,
	}, nil
}

// DecodeGRPCAddContainerToNetworkRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC AddContainerToNetwork request to a messages/network.proto-domain addcontainertonetwork request.
func DecodeGRPCAddContainerToNetworkRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddContainerToNetworkRequest)
	return AddContainerToNetworkRequest{
		RefID:       uint(req.RefID),
		Name:        req.Name,
		ContainerID: req.ContainerID,
	}, nil
}

// DecodeGRPCRemoveContainerFromNetworkRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemoveContainerFromNetwork request to a messages/network.proto-domain removecontainerfromnetwork request.
func DecodeGRPCRemoveContainerFromNetworkRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RemoveContainerFromNetworkRequest)
	return RemoveContainerFromNetworkRequest{
		RefID:       uint(req.RefID),
		Name:        req.Name,
		ContainerID: req.ContainerID,
	}, nil
}

// DecodeGRPCExposePortToContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC ExposePortToContainer request to a messages/network.proto-domain exposeporttocontainer request.
func DecodeGRPCExposePortToContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ExposePortToContainerRequest)
	return ExposePortToContainerRequest{
		RefID:          uint(req.RefID),
		SrcContainerID: req.SrcContainerID,
		Port:           uint16(req.Port),
		Protocol:       req.Protocol,
		DstContainerID: req.DstContainerID,
	}, nil
}

// DecodeGRPCRemovePortFromContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemovePortFromContainer request to a messages/network.proto-domain removeportfromcontainer request.
func DecodeGRPCRemovePortFromContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RemovePortFromContainerRequest)
	return RemovePortFromContainerRequest{
		RefID:          uint(req.RefID),
		SrcContainerID: req.SrcContainerID,
		Port:           uint16(req.Port),
		Protocol:       req.Protocol,
		DstContainerID: req.DstContainerID,
	}, nil
}

// EncodeGRPCCreatePrimaryNetworkForContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain createprimarynetworkforcontainer response to a gRPC CreatePrimaryNetworkForContainer response.
func EncodeGRPCCreatePrimaryNetworkForContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(CreatePrimaryNetworkForContainerResponse)
	gRPCRes := &pb.CreatePrimaryNetworkForContainerResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCCreateNetworkResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain createnetwork response to a gRPC CreateNetwork response.
func EncodeGRPCCreateNetworkResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(CreateNetworkResponse)
	gRPCRes := &pb.CreateNetworkResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemoveNetworkByNameResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain removenetworkbyname response to a gRPC RemoveNetworkByName response.
func EncodeGRPCRemoveNetworkByNameResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveNetworkByNameResponse)
	gRPCRes := &pb.RemoveNetworkByNameResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCAddContainerToNetworkResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain addcontainertonetwork response to a gRPC AddContainerToNetwork response.
func EncodeGRPCAddContainerToNetworkResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(AddContainerToNetworkResponse)
	gRPCRes := &pb.AddContainerToNetworkResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemoveContainerFromNetworkResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain removecontainerfromnetwork response to a gRPC RemoveContainerFromNetwork response.
func EncodeGRPCRemoveContainerFromNetworkResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveContainerFromNetworkResponse)
	gRPCRes := &pb.RemoveContainerFromNetworkResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCExposePortToContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain exposeporttocontainer response to a gRPC ExposePortToContainer response.
func EncodeGRPCExposePortToContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(ExposePortToContainerResponse)
	gRPCRes := &pb.ExposePortToContainerResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemovePortFromContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/network.proto-domain removeportfromcontainer response to a gRPC RemovePortFromContainer response.
func EncodeGRPCRemovePortFromContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemovePortFromContainerResponse)
	gRPCRes := &pb.RemovePortFromContainerResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}
