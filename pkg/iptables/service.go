// Package iptables is a wrapper around the iptables binary
package iptables

import "os/exec"

// Service handles iptables rules
type Service interface {
	// AddRule adds a given iptables rule
	AddRule(refid uint, rule Rule) error
	// RemoveRule removes a given iptables rule
	RemoveRule() error
	// GetRulesForUser returns a list of all rules for a given user
	GetRulesForUser(refid uint) []Rule
	// CreateIPTablesBackup creates an iptables backup file
	CreateIPTablesBackup() string
	// LoadIPTablesBackup restores iptables from backup file
	LoadIPTablesBackup() error
}

type service struct {
	iptPath string
}

func (w *service) AddRule(refid uint, rule Rule) error {
	// TODO: implement
	return nil
}

func (w *service) RemoveRule() error {
	// TODO: implement
	return nil
}

func (w *service) GetRulesForUser(refid uint) []Rule {
	// TODO: implement
	return []Rule{}
}

func (w *service) CreateIPTablesBackup() string {
	// TODO: implement
	return ""
}

func (w *service) LoadIPTablesBackup() error {
	// TODO: implement
	return nil
}

// ExecCommand is a wrapper around exec.Command used for testing
var ExecCommand = exec.Command

// NewService creates a new iptables service
func NewService(iptPath string) (Service, error) {
	cmd := ExecCommand(iptPath, "--version")
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return &service{
		iptPath: iptPath,
	}, nil
}
