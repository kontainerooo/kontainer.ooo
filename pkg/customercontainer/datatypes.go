package customercontainer

import "github.com/kontainerooo/kontainer.ooo/pkg/kmi"

// ContainerConfig defines the configuration of a customer container
type ContainerConfig struct {
	ImageName string
}

// ContainerModule is the DB entry for each container module with its
// corresponding KMI
type ContainerModule struct {
	ImageID       string `gorm:"primary_key"`
	RefID         uint
	ContainerID   string
	ContainerName string
	KMI           kmi.KMI
}
