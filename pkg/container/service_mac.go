// +build !linux

package container

import "github.com/kontainerooo/kontainer.ooo/pkg/kmi"

// Service Container Service
type Service interface {
	// CreateContainer instanciates a container for a User with the id refID and returns its id
	CreateContainer(refID uint, kmiID uint, name string) (id string, err error)

	// RemoveContainer is used to remove a container instance by id
	RemoveContainer(refID uint, id string) error

	// Instances returns a list of container instances of a user by id
	Instances(refID uint) []Container

	// StopContainer stops a container
	StopContainer(refID uint, id string) error

	// Execute executes a command in a given container
	Execute(refID uint, id string, cmd string, env map[string]string) (string, error)

	// GetEnv returns the value to a given environment variable setting. Returns the whole
	// environment as string if key is empty
	GetEnv(refID uint, id string, key string) (string, error)

	// SetEnv sets an environment variable for the container
	SetEnv(refID uint, id string, key string, value string) error

	// IDForName returs the containerID for a given container name and user
	IDForName(refID uint, name string) (string, error)

	// GetContainerKMI returns the KMI for a given container
	GetContainerKMI(containerID string) (kmi.KMI, error)
}
