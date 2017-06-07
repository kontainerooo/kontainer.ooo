package client

import (
	"context"
	"errors"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
	kmiClient "github.com/kontainerooo/kontainer.ooo/pkg/kmi/client"
	kmiPB "github.com/kontainerooo/kontainer.ooo/pkg/kmi/pb"
	"github.com/kontainerooo/kontainer.ooo/pkg/module"
	"github.com/kontainerooo/kontainer.ooo/pkg/module/pb"
)

// New creates a set of endpoints based on a gRPC connection
func New(conn *grpc.ClientConn, logger log.Logger) *module.Endpoints {

	var CreateContainerModuleEndpoint endpoint.Endpoint
	{
		CreateContainerModuleEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"CreateContainerModule",
			EncodeGRPCCreateContainerModuleRequest,
			DecodeGRPCCreateContainerModuleResponse,
			pb.CreateContainerModuleResponse{},
		).Endpoint()
	}
	var SetPublicKeyEndpoint endpoint.Endpoint
	{
		SetPublicKeyEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"SetPublicKey",
			EncodeGRPCSetPublicKeyRequest,
			DecodeGRPCSetPublicKeyResponse,
			pb.SetPublicKeyResponse{},
		).Endpoint()
	}

	var RemoveFileEndpoint endpoint.Endpoint
	{
		RemoveFileEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"RemoveFile",
			EncodeGRPCRemoveFileRequest,
			DecodeGRPCRemoveFileResponse,
			pb.RemoveFileResponse{},
		).Endpoint()
	}

	var RemoveDirectoryEndpoint endpoint.Endpoint
	{
		RemoveDirectoryEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"RemoveDirectory",
			EncodeGRPCRemoveDirectoryRequest,
			DecodeGRPCRemoveDirectoryResponse,
			pb.RemoveDirectoryResponse{},
		).Endpoint()
	}

	var GetFilesEndpoint endpoint.Endpoint
	{
		GetFilesEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"GetFiles",
			EncodeGRPCGetFilesRequest,
			DecodeGRPCGetFilesResponse,
			pb.GetFilesResponse{},
		).Endpoint()
	}

	var GetFileEndpoint endpoint.Endpoint
	{
		GetFileEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"GetFile",
			EncodeGRPCGetFileRequest,
			DecodeGRPCGetFileResponse,
			pb.GetFileResponse{},
		).Endpoint()
	}

	var UploadFileEndpoint endpoint.Endpoint
	{
		UploadFileEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"UploadFile",
			EncodeGRPCUploadFileRequest,
			DecodeGRPCUploadFileResponse,
			pb.UploadFileResponse{},
		).Endpoint()
	}

	var GetModuleConfigEndpoint endpoint.Endpoint
	{
		GetModuleConfigEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"GetModuleConfig",
			EncodeGRPCGetModuleConfigRequest,
			DecodeGRPCGetModuleConfigResponse,
			pb.GetModuleConfigResponse{},
		).Endpoint()
	}

	var SendCommandEndpoint endpoint.Endpoint
	{
		SendCommandEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"SendCommand",
			EncodeGRPCSendCommandRequest,
			DecodeGRPCSendCommandResponse,
			pb.SendCommandResponse{},
		).Endpoint()
	}

	var SetEnvEndpoint endpoint.Endpoint
	{
		SetEnvEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"SetEnv",
			EncodeGRPCSetEnvRequest,
			DecodeGRPCSetEnvResponse,
			pb.SetEnvResponse{},
		).Endpoint()
	}

	var GetEnvEndpoint endpoint.Endpoint
	{
		GetEnvEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"GetEnv",
			EncodeGRPCGetEnvRequest,
			DecodeGRPCGetEnvResponse,
			pb.GetEnvResponse{},
		).Endpoint()
	}

	var SetLinkEndpoint endpoint.Endpoint
	{
		SetLinkEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"SetLink",
			EncodeGRPCSetLinkRequest,
			DecodeGRPCSetLinkResponse,
			pb.SetLinkResponse{},
		).Endpoint()
	}

	var RemoveLinkEndpoint endpoint.Endpoint
	{
		RemoveLinkEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"RemoveLink",
			EncodeGRPCRemoveLinkRequest,
			DecodeGRPCRemoveLinkResponse,
			pb.RemoveLinkResponse{},
		).Endpoint()
	}

	var GetModulesEndpoint endpoint.Endpoint
	{
		GetModulesEndpoint = grpctransport.NewClient(
			conn,
			"module.ModuleService",
			"GetModules",
			EncodeGRPCGetModulesRequest,
			DecodeGRPCGetModulesResponse,
			pb.GetModulesResponse{},
		).Endpoint()
	}

	return &module.Endpoints{
		CreateContainerModuleEndpoint: CreateContainerModuleEndpoint,
		SetPublicKeyEndpoint:          SetPublicKeyEndpoint,
		RemoveFileEndpoint:            RemoveFileEndpoint,
		RemoveDirectoryEndpoint:       RemoveDirectoryEndpoint,
		GetFilesEndpoint:              GetFilesEndpoint,
		GetFileEndpoint:               GetFileEndpoint,
		UploadFileEndpoint:            UploadFileEndpoint,
		GetModuleConfigEndpoint:       GetModuleConfigEndpoint,
		SendCommandEndpoint:           SendCommandEndpoint,
		SetEnvEndpoint:                SetEnvEndpoint,
		GetEnvEndpoint:                GetEnvEndpoint,
		SetLinkEndpoint:               SetLinkEndpoint,
		RemoveLinkEndpoint:            RemoveLinkEndpoint,
		GetModulesEndpoint:            GetModulesEndpoint,
	}
}

func getError(e string) error {
	if e != "" {
		return errors.New(e)
	}
	return nil
}

func convertFrontendModule(f *kmiPB.FrontendModule) *kmi.FrontendModule {
	return &kmi.FrontendModule{
		Template:   f.Template,
		Parameters: abstraction.NewJSONFromMap(f.Parameters),
	}
}

func convertFrontendModuleArray(f []*kmiPB.FrontendModule) kmi.FrontendArray {
	a := make(kmi.FrontendArray, len(f))
	for i, m := range f {
		a[i] = convertFrontendModule(m)
	}
	return a
}

func convertKMDI(k *kmiPB.KMDI) kmi.KMDI {
	return kmi.KMDI{
		ID:          uint(k.ID),
		Name:        k.Name,
		Version:     k.Version,
		Description: k.Description,
	}
}

func convertKMI(k *kmiPB.KMI) *kmi.KMI {
	return &kmi.KMI{
		KMDI:            convertKMDI(k.KMDI),
		ProvisionScript: k.ProvisionScript,
		Commands:        abstraction.NewJSONFromMap(k.Commands),
		Environment:     abstraction.NewJSONFromMap(k.Environment),
		Frontend:        convertFrontendModuleArray(k.Frontend),
		Imports:         pq.StringArray(k.Imports),
		Interfaces:      abstraction.NewJSONFromMap(k.Interfaces),
		Resources:       abstraction.NewJSONFromMap(k.Resources),
	}
}

func convertKMDIArray(k []*kmiPB.KMDI) *[]kmi.KMDI {
	a := make([]kmi.KMDI, len(k))
	for i, d := range k {
		a[i] = convertKMDI(d)
	}
	return &a
}

// EncodeGRPCCreateContainerModuleRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain createcontainermodule request to a gRPC CreateContainerModule request.
func EncodeGRPCCreateContainerModuleRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.CreateContainerModuleRequest)
	return &pb.CreateContainerModuleRequest{
		RefID: uint32(req.RefID),
		KmiID: uint32(req.KmiID),
		Name:  req.Name,
	}, nil
}

// DecodeGRPCCreateContainerModuleResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC CreateContainerModule response to a module.proto-domain createcontainermodule response.
func DecodeGRPCCreateContainerModuleResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.CreateContainerModuleResponse)
	return &module.CreateContainerModuleResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCSetPublicKeyRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain setpublickey request to a gRPC SetPublicKey request.
func EncodeGRPCSetPublicKeyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.SetPublicKeyRequest)
	return &pb.SetPublicKeyRequest{
		RefID:         uint32(req.RefID),
		ContainerName: req.ContainerName,
		Key:           req.Key,
	}, nil
}

// DecodeGRPCSetPublicKeyResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC SetPublicKey response to a module.proto-domain setpublickey response.
func DecodeGRPCSetPublicKeyResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.SetPublicKeyResponse)
	return &module.SetPublicKeyResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveFileRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain removefile request to a gRPC RemoveFile request.
func EncodeGRPCRemoveFileRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.RemoveFileRequest)
	return &pb.RemoveFileRequest{
		RefID:         uint32(req.RefID),
		ContainerName: req.ContainerName,
		Filename:      req.Filename,
	}, nil
}

// DecodeGRPCRemoveFileResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveFile response to a module.proto-domain removefile response.
func DecodeGRPCRemoveFileResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.RemoveFileResponse)
	return &module.RemoveFileResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveDirectoryRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain removedirectory request to a gRPC RemoveDirectory request.
func EncodeGRPCRemoveDirectoryRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.RemoveDirectoryRequest)
	return &pb.RemoveDirectoryRequest{
		RefID:         uint32(req.RefID),
		ContainerName: req.ContainerName,
		Path:          req.Path,
	}, nil
}

// DecodeGRPCRemoveDirectoryResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveDirectory response to a module.proto-domain removedirectory response.
func DecodeGRPCRemoveDirectoryResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.RemoveDirectoryResponse)
	return &module.RemoveDirectoryResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCGetFilesRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain getfiles request to a gRPC GetFiles request.
func EncodeGRPCGetFilesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.GetFilesRequest)
	return &pb.GetFilesRequest{
		RefID:         uint32(req.RefID),
		ContainerName: req.ContainerName,
		Path:          req.Path,
	}, nil
}

// DecodeGRPCGetFilesResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetFiles response to a module.proto-domain getfiles response.
func DecodeGRPCGetFilesResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.GetFilesResponse)
	return &module.GetFilesResponse{
		Files: response.Files,
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCGetFileRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain getfile request to a gRPC GetFile request.
func EncodeGRPCGetFileRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.GetFileRequest)
	return &pb.GetFileRequest{
		RefID:         uint32(req.RefID),
		ContainerName: req.ContainerName,
		Path:          req.Path,
	}, nil
}

// DecodeGRPCGetFileResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetFile response to a module.proto-domain getfile response.
func DecodeGRPCGetFileResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.GetFileResponse)
	return &module.GetFileResponse{
		Content: response.Content,
		Error:   getError(response.Error),
	}, nil
}

// EncodeGRPCUploadFileRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain uploadfile request to a gRPC UploadFile request.
func EncodeGRPCUploadFileRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.UploadFileRequest)
	return &pb.UploadFileRequest{
		RefID:         uint32(req.RefID),
		ContainerName: req.ContainerName,
		Path:          req.Path,
		Content:       req.Content,
		Override:      req.Override,
	}, nil
}

// DecodeGRPCUploadFileResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC UploadFile response to a module.proto-domain uploadfile response.
func DecodeGRPCUploadFileResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.UploadFileResponse)
	return &module.UploadFileResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCGetModuleConfigRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain getmoduleconfig request to a gRPC GetModuleConfig request.
func EncodeGRPCGetModuleConfigRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.GetModuleConfigRequest)
	return &pb.GetModuleConfigRequest{
		RefID:         uint32(req.RefID),
		ContainerName: req.ContainerName,
	}, nil
}

// DecodeGRPCGetModuleConfigResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetModuleConfig response to a module.proto-domain getmoduleconfig response.
func DecodeGRPCGetModuleConfigResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.GetModuleConfigResponse)

	arrayMap := make(map[string][]string)
	for k, v := range response.Links {
		arrayMap[k] = strings.Split(v, ",")
	}

	return &module.GetModuleConfigResponse{
		ContainerKMI: *convertKMI(response.Kmi),
		Links:        arrayMap,
		Error:        getError(response.Error),
	}, nil
}

// EncodeGRPCSendCommandRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain sendcommand request to a gRPC SendCommand request.
func EncodeGRPCSendCommandRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.SendCommandRequest)
	return &pb.SendCommandRequest{
		RefID:         uint32(req.RefID),
		ContainerName: req.ContainerName,
		Command:       req.Command,
		Env:           req.Env,
	}, nil
}

// DecodeGRPCSendCommandResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC SendCommand response to a module.proto-domain sendcommand response.
func DecodeGRPCSendCommandResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.SendCommandResponse)
	return &module.SendCommandResponse{
		Response: response.Response,
		Error:    getError(response.Error),
	}, nil
}

// EncodeGRPCSetEnvRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain setenv request to a gRPC SetEnv request.
func EncodeGRPCSetEnvRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.SetEnvRequest)
	return &pb.SetEnvRequest{
		RefID:         uint32(req.RefID),
		ContainerName: req.ContainerName,
		Key:           req.Key,
		Value:         req.Value,
	}, nil
}

// DecodeGRPCSetEnvResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC SetEnv response to a module.proto-domain setenv response.
func DecodeGRPCSetEnvResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.SetEnvResponse)
	return &module.SetEnvResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCGetEnvRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain getenv request to a gRPC GetEnv request.
func EncodeGRPCGetEnvRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.GetEnvRequest)
	return &pb.GetEnvRequest{
		RefID:         uint32(req.RefID),
		ContainerName: req.ContainerName,
		Key:           req.Key,
	}, nil
}

// DecodeGRPCGetEnvResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetEnv response to a module.proto-domain getenv response.
func DecodeGRPCGetEnvResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.GetEnvResponse)
	return &module.GetEnvResponse{
		Value: response.Value,
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCSetLinkRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain setlink request to a gRPC SetLink request.
func EncodeGRPCSetLinkRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.SetLinkRequest)
	return &pb.SetLinkRequest{
		RefID:         uint32(req.RefID),
		ContainerName: req.ContainerName,
		LinkInterface: req.LinkInterface,
		LinkName:      req.LinkName,
	}, nil
}

// DecodeGRPCSetLinkResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC SetLink response to a module.proto-domain setlink response.
func DecodeGRPCSetLinkResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.SetLinkResponse)
	return &module.SetLinkResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCRemoveLinkRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain removelink request to a gRPC RemoveLink request.
func EncodeGRPCRemoveLinkRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.RemoveLinkRequest)
	return &pb.RemoveLinkRequest{
		RefID:         uint32(req.RefID),
		ContainerName: req.ContainerName,
		LinkInterface: req.LinkInterface,
		LinkName:      req.LinkName,
	}, nil
}

// DecodeGRPCRemoveLinkResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC RemoveLink response to a module.proto-domain removelink response.
func DecodeGRPCRemoveLinkResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.RemoveLinkResponse)
	return &module.RemoveLinkResponse{
		Error: getError(response.Error),
	}, nil
}

// EncodeGRPCGetModulesRequest is a transport/grpc.EncodeRequestFunc that converts a
// module.proto-domain getmodules request to a gRPC GetModules request.
func EncodeGRPCGetModulesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*module.GetModulesRequest)
	return &pb.GetModulesRequest{
		RefID: uint32(req.RefID),
	}, nil
}

func pbToModules(mod []*pb.Module) []module.Module {
	mods := []module.Module{}
	for _, v := range mod {
		mods = append(mods, module.Module{
			ContainerName: v.ContainerName,
			KMDI:          kmiClient.ConvertKMDI(v.Kmdi),
		})
	}

	return mods
}

// DecodeGRPCGetModulesResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetModules response to a module.proto-domain getmodules response.
func DecodeGRPCGetModulesResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.GetModulesResponse)
	return &module.GetModulesResponse{
		Error:   getError(response.Error),
		Modules: pbToModules(response.Modules),
	}, nil
}
