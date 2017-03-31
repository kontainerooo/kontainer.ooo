package testutils

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/firewall"
)

// MockFirewallClient simulates a firewall client for testing purposes
type MockFirewallClient struct {
}

// NewMockFirewallEndpoints creates a new MockFirewallClientEndpoints
func NewMockFirewallEndpoints(logger log.Logger, m MockFirewallClient) *firewall.Endpoints {
	return &firewall.Endpoints{
		AllowPortEndpoint: m.AllowPortEndpoint,
	}
}

// AllowPortEndpoint is a mock endpoint
func (m *MockFirewallClient) AllowPortEndpoint(ctx context.Context, req interface{}) (interface{}, error) {
	_ = req.(*firewall.AllowPortRequest)

	return &firewall.AllowPortResponse{
		Error: nil,
	}, nil
}

// NewMockFirewallClient creates a new MockFirewallClient
func NewMockFirewallClient() *MockFirewallClient {
	return &MockFirewallClient{}
}
