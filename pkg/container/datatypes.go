package container

import "github.com/kontainerooo/kontainer.ooo/pkg/kmi"

// Container is the DB entry for each container module with its
// corresponding KMI
type Container struct {
	RefID         uint
	ContainerID   string `gorm:"primary_key"`
	ContainerName string
	KMI           kmi.KMI
}
