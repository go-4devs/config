package key

import (
	"context"
	"fmt"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/value"
)

type pkey int

const paramKey pkey = iota

func WithKey(in string) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, paramKey, in)
	}
}

func Key(_ context.Context, in config.Value, opts ...param.Option) (config.Value, error) {
	data := make(map[string]any, 0)
	if err := in.Unmarshal(&data); err != nil {
		return nil, fmt.Errorf("unmarshal:%w", err)
	}

	key, ok := getKey(opts...)
	if !ok {
		return nil, fmt.Errorf("key is %w", config.ErrRequired)
	}

	val, vok := data[key]
	if !vok {
		return nil, fmt.Errorf("value by key[%v]: %w", key, config.ErrNotFound)
	}

	return value.New(val), nil
}

func getKey(opts ...param.Option) (string, bool) {
	params := param.New(opts...)

	if name, ok := param.String(params, paramKey); ok {
		return name, ok
	}

	return "", false
}
