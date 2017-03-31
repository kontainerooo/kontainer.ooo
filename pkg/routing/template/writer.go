package template

import (
	"errors"
	"os"
	"text/template"

	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
)

// Router is an enum of possible Routers
type Router uint

const (
	// Nginx router
	Nginx Router = iota
)

var router = [...]string{
	"nginx.conf",
}

// Writer is
type Writer interface {
	CreateFile(*routing.RouterConfig) error
}

type writer struct {
	template *template.Template
	path     string
}

func (w writer) CreateFile(c *routing.RouterConfig) error {
	return nil
}

// NewWriter returns a new writer
func NewWriter(r Router, p string) (Writer, error) {
	t, err := template.ParseFiles(router[r])
	if err != nil {
		return nil, err
	}

	dir, err := os.Stat(p)
	if err != nil {
		return nil, err
	} else if !dir.IsDir() {
		return nil, errors.New("path is no dir")
	}

	return &writer{
		template: t,
		path:     p,
	}, nil
}
