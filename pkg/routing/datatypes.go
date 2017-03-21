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
	ipAddress inet `sql:"type:inet"`
	port      int
	keyword   string
}

// Log combines a path and a keyword
type Log struct {
	path    string
	keyword string
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
	location string
	rules    map[string][]string
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
	RefID           uint   `gorm:"primary_key"`
	Name            string `gorm:"primary_key"`
	ListenStatement ListenStatement
	ServerName      pq.StringArray
	AccessLog       Log
	ErrorLog        Log
	EootPath        string
	SSLSettings     SSLSettings
	LocationRules   LocationRules
}

// TableName sets RouterConfig's database table name
func (RouterConfig) TableName() string {
	return "Routing"
}
