// Package iptables is a wrapper around the iptables binary
package iptables

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
)

// Service handles iptables rules
type Service interface {
	// CreateRule creates a rule with a given type and data
	CreateRule(ruleType int, ruleData interface{}) error

	// RemoveRule removes a rule in case it exists
	RemoveRule(ruleType int, ruleData interface{}) error

	// CreateRuleEntryString creates a rule entry and the command string from a type and data
	CreateRuleEntryString(ruleType int, ruleData interface{}) (RuleEntry, string, error)
}

type dbAdapter interface {
	abstraction.DBAdapter
	AutoMigrate(...interface{}) error
	Create(interface{}) error
	Where(interface{}, ...interface{}) error
	Delete(interface{}, ...interface{}) error
}

type service struct {
	iptPath string
	db      dbAdapter
}

func (s *service) executeIPTableCommand(c string) error {
	cmd := ExecCommand(s.iptPath, c)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (s *service) InitializeDatabases() error {
	return s.db.AutoMigrate(&RuleEntry{})
}

func (s *service) ruleExists(id string) bool {
	s.db.Where("ID = ?", id)
	res := s.db.GetValue()
	if res != nil {
		return true
	}
	return false
}

func createHash(cmd string) string {
	sum := sha256.Sum256([]byte(cmd))
	return fmt.Sprintf("%x", sum)
}

func (s *service) CreateRuleEntryString(ruleType int, ruleData interface{}) (RuleEntry, string, error) {
	errInvalidData := errors.New("Invalid rule data")
	re := RuleEntry{}
	cmdStr := ""

	switch ruleType {
	case CreateChainRuleType:
		rd, ok := ruleData.(CreateChainRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		if rd.Table == "" {
			rd.Table = "filter"
		}
		rule := Rule{
			Data:     rd,
			RuleType: CreateChainRuleType,
		}
		re.rule = rule
		re.setRefs("", "", abstraction.Inet(""), abstraction.Inet(""))

		var buf bytes.Buffer
		err := CreateChainRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case JumpToChainRuleType:
		rd, ok := ruleData.(JumpToChainRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		if rd.Table == "" {
			rd.Table = "filter"
		}
		rule := Rule{
			Data:     rd,
			RuleType: JumpToChainRuleType,
		}
		re.rule = rule
		re.setRefs("", "", abstraction.Inet(""), abstraction.Inet(""))

		var buf bytes.Buffer
		err := JumpToChainRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case IsolationRuleType:
		rd, ok := ruleData.(IsolationRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		rule := Rule{
			Data:     rd,
			RuleType: IsolationRuleType,
		}
		re.rule = rule
		re.setRefs(rd.SrcNetwork, "", abstraction.Inet(""), abstraction.Inet(""))

		var buf bytes.Buffer
		err := IsolationRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case OutgoingOutRuleType:
		rd, ok := ruleData.(OutgoingOutRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		rule := Rule{
			Data:     rd,
			RuleType: OutgoingOutRuleType,
		}
		re.rule = rule
		re.setRefs(rd.SrcNetwork, "", rd.SrcIP, abstraction.Inet(""))

		var buf bytes.Buffer
		err := OutgoingOutRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case OutgoingInRuleType:
		rd, ok := ruleData.(OutgoingInRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		rule := Rule{
			Data:     rd,
			RuleType: OutgoingInRuleType,
		}
		re.rule = rule
		re.setRefs(rd.SrcNetwork, "", rd.SrcIP, abstraction.Inet(""))

		var buf bytes.Buffer
		err := OutgoingInRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case LinkContainerPortToRuleType:
		rd, ok := ruleData.(LinkContainerPortToRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		rule := Rule{
			Data:     rd,
			RuleType: LinkContainerPortToRuleType,
		}
		re.rule = rule
		re.setRefs(rd.SrcNetwork, rd.DstNetwork, rd.SrcIP, rd.DstIP)

		var buf bytes.Buffer
		err := LinkContainerPortToRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case LinkContainerPortFromRuleType:
		rd, ok := ruleData.(LinkContainerPortFromRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		rule := Rule{
			Data:     rd,
			RuleType: LinkContainerPortFromRuleType,
		}
		re.rule = rule
		re.setRefs(rd.SrcNetwork, rd.DstNetwork, rd.SrcIP, rd.DstIP)

		var buf bytes.Buffer
		err := LinkContainerPortFromRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case LinkContainerToRuleType:
		rd, ok := ruleData.(LinkContainerToRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		rule := Rule{
			Data:     rd,
			RuleType: LinkContainerToRuleType,
		}
		re.rule = rule
		re.setRefs(rd.SrcNetwork, rd.DstNetwork, rd.SrcIP, rd.DstIP)

		var buf bytes.Buffer
		err := LinkContainerToRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case LinkContainerFromRuleType:
		rd, ok := ruleData.(LinkContainerFromRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		rule := Rule{
			Data:     rd,
			RuleType: LinkContainerFromRuleType,
		}
		re.rule = rule
		re.setRefs(rd.SrcNetwork, rd.DstNetwork, rd.SrcIP, rd.DstIP)

		var buf bytes.Buffer
		err := LinkContainerFromRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case ConnectContainerFromRuleType:
		rd, ok := ruleData.(ConnectContainerFromRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		rule := Rule{
			Data:     rd,
			RuleType: ConnectContainerFromRuleType,
		}
		re.rule = rule
		re.setRefs(rd.SrcNetwork, rd.DstNetwork, rd.SrcIP, rd.DstIP)

		var buf bytes.Buffer
		err := ConnectContainerFromRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case ConnectContainerToRuleType:
		rd, ok := ruleData.(ConnectContainerToRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		rule := Rule{
			Data:     rd,
			RuleType: ConnectContainerToRuleType,
		}
		re.rule = rule
		re.setRefs(rd.SrcNetwork, rd.DstNetwork, rd.SrcIP, rd.DstIP)

		var buf bytes.Buffer
		err := ConnectContainerToRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case AllowPortInRuleType:
		rd, ok := ruleData.(AllowPortInRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		rule := Rule{
			Data:     rd,
			RuleType: AllowPortInRuleType,
		}
		re.rule = rule
		re.setRefs("", "", abstraction.Inet(""), abstraction.Inet(""))

		var buf bytes.Buffer
		err := AllowPortInRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case AllowPortOutRuleType:
		rd, ok := ruleData.(AllowPortOutRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		rule := Rule{
			Data:     rd,
			RuleType: AllowPortOutRuleType,
		}
		re.rule = rule
		re.setRefs("", "", abstraction.Inet(""), abstraction.Inet(""))

		var buf bytes.Buffer
		err := AllowPortOutRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case NatOutRuleType:
		rd, ok := ruleData.(NatOutRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		rule := Rule{
			Data:     rd,
			RuleType: NatOutRuleType,
		}
		re.rule = rule
		re.setRefs("", "", abstraction.Inet(""), abstraction.Inet(""))

		var buf bytes.Buffer
		err := NatOutRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	case NatMaskRuleType:
		rd, ok := ruleData.(NatMaskRule)
		if !ok {
			return RuleEntry{}, "", errInvalidData
		}
		rule := Rule{
			Data:     rd,
			RuleType: NatMaskRuleType,
		}
		re.rule = rule
		re.setRefs(rd.SrcNetwork, "", rd.SrcIP, abstraction.Inet(""))

		var buf bytes.Buffer
		err := NatMaskRuleTmpl.Execute(&buf, rd)
		if err != nil {
			return RuleEntry{}, "", err
		}

		cmdStr = buf.String()
		re.ID = createHash(cmdStr)
	default:
		return RuleEntry{}, "", errors.New("pq: cannot convert input src to FrontendArray")
	}

	return re, cmdStr, nil
}

func (s *service) CreateRule(ruleType int, ruleData interface{}) error {

	re, cmdStr, err := s.CreateRuleEntryString(ruleType, ruleData)
	if err != nil {
		return err
	}

	if s.ruleExists(re.ID) {
		return errors.New("Rule already exists")
	}

	err = s.executeIPTableCommand(cmdStr)
	if err != nil {
		return err
	}

	err = s.db.Create(&re)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) RemoveRule(ruleType int, ruleData interface{}) error {
	re, cmdStr, err := s.CreateRuleEntryString(ruleType, ruleData)
	if err != nil {
		return err
	}

	if !s.ruleExists(re.ID) {
		return errors.New("Rule does not exist")
	}

	if !strings.Contains(cmdStr, "-A") {
		return errors.New("Rule cannot be removed (no -A present)")
	}

	strings.Replace(cmdStr, "-A", "-D", -1)

	err = s.executeIPTableCommand(cmdStr)
	if err != nil {
		return err
	}

	err = s.db.Delete(&re)
	if err != nil {
		return err
	}

	return nil
}

// ExecCommand is a wrapper around exec.Command used for testing
var ExecCommand = exec.Command

// NewService creates a new iptables service
func NewService(iptPath string, db dbAdapter) (Service, error) {
	s := &service{
		iptPath: iptPath,
		db:      db,
	}

	cmd := ExecCommand(iptPath, "--version")
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	err = s.InitializeDatabases()
	if err != nil {
		return nil, err
	}

	return s, nil
}
