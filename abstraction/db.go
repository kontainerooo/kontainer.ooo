/*
 * Package dbwrap provides a Wrapper for gorm.DB
 * With the help of that wrapper, services can become independent of an external db dependency
 * Furthermore it makes unit testing far easier, because db abstractions don't have to depend on gorm either
 */

package abstraction

import (
	"github.com/jinzhu/gorm"
)

// DB is an Interface to abstract gorm Database Functions
type DB interface {
	// GetAffectedRows returns gorm.DB's RowsAffected property
	GetAffectedRows() int64

	// GetValue returns gorm.DB's Value property
	GetValue() interface{}

	// Every function below invokes gorm.DB's function and returns gorm.DB's Error property
	AutoMigrate(values ...interface{}) error
	Where(query interface{}, args ...interface{}) error
	First(out interface{}, where ...interface{}) error
	Create(value interface{}) error
	Delete(value interface{}, where ...interface{}) error
	Update(attrs ...interface{}) error
}

type dbWrapper struct {
	db    *gorm.DB
	scope *gorm.DB
}

func (w *dbWrapper) GetAffectedRows() int64 {
	return w.db.RowsAffected
}

func (w *dbWrapper) GetValue() interface{} {
	return w.db.Value
}

func (w *dbWrapper) AutoMigrate(values ...interface{}) error {
	w.scope = w.db.AutoMigrate(values...)
	return w.scope.Error
}

func (w *dbWrapper) Where(query interface{}, args ...interface{}) error {
	w.scope = w.db.Where(query, args...)
	return w.scope.Error
}

func (w *dbWrapper) First(out interface{}, where ...interface{}) error {
	w.scope = w.scope.First(out, where...)
	return w.scope.Error
}

func (w *dbWrapper) Create(value interface{}) error {
	w.db = w.db.Create(value)
	return w.db.Error
}

func (w *dbWrapper) Delete(value interface{}, where ...interface{}) error {
	w.scope = w.db.Delete(value, where...)
	return w.scope.Error
}

func (w *dbWrapper) Update(attrs ...interface{}) error {
	w.scope = w.db.Update(attrs...)
	return w.db.Error
}

// NewDB returns an new Wrapper instance
func NewDB(db *gorm.DB) DB {
	return &dbWrapper{
		db: db,
	}
}
