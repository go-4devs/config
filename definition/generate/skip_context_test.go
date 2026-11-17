package generate_test

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/go-4devs/config/definition"
	"github.com/go-4devs/config/definition/generate"
	"github.com/go-4devs/config/definition/generate/view"
	"github.com/go-4devs/config/definition/group"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/param"
	"github.com/go-4devs/config/test/require"
)

func TestGenerateSkipContext(t *testing.T) {
	t.Parallel()

	t.Run("nested", func(t *testing.T) {
		t.Parallel()

		out := generated(t, definition.New(
			group.New("live", "live",
				option.String("id", "id"),
				group.New("server", "server",
					option.String("address", "address"),
				).With(view.WithSkipContext),
			),
		))

		contains(t, out, "func NewRoot( ctx context.Context,prov config.Provider, opts ...func(*Root)) Root {")
		contains(t, out, "func (i Live) Id( ctx context.Context ) string {")
		contains(t, out, "func (i LiveServer) Address() string {")
		contains(t, out, "i.readAddress(i.ctx)")
		compileGenerated(t, out)
	})

	t.Run("direct", func(t *testing.T) {
		t.Parallel()

		out := generated(t, definition.New(
			option.String("id", "id", view.WithSkipContext),
		))

		contains(t, out, "func NewRoot( ctx context.Context,prov config.Provider, opts ...func(*Root)) Root {")
		contains(t, out, "func (i Root) Id() string {")
		contains(t, out, "i.readId(i.ctx)")
		compileGenerated(t, out)
	})

	t.Run("none", func(t *testing.T) {
		t.Parallel()

		out := generated(t, definition.New(
			group.New("live", "live",
				option.String("id", "id"),
				group.New("server", "server",
					option.String("address", "address"),
				),
			),
		))

		contains(t, out, "func NewRoot(prov config.Provider, opts ...func(*Root)) Root {")
		contains(t, out, "func (i Live) Id( ctx context.Context ) string {")
		contains(t, out, "func (i LiveServer) Address( ctx context.Context ) string {")
		missing(t, out, "\tctx context.Context")
		missing(t, out, "ctx: ctx")
		compileGenerated(t, out)
	})
}

func generated(t *testing.T, def *definition.Definition) string {
	t.Helper()

	def = def.With(param.New(view.WithStructName("Root")))

	var buf bytes.Buffer

	err := generate.Run(context.Background(), "example.com/skipctx", &buf, def)
	require.NoError(t, err)

	return buf.String()
}

func contains(t *testing.T, out, part string) {
	t.Helper()

	if !strings.Contains(out, part) {
		t.Fatalf("missing %q in:\n%s", part, out)
	}
}

func missing(t *testing.T, out, part string) {
	t.Helper()

	if strings.Contains(out, part) {
		t.Fatalf("unexpected %q in:\n%s", part, out)
	}
}

func compileGenerated(t *testing.T, src string) {
	t.Helper()

	dir := t.TempDir()

	err := os.WriteFile(filepath.Join(dir, "config.go"), []byte(src), 0o600)
	require.NoError(t, err)

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller")
	}

	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
	mod := "module example.com/skipctx\n\ngo 1.27.3\n\nrequire github.com/go-4devs/config v0.0.0\n\nreplace github.com/go-4devs/config => " +
		filepath.ToSlash(root) + "\n"

	err = os.WriteFile(filepath.Join(dir, "go.mod"), []byte(mod), 0o600)
	require.NoError(t, err)

	cmd := exec.CommandContext(t.Context(), "go", "build", ".")
	cmd.Dir = dir

	cmd.Env = append(os.Environ(), "GOTOOLCHAIN=local", "GOPROXY=off", "GOSUMDB=off", "GOWORK=off")

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("build: %v\n%s\n%s", err, out, src)
	}
}
