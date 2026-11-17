package bootstrap

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/go-4devs/config/definition/generate/pkg"
)

func TestStructNames(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "cfg.go")
	src := `package cfg

// Example configures the application.
//
//fconfig:structname: ConfigExample
func Example() error { return nil }

func Plain() error { return nil }
`
	if err := os.WriteFile(path, []byte(src), 0o600); err != nil {
		t.Fatal(err)
	}

	names, err := structNames(path)
	if err != nil {
		t.Fatal(err)
	}

	if names["Example"] != "ConfigExample" {
		t.Fatalf("Example: %q", names["Example"])
	}

	if _, ok := names["Plain"]; ok {
		t.Fatalf("Plain should keep the method name, got %q", names["Plain"])
	}
}

func TestStructNamesRejectsEmpty(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "cfg.go")
	src := `package cfg

//fconfig:structname:
func Example() error { return nil }
`
	if err := os.WriteFile(path, []byte(src), 0o600); err != nil {
		t.Fatal(err)
	}

	if _, err := structNames(path); err == nil {
		t.Fatal("expected error")
	}
}

func TestBootstrapUsesStructName(t *testing.T) {
	t.Parallel()

	tpl, err := template.ParseFS(tpls, "bootstrap.go.tpl")
	if err != nil {
		t.Fatal(err)
	}

	data := Boot{
		Packages: pkg.NewImports("main"),
		Configure: []configure{
			{Name: "Example", StructName: "ConfigExample"},
			{Name: "Plain"},
		},
		Config: stubConfig{prefix: "App", suffix: "Settings"},
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		t.Fatal(err)
	}

	out := buf.String()
	for _, part := range []string{
		`view.WithStructName("App_ConfigExample_Settings")`,
		`view.WithStructName("App_Plain_Settings")`,
		`app.Example(ctx, defExample)`,
		`app.Plain(ctx, defPlain)`,
	} {
		if !strings.Contains(out, part) {
			t.Fatalf("missing %q in:\n%s", part, out)
		}
	}
}

type stubConfig struct {
	prefix string
	suffix string
}

func (stubConfig) File() string      { return "" }
func (stubConfig) Methods() []string { return nil }
func (stubConfig) SkipContext() bool { return false }
func (s stubConfig) Prefix() string  { return s.prefix }
func (s stubConfig) Suffix() string  { return s.suffix }
func (stubConfig) FullPkg() string   { return "example.com/app" }
