package abstraction

import (
	"database/sql"
	"database/sql/driver"
	"strconv"

	"github.com/lib/pq/hstore"
)

// JSON is a json abstraction
type JSON map[string]interface{}

// Value get value of JSON
func (j JSON) Value() (driver.Value, error) {
	hstore := hstore.Hstore{Map: map[string]sql.NullString{}}
	if len(j) == 0 {
		return nil, nil
	}

	for key, value := range j {
		var s sql.NullString
		if value != nil {
			s.String = value.(string)
			s.Valid = true
		}
		hstore.Map[key] = s
	}

	return hstore.Value()
}

// Scan scan value into JSON
func (j *JSON) Scan(value interface{}) error {
	hstore := hstore.Hstore{}

	if err := hstore.Scan(value); err != nil {
		return err
	}

	if len(hstore.Map) == 0 {
		return nil
	}

	*j = JSON{}
	for k := range hstore.Map {
		if hstore.Map[k].Valid {
			s := hstore.Map[k].String
			i, err := strconv.ParseInt(s, 10, 0)
			if err == nil {
				(*j)[k] = int(i)
			} else {
				(*j)[k] = &s
			}
		} else {
			(*j)[k] = nil
		}
	}

	return nil
}

// NewJSONFromMap creates a new JSON given a string->string map
func NewJSONFromMap(m map[string]string) JSON {
	j := make(JSON)
	for k, v := range m {
		i, err := strconv.ParseInt(k, 10, 0)
		if err != nil {
			j[k] = v
		}
		j[k] = int(i)
	}
	return j
}
