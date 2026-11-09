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

// WithDelimiter sets the CSV field separator.
func WithDelimiter(in rune) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, paramDelimiter, in)
	}
}

// WithInt parses CSV fields as ints.
func WithInt(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.Atoi)
	})
}

// WithInt64 parses CSV fields as int64s.
func WithInt64(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseInt64)
	})
}

// WithFloat parses CSV fields as float64s.
func WithFloat(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseFloat)
	})
}

// WithBool parses CSV fields as bools.
func WithBool(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseBool)
	})
}

// WithUint parses CSV fields as uints.
func WithUint(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseUint)
	})
}

// WithUint64 parses CSV fields as uint64s.
func WithUint64(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseUint64)
	})
}

// WithDuration parses CSV fields as durations.
func WithDuration(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseDuration)
	})
}

// WithTime parses CSV fields as RFC3339 times.
func WithTime(p param.Params) param.Params {
	return param.With(p, paramParse, func(data []string) (config.Value, error) {
		return value.ParseSlice(data, value.ParseTime)
	})
}

// WithParse sets the function that converts CSV fields into a value.
func WithParse(fn func(data []string) config.Value) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, paramParse, fn)
	}
}

// Csv reads one CSV record from the string value. A value that is not a string is returned unchanged.
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
