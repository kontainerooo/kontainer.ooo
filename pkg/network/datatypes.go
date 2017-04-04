// Package network handles container networks and interconnections
package network

// Networks stores the networks belonging to a user
type Networks struct {
	UserID      uint
	NetworkID   string `gorm:"primary_key"`
	NetworkName string
}

// Config describes configuration options for Networks
type Config struct {
	Name   string
	Driver string
}
