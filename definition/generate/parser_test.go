package generate_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/go-4devs/config/definition/generate"
)

func TestParseMethodsChecksEveryArg(t *testing.T) {
	t.Parallel()

	src := `package cfg

func Good(a string, b int) error { return nil }

func Bad(a string, b int64) error { return nil }
`
	methods := parseMethods(t, src)

	if len(methods) != 1 || methods[0] != "Good" {
		t.Fatalf("methods: %v", methods)
	}
}

func TestParseMethodsIgnoresDirective(t *testing.T) {
	t.Parallel()

	src := `package cfg

func Good(a string, b int) error { return nil }

// Good stays. This one is not generated.
//
//fconfig:ignore
func Skip(a string, b int) error { return nil }
`
	methods := parseMethods(t, src)

	if len(methods) != 1 || methods[0] != "Good" {
		t.Fatalf("methods: %v", methods)
	}
}

func parseMethods(t *testing.T, src string) []string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "cfg.go")
	if err := os.WriteFile(path, []byte(src), 0o600); err != nil {
		t.Fatal(err)
	}

	methods, err := generate.NewParseMethods(
		path,
		[]reflect.Type{reflect.TypeFor[string](), reflect.TypeFor[int]()},
		[]reflect.Type{reflect.TypeFor[error]()},
	)
	if err != nil {
		t.Fatal(err)
	}

	return methods
}
