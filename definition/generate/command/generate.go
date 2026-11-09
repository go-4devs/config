package command

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/go-4devs/config/definition/generate"
	"github.com/go-4devs/config/provider/memory"
	gconfig "gitoa.ru/go-4devs/config"
	ggenerate "gitoa.ru/go-4devs/config/definition/generate"
	gchain "gitoa.ru/go-4devs/config/provider/chain"
	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/console/output"
)

const Name = "config:generate"

// Handle fills generator settings from the source file.
func Handle(ctx context.Context, in gconfig.Provider, out output.Output, next command.ExecuteFn) error {
	var name string

	value, err := in.Value(ctx, ggenerate.OptionFile)
	if err == nil {
		name = value.String()
	}

	if name == "" {
		name = os.Getenv("GOFILE")
	}

	parser, err := generate.Parse(ctx, name)
	if err != nil {
		return fmt.Errorf("parse:%w", err)
	}

	mem, merr := ggenerate.NewMemoryProvider(name,
		ggenerate.WithOutName(parser.OutName()),
		ggenerate.WithFullPkg(parser.FullPkg()),
		ggenerate.WithMethods(parser.Methods()...),
	)
	if merr != nil {
		return fmt.Errorf("mem provider:%w", merr)
	}

	return next(ctx, gchain.New(in, mem), out)
}

// Execute runs the generator.
func Execute(ctx context.Context, in gconfig.Provider, _ output.Output) error {
	mem, err := localSettings(ctx, in)
	if err != nil {
		return err
	}

	cfg := generate.NewConfigureConfig(ctx, mem, generate.WithConfigureConfigHandle(func(context.Context, error) {}))
	if err = generate.Generate(ctx, cfg); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Command returns the config:generate console command.
func Command() command.Command {
	return command.New(
		Name,
		"generate helper for configure command",
		Execute,
		command.Configure(ggenerate.Configure),
		command.Handle(Handle),
	)
}

func localSettings(ctx context.Context, in gconfig.Provider) (*memory.Map, error) {
	mem := memory.NewMap("generate")

	for _, key := range []string{"file", "prefix", "suffix", "build-tags", "out-name", "full-pkg"} {
		if err := copyString(ctx, in, mem, key); err != nil {
			return nil, err
		}
	}

	if err := copyBool(ctx, in, mem, "skip-context"); err != nil {
		return nil, err
	}

	for _, key := range []string{"method", "leave-temp"} {
		if err := copyStrings(ctx, in, mem, key); err != nil {
			return nil, err
		}
	}

	return mem, nil
}

func copyString(ctx context.Context, in gconfig.Provider, mem *memory.Map, key string) error {
	val, err := in.Value(ctx, key)
	if errors.Is(err, gconfig.ErrNotFound) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("%s:%w", key, err)
	}

	text, err := val.ParseString()
	if err != nil {
		return fmt.Errorf("%s:%w", key, err)
	}

	mem.SetOption(text, key)

	return nil
}

func copyBool(ctx context.Context, in gconfig.Provider, mem *memory.Map, key string) error {
	val, err := in.Value(ctx, key)
	if errors.Is(err, gconfig.ErrNotFound) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("%s:%w", key, err)
	}

	flag, err := val.ParseBool()
	if err != nil {
		return fmt.Errorf("%s:%w", key, err)
	}

	mem.SetOption(flag, key)

	return nil
}

func copyStrings(ctx context.Context, in gconfig.Provider, mem *memory.Map, key string) error {
	val, err := in.Value(ctx, key)
	if errors.Is(err, gconfig.ErrNotFound) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("%s:%w", key, err)
	}

	var list []string
	if err = val.Unmarshal(&list); err != nil {
		return fmt.Errorf("%s:%w", key, err)
	}

	mem.SetOption(list, key)

	return nil
}
