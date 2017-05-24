// Package module is the module service that talks to dashboard templates
package module

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/container"
	"github.com/kontainerooo/kontainer.ooo/pkg/util"
)

// Service Template
type Service interface {
	// SetPublicKey sets a public key for ssh-ing into the container
	SetPublicKey(refID uint, containerName string, key string) error

	// RemoveFile removes a file from the customer-container-path
	RemoveFile(refID uint, containerName string, filename string) error

	// RemoveDirectory removes a directory from the customer-container-path
	RemoveDirectory(refID uint, containerName string, path string) error

	// GetFiles lists files from the customer-container-path
	GetFiles(refID uint, containerName string, path string) (map[string]string, error)

	// GetFile gets the contents of a file from the customer-container-path
	GetFile(refID uint, containerName string, path string) (string, error)

	// GetModuleConfig returns the configuration for the module
	GetModuleConfig(refID uint, containerName string, moduleName string) (abstraction.JSON, error)

	// SendCommand sends a command to the customer-container, env overrides environment variables
	// that are already globally defined in the container
	SendCommand(refID uint, containerName string, command string, env []string) (string, error)

	// SetEnv sets a permanent environment variable in the container
	SetEnv(refID uint, containerName string, key string, value string) error

	// GetEnv gets the value of a permanent environment varibale in the container
	GetEnv(refID uint, containerName string, key string) (string, error)
}

type dbAdapter interface {
	abstraction.DBAdapter
	AutoMigrate(...interface{}) error
}

type service struct {
	db        dbAdapter
	container *container.Endpoints
	logger    log.Logger
	config    util.ConfigFile
}

func (s *service) InitializeDatabases() error {
	return s.db.AutoMigrate()
}

func (s *service) makePath(refID uint, containerName string) (string, error) {
	cuPath := path.Join(s.config.CustomerPath, fmt.Sprintf("%d", refID))
	_, err := os.Stat(cuPath)
	if err != nil {
		return "", errors.New("customer does not exist")
	}

	coPath := path.Join(s.config.CustomerPath, fmt.Sprintf("%d", refID), containerName)
	_, err = os.Stat(coPath)
	if err != nil {
		return "", errors.New("container does not exist")
	}

	return coPath, nil
}

func (s *service) SetPublicKey(refID uint, containerName string, key string) error {
	coPath, err := s.makePath(refID, containerName)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path.Join(coPath, "ssh.pub"), []byte(key), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) RemoveFile(refID uint, containerName string, filename string) error {
	coPath, err := s.makePath(refID, containerName)
	if err != nil {
		return err
	}

	// TODO: improve path traversal mitigation (let people do a/../b)
	// and count the amount of back traversals
	coPath = path.Join(coPath, filename)
	coPath = strings.Replace(coPath, "../", "", -1)
	err = os.Remove(coPath)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) RemoveDirectory(refID uint, containerName string, dir string) error {
	coPath, err := s.makePath(refID, containerName)
	if err != nil {
		return err
	}

	if len(dir) < 1 || dir == "." {
		return errors.New("dir cannot be empty")
	}

	// TODO: improve path traversal mitigation (let people do a/../b)
	// and count the amount of back traversals
	coPath = path.Join(coPath, dir)
	coPath = strings.Replace(coPath, "../", "", -1)
	err = os.RemoveAll(coPath)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetFiles(refID uint, containerName string, dir string) (map[string]string, error) {
	flist := make(map[string]string)

	coPath, err := s.makePath(refID, containerName)
	if err != nil {
		return flist, err
	}

	if len(dir) < 1 {
		return flist, errors.New("dir cannot be empty")
	}

	// TODO: improve path traversal mitigation (let people do a/../b)
	// and count the amount of back traversals
	coPath = path.Join(coPath, dir)
	coPath = strings.Replace(coPath, "../", "", -1)

	files, _ := ioutil.ReadDir(coPath)
	for _, f := range files {
		if f.IsDir() {
			flist[f.Name()] = "dir"
		} else {
			flist[f.Name()] = "file"
		}
	}

	return flist, nil
}

func (s *service) GetFile(refID uint, containerName string, filepath string) (string, error) {
	coPath, err := s.makePath(refID, containerName)
	if err != nil {
		return "", err
	}

	// TODO: improve path traversal mitigation (let people do a/../b)
	// and count the amount of back traversals
	coPath = path.Join(coPath, filepath)
	coPath = strings.Replace(coPath, "../", "", -1)
	f, err := os.Stat(coPath)
	if err != nil {
		return "", err
	}

	if f.IsDir() {
		return "", errors.New("cannot get a directory")
	}

	content, err := ioutil.ReadFile(coPath)
	if err != nil {
		return "", err
	}

	return string(content), err
}

func (s *service) GetModuleConfig(refid uint, containerName string, moduleName string) (abstraction.JSON, error) {
	return abstraction.JSON{}, nil
}

func (s *service) SendCommand(refID uint, containerName string, command string, env []string) (string, error) {
	id, err := s.getContainerIDForName(refID, containerName)
	if err != nil {
		return "", err
	}

	res, err := s.container.ExecuteEndpoint(context.Background(), &container.ExecuteRequest{
		RefID: refID,
		ID:    id,
		CMD:   command,
	})
	if err != nil {
		return "", err
	}

	execRes, ok := res.(container.ExecuteResponse)
	if !ok {
		return "", errors.New("service returned unexpected response")
	}

	return execRes.Response, nil
}

func (s *service) SetEnv(refID uint, containerName string, key string, value string) error {
	id, err := s.getContainerIDForName(refID, containerName)
	if err != nil {
		return err
	}

	_, err = s.container.SetEnvEndpoint(context.Background(), &container.SetEnvRequest{
		RefID: refID,
		ID:    id,
		Key:   key,
		Value: value,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetEnv(refID uint, containerName string, key string) (string, error) {
	id, err := s.getContainerIDForName(refID, containerName)
	if err != nil {
		return "", err
	}

	res, err := s.container.GetEnvEndpoint(context.Background(), &container.GetEnvRequest{
		RefID: refID,
		ID:    id,
		Key:   key,
	})
	if err != nil {
		return "", err
	}

	val, ok := res.(container.GetEnvResponse)
	if !ok {
		return "", errors.New("service returned unexpected response")
	}

	return val.Value, nil
}

func (s *service) getContainerIDForName(refID uint, containerName string) (string, error) {
	res, err := s.container.IDForNameEndpoint(context.Background(), &container.IDForNameRequest{
		RefID: refID,
		Name:  containerName,
	})
	if err != nil {
		return "", err
	}

	cnt, ok := res.(container.IDForNameResponse)
	if !ok {
		return "", errors.New("service returned unexpected response")
	}

	return cnt.ID, nil
}

// NewService creates a new module service
func NewService(db dbAdapter, ce *container.Endpoints, l log.Logger) (Service, error) {
	conf, err := util.GetConfig()
	if err != nil {
		return &service{}, err
	}

	s := &service{
		db:        db,
		container: ce,
		logger:    l,
		config:    conf,
	}

	err = s.InitializeDatabases()
	if err != nil {
		return s, err
	}

	return s, nil
}
