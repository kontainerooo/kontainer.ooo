package kmi

// The KMDI struct contains the displaying information for the frontend
type KMDI struct {
	Name        string
	Version     string
	Description string
	Type        int
}

type frontendModule struct {
	Template   string
	parameters map[string]string
}

// The KMI struct is used to represent every information included in a kmi-file
type KMI struct {
	KMDI
	Commands    map[string]string
	Environment map[string]interface{}
	Frontend    []frontendModule
	Imports     []string
	Interfaces  map[string]uint16
}

// TableName sets KMI's tablename
func (KMI) TableName() string {
	return "kontainer_module_information"
}
