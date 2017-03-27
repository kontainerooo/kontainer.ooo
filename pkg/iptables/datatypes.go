// Package iptables is a wrapper around the iptables binary
package iptables

import "github.com/kontainerooo/kontainer.ooo/pkg/abstraction"

// Rule represents a rule entry in iptables
type Rule struct {
	Operation       string
	Name            string
	Target          string
	Chain           string
	Protocol        string
	In              string
	Out             string
	Source          abstraction.Inet `sql:"type:inet"`
	Destination     abstraction.Inet `sql:"type:inet"`
	SourcePort      uint16
	DestinationPort uint16
	State           string
}

type iptablesEntry struct {
	ID    string
	Refid string
	Rule
}
