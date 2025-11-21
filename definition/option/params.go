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
	return func(v param.Param) param.Param {
		return param.With(v, paramShort, string(in))
	}
}

func ParamShort(fn param.Param) (string, bool) {
	data, ok := param.String(paramShort, fn)

	return data, ok
}

func HasShort(short string, fn param.Param) bool {
	data, ok := param.String(paramShort, fn)

	return ok && data == short
}

func WithType(in any) param.Option {
	return func(v param.Param) param.Param {
		out := param.With(v, paramType, in)
		if _, ok := in.(bool); ok {
			return param.With(out, paramBool, ok)
		}

		return out
	}
}

func Position(pos uint64) param.Option {
	return func(p param.Param) param.Param {
		return param.With(p, paramPos, pos)
	}
}

func Hidden(v param.Param) param.Param {
	return param.With(v, paramHidden, true)
}

func Required(v param.Param) param.Param {
	return param.With(v, paramRequired, true)
}

func Slice(v param.Param) param.Param {
	return param.With(v, paramSlice, true)
}

func Default(in any) param.Option {
	return func(v param.Param) param.Param {
		return param.With(v, paramDefault, in)
	}
}

func Description(in string) param.Option {
	return func(v param.Param) param.Param {
		return param.With(v, paramDesc, in)
	}
}

func HasDefaut(fn param.Param) bool {
	_, ok := fn.Value(paramDefault)

	return ok
}

func DataPosition(fn param.Param) (uint64, bool) {
	return param.Uint64(paramPos, fn)
}

func DataDefaut(fn param.Param) (any, bool) {
	data, ok := fn.Value(paramDefault)

	return data, ok
}

func IsSlice(fn param.Param) bool {
	data, ok := param.Bool(paramSlice, fn)

	return ok && data
}

func IsBool(fn param.Param) bool {
	data, ok := param.Bool(paramBool, fn)

	return ok && data
}

func IsHidden(fn param.Param) bool {
	data, ok := param.Bool(paramHidden, fn)

	return ok && data
}

func IsRequired(fn param.Param) bool {
	data, ok := param.Bool(paramRequired, fn)

	return ok && data
}

func DataType(fn param.Param) any {
	param, _ := fn.Value(paramType)

	return param
}

func DataDescription(fn param.Param) string {
	data, _ := param.String(paramDesc, fn)

	return data
}
