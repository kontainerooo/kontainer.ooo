package routing

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
)

type inet string

// ListenStatement combines an ipAddress with a port and a keyword
type ListenStatement struct {
	IPAddress inet `sql:"type:inet"`
	Port      int
	Keyword   string
}

// Scan implements the sql.Scanner interface.
func (l *ListenStatement) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return l.scanBytes(src)
	case string:
		return l.scanBytes([]byte(src))
	case nil:
		*l = ListenStatement{}
		return nil
	}

	return fmt.Errorf("pq: cannot convert %T to Frontend Array", src)
}

func (l *ListenStatement) scanBytes(src []byte) error {
	return json.Unmarshal(src, l)
}

// Value implements the driver.Valuer interface.
func (l ListenStatement) Value() (driver.Value, error) {
	if l == (ListenStatement{}) {
		return nil, nil
	}
	b, err := json.Marshal(l)

	return string(b), err
}

// Log combines a path and a keyword
type Log struct {
	path    string
	keyword string
}

// Scan implements the sql.Scanner interface.
func (l *Log) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return l.scanBytes(src)
	case string:
		return l.scanBytes([]byte(src))
	case nil:
		*l = Log{}
		return nil
	}

	return fmt.Errorf("pq: cannot convert %T to Frontend Array", src)
}

func (l *Log) scanBytes(src []byte) error {
	return json.Unmarshal(src, l)
}

// Value implements the driver.Valuer interface.
func (l Log) Value() (driver.Value, error) {
	if l == (Log{}) {
		return nil, nil
	}
	b, err := json.Marshal(l)

	return string(b), err
}

// SSLSettings represents a set of configurations in terms of ssl transport
type SSLSettings struct {
	protocols           []string
	ciphers             []string
	preferServerCiphers string
	certificate         string
	certificateKey      string
}

// LocationRule is a struct which combines a single location (URL path) with a set of rules
type LocationRule struct {
	Location string
	Rules    map[string][]string
}

// Scan implements the sql.Scanner interface.
func (l *LocationRule) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return l.scanBytes(src)
	case string:
		return l.scanBytes([]byte(src))
	case nil:
		*l = LocationRule{
			Rules: make(map[string][]string),
		}
		return nil
	}

	return fmt.Errorf("pq: cannot convert %T to Frontend Array", src)
}

func (l *LocationRule) scanBytes(src []byte) error {
	return json.Unmarshal(src, l)
}

// Value implements the driver.Valuer interface.
func (l LocationRule) Value() (driver.Value, error) {
	if l.Location == "" {
		return nil, nil
	}
	b, err := json.Marshal(l)

	return string(b), err
}

// LocationRules is an array of LocationRules
type LocationRules []*LocationRule

// Scan implements the sql.Scanner interface.
func (l *LocationRules) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return l.scanBytes(src)
	case string:
		return l.scanBytes([]byte(src))
	case nil:
		*l = nil
		return nil
	}

	return fmt.Errorf("pq: cannot convert %T to Frontend Array", src)
}

func (l *LocationRules) scanBytes(src []byte) error {
	return json.Unmarshal(src, l)
}

// Value implements the driver.Valuer interface.
func (l LocationRules) Value() (driver.Value, error) {
	if l == nil {
		return nil, nil
	}
	b, err := json.Marshal(l)

	return string(b), err
}

// The RouterConfig struct represents the collected information needed to configurate an http router
type RouterConfig struct {
	RefID           uint            `gorm:"primary_key"`
	Name            string          `gorm:"primary_key"`
	ListenStatement ListenStatement `sql:"type:jsonb"`
	ServerName      pq.StringArray  `sql:"type:text[]"`
	AccessLog       Log             `sql:"type:jsonb"`
	ErrorLog        Log             `sql:"type:jsonb"`
	RootPath        string
	SSLSettings     SSLSettings   `sql:"type:jsonb"`
	LocationRules   LocationRules `sql:"type:jsonb[]"`
}

// TableName sets RouterConfig's database table name
func (RouterConfig) TableName() string {
	return "routing"
}
