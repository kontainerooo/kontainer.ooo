/*
 * Package dbwrap provides a Wrapper for gorm.DB
 * With the help of that wrapper, services can become independent of an external db dependency
 * Furthermore it makes unit testing far easier, because db abstractions don't have to depend on gorm either
 */

package abstraction

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
)

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

	AppendToArray(query interface{}, target string, values interface{}) error
	RemoveFromArray(query interface{}, target string, index int) error

	// Every function below invokes gorm.DB's function and returns gorm.DB's Error property
	Begin()
	Rollback()
	Commit()
	AutoMigrate(values ...interface{}) error
	Where(query interface{}, args ...interface{}) error
	First(out interface{}, where ...interface{}) error
	Find(out interface{}, where ...interface{}) error
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

func (w *dbWrapper) getTableName(v reflect.Value) string {
	t := v.Type()

	// Get TableName, either from type name or if it exists from TableName function
	nameFunc, exists := t.MethodByName("TableName")
	if exists {
		nameResult := nameFunc.Func.Call([]reflect.Value{v})
		return nameResult[0].String()
	}
	return gorm.ToDBName(t.Name())
}

func (w *dbWrapper) generateWhereClause(query *string, v reflect.Value) error {
	var primary []string

	t := v.Type()

	// Search for primary keys in the reference type
	idTag := "ID"
	_, found := t.FieldByName(idTag)
	if found {
		primary = append(primary, idTag)
	}

	primaryRegExp := regexp.MustCompile("primary_key")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		v, ok := f.Tag.Lookup("gorm")
		if ok {
			isPrimary := primaryRegExp.MatchString(v)
			if isPrimary {
				primary = append(primary, f.Name)
			}
		}
	}

	if len(primary) == 0 {
		return errors.New("no primary key found")
	}

	// Set WHERE clause
	*query = fmt.Sprintf("%s WHERE", *query)
	for i, key := range primary {
		if i > 0 {
			*query = fmt.Sprintf("%s AND", *query)
		}
		value := v.FieldByName(key)
		key = gorm.ToDBName(key)
		k := value.Kind()
		if k == reflect.String {
			*query = fmt.Sprintf("%s %s='%s'", *query, key, value.String())
		} else if k == reflect.Uint {
			*query = fmt.Sprintf("%s %s=%d", *query, key, value.Uint())
		} else {
			return fmt.Errorf("unexpected %s", k)
		}

	}

	return nil
}

func (w *dbWrapper) AppendToArray(elem interface{}, target string, value interface{}) error {
	var (
		tableName string
		stringVal string
		query     string
		typeCast  string
	)

	v := reflect.ValueOf(elem)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()

	// check if target is field of reference type
	f, isPart := t.FieldByName(strings.Title(target))
	if !isPart {
		return fmt.Errorf("%s is not a column in destination table", target)
	}
	target = gorm.ToDBName(target)

	tag := f.Tag.Get("sql")
	re := regexp.MustCompile(`type:(.+)\[\]`)
	match := re.FindStringSubmatch(tag)
	if len(match) < 2 {
		return fmt.Errorf("unsupported type %s", tag)
	}
	typeCast = match[1]

	tableName = w.getTableName(v)

	valuer, exists := reflect.TypeOf(value).MethodByName("Value")

	if exists {
		valuerResult := valuer.Func.Call([]reflect.Value{reflect.ValueOf(value)})
		stringVal = valuerResult[0].Interface().(string)
	} else {
		var ok bool
		stringVal, ok = value.(string)
		if !ok {
			return errors.New("type not supported")
		}
	}

	// Set Update parameters
	query = fmt.Sprintf("UPDATE %s SET %s = %s || $1::%s", tableName, target, target, typeCast)

	err := w.generateWhereClause(&query, v)
	if err != nil {
		return err
	}

	query = fmt.Sprintf("%s RETURNING cardinality(%s) AS index", query, target)

	if w.tx != nil {
		_, err = w.tx.CommonDB().Exec(query, stringVal)
		return err
	}

	_, err = w.db.DB().Exec(query, stringVal)
	return err
}

func (w *dbWrapper) RemoveFromArray(elem interface{}, target string, index int) error {
	var query string

	v := reflect.ValueOf(elem)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()

	_, isPart := t.FieldByName(strings.Title(target))
	if !isPart {
		return fmt.Errorf("%s is not a column in destination table", target)
	}
	target = gorm.ToDBName(target)

	tableName := w.getTableName(v)

	query = fmt.Sprintf("UPDATE %s SET %s = array_remove(%s, %s[$1])", tableName, target, target, target)

	err := w.generateWhereClause(&query, v)
	if err != nil {
		return err
	}

	if w.tx != nil {
		_, err = w.tx.CommonDB().Exec(query, index)
		return err
	}

	_, err = w.db.DB().Exec(query, index)
	return err
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

func (w *dbWrapper) Find(out interface{}, where ...interface{}) error {
	if w.tx != nil {
		return w.tx.Find(out, where...).Error
	}
	return w.db.Find(out, where...).Error
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
