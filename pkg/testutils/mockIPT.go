package testutils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/kontainerooo/kontainer.ooo/pkg/firewall/iptables"
)

// MockIPTService simulates a iptables service for testing purposes
type MockIPTService struct {
	rules map[string]iptables.RuleEntry
	s     iptables.Service
}

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

// CreateRule creates a new iptables rule
func (m *MockIPTService) CreateRule(ruleType int, ruleData interface{}) error {
	re, _, err := m.s.CreateRuleEntryString(ruleType, ruleData)
	if err != nil {
		return err
	}

	fmt.Println(m.rules)

	_, ok := m.rules[re.ID]
	if ok {
		return errors.New("Rule already exists")
	}

	m.rules[re.ID] = re

	return nil
}

// RemoveRule removes an iptables rule
func (m *MockIPTService) RemoveRule(ruleType int, ruleData interface{}) error {
	re, _, err := m.s.CreateRuleEntryString(ruleType, ruleData)
	if err != nil {
		return err
	}

	_, ok := m.rules[re.ID]
	if !ok {
		return errors.New("Rule does not exist")
	}

	delete(m.rules, re.ID)

	return nil
}

// CreateRuleEntryString calls the original CreateRuleEntryString
func (m *MockIPTService) CreateRuleEntryString(ruleType int, ruleData interface{}) (iptables.RuleEntry, string, error) {
	return m.s.CreateRuleEntryString(ruleType, ruleData)
}

// ListRules lists all created rules as string
func (m *MockIPTService) ListRules() []string {
	return []string{}
}

// RestoreRules is not mocked
func (m *MockIPTService) RestoreRules() error {
	return nil
}

// NewMockIPTService creates a new MockIPTServicet
func NewMockIPTService() (*MockIPTService, error) {
	db := NewMockDB()
	iptables.ExecCommand = fakeExecCommand
	ipts, err := iptables.NewService("iptables", "iptables-restore", db)

	return &MockIPTService{
		rules: make(map[string]iptables.RuleEntry),
		s:     ipts,
	}, err
}
