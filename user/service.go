// Package user provides functionalities to handle Users in the context of kontainer.io
package user

import "errors"

// The Service interface describes the function necessary for kontainer.io user handling
type Service interface {
	// CreateUser creates a new User and returns its id
	CreateUser(username string, cfg *Config, adr *Address) (uint, error)

	// EditUser is used to alter user information by id
	EditUser(id uint, cfg *Config) error

	// ChangeUsername is used to change a users username by id
	ChangeUsername(id uint, username string) error

	// DeleteUser is used to remove a user by id
	DeleteUser(id uint) error

	// ResetPassword is used to reset a users password and issue a reset Mail
	ResetPassword(email string) error

	// GetUser is used to gather a users data set by id
	GetUser(id uint, user *User) error
}

type dbAdapter interface {
	GetValue() interface{}
	AutoMigrate(values ...interface{}) error
	Where(query interface{}, args ...interface{}) error
	First(out interface{}, where ...interface{}) error
	Create(value interface{}) error
	Delete(value interface{}, where ...interface{}) error
	Update(attrs ...interface{}) error
}

type service struct {
	db dbAdapter
}

func (s *service) InitializeDatabases() error {
	return s.db.AutoMigrate(&Address{}, &User{}, &Customer{})
}

func (s *service) CreateUser(username string, cfg *Config, adr *Address) (uint, error) {
	s.db.Where("username = ?", username)
	res := s.db.GetValue()
	if res != nil && res != (&User{}) {
		return 0, errors.New("username already used")
	}

	err := s.db.Create(adr)
	if err != nil {
		return 0, err
	}

	cfg.AddressID = adr.ID
	user := &User{Username: username}
	user.setConfig(cfg)
	err = s.db.Create(user)
	if err != nil {
		//TODO: Delete Address/ Transactions?
		return 0, err
	}
	return user.ID, nil
}

func (s *service) EditUser(id uint, cfg *Config) error {
	err := s.db.Where("ID = ?", id)
	if err != nil {
		return err
	}

	err = s.db.Update((&User{}).setConfig(cfg))
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ChangeUsername(id uint, username string) error {
	err := s.db.Where("ID = ?", id)
	if err != nil {
		return err
	}

	err = s.db.Update(&User{Username: username})
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteUser(id uint) error {
	err := s.db.Delete(&User{
		ID: id,
	})
	return err
}

func (s *service) ResetPassword(email string) error {
	// TODO: implement functionality
	return nil
}

func (s *service) GetUser(id uint, user *User) error {
	err := s.db.Where("ID = ?", id)
	if err != nil {
		return err
	}

	err = s.db.First(user)
	if err != nil {
		return err
	}
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
