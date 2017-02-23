package testutils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// ErrDBFailure is returned, when the database should return an error
var ErrDBFailure = errors.New("database failure")

// MockDB simulates a database for testing purposes
type MockDB struct {
	Error  error
	tables map[string]*table
	err    int
	value  interface{}
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

// GetValue returns mockDB's value property
func (m *MockDB) GetValue() interface{} {
	return m.value
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

// Where is
func (m *MockDB) Where(query interface{}, args ...interface{}) error {
	if m.produceError() {
		return ErrDBFailure
	}
	s := strings.Split(query.(string), " ")
	field := strings.ToLower(s[0])
	field = strings.Title(field)
	for _, table := range m.tables {
		if table.checkForField(field) {
			result, err := table.find(field, args[0])
			m.value = result
			return err
		}
	}
	return nil
}

// First is
func (m MockDB) First(out interface{}, where ...interface{}) error {
	if m.produceError() {
		return ErrDBFailure
	}
	return nil
}

// Create is
func (m *MockDB) Create(value interface{}) error {
	if m.produceError() {
		return ErrDBFailure
	}
	ref := reflect.TypeOf(value).Elem()
	name := ref.String()
	return m.tables[name].insert(value)
}

// NewMockDB returns new MockDB
func NewMockDB() *MockDB {
	return &MockDB{
		tables: make(map[string]*table),
	}
}
