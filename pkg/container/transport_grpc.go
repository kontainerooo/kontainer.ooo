package container

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/kontainerooo/kontainer.ooo/pkg/container/pb"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
	kmiPB "github.com/kontainerooo/kontainer.ooo/pkg/kmi/pb"
	oldcontext "golang.org/x/net/context"
)

// MakeGRPCServer makes a set of Endpoints available as a gRPC containerServer
func MakeGRPCServer(ctx context.Context, endpoints Endpoints, logger log.Logger) pb.ContainerServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcServer{
		createcontainer: grpctransport.NewServer(
			endpoints.CreateContainerEndpoint,
			DecodeGRPCCreateContainerRequest,
			EncodeGRPCCreateContainerResponse,
			options...,
		),
		removecontainer: grpctransport.NewServer(
			endpoints.RemoveContainerEndpoint,
			DecodeGRPCRemoveContainerRequest,
			EncodeGRPCRemoveContainerResponse,
			options...,
		),
		instances: grpctransport.NewServer(
			endpoints.InstancesEndpoint,
			DecodeGRPCInstancesRequest,
			EncodeGRPCInstancesResponse,
			options...,
		),
		stopcontainer: grpctransport.NewServer(
			endpoints.StopContainerEndpoint,
			DecodeGRPCStopContainerRequest,
			EncodeGRPCStopContainerResponse,
			options...,
		),
		execute: grpctransport.NewServer(
			endpoints.ExecuteEndpoint,
			DecodeGRPCExecuteRequest,
			EncodeGRPCExecuteResponse,
			options...,
		),
		getenv: grpctransport.NewServer(
			endpoints.GetEnvEndpoint,
			DecodeGRPCGetEnvRequest,
			EncodeGRPCGetEnvResponse,
			options...,
		),
		setenv: grpctransport.NewServer(
			endpoints.SetEnvEndpoint,
			DecodeGRPCSetEnvRequest,
			EncodeGRPCSetEnvResponse,
			options...,
		),
		idforname: grpctransport.NewServer(
			endpoints.IDForNameEndpoint,
			DecodeGRPCIDForNameRequest,
			EncodeGRPCIDForNameResponse,
			options...,
		),
		getcontainerkmi: grpctransport.NewServer(
			endpoints.GetContainerKMIEndpoint,
			DecodeGRPCGetContainerKMIRequest,
			EncodeGRPCGetContainerKMIResponse,
			options...,
		),
	}
}

type grpcServer struct {
	createcontainer grpctransport.Handler
	removecontainer grpctransport.Handler
	instances       grpctransport.Handler
	startcontainer  grpctransport.Handler
	stopcontainer   grpctransport.Handler
	execute         grpctransport.Handler
	getenv          grpctransport.Handler
	setenv          grpctransport.Handler
	idforname       grpctransport.Handler
	getcontainerkmi grpctransport.Handler
}

func convertPBFrontendModule(f *kmi.FrontendModule) *kmiPB.FrontendModule {
	return &kmiPB.FrontendModule{
		Template:   f.Template,
		Parameters: f.Parameters.ToStringMap(),
	}
}

func convertPBFrontendModuleArray(f kmi.FrontendArray) []*kmiPB.FrontendModule {
	a := make([]*kmiPB.FrontendModule, len(f))
	for i, m := range f {
		a[i] = convertPBFrontendModule(m)
	}
	return a
}

func convertPBKMDI(k kmi.KMDI) *kmiPB.KMDI {
	return &kmiPB.KMDI{
		ID:          uint32(k.ID),
		Name:        k.Name,
		Version:     k.Version,
		Description: k.Description,
	}
}

func convertPBKMI(k *CKMI) *kmiPB.KMI {
	return &kmiPB.KMI{
		KMDI:            convertPBKMDI(k.KMDI),
		ProvisionScript: k.ProvisionScript,
		Commands:        k.Commands.ToStringMap(),
		Environment:     k.Environment.ToStringMap(),
		Frontend:        convertPBFrontendModuleArray(k.Frontend),
		Imports:         k.Imports,
		Interfaces:      k.Interfaces.ToStringMap(),
		Resources:       k.Resources.ToStringMap(),
	}
}

func convertPBKMDIArray(k *[]kmi.KMDI) []*kmiPB.KMDI {
	a := make([]*kmiPB.KMDI, len(*k))
	for i, d := range *k {
		a[i] = convertPBKMDI(d)
	}
	return a
}

func (s *grpcServer) CreateContainer(ctx oldcontext.Context, req *pb.CreateContainerRequest) (*pb.CreateContainerResponse, error) {
	_, res, err := s.createcontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.CreateContainerResponse), nil
}

func (s *grpcServer) RemoveContainer(ctx oldcontext.Context, req *pb.RemoveContainerRequest) (*pb.RemoveContainerResponse, error) {
	_, res, err := s.removecontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.RemoveContainerResponse), nil
}

func (s *grpcServer) Instances(ctx oldcontext.Context, req *pb.InstancesRequest) (*pb.InstancesResponse, error) {
	_, res, err := s.instances.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.InstancesResponse), nil
}

func (s *grpcServer) StopContainer(ctx oldcontext.Context, req *pb.StopContainerRequest) (*pb.StopContainerResponse, error) {
	_, res, err := s.stopcontainer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StopContainerResponse), nil
}

func (s *grpcServer) Execute(ctx oldcontext.Context, req *pb.ExecuteRequest) (*pb.ExecuteResponse, error) {
	_, res, err := s.execute.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ExecuteResponse), nil
}

func (s *grpcServer) GetEnv(ctx oldcontext.Context, req *pb.GetEnvRequest) (*pb.GetEnvResponse, error) {
	_, res, err := s.getenv.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetEnvResponse), nil
}

func (s *grpcServer) SetEnv(ctx oldcontext.Context, req *pb.SetEnvRequest) (*pb.SetEnvResponse, error) {
	_, res, err := s.setenv.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SetEnvResponse), nil
}

func (s *grpcServer) IDForName(ctx oldcontext.Context, req *pb.IDForNameRequest) (*pb.IDForNameResponse, error) {
	_, res, err := s.idforname.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.IDForNameResponse), nil
}

func (s *grpcServer) GetContainerKMI(ctx oldcontext.Context, req *pb.GetContainerKMIRequest) (*pb.GetContainerKMIResponse, error) {
	_, res, err := s.getcontainerkmi.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetContainerKMIResponse), nil
}

// DecodeGRPCCreateContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateContainer request to a messages/container.proto-domain createcontainer request.
func DecodeGRPCCreateContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateContainerRequest)
	return CreateContainerRequest{
		RefID: uint(req.RefID),
		KmiID: uint(req.KmiID),
		Name:  req.Name,
	}, nil
}

// DecodeGRPCRemoveContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemoveContainer request to a messages/container.proto-domain removecontainer request.
func DecodeGRPCRemoveContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RemoveContainerRequest)
	return RemoveContainerRequest{
		RefID: uint(req.RefID),
		ID:    req.ID,
	}, nil
}

// DecodeGRPCInstancesRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC Instances request to a messages/container.proto-domain instances request.
func DecodeGRPCInstancesRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.InstancesRequest)
	return InstancesRequest{
		RefID: uint(req.RefID),
	}, nil
}

// DecodeGRPCStopContainerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC StopContainer request to a messages/container.proto-domain stopcontainer request.
func DecodeGRPCStopContainerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.StopContainerRequest)
	return StopContainerRequest{
		RefID: uint(req.RefID),
		ID:    req.ID,
	}, nil
}

// DecodeGRPCExecuteRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC Execute request to a messages/container.proto-domain execute request.
func DecodeGRPCExecuteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ExecuteRequest)
	return ExecuteRequest{
		RefID: uint(req.RefID),
		ID:    req.ID,
		CMD:   req.Cmd,
	}, nil
}

// DecodeGRPCGetEnvRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetEnv request to a messages/container.proto-domain getenv request.
func DecodeGRPCGetEnvRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetEnvRequest)
	return GetEnvRequest{
		RefID: uint(req.RefID),
		ID:    req.ID,
		Key:   req.Key,
	}, nil
}

// DecodeGRPCSetEnvRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC SetEnv request to a messages/container.proto-domain setenv request.
func DecodeGRPCSetEnvRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SetEnvRequest)
	return SetEnvRequest{
		RefID: uint(req.RefID),
		ID:    req.ID,
		Key:   req.Key,
		Value: req.Value,
	}, nil
}

// DecodeGRPCIDForNameRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC IDForName request to a messages/container.proto-domain idforname request.
func DecodeGRPCIDForNameRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.IDForNameRequest)
	return IDForNameRequest{
		RefID: uint(req.RefID),
		Name:  req.Name,
	}, nil
}

// DecodeGRPCGetContainerKMIRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetContainerKMI request to a messages/container.proto-domain getcontainerkmi request.
func DecodeGRPCGetContainerKMIRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetContainerKMIRequest)
	return GetContainerKMIRequest{
		ContainerID: req.ContainerID,
	}, nil
}

// EncodeGRPCCreateContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain createcontainer response to a gRPC CreateContainer response.
func EncodeGRPCCreateContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(CreateContainerResponse)
	gRPCRes := &pb.CreateContainerResponse{
		ID: res.ID,
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemoveContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain removecontainer response to a gRPC RemoveContainer response.
func EncodeGRPCRemoveContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveContainerResponse)
	gRPCRes := &pb.RemoveContainerResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCInstancesResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain instances response to a gRPC Instances response.
func EncodeGRPCInstancesResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(InstancesResponse)
	cts := []*pb.Container{}
	for _, v := range res.Containers {
		c := &pb.Container{
			ContainerID:   v.ContainerID,
			ContainerName: v.ContainerName,
			Kmi:           convertPBKMI(&v.KMI),
			RefID:         uint32(v.RefID),
		}
		cts = append(cts, c)
	}
	gRPCRes := &pb.InstancesResponse{
		Instances: cts,
	}
	return gRPCRes, nil
}

// EncodeGRPCStopContainerResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain stopcontainer response to a gRPC StopContainer response.
func EncodeGRPCStopContainerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(StopContainerResponse)
	gRPCRes := &pb.StopContainerResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCExecuteResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain execute response to a gRPC Execute response.
func EncodeGRPCExecuteResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(ExecuteResponse)
	gRPCRes := &pb.ExecuteResponse{
		Response: res.Response,
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCGetEnvResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain execute response to a gRPC GetEnv response.
func EncodeGRPCGetEnvResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetEnvResponse)
	gRPCRes := &pb.GetEnvResponse{
		Value: res.Value,
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCSetEnvResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain execute response to a gRPC SetEnv response.
func EncodeGRPCSetEnvResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(SetEnvResponse)
	gRPCRes := &pb.SetEnvResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCIDForNameResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain execute response to a gRPC IDForName response.
func EncodeGRPCIDForNameResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(IDForNameResponse)
	gRPCRes := &pb.IDForNameResponse{
		ID: res.ID,
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCGetContainerKMIResponse is a transport/grpc.EncodeRequestFunc that converts a
// messages/container.proto-domain getcontainerkmi response to a gRPC GetContainerKMI response.
func EncodeGRPCGetContainerKMIResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetContainerKMIResponse)
	gRPCRes := &pb.GetContainerKMIResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}
