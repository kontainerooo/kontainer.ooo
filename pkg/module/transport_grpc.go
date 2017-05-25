package module

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
	kmiPB "github.com/kontainerooo/kontainer.ooo/pkg/kmi/pb"
	modulePB "github.com/kontainerooo/kontainer.ooo/pkg/module/pb"
	oldcontext "golang.org/x/net/context"
)

// MakeGRPCServer makes a set of Endpoints available as a gRPC moduleServer
func MakeGRPCServer(ctx context.Context, endpoints Endpoints, logger log.Logger) modulePB.ModuleServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		setpublickey: grpctransport.NewServer(
			endpoints.SetPublicKeyEndpoint,
			DecodeGRPCSetPublicKeyRequest,
			EncodeGRPCSetPublicKeyResponse,
			options...,
		),

		removefile: grpctransport.NewServer(
			endpoints.RemoveFileEndpoint,
			DecodeGRPCRemoveFileRequest,
			EncodeGRPCRemoveFileResponse,
			options...,
		),

		removedirectory: grpctransport.NewServer(
			endpoints.RemoveDirectoryEndpoint,
			DecodeGRPCRemoveDirectoryRequest,
			EncodeGRPCRemoveDirectoryResponse,
			options...,
		),

		getfiles: grpctransport.NewServer(
			endpoints.GetFilesEndpoint,
			DecodeGRPCGetFilesRequest,
			EncodeGRPCGetFilesResponse,
			options...,
		),

		getfile: grpctransport.NewServer(
			endpoints.GetFileEndpoint,
			DecodeGRPCGetFileRequest,
			EncodeGRPCGetFileResponse,
			options...,
		),

		uploadfile: grpctransport.NewServer(
			endpoints.UploadFileEndpoint,
			DecodeGRPCUploadFileRequest,
			EncodeGRPCUploadFileResponse,
			options...,
		),

		getmoduleconfig: grpctransport.NewServer(
			endpoints.GetModuleConfigEndpoint,
			DecodeGRPCGetModuleConfigRequest,
			EncodeGRPCGetModuleConfigResponse,
			options...,
		),

		sendcommand: grpctransport.NewServer(
			endpoints.SendCommandEndpoint,
			DecodeGRPCSendCommandRequest,
			EncodeGRPCSendCommandResponse,
			options...,
		),

		setenv: grpctransport.NewServer(
			endpoints.SetEnvEndpoint,
			DecodeGRPCSetEnvRequest,
			EncodeGRPCSetEnvResponse,
			options...,
		),

		getenv: grpctransport.NewServer(
			endpoints.GetEnvEndpoint,
			DecodeGRPCGetEnvRequest,
			EncodeGRPCGetEnvResponse,
			options...,
		),
	}
}

type grpcServer struct {
	setpublickey    grpctransport.Handler
	removefile      grpctransport.Handler
	removedirectory grpctransport.Handler
	getfiles        grpctransport.Handler
	getfile         grpctransport.Handler
	uploadfile      grpctransport.Handler
	getmoduleconfig grpctransport.Handler
	sendcommand     grpctransport.Handler
	setenv          grpctransport.Handler
	getenv          grpctransport.Handler
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

func convertPBKMI(k *kmi.KMI) *kmiPB.KMI {
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

func (s *grpcServer) SetPublicKey(ctx oldcontext.Context, req *modulePB.SetPublicKeyRequest) (*modulePB.SetPublicKeyResponse, error) {
	_, res, err := s.setpublickey.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.SetPublicKeyResponse), nil
}

func (s *grpcServer) RemoveFile(ctx oldcontext.Context, req *modulePB.RemoveFileRequest) (*modulePB.RemoveFileResponse, error) {
	_, res, err := s.removefile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.RemoveFileResponse), nil
}

func (s *grpcServer) RemoveDirectory(ctx oldcontext.Context, req *modulePB.RemoveDirectoryRequest) (*modulePB.RemoveDirectoryResponse, error) {
	_, res, err := s.removedirectory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.RemoveDirectoryResponse), nil
}

func (s *grpcServer) GetFiles(ctx oldcontext.Context, req *modulePB.GetFilesRequest) (*modulePB.GetFilesResponse, error) {
	_, res, err := s.getfiles.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.GetFilesResponse), nil
}

func (s *grpcServer) GetFile(ctx oldcontext.Context, req *modulePB.GetFileRequest) (*modulePB.GetFileResponse, error) {
	_, res, err := s.getfile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.GetFileResponse), nil
}

func (s *grpcServer) UploadFile(ctx oldcontext.Context, req *modulePB.UploadFileRequest) (*modulePB.UploadFileResponse, error) {
	_, res, err := s.uploadfile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.UploadFileResponse), nil
}

func (s *grpcServer) GetModuleConfig(ctx oldcontext.Context, req *modulePB.GetModuleConfigRequest) (*modulePB.GetModuleConfigResponse, error) {
	_, res, err := s.getmoduleconfig.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.GetModuleConfigResponse), nil
}

func (s *grpcServer) SendCommand(ctx oldcontext.Context, req *modulePB.SendCommandRequest) (*modulePB.SendCommandResponse, error) {
	_, res, err := s.sendcommand.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.SendCommandResponse), nil
}

func (s *grpcServer) SetEnv(ctx oldcontext.Context, req *modulePB.SetEnvRequest) (*modulePB.SetEnvResponse, error) {
	_, res, err := s.setenv.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.SetEnvResponse), nil
}

func (s *grpcServer) GetEnv(ctx oldcontext.Context, req *modulePB.GetEnvRequest) (*modulePB.GetEnvResponse, error) {
	_, res, err := s.getenv.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.GetEnvResponse), nil
}

// DecodeGRPCSetPublicKeyRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC SetPublicKey request to a module.proto-domain setpublickey request.
func DecodeGRPCSetPublicKeyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.SetPublicKeyRequest)
	return SetPublicKeyRequest{
		RefID:         uint(req.RefID),
		ContainerName: req.ContainerName,
		Key:           req.Key,
	}, nil
}

// DecodeGRPCRemoveFileRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemoveFile request to a module.proto-domain removefile request.
func DecodeGRPCRemoveFileRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.RemoveFileRequest)
	return RemoveFileRequest{
		RefID:         uint(req.RefID),
		ContainerName: req.ContainerName,
		Filename:      req.Filename,
	}, nil
}

// DecodeGRPCRemoveDirectoryRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemoveDirectory request to a module.proto-domain removedirectory request.
func DecodeGRPCRemoveDirectoryRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.RemoveDirectoryRequest)
	return RemoveDirectoryRequest{
		RefID:         uint(req.RefID),
		ContainerName: req.ContainerName,
		Path:          req.Path,
	}, nil
}

// DecodeGRPCGetFilesRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetFiles request to a module.proto-domain getfiles request.
func DecodeGRPCGetFilesRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.GetFilesRequest)
	return GetFilesRequest{
		RefID:         uint(req.RefID),
		ContainerName: req.ContainerName,
		Path:          req.Path,
	}, nil
}

// DecodeGRPCGetFileRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetFile request to a module.proto-domain getfile request.
func DecodeGRPCGetFileRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.GetFileRequest)
	return GetFileRequest{
		RefID:         uint(req.RefID),
		ContainerName: req.ContainerName,
		Path:          req.Path,
	}, nil
}

// DecodeGRPCUploadFileRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC UploadFile request to a module.proto-domain uploadfile request.
func DecodeGRPCUploadFileRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.UploadFileRequest)
	return UploadFileRequest{
		RefID:         uint(req.RefID),
		ContainerName: req.ContainerName,
		Path:          req.Path,
		Content:       req.Content,
		Override:      req.Override,
	}, nil
}

// DecodeGRPCGetModuleConfigRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetModuleConfig request to a module.proto-domain getmoduleconfig request.
func DecodeGRPCGetModuleConfigRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.GetModuleConfigRequest)
	return GetModuleConfigRequest{
		RefID:         uint(req.RefID),
		ContainerName: req.ContainerName,
	}, nil
}

// DecodeGRPCSendCommandRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC SendCommand request to a module.proto-domain sendcommand request.
func DecodeGRPCSendCommandRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.SendCommandRequest)
	return SendCommandRequest{
		RefID:         uint(req.RefID),
		ContainerName: req.ContainerName,
		Command:       req.Command,
		Env:           req.Env,
	}, nil
}

// DecodeGRPCSetEnvRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC SetEnv request to a module.proto-domain setenv request.
func DecodeGRPCSetEnvRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.SetEnvRequest)
	return SetEnvRequest{
		RefID:         uint(req.RefID),
		ContainerName: req.ContainerName,
		Key:           req.Key,
		Value:         req.Value,
	}, nil
}

// DecodeGRPCGetEnvRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetEnv request to a module.proto-domain getenv request.
func DecodeGRPCGetEnvRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.GetEnvRequest)
	return GetEnvRequest{
		RefID:         uint(req.RefID),
		ContainerName: req.ContainerName,
		Key:           req.Key,
	}, nil
}

// EncodeGRPCSetPublicKeyResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain setpublickey response to a gRPC SetPublicKey response.
func EncodeGRPCSetPublicKeyResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(SetPublicKeyResponse)
	gRPCRes := &modulePB.SetPublicKeyResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemoveFileResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain removefile response to a gRPC RemoveFile response.
func EncodeGRPCRemoveFileResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveFileResponse)
	gRPCRes := &modulePB.RemoveFileResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemoveDirectoryResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain removedirectory response to a gRPC RemoveDirectory response.
func EncodeGRPCRemoveDirectoryResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveDirectoryResponse)
	gRPCRes := &modulePB.RemoveDirectoryResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCGetFilesResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain getfiles response to a gRPC GetFiles response.
func EncodeGRPCGetFilesResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetFilesResponse)
	gRPCRes := &modulePB.GetFilesResponse{
		Files: res.Files,
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCGetFileResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain getfile response to a gRPC GetFile response.
func EncodeGRPCGetFileResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetFileResponse)
	gRPCRes := &modulePB.GetFileResponse{
		Content: res.Content,
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCUploadFileResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain uploadfile response to a gRPC UploadFile response.
func EncodeGRPCUploadFileResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(UploadFileResponse)
	gRPCRes := &modulePB.UploadFileResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCGetModuleConfigResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain getmoduleconfig response to a gRPC GetModuleConfig response.
func EncodeGRPCGetModuleConfigResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetModuleConfigResponse)
	cKMI := convertPBKMI(&res.ContainerKMI)
	gRPCRes := &modulePB.GetModuleConfigResponse{
		Kmi: cKMI,
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCSendCommandResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain sendcommand response to a gRPC SendCommand response.
func EncodeGRPCSendCommandResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(SendCommandResponse)
	gRPCRes := &modulePB.SendCommandResponse{
		Response: res.Response,
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCSetEnvResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain setenv response to a gRPC SetEnv response.
func EncodeGRPCSetEnvResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(SetEnvResponse)
	gRPCRes := &modulePB.SetEnvResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCGetEnvResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain getenv response to a gRPC GetEnv response.
func EncodeGRPCGetEnvResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetEnvResponse)
	gRPCRes := &modulePB.GetEnvResponse{
		Value: res.Value,
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}
