package testutils

import (
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
	return m.s.CreateRule(ruleType, ruleData)
}

// ListRules lists all created rules as string
func (m *MockIPTService) ListRules() []string {
	return []string{}
}

// NewMockIPTService creates a new MockIPTServicet
func NewMockIPTService() (*MockIPTService, error) {
	db := NewMockDB()
	iptables.ExecCommand = fakeExecCommand
	ipts, err := iptables.NewService("iptables", db)

	return &MockIPTService{
		rules: make(map[string]iptables.RuleEntry),
		s:     ipts,
	}, err
}
