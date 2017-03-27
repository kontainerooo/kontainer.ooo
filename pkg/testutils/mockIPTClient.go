package testutils

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/iptables"
)

// MockIPTClient simulates an iptables client for testing purposes
type MockIPTClient struct {
	cliCommands map[string]mockEntry
}

type mockEntry struct {
	rule  iptables.Rule
	refid uint
}

// NewMockIPTEndpoints creates a new MockIPTClientEndpoints
func NewMockIPTEndpoints(logger log.Logger, m MockIPTClient) *iptables.Endpoints {
	return &iptables.Endpoints{
		AddRuleEndpoint: m.AddRuleEndpoint,
	}
}

// AddRuleEndpoint adds a new rule to the mockclient
func (m *MockIPTClient) AddRuleEndpoint(ctx context.Context, req interface{}) (interface{}, error) {
	iptReq := req.(*iptables.AddRuleRequest)

	if err := iptReq.Rule.IsValid(); err != nil {
		return &iptables.AddRuleResponse{}, err
	}

	_, ok := m.cliCommands[iptReq.Rule.GetHash()]
	if ok {
		return &iptables.AddRuleResponse{}, errors.New("Rule already exists")

	}

	m.cliCommands[iptReq.Rule.GetHash()] = mockEntry{
		refid: iptReq.Refid,
		rule:  iptReq.Rule,
	}

	return &iptables.AddRuleResponse{
		Error: nil,
	}, nil
}

// ListRuleStrings prints the rules as iptable commands
func (m *MockIPTClient) ListRuleStrings() {
	for _, v := range m.cliCommands {
		str, _ := v.rule.ToString()
		fmt.Printf("iptables %s\n", str)
	}
}

// NewMockIPTClient creates a new MockKMIClient
func NewMockIPTClient() *MockIPTClient {
	return &MockIPTClient{
		cliCommands: make(map[string]mockEntry),
	}
}
