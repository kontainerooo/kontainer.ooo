// Package iptables is a wrapper around the iptables binary
package iptables

import (
	"fmt"
	"html/template"

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
)

var (
	createChainRuleStr = "-N {{.Name}}"

	jumpToChainRuleStr = "-A {{.From}} -j {{.To}}"

	isolationRuleStr = fmt.Sprintf("-A %s ! -i {{.SrcInterface}} -o {{.SrcInterface}} -j DROP", IptIsolationChain)

	outgoingOutRuleStr = fmt.Sprintf("-A %s -s {{.SrcIP}} ! -d 172.16.0.0/12 -i {{.SrcNetwork}} ! -o {{.SrcNetwork}} -j ACCEPT", IptOutboundChain)
	outgoingInRuleStr  = fmt.Sprintf("-A %s ! -s 172.16.0.0/12 - d {{.SrcIP}} ! -i {{.SrcNetwork}} -o {{.SrcNetwork}} -j ACCEPT", IptOutboundChain)

	linkContainerPortToStr   = fmt.Sprintf("-A %s -s {{.SrcIP}} -d {{.DstIP}} -i {{.SrcNetwork}} -o {{.DstNetwork}} -p {{.Protocol}} --dport {{.DstPort}} -j ACCEPT", IptLinkChain)
	linkContainerPortFromStr = fmt.Sprintf("-A %s -s {{.DstIP}} -d {{.SrcIP}} -i {{.DstNetwork}} -o {{.SrcNetwork}} -p {{.Protocol}} -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT", IptLinkChain)

	linkContainerToStr   = fmt.Sprintf("-A %s -s {{.SrcIP}} -d {{.DstIP}} -i {{.SrcNetwork}} -o {{.DstNetwork}} -j ACCEPT", IptLinkChain)
	linkContainerFromStr = fmt.Sprintf("-A %s -s {{.DstIP}} -d {{.SrcIP}} -i {{.DstNetwork}} -o {{.SrcNetwork}} -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT", IptLinkChain)

	connectContainerFromStr = fmt.Sprintf("-A %s -s {{.SrcIP}} -d {{.DstIP}} -i {{.SrcNetwork}} -o {{.DstNetwork}} -j ACCEPT", IptLinkChain)
	connectContainerToStr   = fmt.Sprintf("-A %s -s {{.DstIP}} -d {{.SrcIP}} -i {{.DstNetwork}} -o {{.SrcNetwork}} -j ACCEPT", IptLinkChain)
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

	// LinkContainerPortToTmpl is the template for the container network link rule for outgoing traffic with port
	LinkContainerPortToTmpl = template.Must(template.New("linkContainerPortTo").Parse(linkContainerPortToStr))

	// LinkContainerPortFromTmpl is the template for the container network link rule for incoming traffic with port
	LinkContainerPortFromTmpl = template.Must(template.New("linkContainerPortFrom").Parse(linkContainerPortFromStr))

	// LinkContainerToTmpl is the template for the container network link rule for outgoing traffic
	LinkContainerToTmpl = template.Must(template.New("linkContainerTo").Parse(linkContainerToStr))

	// LinkContainerFromTmpl is the template for the container network link rule for incoming traffic
	LinkContainerFromTmpl = template.Must(template.New("linkContainerFrom").Parse(linkContainerFromStr))

	// ConnectContainerFromTmpl is the template for the container connection rule for outgoing traffic
	ConnectContainerFromTmpl = template.Must(template.New("connectContainerFrom").Parse(connectContainerFromStr))

	// ConnectContainerToTmpl is the template for the container connection rule for incoming traffic
	ConnectContainerToTmpl = template.Must(template.New("connectContainerTo").Parse(connectContainerToStr))
)

type ruleEntry struct {
	ID         string
	bridgeRef1 string
	bridgeRef2 string
	ipRef1     abstraction.Inet
	ipRef2     abstraction.Inet
	Rule
}

// Rule specifies a rule with its corresponding type and data
type Rule struct {
	RuleType string
	RuleData interface{}
}

// CreateChainRule represents rule data for a CreateChainRuleType
type CreateChainRule struct {
	Name string
}

// JumpToChainRule represents rule data for a JumpToChainRuleType
type JumpToChainRule struct {
	From string
	To   string
}

// IsolationRule represents rula data for an IsolationRuleType
type IsolationRule struct {
	SrcNetwork string
}

// OutgoingOutRule represents rula data for an OutgoingOutRuleType
type OutgoingOutRule struct {
	SrcNetwork string
	SrcIP      abstraction.Inet
}

// OutgoingInRule represents rula data for an OutgoingInRuleType
type OutgoingInRule struct {
	SrcNetwork string
	SrcIP      abstraction.Inet
}

// LinkContainerPortToRule represents rula data for an LinkContainerPortToRuleType
type LinkContainerPortToRule struct {
	SrcIP      abstraction.Inet
	DstIP      abstraction.Inet
	SrcNetwork string
	DstNetwork string
	Protocol   string
	DstPort    string
}

// LinkContainerPortFromRule represents rula data for an LinkContainerPortFromRuleType
type LinkContainerPortFromRule struct {
	DstIP      abstraction.Inet
	SrcIP      abstraction.Inet
	DstNetwork string
	SrcNetwork string
	Protocol   string
}

// LinkContainerToRule represents rula data for an LinkContainerToRuleType
type LinkContainerToRule struct {
	DstIP      abstraction.Inet
	SrcIP      abstraction.Inet
	DstNetwork string
	SrcNetwork string
}

// LinkContainerFromRule represents rula data for an LinkContainerFromRuleType
type LinkContainerFromRule struct {
	DstIP      abstraction.Inet
	SrcIP      abstraction.Inet
	DstNetwork string
	SrcNetwork string
}

// ConnectContainerFromRule represents rula data for an ConnectContainerFromRuleType
type ConnectContainerFromRule struct {
	DstIP      abstraction.Inet
	SrcIP      abstraction.Inet
	DstNetwork string
	SrcNetwork string
}

// ConnectContainerToRule represents rula data for an ConnectContainerToRuleType
type ConnectContainerToRule struct {
	DstIP      abstraction.Inet
	SrcIP      abstraction.Inet
	DstNetwork string
	SrcNetwork string
}
