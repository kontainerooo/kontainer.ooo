// Package module is the module service that talks to dashboard templates
package module

import "github.com/kontainerooo/kontainer.ooo/pkg/kmi"

// Module contains information about a container module
type Module struct {
	ContainerName string
	KMDI          kmi.KMDI
}
