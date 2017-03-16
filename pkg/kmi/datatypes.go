package kmi

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
	"github.com/ttdennis/kontainer.io/pkg/abstraction"
)

// The KMDI struct contains the displaying information for the frontend
type KMDI struct {
	ID          uint
	Name        string
	Version     string
	Description string
	Type        int
}

type frontendModule struct {
	template   string
	parameters abstraction.JSON `sql:"type:jsonb"`
}

// FrontendArray represents a frontendModule Array
type FrontendArray []frontendModule

// Scan implements the sql.Scanner interface.
func (f *FrontendArray) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return f.scanBytes(src)
	case string:
		return f.scanBytes([]byte(src))
	case nil:
		*f = nil
		return nil
	}

	return fmt.Errorf("pq: cannot convert %T to Frontend Array", src)
}

func (f *FrontendArray) scanBytes(src []byte) error {
	return GetFrontend(src, f)
}

// Value implements the driver.Valuer interface.
func (f FrontendArray) Value() (driver.Value, error) {
	if f == nil {
		return nil, nil
	}
	b, err := json.Marshal(f)

	return string(b), err
}

// The KMI struct is used to represent every information included in a kmi-file
type KMI struct {
	KMDI
	Dockerfile  string
	Container   string
	Commands    abstraction.JSON `sql:"type:jsonb"`
	Environment abstraction.JSON `sql:"type:jsonb"`
	Frontend    FrontendArray    `sql:"type:jsonb"`
	Imports     pq.StringArray   `sql:"type:text[]"`
	Interfaces  abstraction.JSON `sql:"type:jsonb"`
	Mounts      pq.StringArray   `sql:"type:text[]"`
	Variables   pq.StringArray   `sql:"type:text[]"`
	Resources   abstraction.JSON `sql:"type:jsonb"`
}

// TableName sets KMI's tablename
func (KMI) TableName() string {
	return "kontainer_module_information"
}
