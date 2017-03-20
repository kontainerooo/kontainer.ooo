// Package network handles container networks and interconnections
package network

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/ttdennis/kontainer.io/pkg/abstraction"
)

// Service NetworkService
type Service interface {
	// CreateNetwork creates a new network for a given user and returns its ID and name
	CreateNetwork(refid int, cfg *Config) (name string, id string, err error)

	// RemoveNetwork removes a network with a given name
	RemoveNetworkByName(refid int, name string) error

	// AddContainerToNetwork joins a given container to a given network
	AddContainerToNetwork(refid int, name string, containerID string) error

	// RemoveContainerFromNetwork removes a container from a given network
	RemoveContainerFromNetwork(refid int, name string, containerID string) error

	// ExposePortToContainer exposes a port from one container to another
	ExposePortToContainer(refid int, srcContainerID string, port uint32, destContainerID string) error

	// RemovePortFromContainer removes an exposed port from a container
	RemovePortFromContainer(refid int, srcContainerID string, port uint32, destContainerID string) error
}

type dbAdapter interface {
	abstraction.DBAdapter
	AutoMigrate(...interface{}) error
	Where(interface{}, ...interface{}) error
	First(interface{}, ...interface{}) error
	Find(interface{}, ...interface{}) error
	Create(interface{}) error
	Delete(interface{}, ...interface{}) error
}

type service struct {
	dcli   abstraction.DCli
	db     dbAdapter
	logger log.Logger
}

func (s *service) InitializeDatabases() error {
	return s.db.AutoMigrate(&Networks{})
}

func (s *service) CreateNetwork(refid int, cfg *Config) (name string, id string, err error) {

	// Generate a 128 byte unique name
	b := make([]byte, 128)
	_, err = rand.Read(b)
	if err != nil {
		return "", "", err
	}
	name = base64.URLEncoding.EncodeToString(b)

	res, err := s.dcli.NetworkCreate(context.Background(), name, types.NetworkCreate{
		Driver: cfg.Driver,
	})
	if err != nil {
		return "", "", err
	}

	nw := Networks{
		UserID:      uint(refid),
		NetworkName: name,
		NetworkID:   res.ID,
	}

	err = s.db.Create(nw)
	if err != nil {
		return "", "", nil
	}

	return name, res.ID, nil
}

func (s *service) RemoveNetworkByName(refid int, name string) error {
	// TODO: implement
	return nil
}

func (s *service) AddContainerToNetwork(refid int, name string, containerID string) error {
	// TODO: implement
	return nil
}

func (s *service) RemoveContainerFromNetwork(refid int, name string, containerID string) error {
	// TODO: implement
	return nil
}

func (s *service) ExposePortToContainer(refid int, srcContainerID string, port uint32, destContainerID string) error {
	// TODO: implement
	return nil
}

func (s *service) RemovePortFromContainer(refid int, srcContainerID string, port uint32, destContainerID string) error {
	// TODO: implement
	return nil
}

// NewService creates a new network service
func NewService(dcli abstraction.DCli, db dbAdapter) (Service, error) {
	s := &service{
		dcli: dcli,
		db:   db,
	}

	err := s.InitializeDatabases()
	if err != nil {
		return s, err
	}

	return s, nil
}
