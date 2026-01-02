package csv

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/value"
)

type pkey int

const (
	paramDelimiter pkey = iota + 1
	paramParse
)

const defaultDelimiter = ','

func WithDelimiter(in rune) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, paramDelimiter, in)
	}
}

func WithInt(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.Atoi)
	})
}

func WithInt64(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseInt64)
	})
}

func WithFloat(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseFloat)
	})
}

func WithBool(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseBool)
	})
}

func WithUint(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseUint)
	})
}

func WithUint64(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseUint64)
	})
}

func WithDuration(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseDuration)
	})
}

func WithTime(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseTime)
	})
}

func WithParse(fn func(data []string) config.Value) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, paramParse, fn)
	}
}

func Csv(_ context.Context, in config.Value, opts ...param.Option) (config.Value, error) {
	sval, serr := in.ParseString()
	if serr != nil {
		return in, nil //nolint:nilerr
	}

	params := param.New(opts...)

	reader := csv.NewReader(bytes.NewBufferString(sval))
	reader.Comma = getDelimiter(params)

	data, rerr := reader.Read()
	if rerr != nil {
		return nil, fmt.Errorf("read csv:%w", rerr)
	}

	return csvValue(params, data)
}

func csvValue(params param.Params, data []string) (config.Value, error) {
	fn, ok := params.Param(paramParse)
	if !ok {
		return stringsValue(data)
	}

	parse, _ := fn.(func([]string) (config.Value, error))

	return parse(data)
}

func stringsValue(data []string) (config.Value, error) {
	return value.New(data), nil
}

func getDelimiter(params param.Params) rune {
	if name, ok := param.Rune(params, paramDelimiter); ok {
		return name
	}

	return defaultDelimiter
}
