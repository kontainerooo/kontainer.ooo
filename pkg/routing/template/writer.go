package template

import (
	"errors"
	"fmt"
	"os"
	"strings"
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
	name     string
}

func (w writer) CreateFile(c *routing.RouterConfig) error {
	path := fmt.Sprintf("%s/%d_%s.conf", w.path, c.RefID, c.Name)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = w.template.ExecuteTemplate(f, w.name, c)
	if err != nil {
		return err
	}

	return nil
}

// NewWriter returns a new writer
func NewWriter(r Router, p string) (Writer, error) {
	t := template.New("routing").Funcs(template.FuncMap{
		"join": strings.Join,
	})

	if !(int(r) < len(router)) {
		return nil, fmt.Errorf("Router with id %d does not exist", r)
	}
	name := router[r]
	t, err := t.ParseFiles(name)
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
		name:     name,
	}, nil
}
