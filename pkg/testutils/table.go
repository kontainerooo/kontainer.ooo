package testutils

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"

	"github.com/jinzhu/gorm"
)

// Table simulates a Table in the MockDb
type table struct {
	Name       string
	PrimaryKey []string
	rows       []interface{}
	ref        reflect.Type
	idx        []string
	count      uint
	columns    map[string]string
}

func (t *table) String() string {
	s := fmt.Sprintf("Table: %s\n", t.Name)
	for i, r := range t.rows {
		s += fmt.Sprintf("ID: %d - %v\n", i, r)
	}
	return s
}

func (t *table) getRef() reflect.Type {
	return t.ref
}

func (t *table) copy() *table {
	nt := &table{
		Name:       t.Name,
		PrimaryKey: t.PrimaryKey,
		ref:        t.ref,
		idx:        t.idx,
		count:      t.count,
		columns:    t.columns,
	}

	for _, r := range t.rows {
		nt.rows = append(nt.rows, *(&r))
	}

	return nt
}

func (t *table) all(out interface{}) error {
	outVal := reflect.ValueOf(out).Elem()
	for _, row := range t.rows {
		outVal.Set(reflect.Append(outVal, reflect.ValueOf(row).Elem()))
	}
	return nil
}

func (t *table) checkForField(f string) bool {
	if f == gorm.ToDBName(f) {
		f = t.columns[f]
	}
	_, found := t.ref.FieldByName(f)
	return found
}

func (t *table) typeCheck(i interface{}) error {
	typ := reflect.TypeOf(i)
	if typ != reflect.PtrTo(t.ref) && typ != t.ref {
		return ErrTypeMismatch
	}
	return nil
}

func (t *table) insert(row interface{}) error {
	err := t.typeCheck(row)
	if err != nil {
		return err
	}

	t.count++
	for _, k := range t.idx {
		s := reflect.ValueOf(row).Elem()
		idField := s.FieldByName(k)
		value := reflect.ValueOf(t.count)
		idField.Set(value)
	}

	t.rows = append(t.rows, row)
	return nil
}

func (t *table) find(field string, value interface{}) (reflect.Value, error) {
	if field == gorm.ToDBName(field) {
		field = t.columns[field]
	}

	_, found := t.ref.FieldByName(field)
	if !found {
		return RNil, errors.New("field name not in struct")
	}

	result := reflect.MakeSlice(reflect.SliceOf(reflect.PtrTo(t.ref)), 0, 0)
	for _, row := range t.rows {
		if reflect.TypeOf(value).Kind() == reflect.Uint {
			v := reflect.ValueOf(row).Elem().FieldByName(field).Uint()
			if uint(v) == value {
				result = reflect.Append(result, reflect.ValueOf(row))
			}
		} else if reflect.TypeOf(value).Kind() == reflect.Bool {
			v := reflect.ValueOf(row).Elem().FieldByName(field).Bool()
			if bool(v) == value {
				result = reflect.Append(result, reflect.ValueOf(row))
			}
		} else {
			v := reflect.ValueOf(row).Elem().FieldByName(field).String()
			if v == value {
				result = reflect.Append(result, reflect.ValueOf(row))
			}
		}
	}

	if result.Len() == 0 {
		return reflect.ValueOf(nil), nil
	}
	return result, nil
}

func (t *table) delete(ids map[string]reflect.Value) error {
	for i, row := range t.rows {
		match := true
		for k, v := range ids {
			if k == gorm.ToDBName(k) {
				k = t.columns[k]
			}

			if v.Interface() != reflect.ValueOf(row).Elem().FieldByName(k).Interface() {
				match = false
				break
			}
		}
		if match {
			t.rows = append(t.rows[:i], t.rows[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

func (t *table) appendToArray(query reflect.Value, target string, value interface{}) error {
	if query.Type() != t.ref {
		return errors.New("wrong query type")
	}

	key := t.PrimaryKey[0]
	res, err := t.find(key, query.FieldByName(key).Interface())
	if err != nil {
		return err
	}
	if res == RNil {
		return ErrNotFound
	}

	length := res.Len()

	for id, k := range t.PrimaryKey {
		if id == 0 {
			continue
		}

		for i := 0; i < length; i++ {
			r := res.Index(i).Elem()
			if r.FieldByName(k).Interface() == query.FieldByName(k).Interface() {
				continue
			}
			if length == 1 {
				return errors.New("row not found")
			}
			res = reflect.AppendSlice(res.Slice(0, i), res.Slice(i+1, length))
			length--
		}
	}

	row := res.Index(0).Elem()
	f := row.FieldByName(target)

	if f.Kind() != reflect.Slice {
		return errors.New("target is no slice")
	}

	f.Set(reflect.Append(f, reflect.ValueOf(value)))

	return nil
}

func (t *table) removeFromArray(query reflect.Value, target string, index int) error {
	if query.Type() != t.ref {
		return errors.New("wrong query type")
	}

	key := t.PrimaryKey[0]
	res, err := t.find(key, query.FieldByName(key).Interface())
	if err != nil {
		return err
	}
	if res == RNil {
		return ErrNotFound
	}

	length := res.Len()

	for id, k := range t.PrimaryKey {
		if id == 0 {
			continue
		}

		for i := 0; i < length; i++ {
			r := res.Index(i).Elem()
			if r.FieldByName(k).Interface() == query.FieldByName(k).Interface() {
				continue
			}
			if length == 1 {
				return errors.New("row not found")
			}
			res = reflect.AppendSlice(res.Slice(0, i), res.Slice(i+1, length))
			length--
		}
	}

	row := res.Index(0).Elem()
	f := row.FieldByName(target)

	if f.Kind() != reflect.Slice {
		return errors.New("target is no slice")
	}

	f.Set(reflect.AppendSlice(f.Slice(0, index), f.Slice(index+1, f.Len())))

	return nil
}

func getPrimaryKeys(ref reflect.Type, prime, idx *[]string, m map[string]string) {
	primaryRegExp := regexp.MustCompile("primary_key")
	refRegExp := regexp.MustCompile("Ref")

	for i := 0; i < ref.NumField(); i++ {
		f := ref.Field(i)
		t := f.Type
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		if t.Kind() == reflect.Struct {
			getPrimaryKeys(t, prime, idx, m)
			continue
		}

		dbName := gorm.ToDBName(f.Name)
		m[dbName] = f.Name

		v, ok := f.Tag.Lookup("gorm")
		if ok {
			isPrimary := primaryRegExp.MatchString(v)
			if isPrimary {
				*prime = append(*prime, f.Name)
				if f.Type.Kind() == reflect.Uint && !refRegExp.MatchString(f.Name) {
					*idx = append(*idx, f.Name)
				}
			}
		}
	}
}

func newTable(ref reflect.Type, name string) *table {
	var (
		idx     []string
		primary []string
		columns = make(map[string]string)
	)

	idTag := "ID"
	f, found := ref.FieldByName(idTag)
	if found {
		primary = append(primary, idTag)
		if f.Type.Kind() == reflect.Uint {
			idx = append(idx, idTag)
		}
	}

	getPrimaryKeys(ref, &primary, &idx, columns)

	return &table{
		Name:       name,
		PrimaryKey: primary,
		columns:    columns,
		ref:        ref,
		idx:        idx,
	}
}
