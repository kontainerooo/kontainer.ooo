// Package template is the template service that talks to dashboard templates
package template

import (
	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/user"
)

const (
	// TemplateFileTypeText represents a text file type
	TemplateFileTypeText = "txt"

	// TemplateFileTypeBinary represents a binary file type
	TemplateFileTypeBinary = "bin"
)

// Service Template
type Service interface {
	// SetPublicKey sets a public key for ssh-ing into the container
	SetPublicKey(refid uint, containerName string, key string) error

	// RemoveFile removes a file from the customer-container-path
	RemoveFile(refid uint, containerName string, filename string) error

	// RemoveDirectory removes a directory from the customer-container-path
	RemoveDirectory(refid uint, containerName string, path string) error

	// GetFiles lists files from the customer-container-path
	GetFiles(refid uint, containerName string, path string) ([]string, error)

	// GetFile gets the contents of a file from the customer-container-path
	GetFile(refid uint, containerName string, path string) (string, string, error)

	// GetModuleConfig returns the configuration for the module
	GetModuleConfig(refid uint, containerName string, moduleName string) (abstraction.JSON, error)

	// SendCommand sends a command to the customer-container
	SendCommand(refid uint, containerName string, command string) (string, error)
}

type dbAdapter interface {
	abstraction.DBAdapter
	AutoMigrate(...interface{}) error
}

type service struct {
	db         dbAdapter
	userClient *user.Endpoints
	logger     log.Logger
}

func (s *service) InitializeDatabases() error {
	return s.db.AutoMigrate()
}

func (s *service) SetPublicKey(refid uint, containerName string, key string) error {
	return nil
}

func (s *service) RemoveFile(refid uint, containerName string, filename string) error {
	return nil
}

func (s *service) RemoveDirectory(refid uint, containerName string, path string) error {
	return nil
}

func (s *service) GetFiles(refid uint, containerName string, path string) ([]string, error) {
	return []string{}, nil
}

func (s *service) GetFile(refid uint, containerName string, path string) (string, string, error) {
	return "", "", nil
}

func (s *service) GetModuleConfig(refid uint, containerName string, moduleName string) (abstraction.JSON, error) {
	return abstraction.JSON{}, nil
}

func (s *service) SendCommand(refid uint, containerName string, command string) (string, error) {
	return "", nil
}

// NewService creates a new template service
func NewService(db dbAdapter, uc *user.Endpoints, l log.Logger) (Service, error) {
	s := &service{
		db:         db,
		userClient: uc,
		logger:     l,
	}

	err := s.InitializeDatabases()
	if err != nil {
		return s, err
	}

	return s, nil
}
