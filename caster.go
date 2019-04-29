package caster

import (
	"fmt"
	"html/template"
	"io"
)

type Caster interface {
	ExtendAll(tsetMap map[string]*TemplateSet) error
	Extend(key string, tset *TemplateSet) error
	Cast(w io.Writer, key string, data interface{}) error
}

type TemplateSet struct {
	Filenames []string
	FuncMap   template.FuncMap
}

type caster struct {
	tmpls      map[string]*template.Template
	layoutTmpl *template.Template
}

func New(tset *TemplateSet) (Caster, error) {
	t := template.New("")
	layoutTmpl, err := t.Funcs(tset.FuncMap).ParseFiles(tset.Filenames...)
	if err != nil {
		return nil, fmt.Errorf("faild to create new caster: %s\n", err)
	}

	return &caster{
		tmpls:      make(map[string]*template.Template),
		layoutTmpl: layoutTmpl,
	}, nil
}

func (c *caster) ExtendAll(tsetMap map[string]*TemplateSet) error {
	for key, tset := range tsetMap {
		if err := c.Extend(key, tset); err != nil {
			return err
		}
	}

	return nil
}

func (c *caster) Extend(key string, tset *TemplateSet) error {
	tmpl, err := c.layoutTmpl.Clone()
	if err != nil {
		return fmt.Errorf("faild to extend template: %s\n", err)
	}
	tmpl, err = tmpl.Funcs(tset.FuncMap).ParseFiles(tset.Filenames...)
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
