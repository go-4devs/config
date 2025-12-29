package generate_test

import (
	"context"
	"os"
	"testing"

	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/definition/generate"
	"gitoa.ru/go-4devs/config/definition/generate/bootstrap"
	"gitoa.ru/go-4devs/config/definition/generate/view"
	"gitoa.ru/go-4devs/config/definition/group"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/definition/proto"
	"gitoa.ru/go-4devs/config/test/require"
)

type LogLevel string

func (l *LogLevel) UnmarshalText(in []byte) error {
	data := string(in)
	*l = LogLevel(data)

	return nil
}

func Configure(_ context.Context, def *definition.Definition) error {
	def.Add(
		option.String("test", "test string", view.WithSkipContext),
		group.New("user", "configure user",
			option.String("name", "name", option.Default("4devs")),
			option.String("pass", "password"),
		),

		group.New("log", "configure logger",
			option.New("level", "log level", LogLevel("")),
			proto.New("service", "servise logger", option.New("level", "log level", LogLevel(""))),
		),
	)

	return nil
}

func TestGenerate_Bootstrap(t *testing.T) {
	t.SkipNow()
	t.Parallel()

	ctx := context.Background()
	options := definition.New()
	err := Configure(ctx, options)
	require.NoError(t, err)

	cfg, _ := generate.NewMemoryProvider("generate_test.go",
		generate.WithMethods("Config"),
		generate.WithFullPkg("gitoa.ru/go-4devs/config/definition/generate_test"),
	)

	path, err := bootstrap.Bootstrap(ctx, generate.NewConfigure(ctx, cfg))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	os.Remove(path)

	t.Log(path)
	t.FailNow()
}

func TestGenerate_Genereate(t *testing.T) {
	t.SkipNow()
	t.Parallel()

	ctx := context.Background()
	options := definition.New()
	err := Configure(ctx, options)
	require.NoError(t, err)

	cfg, _ := generate.NewMemoryProvider("generate_test.go",
		generate.WithMethods("Config"),
		generate.WithFullPkg("gitoa.ru/go-4devs/config/definition/generate_test"),
	)

	err = generate.Generate(ctx, generate.NewConfigure(ctx, cfg))
	require.NoError(t, err)

	t.FailNow()
}
