package command

import (
	"context"
	"fmt"
	"os"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/generate"
	"gitoa.ru/go-4devs/config/provider/chain"
	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/console/output"
)

const Name = "config:generate"

func Handle(ctx context.Context, in config.Provider, out output.Output, next command.ExecuteFn) error {
	var name string

	value, err := in.Value(ctx, generate.OptionFile)
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

	mem, merr := generate.NewMemoryProvider(name,
		generate.WithOutName(parser.OutName()),
		generate.WithFullPkg(parser.FullPkg()),
		generate.WithMethods(parser.Methods()...),
	)
	if merr != nil {
		return fmt.Errorf("mem provider:%w", merr)
	}

	return next(ctx, chain.New(in, mem), out)
}

func Execute(ctx context.Context, in config.Provider, _ output.Output) error {
	cfg := generate.NewConfigureConfig(ctx, in, generate.WithConfigureConfigHandle(func(context.Context, error) {}))

	if err := generate.Generate(ctx, cfg); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func Command() command.Command {
	return command.New(
		Name,
		"generate helper for configure command",
		Execute,
		command.Configure(generate.Configure),
		command.Handle(Handle),
	)
}
