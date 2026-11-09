package generate_test

import (
	"bytes"
	"context"
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

type UserAPI interface {
	Name(ctx context.Context) string
}

func TestGroupReturnsInterface(t *testing.T) {
	t.Parallel()

	def := definition.New(
		group.New("user", "configure user",
			option.String("name", "name"),
		).With(view.WithInterface[UserAPI]()),
		group.New("log", "configure logger",
			option.String("level", "level"),
		),
	).With(param.New(
		view.WithStructName("Example_Config"),
		view.WithStructSuffix("Config"),
	))

	var buf bytes.Buffer

	err := generate.Run(context.Background(), "github.com/go-4devs/config/definition/generate", &buf, def)
	require.NoError(t, err)

	out := buf.String()
	if !strings.Contains(out, "func (i ExampleConfig) User() generateTest.UserAPI {") {
		t.Fatalf("user return:\n%s", out)
	}

	if !strings.Contains(out, "return UserConfig{i}") {
		t.Fatalf("user value:\n%s", out)
	}

	if !strings.Contains(out, "func (i ExampleConfig) Log() LogConfig {") {
		t.Fatalf("log return:\n%s", out)
	}
}
