package module

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/kontainerooo/kontainer.ooo/pkg/module/pb"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
)

// MakeWebsocketService makes a set of module Endpoints available as a websocket Service
func MakeWebsocketService(endpoints Endpoints) *ws.ServiceDescription {
	service, _ := ws.NewServiceDescription("moduleService", ws.ProtoIDFromString("MDL"))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"SetPublicKey",
		ws.ProtoIDFromString("SPK"),
		endpoints.SetPublicKeyEndpoint,
		DecodeWSSetPublicKeyRequest,
		EncodeGRPCSetPublicKeyResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"RemoveFile",
		ws.ProtoIDFromString("RMF"),
		endpoints.RemoveFileEndpoint,
		DecodeWSRemoveFileRequest,
		EncodeGRPCRemoveFileResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"RemoveDirectory",
		ws.ProtoIDFromString("RMD"),
		endpoints.RemoveDirectoryEndpoint,
		DecodeWSRemoveDirectoryRequest,
		EncodeGRPCRemoveDirectoryResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"GetFiles",
		ws.ProtoIDFromString("GFS"),
		endpoints.GetFilesEndpoint,
		DecodeWSGetFilesRequest,
		EncodeGRPCGetFilesResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"GetFile",
		ws.ProtoIDFromString("GTF"),
		endpoints.GetFileEndpoint,
		DecodeWSGetFileRequest,
		EncodeGRPCGetFileResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"UploadFile",
		ws.ProtoIDFromString("ULF"),
		endpoints.UploadFileEndpoint,
		DecodeWSUploadFileRequest,
		EncodeGRPCUploadFileResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"GetModuleConfig",
		ws.ProtoIDFromString("GMC"),
		endpoints.GetModuleConfigEndpoint,
		DecodeWSGetModuleConfigRequest,
		EncodeGRPCGetModuleConfigResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"SendCommand",
		ws.ProtoIDFromString("SCM"),
		endpoints.SendCommandEndpoint,
		DecodeWSSendCommandRequest,
		EncodeGRPCSendCommandResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"SetEnv",
		ws.ProtoIDFromString("SEV"),
		endpoints.SetEnvEndpoint,
		DecodeWSSetEnvRequest,
		EncodeGRPCSetEnvResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"GetEnv",
		ws.ProtoIDFromString("GEV"),
		endpoints.GetEnvEndpoint,
		DecodeWSGetEnvRequest,
		EncodeGRPCGetEnvResponse,
	))

	return service
}

// DecodeWSSetPublicKeyRequest is a websocket.DecodeRequestFunc that converts a
// WS SetPublicKey request to a module.proto-domain setpublickey request.
func DecodeWSSetPublicKeyRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.SetPublicKeyRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCSetPublicKeyRequest(ctx, req)
}

// DecodeWSRemoveFileRequest is a websocket.DecodeRequestFunc that converts a
// WS RemoveFile request to a module.proto-domain removefile request.
func DecodeWSRemoveFileRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.RemoveFileRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCRemoveFileRequest(ctx, req)
}

// DecodeWSRemoveDirectoryRequest is a websocket.DecodeRequestFunc that converts a
// WS RemoveDirectory request to a module.proto-domain removedirectory request.
func DecodeWSRemoveDirectoryRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.RemoveDirectoryRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCRemoveDirectoryRequest(ctx, req)
}

// DecodeWSGetFilesRequest is a websocket.DecodeRequestFunc that converts a
// WS GetFiles request to a module.proto-domain getfiles request.
func DecodeWSGetFilesRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.GetFilesRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCGetFilesRequest(ctx, req)
}

// DecodeWSGetFileRequest is a websocket.DecodeRequestFunc that converts a
// WS GetFile request to a module.proto-domain getfile request.
func DecodeWSGetFileRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.GetFileRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCGetFileRequest(ctx, req)
}

// DecodeWSUploadFileRequest is a websocket.DecodeRequestFunc that converts a
// WS UploadFile request to a module.proto-domain uploadfile request.
func DecodeWSUploadFileRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.UploadFileRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCUploadFileRequest(ctx, req)
}

// DecodeWSGetModuleConfigRequest is a websocket.DecodeRequestFunc that converts a
// WS GetModuleConfig request to a module.proto-domain getmoduleconfig request.
func DecodeWSGetModuleConfigRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.GetModuleConfigRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCGetModuleConfigRequest(ctx, req)
}

// DecodeWSSendCommandRequest is a websocket.DecodeRequestFunc that converts a
// WS SendCommand request to a module.proto-domain sendcommand request.
func DecodeWSSendCommandRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.SendCommandRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCSendCommandRequest(ctx, req)
}

// DecodeWSSetEnvRequest is a websocket.DecodeRequestFunc that converts a
// WS SetEnv request to a module.proto-domain setenv request.
func DecodeWSSetEnvRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.SetEnvRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCSetEnvRequest(ctx, req)
}

// DecodeWSGetEnvRequest is a websocket.DecodeRequestFunc that converts a
// WS GetEnv request to a module.proto-domain getenv request.
func DecodeWSGetEnvRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.GetEnvRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCGetEnvRequest(ctx, req)
}
