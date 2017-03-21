package routing

import (
	"errors"
	"fmt"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
)

// The Service interface describes the functions necessary for kontainer.ooo routing service
type Service interface {
	// Insert a new configuration into the DB, write to a file, update the router
	CreateRouterConfig(r *RouterConfig) error

	// Edit an existing configuration by refID and name, update file and router
	EditRouterConfig(refID uint, name string, r *RouterConfig) error

	// Get an existing configuration by refID and name
	GetRouterConfig(refID uint, name string, r *RouterConfig) error

	// Remove a configuration by id, remove file, update router
	RemoveRouterConfig(refID uint, name string) error

	// Add a location rule to a configuration by id, update file and router
	AddLocationRule(id uint, lr *LocationRule) (int, error)

	// Remove a location rule by its id in a configuration by id, update file and router
	RemoveLocationRule(id uint, lid int) error

	// Chante the listen statement of a configuration by id, update file and router
	ChangeListenStatement(id uint, ls *ListenStatement) error

	// Change the server name(s) of a configuration by id, update file and router
	ChangeServerName(id uint, sn []string) error

	// Configuration returns all Configurations
	Configurations() []RouterConfig
}

type dbAdapter interface {
	abstraction.DBAdapter
	AutoMigrate(...interface{}) error
	Where(interface{}, ...interface{}) error
	First(interface{}, ...interface{}) error
	Create(interface{}) error
	Update(interface{}, ...interface{}) error
	Delete(interface{}, ...interface{}) error
}

type service struct {
	db dbAdapter
}

func (s service) InitializeDatabases() error {
	return s.db.AutoMigrate(&RouterConfig{})
}

func (s *service) CreateRouterConfig(r *RouterConfig) error {
	s.db.Where("RefID = ? AND Name = ?", r.RefID, r.Name)
	res := s.db.GetValue()
	if res != nil && res != (&RouterConfig{}) {
		return fmt.Errorf("config with name %s for user %d already exists", r.Name, r.RefID)
	}

	err := s.db.Create(r)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) EditRouterConfig(refID uint, name string, r *RouterConfig) error {
	if r.RefID != 0 && refID != r.RefID {
		return errors.New("can not change reference id")
	}

	s.db.Begin()
	err := s.db.Where("RefID = ? AND Name = ?", refID, name)
	if err != nil {
		s.db.Rollback()
		return err
	}

	err = s.db.Update(&RouterConfig{}, r)
	if err != nil {
		s.db.Rollback()
		return err
	}
	s.db.Commit()
	return nil
}

func (s *service) GetRouterConfig(refID uint, name string, r *RouterConfig) error {
	s.db.Begin()
	err := s.db.Where("RefID = ? AND Name = ?", refID, name)
	if err != nil {
		s.db.Rollback()
		return err
	}

	err = s.db.First(r)
	if err != nil {
		s.db.Rollback()
		return err
	}

	s.db.Commit()
	return nil
}

func (s *service) RemoveRouterConfig(refID uint, name string) error {
	if refID == 0 || name == "" {
		return errors.New("refID and name have to be set")
	}

	return s.db.Delete(&RouterConfig{
		RefID: refID,
		Name:  name,
	})
}

func (s *service) AddLocationRule(id uint, lr *LocationRule) (int, error) {
	// TODO: implement
	return 0, nil
}

func (s *service) RemoveLocationRule(id uint, lid int) error {
	// TODO: implement
	return nil
}

func (s *service) ChangeListenStatement(id uint, ls *ListenStatement) error {
	// TODO: implement
	return nil
}

func (s *service) ChangeServerName(id uint, sn []string) error {
	// TODO: implement
	return nil
}

func (s *service) Configurations() []RouterConfig {
	// TODO: implement
	return nil
}

// NewService creates a UserService with necessary dependencies.
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