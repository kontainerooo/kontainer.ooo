package testutils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	// ErrDBFailure is returned, when the database should return an error
	ErrDBFailure = errors.New("database failure")

	// ErrNotFound is returned, when the database couldn't find the requested entity in functions which rely on finding it
	ErrNotFound = errors.New("entity not found")

	// ErrTypeMismatch is returned, when the type passed to the function isn't compatible to the type of the data addressed
	ErrTypeMismatch = errors.New("type mismatch")

	// RNil is the value reflection of nil
	RNil = reflect.ValueOf(nil)
)

type result struct {
	table  string
	values reflect.Value
}

// MockDB simulates a database for testing purposes
type MockDB struct {
	Error      error
	tables     map[string]*table
	rollback   map[string]*table
	err        int
	value      reflect.Value
	multiValue []*result
}

// SetError sets the err property to true, causing the next function to be invoked next to return an error
func (m *MockDB) SetError(i int) {
	m.err = i
}

func (m *MockDB) produceError() bool {
	if m.err == 1 {
		m.err--
		return true
	} else if m.err > 0 {
		m.err--
	}
	return false
}

// PrintTables prints every table in the database with its current rows
func (m *MockDB) PrintTables() {
	fmt.Print("MockDB\n")
	for _, t := range m.tables {
		fmt.Printf("%v\n", t)
	}
}

// Begin begin transaction
func (m *MockDB) Begin() {
	m.rollback = make(map[string]*table)
	for name, t := range m.tables {
		m.rollback[name] = t.copy()
	}
}

// Rollback rollback transaction
func (m *MockDB) Rollback() {
	m.tables = m.rollback
	m.rollback = nil
}

// Commit commit transaction
func (m *MockDB) Commit() {
	m.rollback = nil
}

// GetValue returns mockDB's value property
func (m *MockDB) GetValue() interface{} {
	if m.value == RNil {
		return nil
	}
	return m.value
}

// GetAffectedRows returns 0
func (m *MockDB) GetAffectedRows() int64 {
	return 0
}

// AutoMigrate creates new Tables in the database based on the values array
func (m *MockDB) AutoMigrate(values ...interface{}) error {
	if m.produceError() {
		return ErrDBFailure
	}
	for _, model := range values {
		ref := reflect.TypeOf(model).Elem()
		name := ref.String()
		m.tables[name] = newTable(ref, name)
	}
	return nil
}

// Where mocks gorm.DBs Where function
func (m *MockDB) Where(query interface{}, args ...interface{}) error {
	if m.produceError() {
		return ErrDBFailure
	}
	m.value = RNil
	m.multiValue = nil

	s := strings.Split(query.(string), " ")
	field := strings.Title(s[0])
	for _, table := range m.tables {
		if table.checkForField(field) {
			res, err := table.find(field, args[0])
			if err != nil {
				m.value = RNil
				m.multiValue = nil
				return err
			}
			if res != reflect.ValueOf(nil) {
				m.multiValue = append(m.multiValue, &result{
					table:  table.Name,
					values: res,
				})
			}
		}
	}
	if len(m.multiValue) == 1 {
		m.value = m.multiValue[0].values
	}
	return nil
}

// First mocks gorm.DBs First function
func (m *MockDB) First(out interface{}, where ...interface{}) error {
	if m.produceError() {
		return ErrDBFailure
	}

	ref := reflect.TypeOf(out).Elem()
	name := ref.String()

	if m.multiValue == nil || len(m.multiValue) == 0 {
		return ErrNotFound
	}

	for _, res := range m.multiValue {
		if res.table == name {
			src := res.values.Slice(0, 1).Index(0).Elem()
			dst := reflect.ValueOf(out).Elem()
			err := merge(dst, src, true, 0)
			if err != nil {
				return err
			}
			break
		}
	}

	m.value, m.multiValue = RNil, nil
	return nil
}

// Create mocks gorm.DBs Create function
func (m *MockDB) Create(value interface{}) error {
	if m.produceError() {
		return ErrDBFailure
	}
	ref := reflect.TypeOf(value).Elem()
	name := ref.String()
	return m.tables[name].insert(value)
}

// Delete mocks gorm.DBs Delete function
func (m *MockDB) Delete(value interface{}, where ...interface{}) error {
	if m.produceError() {
		return ErrDBFailure
	}
	id := reflect.ValueOf(value).Elem().FieldByName("ID").Uint()
	if id != 0 {
		ref := reflect.TypeOf(value).Elem()
		name := ref.String()
		return m.tables[name].delete(id)
	}
	return ErrDBFailure
}

// Update mocks gorm.DBs Update function
func (m *MockDB) Update(model interface{}, attrs ...interface{}) error {
	if m.produceError() {
		return ErrDBFailure
	}

	if m.value == RNil && m.multiValue != nil {
		ref := reflect.TypeOf(model).Elem()
		name := ref.String()
		for _, res := range m.multiValue {
			if res.table == name {
				m.value = res.values
			}
		}
	}

	if m.value != RNil && m.value.Len() > 0 {
		len := m.value.Len()
		slice := m.value.Slice(0, len)
		for i := 0; i < len; i++ {
			err := merge(slice.Index(i), reflect.ValueOf(attrs[0]).Elem(), false, 0) // apply to table?
			if err != nil {
				return err
			}
		}

	}
	return nil
}

// NewMockDB returns new MockDB
func NewMockDB() *MockDB {
	return &MockDB{
		tables: make(map[string]*table),
	}
}
