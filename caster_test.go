package caster

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"strings"
	"testing"
)

func TestCast(t *testing.T) {
	expected := "<html><head><title>Caster test</title></head><body><header><div><h1>Header</h1></div></header><div><div><h1>Content</h1><h3>hello, world</h3></div></div><footer><div><h6>Footer</h6></div></footer></body></html>"
	tester := &tester{
		tset: &TemplateSet{
			Filenames: []string{
				"testdata/master.html",
				"testdata/header.html",
				"testdata/footer.html",
			},
		},
	}
	if err := tester.extend("var", &TemplateSet{
		Filenames: []string{"testdata/var.html"},
	}); err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	b := &bytes.Buffer{}
	if err := tester.cast(b, "var", map[string]interface{}{
		"message": "hello, world",
	}); err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	actual := trimWhitespaces(b)
	if actual != expected {
		t.Errorf("unexpected output: got %s, expected %s\n", actual, expected)
	}
}

func TestCastWithDefaultFuncs(t *testing.T) {
	expected := "<html><head><title>Caster test</title></head><body><header><div><h1>hello, tomocy</h1></div></header><div><div><h1>Content</h1><h3>hello, world</h3></div></div><footer><div><h6>Footer</h6></div></footer></body></html>"
	tester := &tester{
		tset: &TemplateSet{
			Filenames: []string{
				"testdata/master.html",
				"testdata/header_func.html",
				"testdata/footer.html",
			},
			FuncMap: template.FuncMap{
				"greet": func(to string) string {
					return fmt.Sprintf("hello, %s\n", to)
				},
			},
		},
	}

	if err := tester.extend("var", &TemplateSet{
		Filenames: []string{"testdata/var.html"},
	}); err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	b := &bytes.Buffer{}
	if err := tester.cast(b, "var", map[string]interface{}{
		"to":      "tomocy",
		"message": "hello, world",
	}); err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	actual := trimWhitespaces(b)
	if actual != expected {
		t.Errorf("unexpected output: got %s, expected %s\n", actual, expected)
	}
}
func TestCastWithFuncs(t *testing.T) {
	expected := "<html><head><title>Caster test</title></head><body><header><div><h1>Header</h1></div></header><div><div><h1>Content</h1><h3>hello, tomocy</h3></div></div><footer><div><h6>Footer</h6></div></footer></body></html>"
	tester := newTester()
	if err := tester.extend("func", &TemplateSet{
		Filenames: []string{"testdata/func.html"},
		FuncMap: template.FuncMap{
			"greet": func(to string) string {
				return fmt.Sprintf("hello, %s\n", to)
			},
		},
	}); err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	b := &bytes.Buffer{}
	if err := tester.cast(b, "func", map[string]interface{}{
		"to": "tomocy",
	}); err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	actual := trimWhitespaces(b)
	if actual != expected {
		t.Errorf("unexpected output: got %s, expected %s\n", actual, expected)
	}
}

func trimWhitespaces(r io.Reader) string {
	b := make([]byte, 0, 100)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()
		s = strings.Trim(s, "\n")
		s = strings.Trim(s, "\r")
		s = strings.Trim(s, "\t")
		s = strings.Trim(s, " ")
		b = append(b, s...)
	}

	return string(b)
}

type tester struct {
	caster Caster
	tset   *TemplateSet
}

func newTester() *tester {
	return &tester{
		tset: &TemplateSet{
			Filenames: []string{
				"testdata/master.html",
				"testdata/header.html",
				"testdata/footer.html",
			},
		},
	}
}

func (t *tester) extend(key string, tset *TemplateSet) error {
	caster, err := New(t.tset)
	if err != nil {
		return err
	}
	if err := caster.Extend(key, tset); err != nil {
		return err
	}

	t.caster = caster

	return nil
}

func (t *tester) cast(w io.Writer, key string, data interface{}) error {
	return t.caster.Cast(w, key, data)
}
