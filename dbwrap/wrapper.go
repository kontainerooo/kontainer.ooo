/*
 * Package dbwrap provides a Wrapper for gorm.DB
 * With the help of that wrapper, services can become independent of an external db dependency
 * Furthermore it makes unit testing far easier, because db abstractions don't have to depend on gorm either
 */

package dbwrap

import (
	"github.com/jinzhu/gorm"
)

// Wrapper is an Interface to abstract gorm Database Functions
type Wrapper interface {
	// GetAffectedRows returns gorm.DB's RowsAffected property
	GetAffectedRows() int64

	// GetValue returns gorm.DB's Value property
	GetValue() interface{}

	// Every function below invokes gorm.DB's function and returns gorm.DB's Error property
	AutoMigrate(values ...interface{}) error
	Where(query interface{}, args ...interface{}) error
	First(out interface{}, where ...interface{}) error
	Create(value interface{}) error
}

type wrapper struct {
	db *gorm.DB
}

func (w *wrapper) GetAffectedRows() int64 {
	return w.db.RowsAffected
}

func (w *wrapper) GetValue() interface{} {
	return w.db.Value
}

func (w *wrapper) AutoMigrate(values ...interface{}) error {
	w.db = w.db.AutoMigrate(values...)
	return w.db.Error
}

func (w *wrapper) Where(query interface{}, args ...interface{}) error {
	w.db = w.db.Where(query, args...)
	return w.db.Error
}

func (w *wrapper) First(out interface{}, where ...interface{}) error {
	w.db = w.db.First(out, where...)
	return w.db.Error
}

func (w *wrapper) Create(value interface{}) error {
	w.db = w.db.Create(value)
	return w.db.Error
}

// NewWrapper returns an new Wrapper instance
func NewWrapper(db *gorm.DB) Wrapper {
	return &wrapper{
		db: db,
	}
}
