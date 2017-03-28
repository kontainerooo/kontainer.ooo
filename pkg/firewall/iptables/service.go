// Package iptables is a wrapper around the iptables binary
package iptables

import (
	"os/exec"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
)

// Service handles iptables rules
type Service interface {
	// CreateRule creates a rule with a given type and data
	CreateRule(ruleType int, ruleData interface{}) error
}

type dbAdapter interface {
	abstraction.DBAdapter
	AutoMigrate(...interface{}) error
	Where(interface{}, ...interface{}) error
	Create(interface{}) error
	First(interface{}, ...interface{}) error
	Find(interface{}, ...interface{}) error
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
	return s.db.AutoMigrate(&ruleEntry{})
}

func (s *service) CreateRule(ruleType int, ruleData interface{}) error {
	// TODO: implement
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
