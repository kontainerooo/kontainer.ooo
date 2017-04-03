package template

import (
	"errors"
	"regexp"

	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	"github.com/lib/pq"
)

var (
	// ErrNoRefID is returned, if no ref id is set in a config
	ErrNoRefID = errors.New("no RefID set")

	// ErrNoName is returned, if no name is set in a config
	ErrNoName = errors.New("no Name set")

	// ErrPortRange is returned, if port is <1024
	ErrPortRange = errors.New("port not in acceptable range")

	// ErrKeyword is returned, if the used keyword is not allowed in the used router
	ErrKeyword = errors.New("keyword not allowed")
)

type writingService struct {
	s   routing.Service
	w   Writer
	r   Router
	mem map[uint]map[string]*routing.RouterConfig
}

func (w *writingService) checkListenStatement(r *routing.ListenStatement) error {
	// TODO: get IP pool for check: if !pool.In(inet) return err
	if r.Port <= 1024 {
		return ErrPortRange
	}

	switch w.r {
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

func (w *writingService) checkServerName(s pq.StringArray) error {
	return nil
}

func (w *writingService) checkPath(p string) error {
	return nil
}

func (w *writingService) checkLog(l *routing.Log) error {
	return nil
}

func (w *writingService) checkSSLSettings(s *routing.SSLSettings) error {
	return nil
}

func (w *writingService) checkLocationRules(l *routing.LocationRules) error {
	return nil
}

func (w *writingService) checkConfig(r *routing.RouterConfig) error {
	var err error

	if r.RefID == 0 {
		return ErrNoRefID
	}

	if r.Name == "" {
		return ErrNoName
	}

	err = w.checkListenStatement(r.ListenStatement)
	if err != nil {
		return err
	}

	err = w.checkServerName(r.ServerName)
	if err != nil {
		return err
	}

	err = w.checkLog(&r.AccessLog)
	if err != nil {
		return err
	}

	err = w.checkLog(&r.ErrorLog)
	if err != nil {
		return err
	}

	err = w.checkPath(r.RootPath)
	if err != nil {
		return err
	}

	err = w.checkSSLSettings(&r.SSLSettings)
	if err != nil {
		return err
	}

	err = w.checkLocationRules(&r.LocationRules)
	if err != nil {
		return err
	}

	return nil
}

func (w *writingService) CreateRouterConfig(r *routing.RouterConfig) error {
	err := w.s.CreateRouterConfig(r)
	if err != nil {
		return err
	}
	return nil
}

func (w *writingService) EditRouterConfig(refID uint, name string, r *routing.RouterConfig) error {
	err := w.s.EditRouterConfig(refID, name, r)
	if err != nil {
		return err
	}
	return nil
}

func (w *writingService) GetRouterConfig(refID uint, name string, r *routing.RouterConfig) error {
	err := w.s.GetRouterConfig(refID, name, r)
	if err != nil {
		return err
	}
	return nil
}

func (w *writingService) RemoveRouterConfig(refID uint, name string) error {
	err := w.s.RemoveRouterConfig(refID, name)
	if err != nil {
		return err
	}
	return nil
}

func (w *writingService) AddLocationRule(refID uint, name string, lr *routing.LocationRule) error {
	err := w.s.AddLocationRule(refID, name, lr)
	if err != nil {
		return err
	}
	return nil
}

func (w *writingService) RemoveLocationRule(refID uint, name string, lid int) error {
	err := w.s.RemoveLocationRule(refID, name, lid)
	if err != nil {
		return err
	}
	return nil
}

func (w *writingService) ChangeListenStatement(refID uint, name string, ls *routing.ListenStatement) error {
	err := w.s.ChangeListenStatement(refID, name, ls)
	if err != nil {
		return err
	}
	return nil
}

func (w *writingService) AddServerName(refID uint, name string, sn string) error {
	err := w.s.AddServerName(refID, name, sn)
	if err != nil {
		return err
	}
	return nil
}

func (w *writingService) RemoveServerName(refID uint, name string, id int) error {
	err := w.s.RemoveServerName(refID, name, id)
	if err != nil {
		return err
	}
	return nil
}

func (w *writingService) Configurations(r *[]routing.RouterConfig) {
	w.s.Configurations(r)
}

// NewWritingService creates a writingService with necessary dependencies.
func NewWritingService(s routing.Service, r Router, p string) (routing.Service, error) {
	w, err := NewWriter(r, p)
	if err != nil {
		return nil, err
	}

	return &writingService{
		s:   s,
		w:   w,
		r:   r,
		mem: make(map[uint]map[string]*routing.RouterConfig),
	}, nil
}
