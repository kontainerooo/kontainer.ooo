// Package firewall handles the firewall and forwarding configuration
package firewall

import (
	"context"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/iptables"
)

// Service firewall
type Service interface {
	// InitBridge initializes a bridge network
	InitBridge(ip abstraction.Inet, netIf string) error

	// AllowConnection sets up a rule to let src talk to dst
	AllowConnection(src abstraction.Inet, dst abstraction.Inet) error

	// AllowConnection sets up a rule to block src from talking to dst
	BlockConnection(src abstraction.Inet, dst abstraction.Inet) error

	// AllowPort sets up a rule to let src talk to dst on port port
	AllowPort(src abstraction.Inet, dst abstraction.Inet, port uint32) error

	// AllowPort sets up a rule to block src from talking to dst on port port
	BlockPort(src abstraction.Inet, dst abstraction.Inet, port uint32) error

	// RedirectPort redirects the src port to the dst port on ip
	RedirectPort(ip abstraction.Inet, src uint32, dst uint32) error

	// RemoveRedirectPort removes a port redirection
	RemoveRedirectPort(ip abstraction.Inet, src uint32, dst uint32) error
}

type service struct {
	iptClient *iptables.Endpoints
}

func (s *service) InitBridge(ip abstraction.Inet, netIf string) error {
	// TODO: implement
	return nil
}

func (s *service) AllowConnection(src abstraction.Inet, dst abstraction.Inet) error {
	// TODO: implement
	return nil
}

func (s *service) BlockConnection(src abstraction.Inet, dst abstraction.Inet) error {
	// TODO: implement
	return nil
}

func (s *service) AllowPort(src abstraction.Inet, dst abstraction.Inet, port uint32) error {
	// TODO: implement
	return nil
}

func (s *service) BlockPort(src abstraction.Inet, dst abstraction.Inet, port uint32) error {
	// TODO: implement
	return nil
}

func (s *service) RedirectPort(ip abstraction.Inet, src uint32, dst uint32) error {
	// TODO: implement
	return nil
}

func (s *service) RemoveRedirectPort(ip abstraction.Inet, src uint32, dst uint32) error {
	// TODO: implement
	return nil
}

// NewService creates a new firewall service
func NewService(ipte *iptables.Endpoints) Service {
	s := &service{
		iptClient: ipte,
	}

	return s
}
