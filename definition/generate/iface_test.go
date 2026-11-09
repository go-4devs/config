package generate_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/definition/generate"
	"gitoa.ru/go-4devs/config/definition/generate/view"
	"gitoa.ru/go-4devs/config/definition/group"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/test/require"
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

	err := generate.Run(context.Background(), "gitoa.ru/go-4devs/config/definition/generate", &buf, def)
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
