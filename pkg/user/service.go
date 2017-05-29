// Package user provides functionalities to handle Users in the context of kontainer.io
package user

import (
	"crypto/rand"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
)

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

	// CheckLoginCredentials is used to check the login credentials of a user
	CheckLoginCredentials(username string, password string) uint

	getDB() abstraction.DBAdapter
}

type dbAdapter interface {
	abstraction.DBAdapter
	AutoMigrate(...interface{}) error
	Where(interface{}, ...interface{}) error
	First(interface{}, ...interface{}) error
	Create(interface{}) error
	Delete(interface{}, ...interface{}) error
	Update(interface{}, ...interface{}) error
}

type service struct {
	db         dbAdapter
	bcryptCost int
}

func (s *service) InitializeDatabases() error {
	return s.db.AutoMigrate(&Address{}, &User{}, &Customer{})
}

func (s *service) getDB() abstraction.DBAdapter {
	return s.db
}

func (s *service) CreateUser(username string, cfg *Config, adr *Address) (uint, error) {
	user, err := s.GetUserByUsername(username)
	if user != nil && user.ID != 0 {
		return 0, errors.New("username already used")
	}
	if err != nil {
		return 0, err
	}

	err = s.db.Create(adr)
	if err != nil {
		return 0, err
	}

	cfg.AddressID = adr.ID
	user = &User{Username: username}
	user.setConfig(cfg)

	count := 512
	salt := make([]byte, count)
	_, err = rand.Read(salt)
	if err != nil {
		return 0, err
	}
	user.Salt = string(fmt.Sprintf("%x", salt))

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password+user.Salt), s.bcryptCost)
	if err != nil {
		return 0, err
	}

	user.Password = string(password)

	err = s.db.Create(user)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (s *service) EditUser(id uint, cfg *Config) error {
	err := s.db.Where("ID = ?", id)
	if err != nil {
		return err
	}

	err = s.db.Update(&User{}, (&User{}).setConfig(cfg))
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

	err = s.db.Update(&User{}, &User{Username: username})
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
	err := s.db.First(user, "ID = ?", id)
	if err != nil {
		return err
	}

	err = s.db.First(&user.Address, "ID = ?", user.AddressID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetUserByUsername(username string) (*User, error) {
	user := &User{}
	err := s.db.Where("Username = ?", username)
	if err != nil {
		return nil, err
	}
	err = s.db.First(user)
	if err != nil && !s.db.IsNotFound(err) {
		return nil, err
	}

	return user, nil
}

func (s *service) CheckLoginCredentials(username string, password string) uint {
	user, err := s.GetUserByUsername(username)
	if err != nil || user == nil {
		return 0
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+user.Salt)) == nil {
		return user.ID
	}

	return 0
}

// NewService creates a UserService with necessary dependencies.
func NewService(db dbAdapter, bcryptCost int) (Service, error) {
	s := &service{
		db:         db,
		bcryptCost: bcryptCost,
	}

	err := s.InitializeDatabases()
	if err != nil {
		return nil, err
	}

	return s, nil
}
