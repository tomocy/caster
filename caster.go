package caster

import (
	"fmt"
	"html/template"
	"io"
)

type Caster struct {
	tmpls      map[string]*template.Template
	layoutTmpl *template.Template
}

func New(layouts ...string) (*Caster, error) {
	layoutTmpl, err := template.ParseFiles(layouts...)
	if err != nil {
		return nil, fmt.Errorf("faild to create new caster: %s\n", err)
	}

	return &Caster{
		tmpls:      make(map[string]*template.Template),
		layoutTmpl: layoutTmpl,
	}, nil
}

func (c *Caster) Extend(key string, fnames ...string) error {
	tmpl, err := c.layoutTmpl.Clone()
	if err != nil {
		return fmt.Errorf("faild to extend template: %s\n", err)
	}
	tmpl, err = tmpl.ParseFiles(fnames...)
	if err != nil {
		return fmt.Errorf("faild to extend template: %s\n", err)
	}

	c.tmpls[key] = tmpl
	return nil
}

func (c *Caster) Cast(w io.Writer, key string, data interface{}) error {
	tmpl, ok := c.tmpls[key]
	if !ok {
		return fmt.Errorf("faild to cast template: no template whose key is %s found", key)
	}

	if err := tmpl.ExecuteTemplate(w, "master", data); err != nil {
		return fmt.Errorf("faild to cast template: %s", err)
	}
	return nil
}
