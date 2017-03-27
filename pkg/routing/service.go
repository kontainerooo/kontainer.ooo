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
	AddLocationRule(refID uint, name string, lr *LocationRule) error

	// Remove a location rule by its id in a configuration by id, update file and router
	RemoveLocationRule(refID uint, name string, lid int) error

	// Chante the listen statement of a configuration by id, update file and router
	ChangeListenStatement(refID uint, name string, ls *ListenStatement) error

	// Add something to the server name of a configuration by id, update file and router
	AddServerName(refID uint, name string, sn string) error

	// Add something to the server name of a configuration by id, update file and router
	RemoveServerName(refID uint, name string, id int) error

	// Configuration returns all Configurations
	Configurations(r *[]RouterConfig)
}

type dbAdapter interface {
	abstraction.DBAdapter
	AutoMigrate(...interface{}) error
	Where(interface{}, ...interface{}) error
	First(interface{}, ...interface{}) error
	Find(interface{}, ...interface{}) error
	Create(interface{}) error
	Update(interface{}, ...interface{}) error
	Delete(interface{}, ...interface{}) error
	AppendToArray(interface{}, string, interface{}) error
	RemoveFromArray(interface{}, string, int) error
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

func (s *service) AddLocationRule(refID uint, name string, lr *LocationRule) error {
	s.db.Begin()

	err := s.db.AppendToArray(&RouterConfig{
		RefID: refID,
		Name:  name,
	}, "LocationRules", lr)
	if err != nil {
		s.db.Rollback()
		return err
	}

	s.db.Commit()
	return nil
}

func (s *service) RemoveLocationRule(refID uint, name string, lid int) error {
	s.db.Begin()

	err := s.db.RemoveFromArray(&RouterConfig{
		RefID: refID,
		Name:  name,
	}, "LocationRules", lid)
	if err != nil {
		s.db.Rollback()
		return err
	}

	s.db.Commit()
	return nil
}

func (s *service) ChangeListenStatement(refID uint, name string, ls *ListenStatement) error {
	err := s.db.Where("RefID = ? AND Name = ?", refID, name)
	if err != nil {
		return err
	}

	err = s.db.Update(&RouterConfig{}, &RouterConfig{
		ListenStatement: *ls,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *service) AddServerName(refID uint, name string, sn string) error {
	s.db.Begin()

	err := s.db.AppendToArray(&RouterConfig{
		RefID: refID,
		Name:  name,
	}, "ServerName", sn)
	if err != nil {
		s.db.Rollback()
		return err
	}

	s.db.Commit()
	return nil
}

func (s *service) RemoveServerName(refID uint, name string, id int) error {
	s.db.Begin()

	err := s.db.RemoveFromArray(&RouterConfig{
		RefID: refID,
		Name:  name,
	}, "ServerName", id)
	if err != nil {
		s.db.Rollback()
		return err
	}

	s.db.Commit()
	return nil
}

func (s *service) Configurations(r *[]RouterConfig) {
	s.db.Find(r)
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
