// Package kmi provides functionality to handle kmi files
package kmi

import (
	"fmt"
	"sync"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
)

// The Service interface describes the functions necessary for a KMI Service
type Service interface {
	// AddKMI resolves the path to a kmi file, extracts it and adds its contents to the database as a new kontainer module
	AddKMI(path string) (id uint, err error)

	// RemoveKMI removes the kontainer module information and all files related
	RemoveKMI(id uint) error

	// GetKMI retrieves kontainer module information for a specific module
	GetKMI(id uint, k *KMI) error

	// KMI returns display information for all exisiting kontainer modules
	KMI(*[]KMDI) error
}

type dbAdapter interface {
	abstraction.DBAdapter
	AutoMigrate(...interface{}) error
	Where(interface{}, ...interface{}) error
	First(interface{}, ...interface{}) error
	Find(interface{}, ...interface{}) error
	Create(interface{}) error
	Delete(interface{}, ...interface{}) error
}

type service struct {
	db  dbAdapter
	mtx *sync.Mutex
}

func (s *service) InitializeDatabases() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.initializeDatabases()
}

func (s *service) initializeDatabases() error {
	return s.db.AutoMigrate(&KMI{})
}

func (s *service) AddKMI(path string) (uint, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.addKMI(path)
}

func (s *service) addKMI(path string) (uint, error) {
	kC := NewContent()
	err := Extract(path, kC)
	if err != nil {
		return 0, err
	}

	k := &KMI{}
	err = GetData(kC, k)
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
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.removeKMI(id)
}

func (s *service) removeKMI(id uint) error {
	k := &KMI{}
	k.ID = id
	err := s.db.Delete(k)
	return err
}

func (s *service) GetKMI(id uint, k *KMI) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.getKMI(id, k)
}

func (s *service) getKMI(id uint, k *KMI) error {
	err := s.db.First(k, "ID = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) KMI(out *[]KMDI) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.kmi(out)
}

func (s *service) kmi(out *[]KMDI) error {
	k := []KMI{}
	err := s.db.Find(&k)
	if err != nil {
		return err
	}
	for _, kmi := range k {
		*out = append(*out, kmi.KMDI)
	}
	return nil
}

// NewService creates a KMIService with necessary dependencies.
func NewService(db dbAdapter) (Service, error) {
	s := &service{
		db:  db,
		mtx: &sync.Mutex{},
	}

	err := s.initializeDatabases()
	if err != nil {
		return nil, err
	}

	return s, nil
}
