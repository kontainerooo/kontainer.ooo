// Package firewall handles the firewall and forwarding configuration
package firewall

// Service IPTablesService
type Service interface {

	// InitBridge initializes a bridge network
	InitBridge(ip string) error

	// AllowConnection sets up a rule to let src talk to dst
	AllowConnection(src string, dst string) error

	// AllowConnection sets up a rule to block src from talking to dst
	BlockConnection(src string, dst string) error

	// AllowPort sets up a rule to let src talk to dst on port port
	AllowPort(src string, dst string, port uint32) error

	// AllowPort sets up a rule to block src from talking to dst on port port
	BlockPort(src string, dst string, port uint32) error

	// RedirectPort redirects the src port to the dst port on ip
	RedirectPort(ip string, src uint32, dst uint32) error

	// RemoveRedirectPort removes a port redirection
	RemoveRedirectPort(ip string, src uint32, dst uint32) error
}

type service struct{}

func (s *service) InitBridge(ip string) error {
	// TODO: implement
	return nil
}

func (s *service) AllowConnection(src string, dst string) error {
	// TODO: implement
	return nil
}

func (s *service) BlockConnection(src string, dst string) error {
	// TODO: implement
	return nil
}

func (s *service) AllowPort(src string, dst string, port uint32) error {
	// TODO: implement
	return nil
}

func (s *service) BlockPort(src string, dst string, port uint32) error {
	// TODO: implement
	return nil
}

func (s *service) RedirectPort(ip string, src uint32, dst uint32) error {
	// TODO: implement
	return nil
}

func (s *service) RemoveRedirectPort(ip string, src uint32, dst uint32) error {
	// TODO: implement
	return nil
}

// NewService creates a new firewall service
func NewService() Service {
	return &service{}
}
