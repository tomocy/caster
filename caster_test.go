package caster

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestCast(t *testing.T) {
	expected := "<html><head><title>Caster test</title></head><body><header><div><h1>Header</h1></div></header><div><div><h1>Content</h1><h3>hello, world</h3></div></div><footer><div><h6>Footer</h6></div></footer></body></html>"
	layouts := []string{
		"testdata/master.html",
		"testdata/header.html",
		"testdata/footer.html",
	}
	caster, err := New(layouts...)
	if err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}
	if err := caster.Extend("index", "testdata/index.html"); err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}
	b := &bytes.Buffer{}
	if err := caster.Cast(b, "index", map[string]interface{}{
		"message": "hello, world",
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
