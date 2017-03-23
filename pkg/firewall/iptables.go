// Package firewall handles the firewall and forwarding configuration
package firewall

// IPTablesRule represents a rule entry in iptables
type IPTablesRule struct {
	Target          string
	Protocol        string
	In              string
	Out             string
	Source          string
	Destination     string
	SourcePort      uint16
	DestinationPort uint16
	State           string
}

type iptablesEntry struct {
	ID    uint
	refid uint
	IPTablesRule
}

// IPTablesWrapper handles iptables rules
type IPTablesWrapper interface {
	// AddRule adds a given iptables rule
	AddRule(refid uint, rule IPTablesRule) error
	// RemoveRule removes a given iptables rule
	RemoveRule() error
	// GetRulesForUser returns a list of all rules for a given user
	GetRulesForUser(refid uint) []IPTablesRule
	// CreateIPTablesBackup creates an iptables backup file
	CreateIPTablesBackup() string
	// LoadIPTablesBackup restores iptables from backup file
	LoadIPTablesBackup() error
}

type iptablesWrapper struct {
	iptPath string
}

func (w *iptablesWrapper) AddRule(refid uint, rule IPTablesRule) error {
	// TODO: implement
	return nil
}

func (w *iptablesWrapper) RemoveRule() error {
	// TODO: implement
	return nil
}

func (w *iptablesWrapper) GetRulesForUser(refid uint) []IPTablesRule {
	// TODO: implement
	return []IPTablesRule{}
}

func (w *iptablesWrapper) CreateIPTablesBackup() string {
	// TODO: implement
	return ""
}

func (w *iptablesWrapper) LoadIPTablesBackup() error {
	// TODO: implement
	return nil
}

// NewIPTablesWrapper creates a new iptables wrapper
func NewIPTablesWrapper(iptPath string) IPTablesWrapper {
	return &iptablesWrapper{
		iptPath: iptPath,
	}
}
