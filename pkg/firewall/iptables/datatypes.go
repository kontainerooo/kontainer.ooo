// Package iptables is a wrapper around the iptables binary
package iptables

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"text/template"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
)

const (
	// IptDNSChain is the name of the Dns chain
	IptDNSChain = "KROO-DNS"

	// IptIsolationChain is the name of the isolation chain
	IptIsolationChain = "KROO-ISOLATION"

	// IptLinkChain is the name of the link chain
	IptLinkChain = "KROO-LINK"

	// IptOutboundChain is the name of the outbound chain
	IptOutboundChain = "KROO-OUTBOUND"

	// IptNatChain is the name of the custom chain that is used within the nat table
	IptNatChain = "KROO-NAT"

	// CreateChainRuleType specifies a CreateChainRule
	CreateChainRuleType = iota

	// JumpToChainRuleType specifies a JumpToChainRule
	JumpToChainRuleType = iota

	// IsolationRuleType specifies an IsolationRule
	IsolationRuleType = iota

	// OutgoingOutRuleType specifies an Outgoing rule for outgoing traffic
	OutgoingOutRuleType = iota

	// OutgoingInRuleType specifies an Outgoing rule for incoming traffic
	OutgoingInRuleType = iota

	// LinkContainerPortToRuleType specifies a link with port rule for incoming traffic
	LinkContainerPortToRuleType = iota

	// LinkContainerPortFromRuleType specifies a link with port rule for outgoing traffic
	LinkContainerPortFromRuleType = iota

	// LinkContainerToRuleType specifies a link rule for incoming traffic
	LinkContainerToRuleType = iota
	// LinkContainerFromRuleType specifies a link rule for outgoing traffic
	LinkContainerFromRuleType = iota

	// ConnectContainerFromRuleType specifies a connect rule for incoming traffic
	ConnectContainerFromRuleType = iota

	// ConnectContainerToRuleType specifies a connect rule for outgoing traffic
	ConnectContainerToRuleType = iota

	// AllowPortInRuleType specifies a rule for incoming traffic with a port
	AllowPortInRuleType = iota

	// AllowPortOutRuleType specifies a rule for outgoing traffic with a port
	AllowPortOutRuleType = iota

	// NatOutRuleType specifies the nat rule for outgoing traffic
	NatOutRuleType = iota

	// NatMaskRuleType specifies a rule for masking outgoing traffic
	NatMaskRuleType = iota
)

var (
	createChainRuleStr = "-t {{.Table}} -N {{.Name}}"

	jumpToChainRuleStr = "-t {{.Table}} -A {{.From}} {{if .SrcNetwork}} -i {{.SrcNetwork}} {{end}} {{if .Match}} -m {{.Match}} {{end}} -j {{.To}}"

	isolationRuleStr = fmt.Sprintf("-A %s ! -i {{.SrcNetwork}} -o {{.SrcNetwork}} -j DROP", IptIsolationChain)

	outgoingOutRuleStr = fmt.Sprintf("-A %s -s {{.SrcIP}} ! -d 172.16.0.0/12 -i {{.SrcNetwork}} ! -o {{.SrcNetwork}} -j ACCEPT", IptOutboundChain)
	outgoingInRuleStr  = fmt.Sprintf("-A %s ! -s 172.16.0.0/12 - d {{.SrcIP}} ! -i {{.SrcNetwork}} -o {{.SrcNetwork}} -j ACCEPT", IptOutboundChain)

	linkContainerPortToStr   = fmt.Sprintf("-A %s -s {{.SrcIP}} -d {{.DstIP}} -i {{.SrcNetwork}} -o {{.DstNetwork}} -p {{.Protocol}} --dport {{.DstPort}} -j ACCEPT", IptLinkChain)
	linkContainerPortFromStr = fmt.Sprintf("-A %s -s {{.DstIP}} -d {{.SrcIP}} -i {{.DstNetwork}} -o {{.SrcNetwork}} -p {{.Protocol}} -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT", IptLinkChain)

	linkContainerToStr   = fmt.Sprintf("-A %s -s {{.SrcIP}} -d {{.DstIP}} -i {{.SrcNetwork}} -o {{.DstNetwork}} -j ACCEPT", IptLinkChain)
	linkContainerFromStr = fmt.Sprintf("-A %s -s {{.DstIP}} -d {{.SrcIP}} -i {{.DstNetwork}} -o {{.SrcNetwork}} -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT", IptLinkChain)

	connectContainerFromStr = fmt.Sprintf("-A %s -s {{.SrcIP}} -d {{.DstIP}} -i {{.SrcNetwork}} -o {{.DstNetwork}} -j ACCEPT", IptLinkChain)
	connectContainerToStr   = fmt.Sprintf("-A %s -s {{.DstIP}} -d {{.SrcIP}} -i {{.DstNetwork}} -o {{.SrcNetwork}} -j ACCEPT", IptLinkChain)

	allowPortInStr  = "-A {{.Chain}} -p {{.Protocol}} -m {{.Protocol}} --sport {{.Port}} -m state --state ESTABLISHED -j ACCEPT"
	allowPortOutStr = "-A {{.Chain}} -p {{.Protocol}} -m {{.Protocol}} --dport {{.Port}} -m state --state NEW,ESTABLISHED -j ACCEPT"

	natOutStr = fmt.Sprintf("-A OUTPUT ! -d 127.0.0.0/8 -m addrtype --dst-type LOCAL -j %s", IptNatChain)

	natMaskStr = "-A POSTROUTING -s {{.SrcIP}} ! -o {{.SrcNetwork}} -j MASQUERADE"
)

var (
	// CreateChainRuleTmpl is the template for a rule that creates a new chain
	CreateChainRuleTmpl = template.Must(template.New("createChainRule").Parse(createChainRuleStr))

	// JumpToChainRuleTmpl is the template for a rule that creates a jump to other chain
	JumpToChainRuleTmpl = template.Must(template.New("jumpToChainRule").Parse(jumpToChainRuleStr))

	// IsolationRuleTmpl is the template for the container network isolation rule
	IsolationRuleTmpl = template.Must(template.New("isolationRule").Parse(isolationRuleStr))

	// OutgoingOutRuleTmpl is the template for the container network Outgoing rule for outgoing traffic
	OutgoingOutRuleTmpl = template.Must(template.New("outgoingOutRule").Parse(outgoingOutRuleStr))

	// OutgoingInRuleTmpl is the template for the container network Outgoing rule for incoming traffic
	OutgoingInRuleTmpl = template.Must(template.New("outgoingInRule").Parse(outgoingInRuleStr))

	// LinkContainerPortToRuleTmpl is the template for the container network link rule for outgoing traffic with port
	LinkContainerPortToRuleTmpl = template.Must(template.New("linkContainerPortToRule").Parse(linkContainerPortToStr))

	// LinkContainerPortFromRuleTmpl is the template for the container network link rule for incoming traffic with port
	LinkContainerPortFromRuleTmpl = template.Must(template.New("linkContainerPortFromRule").Parse(linkContainerPortFromStr))

	// LinkContainerToRuleTmpl is the template for the container network link rule for outgoing traffic
	LinkContainerToRuleTmpl = template.Must(template.New("linkContainerToRule").Parse(linkContainerToStr))

	// LinkContainerFromRuleTmpl is the template for the container network link rule for incoming traffic
	LinkContainerFromRuleTmpl = template.Must(template.New("linkContainerFromRule").Parse(linkContainerFromStr))

	// ConnectContainerFromRuleTmpl is the template for the container connection rule for outgoing traffic
	ConnectContainerFromRuleTmpl = template.Must(template.New("connectContainerFromRule").Parse(connectContainerFromStr))

	// ConnectContainerToRuleTmpl is the template for the container connection rule for incoming traffic
	ConnectContainerToRuleTmpl = template.Must(template.New("connectContainerToRule").Parse(connectContainerToStr))

	// AllowPortInRuleTmpl is the template for the port acceptance rule for incoming traffic
	AllowPortInRuleTmpl = template.Must(template.New("allowPortInRule").Parse(allowPortInStr))

	// AllowPortOutRuleTmpl is the template for the port acceptance rule for outgoing traffic
	AllowPortOutRuleTmpl = template.Must(template.New("allowPortOutRule").Parse(allowPortOutStr))

	// NatOutRuleTmpl is the template for the general nat outgoing rule
	NatOutRuleTmpl = template.Must(template.New("natOutRule").Parse(natOutStr))

	// NatMaskRuleTmpl is the template for the nat outgoing mask rule
	NatMaskRuleTmpl = template.Must(template.New("natMaskRule").Parse(natOutStr))
)

// RuleEntry represents a database rule entry
type RuleEntry struct {
	ID         string
	bridgeRef1 string
	bridgeRef2 string
	ipRef1     abstraction.Inet `sql:"type:inet"`
	ipRef2     abstraction.Inet `sql:"type:inet"`
	rule       Rule             `sql:"type:json"`
}

func (re RuleEntry) setRefs(br1 string, br2 string, ipr1 abstraction.Inet, ipr2 abstraction.Inet) {
	if br1 != "" {
		re.bridgeRef1 = br1
	}
	if br2 != "" {
		re.bridgeRef2 = br2
	}
	if string(ipr1) != "" {
		re.ipRef1 = ipr1
	}
	if string(ipr2) != "" {
		re.ipRef2 = ipr2
	}
}

// Rule specifies a rule with its corresponding type and data
type Rule struct {
	RuleType int
	Data     interface{}
}

// Value implements the Valuer interface
func (r Rule) Value() (driver.Value, error) {
	return json.Marshal(r)
}

// Scan implements the Scanner interface
func (r *Rule) Scan(src interface{}) error {
	switch src := src.(type) {
	case string:
		return r.scanBytes([]byte(src))
	case []byte:
		return r.scanBytes(src)
	case nil:
		*r = Rule{}
		return nil
	}

	return errors.New("pq: cannot convert input src to FrontendArray")
}

func (r *Rule) scanBytes(src []byte) error {
	a := Rule{}
	err := json.Unmarshal(src, &a)
	if err != nil {
		return err
	}
	data := a.Data.(anyRule)
	*r = Rule{RuleType: a.RuleType}

	switch a.RuleType {
	case CreateChainRuleType:
		r.Data = CreateChainRule{
			Name:  data.Name,
			Table: data.Table,
		}
	case JumpToChainRuleType:
		r.Data = JumpToChainRule{
			From:  data.From,
			To:    data.To,
			Table: data.Table,
		}
	case IsolationRuleType:
		r.Data = IsolationRule{
			SrcNetwork: data.SrcNetwork,
		}
	case OutgoingOutRuleType:
		srcIP, err := abstraction.NewInet(data.SrcIP)
		if err != nil {
			return err
		}
		r.Data = OutgoingOutRule{
			SrcNetwork: data.SrcNetwork,
			SrcIP:      srcIP,
		}
	case OutgoingInRuleType:
		srcIP, err := abstraction.NewInet(data.SrcIP)
		if err != nil {
			return err
		}
		r.Data = OutgoingInRule{
			SrcNetwork: data.SrcNetwork,
			SrcIP:      srcIP,
		}
	case LinkContainerPortToRuleType:
		srcIP, err := abstraction.NewInet(data.SrcIP)
		if err != nil {
			return err
		}
		dstIP, err := abstraction.NewInet(data.DstIP)
		if err != nil {
			return err
		}

		r.Data = LinkContainerPortToRule{
			SrcIP:      srcIP,
			DstIP:      dstIP,
			SrcNetwork: data.SrcNetwork,
			DstNetwork: data.DstNetwork,
			Protocol:   data.Protocol,
			DstPort:    uint16(data.DstPort),
		}
	case LinkContainerPortFromRuleType:
		srcIP, err := abstraction.NewInet(data.SrcIP)
		if err != nil {
			return err
		}
		dstIP, err := abstraction.NewInet(data.DstIP)
		if err != nil {
			return err
		}

		r.Data = LinkContainerPortFromRule{
			DstIP:      dstIP,
			SrcIP:      srcIP,
			DstNetwork: data.DstNetwork,
			SrcNetwork: data.SrcNetwork,
			Protocol:   data.Protocol,
		}
	case LinkContainerToRuleType:
		srcIP, err := abstraction.NewInet(data.SrcIP)
		if err != nil {
			return err
		}
		dstIP, err := abstraction.NewInet(data.DstIP)
		if err != nil {
			return err
		}

		r.Data = LinkContainerToRule{
			DstIP:      dstIP,
			SrcIP:      srcIP,
			DstNetwork: data.SrcNetwork,
			SrcNetwork: data.DstNetwork,
		}
	case LinkContainerFromRuleType:
		srcIP, err := abstraction.NewInet(data.SrcIP)
		if err != nil {
			return err
		}
		dstIP, err := abstraction.NewInet(data.DstIP)
		if err != nil {
			return err
		}

		r.Data = LinkContainerFromRule{
			DstIP:      dstIP,
			SrcIP:      srcIP,
			DstNetwork: data.DstNetwork,
			SrcNetwork: data.SrcNetwork,
		}
	case ConnectContainerFromRuleType:
		srcIP, err := abstraction.NewInet(data.SrcIP)
		if err != nil {
			return err
		}
		dstIP, err := abstraction.NewInet(data.DstIP)
		if err != nil {
			return err
		}

		r.Data = ConnectContainerFromRule{
			DstIP:      dstIP,
			SrcIP:      srcIP,
			DstNetwork: data.DstNetwork,
			SrcNetwork: data.SrcNetwork,
		}
	case ConnectContainerToRuleType:
		srcIP, err := abstraction.NewInet(data.SrcIP)
		if err != nil {
			return err
		}
		dstIP, err := abstraction.NewInet(data.DstIP)
		if err != nil {
			return err
		}

		r.Data = ConnectContainerToRule{
			DstIP:      dstIP,
			SrcIP:      srcIP,
			DstNetwork: data.DstNetwork,
			SrcNetwork: data.SrcNetwork,
		}
	case AllowPortInRuleType:
		r.Data = AllowPortInRule{
			Protocol: data.Protocol,
			Port:     uint16(data.Port),
			Chain:    data.Chain,
		}
	case AllowPortOutRuleType:
		r.Data = AllowPortOutRule{
			Protocol: data.Protocol,
			Port:     uint16(data.Port),
			Chain:    data.Chain,
		}
	case NatOutRuleType:
		r.Data = NatOutRule{}
	case NatMaskRuleType:
		r.Data = NatMaskRule{}
	default:
		return errors.New("pq: cannot convert input src to FrontendArray")
	}

	return nil
}

type anyRule struct {
	Name       string
	From       string
	To         string
	SrcIP      string
	DstIP      string
	SrcNetwork string
	DstNetwork string
	Protocol   string
	DstPort    float64
	Port       float64
	Chain      string
	Table      string
}

// CreateChainRule represents rule data for a CreateChainRuleType
type CreateChainRule struct {
	Name  string
	Table string
}

// JumpToChainRule represents rule data for a JumpToChainRuleType
type JumpToChainRule struct {
	From       string
	To         string
	Table      string
	SrcNetwork string
	Match      string
}

// IsolationRule represents rule data for an IsolationRuleType
type IsolationRule struct {
	SrcNetwork string
}

// OutgoingOutRule represents rule data for an OutgoingOutRuleType
type OutgoingOutRule struct {
	SrcNetwork string
	SrcIP      abstraction.Inet
}

// OutgoingInRule represents rule data for an OutgoingInRuleType
type OutgoingInRule struct {
	SrcNetwork string
	SrcIP      abstraction.Inet
}

// LinkContainerPortToRule represents rule data for an LinkContainerPortToRuleType
type LinkContainerPortToRule struct {
	SrcIP      abstraction.Inet
	DstIP      abstraction.Inet
	SrcNetwork string
	DstNetwork string
	Protocol   string
	DstPort    uint16
}

// LinkContainerPortFromRule represents rule data for an LinkContainerPortFromRuleType
type LinkContainerPortFromRule struct {
	DstIP      abstraction.Inet
	SrcIP      abstraction.Inet
	DstNetwork string
	SrcNetwork string
	Protocol   string
}

// LinkContainerToRule represents rule data for a LinkContainerToRuleType
type LinkContainerToRule struct {
	DstIP      abstraction.Inet
	SrcIP      abstraction.Inet
	DstNetwork string
	SrcNetwork string
}

// LinkContainerFromRule represents rule data for a LinkContainerFromRuleType
type LinkContainerFromRule struct {
	DstIP      abstraction.Inet
	SrcIP      abstraction.Inet
	DstNetwork string
	SrcNetwork string
}

// ConnectContainerFromRule represents rule data for a ConnectContainerFromRuleType
type ConnectContainerFromRule struct {
	DstIP      abstraction.Inet
	SrcIP      abstraction.Inet
	DstNetwork string
	SrcNetwork string
}

// ConnectContainerToRule represents rule data for a ConnectContainerToRuleType
type ConnectContainerToRule struct {
	DstIP      abstraction.Inet
	SrcIP      abstraction.Inet
	DstNetwork string
	SrcNetwork string
}

// AllowPortInRule represents rule data for an AllowPortInRuleType
type AllowPortInRule struct {
	Protocol string
	Port     uint16
	Chain    string
}

// AllowPortOutRule represents rule data for an AllowPortOutRuleType
type AllowPortOutRule struct {
	Protocol string
	Port     uint16
	Chain    string
}

// NatOutRule represents rule data for a NatOutRuleType
type NatOutRule struct{}

// NatMaskRule represents rule data for a NatMaskRuleType
type NatMaskRule struct {
	SrcIP      abstraction.Inet
	SrcNetwork string
}
