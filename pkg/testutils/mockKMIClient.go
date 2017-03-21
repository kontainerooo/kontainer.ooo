package testutils

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
)

// MockKMIClient simulates a kmi client for testing purposes
type MockKMIClient struct {
	kmis map[uint]kmi.KMI
}

// NewMockKMIEndpoints creates a new MockKMIClientEndpoints
func NewMockKMIEndpoints(logger log.Logger, m MockKMIClient) *kmi.Endpoints {
	return &kmi.Endpoints{
		GetKMIEndpoint: m.GetKMIEndpoint,
	}
}

// GetKMIEndpoint returns a given kmi
func (m *MockKMIClient) GetKMIEndpoint(ctx context.Context, req interface{}) (interface{}, error) {
	kmiReq := req.(*kmi.GetKMIRequest)

	k, ok := m.kmis[kmiReq.ID]
	if !ok {
		return &kmi.GetKMIResponse{}, errors.New("KMI not found")
	}

	return &kmi.GetKMIResponse{
		Error: nil,
		KMI:   &k,
	}, nil
}

// AddMockKMI adds a new KMI entry
func (m *MockKMIClient) AddMockKMI(id uint, kmi kmi.KMI) {
	m.kmis[id] = kmi
}

// NewMockKMIClient creates a new MockKMIClient
func NewMockKMIClient() *MockKMIClient {
	return &MockKMIClient{
		kmis: make(map[uint]kmi.KMI),
	}
}
