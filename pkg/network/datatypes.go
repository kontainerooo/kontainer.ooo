// Package network handles container networks and interconnections
package network

// Networks stores the networks belonging to a user
type Networks struct {
	UserID    uint
	Name      string
	NetworkID string
}

// Config describes configuration options for Networks
type Config struct {
	Driver string
}
