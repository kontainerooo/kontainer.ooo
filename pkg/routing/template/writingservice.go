package template

import "github.com/kontainerooo/kontainer.ooo/pkg/routing"

type writingService struct {
	s   routing.Service
	w   Writer
	r   Router
	mem Cache
	c   Check
}

func (w *writingService) CreateRouterConfig(r *routing.RouterConfig) error {
	var err error
	err = w.c.Config(r)
	if err != nil {
		return err
	}

	err = w.s.CreateRouterConfig(r)
	if err != nil {
		return err
	}

	err = w.w.CreateFile(w.mem.SetConf(r))
	if err != nil {
		return err
	}

	return nil
}

func (w *writingService) EditRouterConfig(refID uint, name string, r *routing.RouterConfig) error {
	var err error
	err = w.c.Config(r)
	if err != nil {
		return err
	}

	err = w.s.EditRouterConfig(refID, name, r)
	if err != nil {
		return err
	}

	err = w.w.CreateFile(w.mem.EditConf(r))
	if err != nil {
		return err
	}

	return nil
}

func (w *writingService) GetRouterConfig(refID uint, name string, r *routing.RouterConfig) error {
	var err error
	err = w.s.GetRouterConfig(refID, name, r)
	if err != nil {
		return err
	}

	return nil
}

func (w *writingService) RemoveRouterConfig(refID uint, name string) error {
	var err error
	err = w.w.RemoveFile(refID, name)
	if err != nil {
		return err
	}

	err = w.s.RemoveRouterConfig(refID, name)
	if err != nil {
		return err
	}

	w.mem.RemoveConf(refID, name)

	return nil
}

func (w *writingService) AddLocationRule(refID uint, name string, lr *routing.LocationRule) error {
	var err error
	err = w.c.LocationRule(lr)
	if err != nil {
		return err
	}

	err = w.s.AddLocationRule(refID, name, lr)
	if err != nil {
		return err
	}

	w.w.CreateFile(w.mem.UpdateConf(refID, name))

	return nil
}

func (w *writingService) RemoveLocationRule(refID uint, name string, lid int) error {
	err := w.s.RemoveLocationRule(refID, name, lid)
	if err != nil {
		return err
	}

	w.w.CreateFile(w.mem.UpdateConf(refID, name))

	return nil
}

func (w *writingService) ChangeListenStatement(refID uint, name string, ls *routing.ListenStatement) error {
	var err error
	err = w.c.ListenStatement(ls)
	if err != nil {
		return err
	}

	err = w.s.ChangeListenStatement(refID, name, ls)
	if err != nil {
		return err
	}

	w.w.CreateFile(w.mem.UpdateConf(refID, name))

	return nil
}

func (w *writingService) AddServerName(refID uint, name string, sn string) error {
	var err error
	err = w.c.ServerName([]string{sn})
	if err != nil {
		return err
	}

	err = w.s.AddServerName(refID, name, sn)
	if err != nil {
		return err
	}

	w.w.CreateFile(w.mem.UpdateConf(refID, name))

	return nil
}

func (w *writingService) RemoveServerName(refID uint, name string, id int) error {
	err := w.s.RemoveServerName(refID, name, id)
	if err != nil {
		return err
	}

	w.w.CreateFile(w.mem.UpdateConf(refID, name))

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
		mem: NewCache(s.GetRouterConfig),
		c:   NewCheck(r),
	}, nil
}
