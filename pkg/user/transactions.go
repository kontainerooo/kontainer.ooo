package user

import (
	"sync"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
)

type transactionBasedService struct {
	s   Service
	db  abstraction.DBAdapter
	mtx *sync.Mutex
}

func (t *transactionBasedService) CreateUser(username string, cfg *Config, adr *Address) (uint, error) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	t.db.Begin()
	id, err := t.s.CreateUser(username, cfg, adr)
	if err != nil {
		t.db.Rollback()
		return 0, err
	}
	t.db.Commit()
	return id, nil
}

func (t *transactionBasedService) EditUser(id uint, cfg *Config) error {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	t.db.Begin()
	err := t.s.EditUser(id, cfg)
	if err != nil {
		t.db.Rollback()
		return err
	}
	t.db.Commit()
	return nil
}

func (t *transactionBasedService) ChangeUsername(id uint, username string) error {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	t.db.Begin()
	err := t.s.ChangeUsername(id, username)
	if err != nil {
		t.db.Rollback()
		return err
	}
	t.db.Commit()
	return nil
}

func (t *transactionBasedService) DeleteUser(id uint) error {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	t.db.Begin()
	err := t.s.DeleteUser(id)
	if err != nil {
		t.db.Rollback()
		return err
	}
	t.db.Commit()
	return nil
}

func (t *transactionBasedService) CheckLoginCredentials(username string, password string) uint {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	return t.s.CheckLoginCredentials(username, password)
}

func (t *transactionBasedService) ResetPassword(email string) error {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	t.db.Begin()
	err := t.s.ResetPassword(email)
	if err != nil {
		t.db.Rollback()
		return err
	}
	t.db.Commit()
	return nil
}

func (t *transactionBasedService) GetUser(id uint, user *User) error {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	err := t.s.GetUser(id, user)
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionBasedService) getDB() abstraction.DBAdapter {
	return t.db
}

// NewTransactionBasedService returns a new transactionBasedService
func NewTransactionBasedService(s Service) Service {
	return &transactionBasedService{
		s:   s,
		db:  s.getDB(),
		mtx: &sync.Mutex{},
	}
}
