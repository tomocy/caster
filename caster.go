package caster

import (
	"fmt"
	"html/template"
	"io"
)

type Caster interface {
	ExtendAll(map[string][]string) error
	Extend(key string, fnames ...string) error
	Cast(w io.Writer, key string, data interface{}) error
}

type caster struct {
	tmpls      map[string]*template.Template
	layoutTmpl *template.Template
}

func New(layouts ...string) (Caster, error) {
	layoutTmpl, err := template.ParseFiles(layouts...)
	if err != nil {
		return nil, fmt.Errorf("faild to create new caster: %s\n", err)
	}

	return &caster{
		tmpls:      make(map[string]*template.Template),
		layoutTmpl: layoutTmpl,
	}, nil
}

func (c *caster) ExtendAll(m map[string][]string) error {
	for key, fnames := range m {
		if err := c.Extend(key, fnames...); err != nil {
			return err
		}
	}

	return nil
}

func (c *caster) Extend(key string, fnames ...string) error {
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

func (c *caster) Cast(w io.Writer, key string, data interface{}) error {
	tmpl, ok := c.tmpls[key]
	if !ok {
		return fmt.Errorf("faild to cast template: no template whose key is %s found", key)
	}

	if err := tmpl.ExecuteTemplate(w, "master", data); err != nil {
		return fmt.Errorf("faild to cast template: %s", err)
	}
	return nil
}
