// Package firewall handles the firewall and forwarding configuration
package firewall

type networkInterface struct {
	ID    uint
	refid uint
	name  string
}
