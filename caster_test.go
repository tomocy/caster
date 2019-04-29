package caster

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"strings"
	"testing"
)

func TestCast(t *testing.T) {
	caster, err := New(&TemplateSet{
		Filenames: []string{
			templateFile("master.html"),
			templateFile("header.html"),
			templateFile("footer.html"),
		},
		FuncMap: template.FuncMap{
			"dear": func(to string) string {
				return fmt.Sprintf("Dear %s,", to)
			},
			"sincerely": func(from string) string {
				return fmt.Sprintf("Sincerely, %s", from)
			},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %s\n", err)
	}

	if err := caster.Extend("test", &TemplateSet{
		Filenames: []string{templateFile("content.html")},
	}); err != nil {
		t.Fatalf("unexpected error: %s\n", err)
	}

	dest := new(bytes.Buffer)
	if err := caster.Cast(dest, "test", map[string]interface{}{
		"to":      "developers",
		"message": "Did you sleep well last night?",
		"from":    "tomocy",
	}); err != nil {
		t.Fatalf("unexpected error: %s\n", err)
	}

	actual := trimWhitespaces(dest)
	if actual != expected() {
		t.Errorf("got %s, expected %s\n", actual, expected())
	}
}

func expected() string {
	return trimWhitespaces(strings.NewReader(`
	<html>

	<head>
		<title>Caster test</title>
	</head>

	<body>
		<header>
			<div>
				<h1>Header</h1>
				<h3>Dear developers,</h3>
		</div>
		</header>
		<main>
			<div>
				<h1>Content</h1>
				<h3>Did you sleep well last night?</h3>
			</div>
		</main>
		<footer>
			<div>
				<h1>Footer</h1>
				<h3>Sincerely, tomocy</h3>
			</div>
		</footer>
	</body>

	</html>
	`))
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

func templateFile(fname string) string {
	return filepath.Join("testdata", fname)
}
