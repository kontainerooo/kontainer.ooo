// Package firewall handles the firewall and forwarding configuration
package firewall

import (
	"errors"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/firewall/iptables"
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
	iptClient iptables.Service
}

func (s *service) InitBridge(ip abstraction.Inet, netIf string) error {
	// Isolate bridge from other bridges and allow outgoing traffic
	err := s.iptClient.CreateRule(iptables.IsolationRuleType, iptables.IsolationRule{
		SrcNetwork: netIf,
	})
	if err == nil {
		return err
	}

	err = s.iptClient.CreateRule(iptables.OutgoingOutRuleType, iptables.OutgoingOutRule{
		SrcNetwork: netIf,
		SrcIP:      ip,
	})
	if err == nil {
		return err
	}

	err = s.iptClient.CreateRule(iptables.OutgoingInRuleType, iptables.OutgoingInRule{
		SrcNetwork: netIf,
		SrcIP:      ip,
	})
	if err == nil {
		return err
	}

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

func (s *service) setUpDNS() error {
	err := s.iptClient.CreateRule(iptables.AllowPortInRuleType, iptables.AllowPortInRule{
		Protocol: "udp",
		Port:     53,
		Chain:    iptables.IptDNSChain,
	})
	if err != nil {
		return err
	}

	err = s.iptClient.CreateRule(iptables.AllowPortOutRuleType, iptables.AllowPortOutRule{
		Protocol: "udp",
		Port:     53,
		Chain:    iptables.IptDNSChain,
	})
	if err != nil {
		return err
	}

	err = s.iptClient.CreateRule(iptables.AllowPortInRuleType, iptables.AllowPortInRule{
		Protocol: "tcp",
		Port:     53,
		Chain:    iptables.IptDNSChain,
	})
	if err != nil {
		return err
	}

	err = s.iptClient.CreateRule(iptables.AllowPortOutRuleType, iptables.AllowPortOutRule{
		Protocol: "tcp",
		Port:     53,
		Chain:    iptables.IptDNSChain,
	})
	if err != nil {
		return err
	}

	return nil
}

// NewService creates a new firewall service
func NewService(ipte iptables.Service) (Service, error) {
	s := &service{
		iptClient: ipte,
	}

	if s.iptClient == nil {
		return &service{}, errors.New("Invalid iptable client")
	}

	// Create predefined chains
	chains := []string{
		"KROO-DNS",
		"KROO-ISOLATION",
		"KROO-LINK",
		"KROO-OUTBOUND",
	}
	for _, v := range chains {
		if err := s.iptClient.CreateRule(iptables.CreateChainRuleType, iptables.CreateChainRule{
			Name: v,
		}); err != nil {
			return &service{}, err
		}
	}
	// Create FORWARD jumps to chains
	for _, v := range chains {
		if err := s.iptClient.CreateRule(iptables.JumpToChainRuleType, iptables.JumpToChainRule{
			From: "FORWARD",
			To:   v,
		}); err != nil {
			return &service{}, err
		}
	}

	return s, nil
}
