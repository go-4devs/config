package example_test

import (
	"context"
	"os"
	"testing"

	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/definition/generate"
	"gitoa.ru/go-4devs/config/definition/generate/bootstrap"
	"gitoa.ru/go-4devs/config/definition/generate/example"
)

func TestGenerate_Bootstrap(t *testing.T) {
	t.SkipNow()
	t.Parallel()

	ctx := context.Background()
	options := definition.New()
	_ = example.Config(ctx, options)

	cfg, _ := generate.NewGConfig("./config.go",
		generate.WithMethods("Config"),
		generate.WithFullPkg("gitoa.ru/go-4devs/config/definition/generate/example"),
	)

	path, err := bootstrap.Bootstrap(ctx, cfg)
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
	_ = example.Config(ctx, options)

	cfg, _ := generate.NewGConfig("./config.go",
		generate.WithMethods("Config"),
		generate.WithFullPkg("gitoa.ru/go-4devs/config/definition/generate/example"),
	)

	err := generate.Generate(ctx, cfg)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.FailNow()
}
