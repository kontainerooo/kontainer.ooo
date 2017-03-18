package abstraction

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
)

// JSON is a json abstraction
type JSON map[string]interface{}

// Value get value of JSON
func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan scan value into JSON
func (j *JSON) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*j, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
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
