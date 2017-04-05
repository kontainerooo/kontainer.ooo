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

func (r *result) intersect(res reflect.Value) bool {
	hasElements := false
	_r := reflect.MakeSlice(r.values.Type(), 0, 0)
	resSlice := res.Slice(0, res.Len())
	rSlice := r.values.Slice(0, r.values.Len())

	for i := 0; i < resSlice.Len(); i++ {
		element := resSlice.Index(i)
		for j := 0; j < rSlice.Len(); j++ {
			if element.Interface() == rSlice.Index(j).Interface() {
				hasElements = true
				_r = reflect.Append(_r, element)
				break
			}
		}
	}

	r.values = _r
	return hasElements
}

// MockDB simulates a database for testing purposes
type MockDB struct {
	Error      error
	tables     map[string]*table
	rollback   map[string]*table
	err        int
	value      reflect.Value
	multiValue []*result
	isQuery    bool
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

// AppendToArray append to array
func (m *MockDB) AppendToArray(query interface{}, target string, values interface{}) error {
	if m.produceError() {
		return ErrDBFailure
	}
	v := reflect.ValueOf(query).Elem()
	t := v.Type()
	name := t.String()
	table, ok := m.tables[name]
	if !ok {
		return errors.New("table does not exist")
	}

	return table.appendToArray(v, target, values)
}

// RemoveFromArray remove from array
func (m *MockDB) RemoveFromArray(query interface{}, target string, index int) error {
	v := reflect.ValueOf(query).Elem()
	t := v.Type()
	name := t.String()
	table, ok := m.tables[name]
	if !ok {
		return errors.New("table does not exist")
	}

	return table.removeFromArray(v, target, index)
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

	// reset value and multiValue
	m.value = RNil
	m.multiValue = nil

	// split query at "AND"
	and := strings.Split(query.(string), " AND ")

	// search for results for each part connected with and
	for i, part := range and {
		// get field name out of query
		s := strings.Split(part, " ")
		field := s[0] //strings.Title(s[0])

		// init result array
		mValue := []*result{}

		// iterate available tables
		for _, table := range m.tables {
			if table.checkForField(field) {
				// get a slice of rows matching the query
				res, err := table.find(field, args[i])
				if err != nil {
					m.multiValue = nil
					return err
				}
				// update result array if the result isn't empty
				if res != RNil {
					mValue = append(mValue, &result{
						table:  table.Name,
						values: res,
					})
				}
			}
		}

		if i == 0 {
			// if nothing is found stop and return
			if mValue == nil || len(mValue) == 0 {
				return nil
			}

			// else set the multiValue property of the mockDB
			m.multiValue = mValue
		} else {
			// after the first iteration the results have to be intersected
			// to store only the ones which match all parts of the query
			for i, existingResult := range m.multiValue {
				intersect := false
				for _, res := range mValue {
					if existingResult.table == res.table {
						// intersect result values if table is in both result arrays
						// return value of intersect() is true if there is something left in the values slice
						intersect = existingResult.intersect(res.values)
						break
					}
				}
				if !intersect {
					// remove table from the multiValue property if it is not part of the new result array
					if len(m.multiValue) == 1 {
						m.multiValue = nil
						return nil
					}

					m.multiValue = append(m.multiValue[:i], m.multiValue[i+1:]...)
				}
			}
		}
	}

	// if there is only one result in multiValue this can be set as the result value (mockDB.value)
	if len(m.multiValue) == 1 {
		m.value = m.multiValue[0].values
	}

	m.isQuery = true

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
	m.isQuery = false
	return nil
}

// Find mocks gorm.DBs Find function
func (m *MockDB) Find(out interface{}, where ...interface{}) error {
	if m.produceError() {
		return ErrDBFailure
	}

	defer func() { m.isQuery = false }()

	ref := reflect.TypeOf(out).Elem()
	for _, t := range m.tables {
		if ref == reflect.SliceOf(t.getRef()) {
			if !m.isQuery {
				err := t.all(out)
				if err != nil {
					return err
				}
				return nil
			}
			if reflect.TypeOf(out).Elem().Kind() == reflect.Slice {
				for k, v := range m.multiValue {
					val := v.values.Slice(0, 1).Index(k).Elem()
					reflect.ValueOf(out).Elem().Set(reflect.Append(reflect.ValueOf(out).Elem(), val))
				}
			}
			break
		}
	}
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
	ids := make(map[string]reflect.Value)

	ref := reflect.TypeOf(value).Elem()
	name := ref.String()

	v := reflect.ValueOf(value).Elem()

	for i := 0; i < ref.NumField(); i++ {
		f := ref.Field(i).Name
		fv := v.Field(i)
		if !isZero(fv) {
			ids[f] = fv
		}
	}

	if len(ids) != 0 {
		return m.tables[name].delete(ids)
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
