package option

import (
	"gitoa.ru/go-4devs/config/param"
)

type key int

const (
	paramHidden key = iota + 1
	paramDefault
	paramDesc
	paramRequired
	paramSlice
	paramBool
	paramType
	paramPos
	paramShort
)

func Short(in rune) param.Option {
	return func(v param.Params) param.Params {
		return param.With(v, paramShort, string(in))
	}
}

func ParamShort(fn param.Params) (string, bool) {
	data, ok := param.String(paramShort, fn)

	return data, ok
}

func HasShort(short string, fn param.Params) bool {
	data, ok := param.String(paramShort, fn)

	return ok && data == short
}

func WithType(in any) param.Option {
	return func(v param.Params) param.Params {
		out := param.With(v, paramType, in)
		if _, ok := in.(bool); ok {
			return param.With(out, paramBool, ok)
		}

		return out
	}
}

func Position(pos uint64) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, paramPos, pos)
	}
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
	return func(v param.Params) param.Params {
		return param.With(v, paramDefault, in)
	}
}

func Description(in string) param.Option {
	return func(v param.Params) param.Params {
		return param.With(v, paramDesc, in)
	}
}

func HasDefaut(fn param.Params) bool {
	_, ok := fn.Param(paramDefault)

	return ok
}

func DataPosition(fn param.Params) (uint64, bool) {
	return param.Uint64(paramPos, fn)
}

func DataDefaut(fn param.Params) (any, bool) {
	data, ok := fn.Param(paramDefault)

	return data, ok
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

func DataType(fn param.Params) any {
	param, _ := fn.Param(paramType)

	return param
}

func DataDescription(fn param.Params) string {
	data, _ := param.String(paramDesc, fn)

	return data
}
