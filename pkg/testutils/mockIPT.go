package testutils

import "github.com/kontainerooo/kontainer.ooo/pkg/firewall/iptables"

// MockIPTService simulates a iptables service for testing purposes
type MockIPTService struct {
	rules map[string]iptables.RuleEntry
}

// CreateRule creates a new iptables rule
func (m *MockIPTService) CreateRule(ruleType int, ruleData interface{}) error {
	return nil
}

// NewMockIPTService creates a new MockIPTServicet
func NewMockIPTService() *MockIPTService {
	return &MockIPTService{
		rules: make(map[string]iptables.RuleEntry),
	}
}
