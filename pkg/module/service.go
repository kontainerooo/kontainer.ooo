// Package module is the module service that talks to dashboard templates
package module

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/container"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
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
	GetFile(refID uint, containerName string, path string) ([]byte, error)

	// UploadFile uploads a file in a given container to a given path
	UploadFile(refID uint, containerName string, filepath string, content []byte, override bool) error

	// GetModuleConfig returns the configuration for the module
	GetModuleConfig(refID uint, containerName string) (kmi.KMI, error)

	// SendCommand sends a command to the customer-container, env overrides environment variables
	// that are already globally defined in the container
	SendCommand(refID uint, containerName string, command string, env map[string]string) (string, error)

	// SetEnv sets a permanent environment variable in the container
	SetEnv(refID uint, containerName string, key string, value string) error

	// GetEnv gets the value of a permanent environment varibale in the container
	GetEnv(refID uint, containerName string, key string) (string, error)
}

type service struct {
	container *container.Endpoints
	logger    log.Logger
	config    util.ConfigFile
	mtx       *sync.Mutex
}

func (s *service) makePath(refID uint, containerName string) (string, error) {
	cuPath := path.Join(s.config.CustomerPath, fmt.Sprintf("%d", refID))
	_, err := os.Stat(cuPath)
	if err != nil {
		return "", errors.New("customer does not exist")
	}

	containerID, err := s.getContainerIDForName(refID, containerName)
	if err != nil {
		return "", err
	}

	coPath := path.Join(s.config.CustomerPath, fmt.Sprintf("%d", refID), containerID, "rootfs")
	_, err = os.Stat(coPath)
	if err != nil {
		return "", errors.New("container does not exist")
	}

	return coPath, nil
}

func (s *service) SetPublicKey(refID uint, containerName string, key string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.setPublicKey(refID, containerName, key)
}

func (s *service) setPublicKey(refID uint, containerName string, key string) error {
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
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.removeFile(refID, containerName, filename)
}

func (s *service) removeFile(refID uint, containerName string, filename string) error {
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
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.removeDirectory(refID, containerName, dir)
}

func (s *service) removeDirectory(refID uint, containerName string, dir string) error {
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
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.getFiles(refID, containerName, dir)
}

func (s *service) getFiles(refID uint, containerName string, dir string) (map[string]string, error) {
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

func (s *service) GetFile(refID uint, containerName string, filepath string) ([]byte, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.getFile(refID, containerName, filepath)
}

func (s *service) getFile(refID uint, containerName string, filepath string) ([]byte, error) {
	coPath, err := s.makePath(refID, containerName)
	if err != nil {
		return []byte{}, err
	}

	// TODO: improve path traversal mitigation (let people do a/../b)
	// and count the amount of back traversals
	coPath = path.Join(coPath, filepath)
	coPath = strings.Replace(coPath, "../", "", -1)
	f, err := os.Stat(coPath)
	if err != nil {
		return []byte{}, err
	}

	if f.IsDir() {
		return []byte{}, errors.New("cannot get a directory")
	}

	content, err := ioutil.ReadFile(coPath)
	if err != nil {
		return []byte{}, err
	}

	return content, err
}

func (s *service) UploadFile(refID uint, containerName string, fpath string, content []byte, override bool) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.uploadFile(refID, containerName, fpath, content, override)
}

func (s *service) uploadFile(refID uint, containerName string, fpath string, content []byte, override bool) error {
	coPath, err := s.makePath(refID, containerName)
	if err != nil {
		return err
	}

	// TODO: improve path traversal mitigation (let people do a/../b)
	// and count the amount of back traversals
	coPath = path.Join(coPath, fpath)
	coPath = strings.Replace(coPath, "../", "", -1)

	// If the file exists only replace if flag is set
	f, err := os.Stat(coPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Make sure the directory exists
			err = os.MkdirAll(filepath.Dir(coPath), os.ModeDir)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(coPath, content, 0644)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	if f.IsDir() {
		return fmt.Errorf("%s is a directory", fpath)
	}

	if override {
		err = ioutil.WriteFile(coPath, content, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) GetModuleConfig(refID uint, containerName string) (kmi.KMI, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.getModuleConfig(refID, containerName)
}

func (s *service) getModuleConfig(refID uint, containerName string) (kmi.KMI, error) {
	id, err := s.getContainerIDForName(refID, containerName)
	if err != nil {
		return kmi.KMI{}, err
	}

	res, err := s.container.GetContainerKMIEndpoint(context.Background(), container.GetContainerKMIRequest{
		ContainerID: id,
	})
	if err != nil {
		return kmi.KMI{}, err
	}

	containerKMI, ok := res.(container.GetContainerKMIResponse)
	if !ok {
		return kmi.KMI{}, errors.New("service returned unexpected response")
	}

	return containerKMI.ContainerKMI, nil
}

func (s *service) SendCommand(refID uint, containerName string, command string, env map[string]string) (string, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.sendCommand(refID, containerName, command, env)
}

func (s *service) sendCommand(refID uint, containerName string, command string, env map[string]string) (string, error) {
	id, err := s.getContainerIDForName(refID, containerName)
	if err != nil {
		return "", err
	}

	res, err := s.container.GetContainerKMIEndpoint(context.Background(), container.GetContainerKMIRequest{
		ContainerID: id,
	})
	if err != nil {
		return "", err
	}

	containerKMI, ok := res.(container.GetContainerKMIResponse)
	if !ok {
		return "", errors.New("service returned unexpected response")
	}

	// Check if the command exists
	cmdMap := containerKMI.ContainerKMI.Commands.ToStringMap()
	cmdString, ok := cmdMap[command]
	if !ok {
		return "", fmt.Errorf("Command %s not found", command)
	}

	res, err = s.container.ExecuteEndpoint(context.Background(), container.ExecuteRequest{
		RefID: refID,
		ID:    id,
		CMD:   cmdString,
		Env:   env,
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
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.setEnv(refID, containerName, key, value)
}

func (s *service) setEnv(refID uint, containerName string, key string, value string) error {
	id, err := s.getContainerIDForName(refID, containerName)
	if err != nil {
		return err
	}

	_, err = s.container.SetEnvEndpoint(context.Background(), container.SetEnvRequest{
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
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.getEnv(refID, containerName, key)
}

func (s *service) getEnv(refID uint, containerName string, key string) (string, error) {
	id, err := s.getContainerIDForName(refID, containerName)
	if err != nil {
		return "", err
	}

	res, err := s.container.GetEnvEndpoint(context.Background(), container.GetEnvRequest{
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
	res, err := s.container.IDForNameEndpoint(context.Background(), container.IDForNameRequest{
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
func NewService(ce *container.Endpoints, l log.Logger) (Service, error) {
	conf, err := util.GetConfig()
	if err != nil {
		return &service{}, err
	}

	s := &service{
		container: ce,
		logger:    l,
		config:    conf,
		mtx:       &sync.Mutex{},
	}

	return s, nil
}
