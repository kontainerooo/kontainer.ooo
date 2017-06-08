package module

import (
	"context"
	"strings"

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
		createcontainermodule: grpctransport.NewServer(
			endpoints.CreateContainerModuleEndpoint,
			DecodeGRPCCreateContainerModuleRequest,
			EncodeGRPCCreateContainerModuleResponse,
			options...,
		),
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
		setlink: grpctransport.NewServer(
			endpoints.SetLinkEndpoint,
			DecodeGRPCSetLinkRequest,
			EncodeGRPCSetLinkResponse,
			options...,
		),
		removelink: grpctransport.NewServer(
			endpoints.RemoveLinkEndpoint,
			DecodeGRPCRemoveLinkRequest,
			EncodeGRPCRemoveLinkResponse,
			options...,
		),
		getmodules: grpctransport.NewServer(
			endpoints.GetModulesEndpoint,
			DecodeGRPCGetModulesRequest,
			EncodeGRPCGetModulesResponse,
			options...,
		),
	}
}

type grpcServer struct {
	createcontainermodule grpctransport.Handler
	setpublickey          grpctransport.Handler
	removefile            grpctransport.Handler
	removedirectory       grpctransport.Handler
	getfiles              grpctransport.Handler
	getfile               grpctransport.Handler
	uploadfile            grpctransport.Handler
	getmoduleconfig       grpctransport.Handler
	sendcommand           grpctransport.Handler
	setenv                grpctransport.Handler
	getenv                grpctransport.Handler
	setlink               grpctransport.Handler
	removelink            grpctransport.Handler
	getmodules            grpctransport.Handler
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

func (s *grpcServer) CreateContainerModule(ctx oldcontext.Context, req *modulePB.CreateContainerModuleRequest) (*modulePB.CreateContainerModuleResponse, error) {
	_, res, err := s.createcontainermodule.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.CreateContainerModuleResponse), nil
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

func (s *grpcServer) SetLink(ctx oldcontext.Context, req *modulePB.SetLinkRequest) (*modulePB.SetLinkResponse, error) {
	_, res, err := s.setlink.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.SetLinkResponse), nil
}

func (s *grpcServer) RemoveLink(ctx oldcontext.Context, req *modulePB.RemoveLinkRequest) (*modulePB.RemoveLinkResponse, error) {
	_, res, err := s.removelink.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.RemoveLinkResponse), nil
}

func (s *grpcServer) GetModules(ctx oldcontext.Context, req *modulePB.GetModulesRequest) (*modulePB.GetModulesResponse, error) {
	_, res, err := s.getmodules.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*modulePB.GetModulesResponse), nil
}

// DecodeGRPCCreateContainerModuleRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateContainerModule request to a module.proto-domain createcontainermodule request.
func DecodeGRPCCreateContainerModuleRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.CreateContainerModuleRequest)
	return CreateContainerModuleRequest{
		RefID: uint(req.RefID),
		KmiID: uint(req.KmiID),
		Name:  req.Name,
	}, nil
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

// DecodeGRPCSetLinkRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC SetLink request to a module.proto-domain setlink request.
func DecodeGRPCSetLinkRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.SetLinkRequest)
	return SetLinkRequest{
		RefID:         uint(req.RefID),
		ContainerName: req.ContainerName,
		LinkName:      req.LinkName,
		LinkInterface: req.LinkInterface,
	}, nil
}

// DecodeGRPCRemoveLinkRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC RemoveLink request to a module.proto-domain removelink request.
func DecodeGRPCRemoveLinkRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.RemoveLinkRequest)
	return RemoveLinkRequest{
		RefID:         uint(req.RefID),
		ContainerName: req.ContainerName,
		LinkName:      req.LinkName,
		LinkInterface: req.LinkInterface,
	}, nil
}

// DecodeGRPCGetModulesRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetModules request to a module.proto-domain getmodules request.
func DecodeGRPCGetModulesRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*modulePB.GetModulesRequest)
	return GetModulesRequest{
		RefID: uint(req.RefID),
	}, nil
}

// EncodeGRPCCreateContainerModuleResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain createcontainermodule response to a gRPC CreateContainerModule response.
func EncodeGRPCCreateContainerModuleResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(CreateContainerModuleResponse)
	gRPCRes := &modulePB.CreateContainerModuleResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
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

	linkMap := make(map[string]string)
	for k, v := range res.Links {
		linkMap[k] = strings.Join(v[:], ",")
	}

	gRPCRes := &modulePB.GetModuleConfigResponse{
		Kmi:   cKMI,
		Links: linkMap,
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

// EncodeGRPCSetLinkResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain setlink response to a gRPC SetLink response.
func EncodeGRPCSetLinkResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(SetLinkResponse)
	gRPCRes := &modulePB.SetLinkResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCRemoveLinkResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain removelink response to a gRPC RemoveLink response.
func EncodeGRPCRemoveLinkResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(RemoveLinkResponse)
	gRPCRes := &modulePB.RemoveLinkResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

func toPBModules(mods []Module) []*modulePB.Module {
	pbMods := []*modulePB.Module{}
	for _, v := range mods {
		pbMods = append(pbMods, &modulePB.Module{
			ContainerName: v.ContainerName,
			Kmdi:          kmi.ConvertPBKMDI(v.KMDI),
		})
	}

	return pbMods
}

// EncodeGRPCGetModulesResponse is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain getmodules response to a gRPC GetModules response.
func EncodeGRPCGetModulesResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetModulesResponse)
	gRPCRes := &modulePB.GetModulesResponse{
		Modules: toPBModules(res.Modules),
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}
