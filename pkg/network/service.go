// Package network handles container networks and interconnections
package network

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	networkTypes "github.com/docker/docker/api/types/network"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/firewall"
)

var (
	// ErrNetworkNotExist occurs when a network does not exist
	ErrNetworkNotExist = errors.New("Network does not exist")

	// ErrNetworkAlreadyExists occurs when a network already exists
	ErrNetworkAlreadyExists = errors.New("Network already exists")
)

// Service NetworkService
type Service interface {
	// CreatePrimaryNetworkForContainer creates a network and assigns it as primary to the given container
	CreatePrimaryNetworkForContainer(refid uint, cfg *Config, containerID string) error

	// CreateNetwork creates a new network for a given user
	CreateNetwork(refid uint, cfg *Config) error

	// RemoveNetwork removes a network with a given name
	RemoveNetworkByName(refid uint, name string) error

	// AddContainerToNetwork joins a given container to a given network
	AddContainerToNetwork(refid uint, name string, containerID string) error

	// RemoveContainerFromNetwork removes a container from a given network
	RemoveContainerFromNetwork(refid uint, name string, containerID string) error

	// ExposePortToContainer exposes a port from one container to another
	ExposePortToContainer(refid uint, srcContainerID string, port uint16, protocol string, destContainerID string) error

	// RemovePortFromContainer removes an exposed port from a container
	RemovePortFromContainer(refid uint, srcContainerID string, port uint16, protocol string, destContainerID string) error
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
	dcli     abstraction.DCli
	db       dbAdapter
	fwClient *firewall.Endpoints
	logger   log.Logger
}

func (s *service) InitializeDatabases() error {
	return s.db.AutoMigrate(&Networks{}, &Containers{})
}

func (s *service) getNetworkByName(refid uint, name string) (Networks, error) {
	nw := Networks{}

	err := s.db.Where("user_id = ? AND network_name = ?", refid, name)
	if err != nil {
		return nw, err
	}

	s.db.First(&nw)

	return nw, nil
}

func (s *service) createNetwork(refid uint, cfg *Config, isPrimary bool) error {
	name := cfg.Name

	nw, err := s.getNetworkByName(refid, name)
	if err != nil {
		return err
	}

	if nw.NetworkID != "" {
		return ErrNetworkAlreadyExists
	}

	res, err := s.dcli.NetworkCreate(context.Background(), fmt.Sprintf("%s-%s", string(refid), name), types.NetworkCreate{
		Driver: cfg.Driver,
	})
	if err != nil {
		return err
	}

	nw = Networks{
		UserID:      uint(refid),
		NetworkName: name,
		NetworkID:   res.ID,
		IsPrimary:   isPrimary,
	}

	err = s.db.Create(&nw)
	if err != nil {
		// Try to remove the actual network on db error
		s.dcli.NetworkRemove(context.Background(), res.ID)
		return err
	}

	return nil
}

func (s *service) CreateNetwork(refid uint, cfg *Config) error {
	return s.createNetwork(refid, cfg, false)
}

func (s *service) CreatePrimaryNetworkForContainer(refid uint, cfg *Config, containerID string) error {
	err := s.createNetwork(refid, cfg, true)
	if err != nil {
		return err
	}

	err = s.AddContainerToNetwork(refid, cfg.Name, containerID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) RemoveNetworkByName(refid uint, name string) error {
	nw, err := s.getNetworkByName(refid, name)
	if err != nil {
		return err
	}

	if nw.NetworkID != "" {
		s.db.Begin()
		err = s.dcli.NetworkRemove(context.Background(), nw.NetworkID)
		if err != nil {
			return err
		}

		err = s.db.Delete(&nw)
		if err != nil {
			s.db.Rollback()
			return err
		}

		cts := &Containers{
			NetworkID: nw.NetworkID,
		}
		err = s.db.Delete(cts)
		if err != nil {
			s.db.Rollback()
			return err
		}
		s.db.Commit()

		return nil
	}
	return ErrNetworkNotExist
}

func (s *service) AddContainerToNetwork(refid uint, name string, containerID string) error {
	nw, err := s.getNetworkByName(refid, name)
	if err != nil {
		return err
	}

	// Check that this network is not someone else's primary network
	if s.isPrimary(nw.NetworkID) {
		cts := s.getContainerForPrimaryNetwork(nw.NetworkID)
		if cts.ContainerID != "" {
			return errors.New("Network is primary for other container")
		}
	}

	if nw.NetworkID != "" {
		err = s.dcli.NetworkConnect(context.Background(), nw.NetworkID, containerID, &networkTypes.EndpointSettings{})
		if err != nil {
			return err
		}

		// Check which IP address we got
		information, err := s.dcli.NetworkInspect(context.Background(), nw.NetworkID, false)
		if err != nil {
			return err
		}

		c, ok := information.Containers[containerID]
		if !ok {
			return errors.New("Container was not connected to network")
		}

		ip, err := abstraction.NewInet(c.IPv4Address)
		if err != nil {
			return err
		}

		s.db.Begin()
		cts := &Containers{
			ContainerID: containerID,
			NetworkID:   nw.NetworkID,
			ContainerIP: ip,
		}

		err = s.db.Create(cts)
		if err != nil {
			s.db.Rollback()
			return err
		}
		s.db.Commit()

	} else {
		return ErrNetworkNotExist
	}

	return nil
}

func (s *service) RemoveContainerFromNetwork(refid uint, name string, containerID string) error {
	nw, err := s.getNetworkByName(refid, name)
	if err != nil {
		return err
	}
	if nw.NetworkID != "" {
		err = s.dcli.NetworkDisconnect(context.Background(), nw.NetworkID, containerID, true)
		if err != nil {
			return err
		}
	} else {
		return ErrNetworkNotExist
	}

	return nil
}

func (s *service) getContainerNetworks(containerID string) ([]Containers, error) {
	cts := []Containers{}
	err := s.db.Where("container_id = ?", containerID)
	if err != nil {
		return []Containers{}, err
	}

	err = s.db.Find(&cts)
	if err != nil {
		return []Containers{}, err
	}

	// TEMPORARY: until mockDB is fixed
	for i, v := range cts {
		if v.ContainerID != containerID {
			cts = append(cts[:i], cts[i+1:]...)
		}
	}

	return cts, nil
}

func (s *service) isPrimary(networkID string) bool {
	err := s.db.Where("network_id = ? AND is_primary = ?", networkID, true)
	if err != nil {
		return false
	}

	if s.db.GetValue() != nil {
		return true
	}

	return false
}

func (s *service) getPrimaryNetworkForContainer(containerID string) Networks {
	cts, err := s.getContainerNetworks(containerID)
	if err != nil {
		return Networks{}
	}

	for _, v := range cts {
		if s.isPrimary(v.NetworkID) {
			s.db.Begin()
			var nw Networks
			s.db.Where("network_id = ?", v.NetworkID)
			s.db.First(&nw)
			s.db.Commit()

			if nw != (Networks{}) {
				return nw
			}
		}
	}

	return Networks{}
}

func (s *service) getContainerForPrimaryNetwork(networkID string) Containers {
	err := s.db.Where("network_id = ?", networkID)
	if err != nil {
		return Containers{}
	}

	cts := Containers{}
	err = s.db.First(&cts)
	if err != nil {
		return Containers{}
	}

	return cts
}

func (s *service) getContainerIPInNetwork(containerID string, networkID string) (abstraction.Inet, error) {
	s.db.Begin()

	err := s.db.Where("container_id = ? AND network_id = ?", containerID, networkID)
	if err != nil {
		return abstraction.Inet(""), err
	}

	cts := &Containers{}
	err = s.db.First(cts)
	if err != nil {
		s.db.Rollback()
		return abstraction.Inet(""), err
	}
	s.db.Commit()

	if cts.ContainerIP != abstraction.Inet("") {
		return cts.ContainerIP, nil
	}

	return abstraction.Inet(""), errors.New("Container is not in network")
}

func (s *service) getExposeData(refid uint, srcContainerID string, port uint16, protocol string, destContainerID string) (exposeData, error) {
	// Check if the containers are in a same network
	srcNetworks, err := s.getContainerNetworks(srcContainerID)
	if err != nil {
		return exposeData{0, abstraction.Inet(""), abstraction.Inet(""), "", "", ""}, err
	}

	dstNetworks, err := s.getContainerNetworks(destContainerID)
	if err != nil {
		return exposeData{0, abstraction.Inet(""), abstraction.Inet(""), "", "", ""}, err
	}

	for _, srcV := range srcNetworks {
		for _, dstV := range dstNetworks {
			if srcV.NetworkID == dstV.NetworkID {
				return exposeData{0, abstraction.Inet(""), abstraction.Inet(""), "", "", ""}, errors.New("Containers are already in the same network")
			}
		}
	}

	srcPrimaryNetwork := s.getPrimaryNetworkForContainer(srcContainerID)
	dstPrimaryNetwork := s.getPrimaryNetworkForContainer(destContainerID)

	if srcPrimaryNetwork == (Networks{}) || dstPrimaryNetwork == (Networks{}) {
		return exposeData{0, abstraction.Inet(""), abstraction.Inet(""), "", "", ""}, errors.New("Both containers must have a primary network")
	}

	srcIP, err := s.getContainerIPInNetwork(srcContainerID, srcPrimaryNetwork.NetworkID)
	if err != nil {
		return exposeData{0, abstraction.Inet(""), abstraction.Inet(""), "", "", ""}, err
	}

	dstIP, err := s.getContainerIPInNetwork(destContainerID, dstPrimaryNetwork.NetworkID)
	if err != nil {
		return exposeData{0, abstraction.Inet(""), abstraction.Inet(""), "", "", ""}, err
	}

	return exposeData{
			port,
			srcIP,
			dstIP,
			dstPrimaryNetwork.NetworkID,
			srcPrimaryNetwork.NetworkID,
			protocol,
		},
		nil
}

func (s *service) ExposePortToContainer(refid uint, srcContainerID string, port uint16, protocol string, destContainerID string) error {
	exp, err := s.getExposeData(refid, srcContainerID, port, protocol, destContainerID)
	if err != nil {
		return err
	}

	_, err = s.fwClient.AllowPortEndpoint(context.Background(), &firewall.AllowPortRequest{
		Port:       exp.Port,
		SrcIP:      exp.SrcIP,
		DstIP:      exp.DstIP,
		DstNetwork: exp.DstNetwork,
		SrcNetwork: exp.SrcNetwork,
		Protocol:   exp.Protocol,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) RemovePortFromContainer(refid uint, srcContainerID string, port uint16, protocol string, destContainerID string) error {
	exp, err := s.getExposeData(refid, srcContainerID, port, protocol, destContainerID)
	if err != nil {
		return err
	}

	_, err = s.fwClient.BlockPortEndpoint(context.Background(), &firewall.BlockPortRequest{
		Port:       exp.Port,
		SrcIP:      exp.SrcIP,
		DstIP:      exp.DstIP,
		DstNetwork: exp.DstNetwork,
		SrcNetwork: exp.SrcNetwork,
		Protocol:   exp.Protocol,
	})
	if err != nil {
		return err
	}

	return nil
}

// NewService creates a new network service
func NewService(dcli abstraction.DCli, db dbAdapter, fw *firewall.Endpoints) (Service, error) {
	s := &service{
		dcli:     dcli,
		db:       db,
		fwClient: fw,
	}

	err := s.InitializeDatabases()
	if err != nil {
		return s, err
	}

	return s, nil
}
