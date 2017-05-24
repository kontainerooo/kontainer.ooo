package module

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
)

// Endpoints is a struct which collects all endpoints for the module service
type Endpoints struct {
	SetPublicKeyEndpoint    endpoint.Endpoint
	RemoveFileEndpoint      endpoint.Endpoint
	RemoveDirectoryEndpoint endpoint.Endpoint
	GetFilesEndpoint        endpoint.Endpoint
	GetFileEndpoint         endpoint.Endpoint
	UploadFileEndpoint      endpoint.Endpoint
	GetModuleConfigEndpoint endpoint.Endpoint
	SendCommandEndpoint     endpoint.Endpoint
	SetEnvEndpoint          endpoint.Endpoint
	GetEnvEndpoint          endpoint.Endpoint
}

// SetPublicKeyRequest is the request struct for the SetPublicKeyEndpoint
type SetPublicKeyRequest struct {
	RefID         uint
	ContainerName string
	Key           string
}

// SetPublicKeyResponse is the response struct for the SetPublicKeyEndpoint
type SetPublicKeyResponse struct {
	Error error
}

// MakeSetPublicKeyEndpoint creates a gokit endpoint which invokes SetPublicKey
func MakeSetPublicKeyEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetPublicKeyRequest)
		err := s.SetPublicKey(req.RefID, req.ContainerName, req.Key)
		return SetPublicKeyResponse{
			Error: err,
		}, nil
	}
}

// RemoveFileRequest is the request struct for the RemoveFileEndpoint
type RemoveFileRequest struct {
	RefID         uint
	ContainerName string
	Filename      string
}

// RemoveFileResponse is the response struct for the RemoveFileEndpoint
type RemoveFileResponse struct {
	Error error
}

// MakeRemoveFileEndpoint creates a gokit endpoint which invokes RemoveFile
func MakeRemoveFileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveFileRequest)
		err := s.RemoveFile(req.RefID, req.ContainerName, req.Filename)
		return RemoveFileResponse{
			Error: err,
		}, nil
	}
}

// RemoveDirectoryRequest is the request struct for the RemoveDirectoryEndpoint
type RemoveDirectoryRequest struct {
	RefID         uint
	ContainerName string
	Path          string
}

// RemoveDirectoryResponse is the response struct for the RemoveDirectoryEndpoint
type RemoveDirectoryResponse struct {
	Error error
}

// MakeRemoveDirectoryEndpoint creates a gokit endpoint which invokes RemoveDirectory
func MakeRemoveDirectoryEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveDirectoryRequest)
		err := s.RemoveDirectory(req.RefID, req.ContainerName, req.Path)
		return RemoveDirectoryResponse{
			Error: err,
		}, nil
	}
}

// GetFilesRequest is the request struct for the GetFilesEndpoint
type GetFilesRequest struct {
	RefID         uint
	ContainerName string
	Path          string
}

// GetFilesResponse is the response struct for the GetFilesEndpoint
type GetFilesResponse struct {
	Files map[string]string
	Error error
}

// MakeGetFilesEndpoint creates a gokit endpoint which invokes GetFiles
func MakeGetFilesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFilesRequest)
		files, err := s.GetFiles(req.RefID, req.ContainerName, req.Path)
		return GetFilesResponse{
			Files: files,
			Error: err,
		}, nil
	}
}

// GetFileRequest is the request struct for the GetFileEndpoint
type GetFileRequest struct {
	RefID         uint
	ContainerName string
	Path          string
}

// GetFileResponse is the response struct for the GetFileEndpoint
type GetFileResponse struct {
	Content []byte
	Error   error
}

// MakeGetFileEndpoint creates a gokit endpoint which invokes GetFile
func MakeGetFileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFileRequest)
		content, err := s.GetFile(req.RefID, req.ContainerName, req.Path)
		return GetFileResponse{
			Content: content,
			Error:   err,
		}, nil
	}
}

// UploadFileRequest is the request struct for the UploadFileEndpoint
type UploadFileRequest struct {
	RefID         uint
	ContainerName string
	Path          string
	Content       []byte
	Override      bool
}

// UploadFileResponse is the response struct for the UploadFileEndpoint
type UploadFileResponse struct {
	Error error
}

// MakeUploadFileEndpoint creates a gokit endpoint which invokes UploadFile
func MakeUploadFileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UploadFileRequest)
		err := s.UploadFile(req.RefID, req.ContainerName, req.Path, req.Content, req.Override)
		return UploadFileResponse{
			Error: err,
		}, nil
	}
}

// GetModuleConfigRequest is the request struct for the GetModuleConfigEndpoint
type GetModuleConfigRequest struct {
	RefID         uint
	ContainerName string
}

// GetModuleConfigResponse is the response struct for the GetModuleConfigEndpoint
type GetModuleConfigResponse struct {
	ContainerKMI kmi.KMI
	Error        error
}

// MakeGetModuleConfigEndpoint creates a gokit endpoint which invokes GetModuleConfig
func MakeGetModuleConfigEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetModuleConfigRequest)
		cKMI, err := s.GetModuleConfig(req.RefID, req.ContainerName)
		return GetModuleConfigResponse{
			ContainerKMI: cKMI,
			Error:        err,
		}, nil
	}
}

// SendCommandRequest is the request struct for the SendCommandEndpoint
type SendCommandRequest struct {
	RefID         uint
	ContainerName string
	Command       string
	Env           map[string]string
}

// SendCommandResponse is the response struct for the SendCommandEndpoint
type SendCommandResponse struct {
	Response string
	Error    error
}

// MakeSendCommandEndpoint creates a gokit endpoint which invokes SendCommand
func MakeSendCommandEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendCommandRequest)
		res, err := s.SendCommand(req.RefID, req.ContainerName, req.Command, req.Env)
		return SendCommandResponse{
			Response: res,
			Error:    err,
		}, nil
	}
}

// SetEnvRequest is the request struct for the SetEnvEndpoint
type SetEnvRequest struct {
	RefID         uint
	ContainerName string
	Key           string
	Value         string
}

// SetEnvResponse is the response struct for the SetEnvEndpoint
type SetEnvResponse struct {
	Error error
}

// MakeSetEnvEndpoint creates a gokit endpoint which invokes SetEnv
func MakeSetEnvEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetEnvRequest)
		err := s.SetEnv(req.RefID, req.ContainerName, req.Key, req.Value)
		return SetEnvResponse{
			Error: err,
		}, nil
	}
}

// GetEnvRequest is the request struct for the GetEnvEndpoint
type GetEnvRequest struct {
	RefID         uint
	ContainerName string
	Key           string
}

// GetEnvResponse is the response struct for the GetEnvEndpoint
type GetEnvResponse struct {
	Value string
	Error error
}

// MakeGetEnvEndpoint creates a gokit endpoint which invokes GetEnv
func MakeGetEnvEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetEnvRequest)
		val, err := s.GetEnv(req.RefID, req.ContainerName, req.Key)
		return GetEnvResponse{
			Value: val,
			Error: err,
		}, nil
	}
}
