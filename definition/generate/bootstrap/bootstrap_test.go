package bootstrap_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-4devs/config/definition/generate/bootstrap"
)

func TestBootstrapUsesStructName(t *testing.T) {
	t.Parallel()

	src := `package cfg

//fconfig:structname: ConfigExample
func Example() error { return nil }

func Plain() error { return nil }
`
	out := bootstrapSource(t, src, []string{"Example", "Plain"})

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

func TestBootstrapRejectsEmptyStructName(t *testing.T) {
	t.Parallel()

	src := `package cfg

//fconfig:structname:
func Example() error { return nil }
`
	path := writeCfg(t, src)

	_, err := bootstrap.Bootstrap(context.Background(), stubConfig{
		file:   path,
		prefix: "App",
		suffix: "Settings",
	})
	if err == nil {
		t.Fatal("expected error")
	}

	left, gerr := filepath.Glob(filepath.Join(filepath.Dir(path), "config-bootstrap-*"))
	if gerr != nil {
		t.Fatal(gerr)
	}

	if len(left) != 0 {
		t.Fatalf("temp files left: %v", left)
	}
}

func bootstrapSource(t *testing.T, src string, methods []string) string {
	t.Helper()

	path := writeCfg(t, src)

	out, err := bootstrap.Bootstrap(context.Background(), stubConfig{
		file:    path,
		methods: methods,
		prefix:  "App",
		suffix:  "Settings",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.Remove(out)
	})

	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}

	return string(data)
}

func writeCfg(t *testing.T, src string) string {
	t.Helper()

	dir := t.Name() + "-src"
	if err := os.Mkdir(dir, 0o700); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.RemoveAll(dir)
	})

	path := filepath.Join(dir, "cfg.go")
	if err := os.WriteFile(path, []byte(src), 0o600); err != nil {
		t.Fatal(err)
	}

	return path
}

type stubConfig struct {
	file    string
	methods []string
	prefix  string
	suffix  string
}

func (s stubConfig) File() string      { return s.file }
func (s stubConfig) Methods() []string { return s.methods }
func (stubConfig) SkipContext() bool   { return false }
func (s stubConfig) Prefix() string    { return s.prefix }
func (s stubConfig) Suffix() string    { return s.suffix }
func (stubConfig) FullPkg() string     { return "example.com/app" }
