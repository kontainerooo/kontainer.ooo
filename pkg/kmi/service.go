// Package kmi provides functionality to handle kmi files
package kmi

import (
	"fmt"

	"github.com/ttdennis/kontainer.io/pkg/abstraction"
)

// The Service interface describes the functions necessary for a KMI Service
type Service interface {
	// AddKMI resolves the path to a kmi file, extracts it and adds its contents to the database as a new kontainer module
	AddKMI(path string) (id uint, err error)

	// RemoveKMI removes the kontainer module information and all files related
	RemoveKMI(id uint) error

	// GetKMI retrieves kontainer module information for a specific module
	GetKMI(id uint) (*KMI, error)

	// KMI returns display information for all exisiting kontainer modules
	KMI() *[]KMDI
}

type dbAdapter interface {
	abstraction.DBAdapter
	AutoMigrate(...interface{}) error
	Where(interface{}, ...interface{}) error
	Create(interface{}) error
	Delete(interface{}, ...interface{}) error
}

type service struct {
	db dbAdapter
}

func (s *service) InitializeDatabases() error {
	return s.db.AutoMigrate(&KMI{})
}

func (s *service) AddKMI(path string) (uint, error) {
	kC := &kmiContent{}
	err := extract(path, kC)
	if err != nil {
		return 0, err
	}

	k := &KMI{}
	err = getData(kC, k)
	if err != nil {
		return 0, err
	}

	s.db.Where("name = ?", k.Name)
	res := s.db.GetValue()
	if res != nil && res != (&KMI{}) {
		return 0, fmt.Errorf("%s already exists", k.Name)
	}

	err = s.db.Create(k)
	if err != nil {
		return 0, err
	}
	return k.ID, nil
}

func (s *service) RemoveKMI(id uint) error {
	k := &KMI{}
	k.ID = id
	err := s.db.Delete(k)
	return err
}

func (s *service) GetKMI(id uint) (*KMI, error) {
	//TODO: implement
	return nil, nil
}

func (s *service) KMI() *[]KMDI {
	//TODO: implement
	return nil
}

// NewService creates a KMIService with necessary dependencies.
func NewService(db dbAdapter) (Service, error) {
	s := &service{
		db: db,
	}

	err := s.InitializeDatabases()
	if err != nil {
		return nil, err
	}

	return s, nil
}
