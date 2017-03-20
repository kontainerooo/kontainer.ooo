// Package network handles container networks and interconnections
package network

// Networks stores the networks belonging to a user
type Networks struct {
	UserID   uint
	Networks []ID
}

// ID stores network's names and their respective IDs
type ID struct {
	Name string
	ID   string
}

// Config describes configuration options for Networks
type Config struct {
	Driver string
}
