package kmi

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
	parameters map[string]interface{}
}

// The KMI struct is used to represent every information included in a kmi-file
type KMI struct {
	KMDI
	Commands    map[string]interface{}
	Environment map[string]interface{}
	Frontend    []frontendModule
	Imports     []string
	Interfaces  map[string]interface{}
}

// TableName sets KMI's tablename
func (KMI) TableName() string {
	return "kontainer_module_information"
}
