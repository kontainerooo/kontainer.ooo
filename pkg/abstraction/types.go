package abstraction

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)

// JSON is a json abstraction
type JSON map[string]interface{}

// ToStringMap returns a string map with the contents of the json
func (j JSON) ToStringMap() map[string]string {
	m := make(map[string]string)
	for k, v := range j {
		switch v.(type) {
		case string:
			m[k] = v.(string)
		case int:
			m[k] = strconv.FormatInt(v.(int64), 10)
		case float64:
			m[k] = strconv.FormatFloat(v.(float64), 'f', -1, 64)
		}
	}
	return m
}

// ToStringArrayMap returns a string array map with the contents of the json
func (j JSON) ToStringArrayMap() map[string][]string {
	m := make(map[string][]string)
	for k, v := range j {
		tmpArr := []string{}
		switch v.(type) {
		case []interface{}:
			for _, b := range v.([]interface{}) {
				tmpArr = append(tmpArr, b.(string))
			}
			m[k] = tmpArr
		}
	}
	return m
}

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
		i, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			j[k] = v
			continue
		}
		j[k] = int(i)
	}
	return j
}

// NewJSONFromMapArray creates a new JSON given a string->[]string map
func NewJSONFromMapArray(m map[string][]string) JSON {
	j := make(JSON)
	for k, v := range m {
		j[k] = v
	}
	return j
}

// Inet represents an IP Address with optional subnet mask
type Inet string

// NewInet creates a new Inet type
func NewInet(s string) (Inet, error) {
	r := regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\/(0|3[0-2]|[1-2][0-9]|[0-9]))?`)
	if r.MatchString(s) {
		return Inet(s), nil
	}

	return "", errors.New("Not a valid IP Address")
}
