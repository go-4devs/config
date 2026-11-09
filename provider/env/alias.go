package env

import (
	"context"
	"fmt"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/param"
	"github.com/go-4devs/config/processor/env"
	"github.com/go-4devs/config/provider/memory"
)

const (
	NameAlias = "ealias"
	NameDepr  = "edeprecated"
)

// NewAlias returns a provider that reads the environment variable named by an alias.
func NewAlias() *memory.Param {
	return memory.NewParam(NameAlias, ParamAlias, memory.WithParamProcess(env.Env))
}

// NewDeprecated returns a provider that reads a deprecated environment variable and calls notice.
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
