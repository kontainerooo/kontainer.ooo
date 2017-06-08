package container

import (
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
)

// Container is the DB entry for each container module with its
// corresponding KMI
type Container struct {
	RefID         uint
	ContainerID   string `gorm:"primary_key"`
	ContainerName string
	KMIID         uint
	KMI           CKMI
}

// CKMI is the database representation for the kmi of a specific instance
type CKMI struct {
	kmi.KMI
	Links abstraction.JSON `sql:"type:jsonb"`
}

// TableName sets the database name for ckmi
func (CKMI) TableName() string {
	return "container_kmis"
}
