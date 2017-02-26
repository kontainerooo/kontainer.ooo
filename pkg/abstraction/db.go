/*
 * Package dbwrap provides a Wrapper for gorm.DB
 * With the help of that wrapper, services can become independent of an external db dependency
 * Furthermore it makes unit testing far easier, because db abstractions don't have to depend on gorm either
 */

package abstraction

import "github.com/jinzhu/gorm"

// DBAdapter includes functions used for Transactions and Getters
type DBAdapter interface {
	Begin()
	Rollback()
	Commit()
	GetValue() interface{}
	GetAffectedRows() int64
}

// DB is an Interface to abstract gorm Database Functions
type DB interface {
	// GetAffectedRows returns gorm.DB's RowsAffected property
	GetAffectedRows() int64

	// GetValue returns gorm.DB's Value property
	GetValue() interface{}

	// Every function below invokes gorm.DB's function and returns gorm.DB's Error property
	Begin()
	Rollback()
	Commit()
	AutoMigrate(values ...interface{}) error
	Where(query interface{}, args ...interface{}) error
	First(out interface{}, where ...interface{}) error
	Create(value interface{}) error
	Delete(value interface{}, where ...interface{}) error
	Update(model interface{}, attrs ...interface{}) error
}

type dbWrapper struct {
	db *gorm.DB
	tx *gorm.DB
}

func (w *dbWrapper) GetAffectedRows() int64 {
	return w.db.RowsAffected
}

func (w *dbWrapper) GetValue() interface{} {
	return w.db.Value
}

func (w *dbWrapper) Begin() {
	w.tx = w.db.Begin()
}

func (w *dbWrapper) resetTransaction() {
	w.tx = nil
}

func (w *dbWrapper) Rollback() {
	defer w.resetTransaction()
	w.tx.Rollback()
}

func (w *dbWrapper) Commit() {
	defer w.resetTransaction()
	w.tx.Commit()
}

func (w *dbWrapper) AutoMigrate(values ...interface{}) error {
	if w.tx != nil {
		return w.tx.AutoMigrate(values...).Error
	}
	return w.db.AutoMigrate(values...).Error
}

func (w *dbWrapper) Where(query interface{}, args ...interface{}) error {
	if w.tx != nil {
		w.tx = w.tx.Where(query, args...)
		return w.tx.Error
	}
	return w.db.Where(query, args...).Error
}

func (w *dbWrapper) First(out interface{}, where ...interface{}) error {
	if w.tx != nil {
		return w.tx.First(out, where...).Error
	}
	return w.db.First(out, where...).Error
}

func (w *dbWrapper) Create(value interface{}) error {
	if w.tx != nil {
		return w.tx.Create(value).Error
	}
	return w.db.Create(value).Error
}

func (w *dbWrapper) Delete(value interface{}, where ...interface{}) error {
	if w.tx != nil {
		return w.tx.Delete(value, where...).Error
	}
	return w.db.Delete(value, where...).Error
}

func (w *dbWrapper) Update(model interface{}, attrs ...interface{}) error {
	if w.tx != nil {
		return w.tx.Model(model).Update(attrs...).Error
	}
	return w.db.Model(model).Update(attrs...).Error
}

// NewDB returns an new Wrapper instance
func NewDB(db *gorm.DB) DB {
	return &dbWrapper{
		db: db,
	}
}
