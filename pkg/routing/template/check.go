package template

import (
	"regexp"

	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	"github.com/lib/pq"
)

// Check defines functions needed to check every part of a routing.RouterConfig
type Check interface {
	ListenStatement(r *routing.ListenStatement) error
	ServerName(s pq.StringArray) error
	Path(p string) error
	Log(l *routing.Log) error
	SSLSettings(s *routing.SSLSettings) error
	LocationRule(l *routing.LocationRule) error
	LocationRules(l *routing.LocationRules) error
	Config(r *routing.RouterConfig) error
}

type check struct {
	r Router
}

func (c *check) ListenStatement(r *routing.ListenStatement) error {
	// TODO: get IP pool for check: if !pool.In(inet) return err
	if r.Port <= 1024 {
		return ErrPortRange
	}

	switch c.r {
	case Nginx:
		regex := regexp.MustCompile(`^ssl$`)
		if !regex.MatchString(r.Keyword) {
			return ErrKeyword
		}
	default:
		if r.Keyword != "" {
			return ErrKeyword
		}
	}

	return nil
}

func (c *check) ServerName(s pq.StringArray) error {
	for _, n := range s {
		if !urlRegex.MatchString(n) {
			return ErrInvalidName
		}
	}
	return nil
}

func (c *check) Path(p string) error {
	return nil
}

func (c *check) Log(l *routing.Log) error {
	return nil
}

func (c *check) SSLSettings(s *routing.SSLSettings) error {
	return nil
}

func (c *check) LocationRule(l *routing.LocationRule) error {
	return nil
}

func (c *check) LocationRules(l *routing.LocationRules) error {
	return nil
}

func (c *check) Config(r *routing.RouterConfig) error {
	var err error

	if r.RefID == 0 {
		return ErrNoRefID
	}

	if r.Name == "" {
		return ErrNoName
	}

	err = c.ListenStatement(r.ListenStatement)
	if err != nil {
		return err
	}

	err = c.ServerName(r.ServerName)
	if err != nil {
		return err
	}

	err = c.Log(&r.AccessLog)
	if err != nil {
		return err
	}

	err = c.Log(&r.ErrorLog)
	if err != nil {
		return err
	}

	err = c.Path(r.RootPath)
	if err != nil {
		return err
	}

	err = c.SSLSettings(&r.SSLSettings)
	if err != nil {
		return err
	}

	err = c.LocationRules(&r.LocationRules)
	if err != nil {
		return err
	}

	return nil
}

// NewCheck returns a new Check
func NewCheck(r Router) Check {
	return &check{r}
}
