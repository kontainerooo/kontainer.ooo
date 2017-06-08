package module

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
)

// Endpoints is a struct which collects all endpoints for the module service
type Endpoints struct {
	CreateContainerModuleEndpoint endpoint.Endpoint
	SetPublicKeyEndpoint          endpoint.Endpoint
	RemoveFileEndpoint            endpoint.Endpoint
	RemoveDirectoryEndpoint       endpoint.Endpoint
	GetFilesEndpoint              endpoint.Endpoint
	GetFileEndpoint               endpoint.Endpoint
	UploadFileEndpoint            endpoint.Endpoint
	GetModuleConfigEndpoint       endpoint.Endpoint
	SendCommandEndpoint           endpoint.Endpoint
	SetEnvEndpoint                endpoint.Endpoint
	GetEnvEndpoint                endpoint.Endpoint
	SetLinkEndpoint               endpoint.Endpoint
	RemoveLinkEndpoint            endpoint.Endpoint
	GetModulesEndpoint            endpoint.Endpoint
}

// CreateContainerModuleRequest is the request struct for the CreateContainerModuleEndpoint
type CreateContainerModuleRequest struct {
	RefID uint
	KmiID uint
	Name  string
}

// CreateContainerModuleResponse is the response struct for the CreateContainerModuleEndpoint
type CreateContainerModuleResponse struct {
	Error error
}

// MakeCreateContainerModuleEndpoint creates a gokit endpoint which invokes CreateContainerModule
func MakeCreateContainerModuleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateContainerModuleRequest)
		err := s.CreateContainerModule(req.RefID, req.KmiID, req.Name)
		return CreateContainerModuleResponse{
			Error: err,
		}, nil
	}
}

// SetPublicKeyRequest is the request struct for the SetPublicKeyEndpoint
type SetPublicKeyRequest struct {
	RefID         uint `bart:"ref"`
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
	RefID         uint `bart:"ref"`
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
	RefID         uint `bart:"ref"`
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
	RefID         uint `bart:"ref"`
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
	RefID         uint `bart:"ref"`
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
	RefID         uint `bart:"ref"`
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
	RefID         uint `bart:"ref"`
	ContainerName string
}

// GetModuleConfigResponse is the response struct for the GetModuleConfigEndpoint
type GetModuleConfigResponse struct {
	ContainerKMI kmi.KMI
	Links        map[string][]string
	Error        error
}

// MakeGetModuleConfigEndpoint creates a gokit endpoint which invokes GetModuleConfig
func MakeGetModuleConfigEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetModuleConfigRequest)
		cKMI, links, err := s.GetModuleConfig(req.RefID, req.ContainerName)
		return GetModuleConfigResponse{
			ContainerKMI: cKMI,
			Links:        links,
			Error:        err,
		}, nil
	}
}

// SendCommandRequest is the request struct for the SendCommandEndpoint
type SendCommandRequest struct {
	RefID         uint `bart:"ref"`
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
	RefID         uint `bart:"ref"`
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
	RefID         uint `bart:"ref"`
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

// SetLinkRequest is the request struct for the SetLinkEndpoint
type SetLinkRequest struct {
	RefID         uint
	ContainerName string
	LinkName      string
	LinkInterface string
}

// SetLinkResponse is the response struct for the SetLinkEndpoint
type SetLinkResponse struct {
	Error error
}

// MakeSetLinkEndpoint creates a gokit endpoint which invokes SetLink
func MakeSetLinkEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetLinkRequest)
		err := s.SetLink(req.RefID, req.ContainerName, req.LinkName, req.LinkInterface)
		return SetLinkResponse{
			Error: err,
		}, nil
	}
}

// RemoveLinkRequest is the request struct for the RemoveLinkEndpoint
type RemoveLinkRequest struct {
	RefID         uint
	ContainerName string
	LinkName      string
	LinkInterface string
}

// RemoveLinkResponse is the response struct for the RemoveLinkEndpoint
type RemoveLinkResponse struct {
	Error error
}

// MakeRemoveLinkEndpoint creates a gokit endpoint which invokes RemoveLink
func MakeRemoveLinkEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveLinkRequest)
		err := s.RemoveLink(req.RefID, req.ContainerName, req.LinkName, req.LinkInterface)
		return RemoveLinkResponse{
			Error: err,
		}, nil
	}
}

// GetModulesRequest is the request struct for the GetModulesEndpoint
type GetModulesRequest struct {
	RefID uint
}

// GetModulesResponse is the response struct for the GetModulesEndpoint
type GetModulesResponse struct {
	Modules []Module
	Error   error
}

// MakeGetModulesEndpoint creates a gokit endpoint which invokes GetModules
func MakeGetModulesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetModulesRequest)
		mods, err := s.GetModules(req.RefID)
		return GetModulesResponse{
			Modules: mods,
			Error:   err,
		}, nil
	}
}
