package kmi

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/lib/pq"
)

// The KMDI struct contains the displaying information for the frontend
type KMDI struct {
	ID          uint
	Name        string
	Version     string
	Description string
	Type        int
}

// FrontendModule contains a template and its parameters
type FrontendModule struct {
	Template   string
	Parameters abstraction.JSON `sql:"type:jsonb"`
}

// FrontendArray represents a frontendModule Array
type FrontendArray []*FrontendModule

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
	return json.Unmarshal(src, f)
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
	ProvisionScript string
	Commands        abstraction.JSON `sql:"type:jsonb"`
	Environment     abstraction.JSON `sql:"type:jsonb"`
	Frontend        FrontendArray    `sql:"type:jsonb"`
	Imports         pq.StringArray   `sql:"type:text[]"`
	Interfaces      abstraction.JSON `sql:"type:jsonb"`
	Resources       abstraction.JSON `sql:"type:jsonb"`
}

// TableName sets KMI's tablename
func (KMI) TableName() string {
	return "kontainer_module_information"
}
