// Package network handles container networks and interconnections
package network

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
	ExposePortToContainer(refid int, srcContainerID string, port uint32, destContainerID string) (int, error)

	// RemovePortFromContainer removes an exposed port from a container
	RemovePortFromContainer(refid int, srcContainerID string, port uint32, destContainerID string) error
}
