// Package iptables is a wrapper around the iptables binary
package iptables

// Rule represents a rule entry in iptables
type Rule struct {
	Target          string
	Chain           string
	Protocol        string
	In              string
	Out             string
	Source          string
	Destination     string
	SourcePort      uint16
	DestinationPort uint16
	State           string
}

type iptablesEntry struct {
	ID    uint
	refid uint
	Rule
}
