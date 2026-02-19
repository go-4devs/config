package env

import (
	"context"
	"fmt"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/processor/env"
	"gitoa.ru/go-4devs/config/provider/memory"
)

const (
	NameAlias = "ealias"
	NameDepr  = "edeprecated"
)

func NewAlias() *memory.Param {
	return memory.NewParam(NameAlias, ParamAlias, memory.WithParamProcess(env.Env))
}

func NewDeprecated(notice func(context.Context, string, ...any)) *memory.Param {
	return memory.NewParam(NameDepr, ParamDeprecated, memory.WithParamProcess(deprNotice(notice)))
}

func deprNotice(notice func(context.Context, string, ...any)) config.ProcessFunc {
	return func(ctx context.Context, in config.Value, opts ...param.Option) (config.Value, error) {
		val, err := env.Env(ctx, in, opts...)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		notice(ctx, "env %v is depreccated", in)

		return val, nil
	}
}
