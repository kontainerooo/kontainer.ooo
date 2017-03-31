// Package network handles container networks and interconnections
package network

import "github.com/kontainerooo/kontainer.ooo/pkg/abstraction"

// Networks stores the networks belonging to a user
type Networks struct {
	UserID      uint
	NetworkID   string `gorm:"primary_key"`
	NetworkName string
	IsPrimary   bool
}

// Containers map containers to their networks
type Containers struct {
	NetworkID   string
	ContainerID string
	ContainerIP abstraction.Inet `gorm:"primary_key"`
}

// Config describes configuration options for Networks
type Config struct {
	Name   string
	Driver string
}

// exposeData is the data needed for an expose rule
type exposeData struct {
	Port       uint16
	SrcIP      abstraction.Inet
	DstIP      abstraction.Inet
	DstNetwork string
	SrcNetwork string
	Protocol   string
}
