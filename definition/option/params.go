package option

import (
	"gitoa.ru/go-4devs/config/param"
)

type key int

const (
	paramHidden key = iota + 1
	paramRequired
	paramSlice
	paramBool
	paramShort
)

func Short(in rune) param.Option {
	return func(v param.Params) param.Params {
		return param.With(v, paramShort, string(in))
	}
}

func ParamShort(fn param.Params) (string, bool) {
	data, ok := param.String(fn, paramShort)

	return data, ok
}

func HasShort(short string) param.Has {
	return func(fn param.Params) bool {
		data, ok := param.String(fn, paramShort)

		return ok && data == short
	}
}

func WithType(in any) param.Option {
	return func(v param.Params) param.Params {
		out := param.WithType(in)(v)
		if _, ok := in.(bool); ok {
			return param.With(out, paramBool, ok)
		}

		return out
	}
}

func Position(pos uint64) param.Option {
	return param.WithPostition(pos)
}

func Hidden(v param.Params) param.Params {
	return param.With(v, paramHidden, true)
}

func Required(v param.Params) param.Params {
	return param.With(v, paramRequired, true)
}

func Slice(v param.Params) param.Params {
	return param.With(v, paramSlice, true)
}

func Default(in any) param.Option {
	return param.WithDefault(in)
}

// Deprecated: use param.WithDescription.
func Description(in string) param.Option {
	return param.WithDescription(in)
}

func HasDefaut(fn param.Params) bool {
	_, ok := param.Default(fn)

	return ok
}

// Deprecated: use param.Position.
func DataPosition(fn param.Params) (uint64, bool) {
	pos := param.Position(fn)

	return pos, pos != 0
}

// Deprecated: use param.Default.
func DataDefaut(fn param.Params) (any, bool) {
	return param.Default(fn)
}

func IsSlice(fn param.Params) bool {
	data, ok := param.Bool(paramSlice, fn)

	return ok && data
}

func IsBool(fn param.Params) bool {
	data, ok := param.Bool(paramBool, fn)

	return ok && data
}

func IsHidden(fn param.Params) bool {
	data, ok := param.Bool(paramHidden, fn)

	return ok && data
}

func IsRequired(fn param.Params) bool {
	data, ok := param.Bool(paramRequired, fn)

	return ok && data
}

// Deprecated: use param.Description.
func DataDescription(fn param.Params) string {
	return param.Description(fn)
}
