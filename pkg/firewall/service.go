// Package firewall handles the firewall and forwarding configuration
package firewall

import (
	"errors"
	"regexp"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/firewall/iptables"
)

// Service firewall
type Service interface {
	// InitBridge initializes a bridge network
	InitBridge(ip abstraction.Inet, netIf string) error

	// AllowConnection sets up a rule to let src talk to dst
	AllowConnection(srcIP abstraction.Inet, srcNw string, dstIP abstraction.Inet, dstNw string) error

	// AllowConnection sets up a rule to block src from talking to dst
	BlockConnection(srcIP abstraction.Inet, srcNw string, dstIP abstraction.Inet, dstNw string) error

	// AllowPort sets up a rule to let src talk to dst on port port
	AllowPort(srcIP abstraction.Inet, srcNw string, dstIP abstraction.Inet, dstNw string, port uint16, protocol string) error

	// AllowPort sets up a rule to block src from talking to dst on port port
	BlockPort(srcIP abstraction.Inet, srcNw string, dstIP abstraction.Inet, dstNw string, port uint16, protocol string) error

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
	if err != nil {
		return err
	}

	err = s.iptClient.CreateRule(iptables.OutgoingOutRuleType, iptables.OutgoingOutRule{
		SrcNetwork: netIf,
		SrcIP:      ip,
	})
	if err != nil {
		return err
	}

	err = s.iptClient.CreateRule(iptables.OutgoingInRuleType, iptables.OutgoingInRule{
		SrcNetwork: netIf,
		SrcIP:      ip,
	})
	if err != nil {
		return err
	}

	err = s.iptClient.CreateRule(iptables.JumpToChainRuleType, iptables.JumpToChainRule{
		Table:      "nat",
		SrcNetwork: netIf,
		From:       iptables.IptNatChain,
		To:         "RETURN",
	})
	if err != nil {
		return err
	}

	err = s.iptClient.CreateRule(iptables.NatOutRuleType, iptables.NatOutRule{})
	if err != nil {
		return err
	}

	err = s.iptClient.CreateRule(iptables.NatMaskRuleType, iptables.NatMaskRule{
		SrcIP:      ip,
		SrcNetwork: netIf,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) AllowConnection(srcIP abstraction.Inet, srcNw string, dstIP abstraction.Inet, dstNw string) error {
	err := s.iptClient.CreateRule(iptables.LinkContainerFromRuleType, iptables.LinkContainerFromRule{
		SrcIP:      srcIP,
		SrcNetwork: srcNw,
		DstIP:      dstIP,
		DstNetwork: dstNw,
	})
	if err != nil {
		return err
	}

	err = s.iptClient.CreateRule(iptables.LinkContainerToRuleType, iptables.LinkContainerToRule{
		SrcIP:      srcIP,
		SrcNetwork: srcNw,
		DstIP:      dstIP,
		DstNetwork: dstNw,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) BlockConnection(srcIP abstraction.Inet, srcNw string, dstIP abstraction.Inet, dstNw string) error {
	err := s.iptClient.RemoveRule(iptables.LinkContainerFromRuleType, iptables.LinkContainerFromRule{
		SrcIP:      srcIP,
		SrcNetwork: srcNw,
		DstIP:      dstIP,
		DstNetwork: dstNw,
	})
	if err != nil {
		return err
	}

	err = s.iptClient.RemoveRule(iptables.LinkContainerToRuleType, iptables.LinkContainerToRule{
		SrcIP:      srcIP,
		SrcNetwork: srcNw,
		DstIP:      dstIP,
		DstNetwork: dstNw,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) isValidProtocol(p string) bool {
	// TODO: which protocols do we support?
	r := regexp.MustCompile("tcp|udp")
	if r.MatchString(p) {
		return true
	}
	return false
}

func (s *service) AllowPort(srcIP abstraction.Inet, srcNw string, dstIP abstraction.Inet, dstNw string, port uint16, protocol string) error {
	if !s.isValidProtocol(protocol) {
		return errors.New("Not a valid protocol")
	}

	err := s.iptClient.CreateRule(iptables.LinkContainerPortFromRuleType, iptables.LinkContainerPortFromRule{
		SrcIP:      srcIP,
		SrcNetwork: srcNw,
		DstIP:      dstIP,
		DstNetwork: dstNw,
		Protocol:   protocol,
	})
	if err != nil {
		return err
	}

	err = s.iptClient.CreateRule(iptables.LinkContainerPortToRuleType, iptables.LinkContainerPortToRule{
		SrcIP:      srcIP,
		SrcNetwork: srcNw,
		DstIP:      dstIP,
		DstNetwork: dstNw,
		Protocol:   protocol,
		DstPort:    port,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) BlockPort(srcIP abstraction.Inet, srcNw string, dstIP abstraction.Inet, dstNw string, port uint16, protocol string) error {
	if !s.isValidProtocol(protocol) {
		return errors.New("Not a valid protocol")
	}

	err := s.iptClient.RemoveRule(iptables.LinkContainerPortFromRuleType, iptables.LinkContainerPortFromRule{
		SrcIP:      srcIP,
		SrcNetwork: srcNw,
		DstIP:      dstIP,
		DstNetwork: dstNw,
		Protocol:   protocol,
	})
	if err != nil {
		return err
	}

	err = s.iptClient.RemoveRule(iptables.LinkContainerPortToRuleType, iptables.LinkContainerPortToRule{
		SrcIP:      srcIP,
		SrcNetwork: srcNw,
		DstIP:      dstIP,
		DstNetwork: dstNw,
		Protocol:   protocol,
		DstPort:    port,
	})
	if err != nil {
		return err
	}

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
		iptables.IptDNSChain,
		iptables.IptOutboundChain,
		iptables.IptLinkChain,
		iptables.IptIsolationChain,
	}
	for _, v := range chains {
		if err := s.iptClient.CreateRule(iptables.CreateChainRuleType, iptables.CreateChainRule{
			Name: v,
		}); err != nil {
			return &service{}, err
		}
	}
	// Create chain in nat table
	if err := s.iptClient.CreateRule(iptables.CreateChainRuleType, iptables.CreateChainRule{
		Name:  iptables.IptNatChain,
		Table: "nat",
	}); err != nil {
		return &service{}, err
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
	// Create nat jump
	if err := s.iptClient.CreateRule(iptables.JumpToChainRuleType, iptables.JumpToChainRule{
		From:  "PREROUTING",
		To:    iptables.IptNatChain,
		Table: "nat",
		Match: "addrtype --dst-type LOCAL",
	}); err != nil {
		return &service{}, err
	}

	err := s.setUpDNS()
	if err != nil {
		return &service{}, err
	}

	return s, nil
}
