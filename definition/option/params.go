package option

import (
	"github.com/go-4devs/config/param"
)

type key int

const (
	paramHidden key = iota + 1
	paramRequired
	paramSlice
	paramBool
	paramShort
)

// Short stores a one-letter flag name.
func Short(in rune) param.Option {
	return func(v param.Params) param.Params {
		return param.With(v, paramShort, string(in))
	}
}

// ParamShort returns the stored one-letter flag name.
func ParamShort(fn param.Params) (string, bool) {
	data, ok := param.String(fn, paramShort)

	return data, ok
}

// HasShort reports whether the option has that one-letter flag name.
func HasShort(short string) param.Has {
	return func(fn param.Params) bool {
		data, ok := param.String(fn, paramShort)

		return ok && data == short
	}
}

// WithType stores the option value type. A bool type is also marked as a bool flag.
func WithType(in any) param.Option {
	return func(v param.Params) param.Params {
		out := param.WithType(in)(v)
		if _, ok := in.(bool); ok {
			return param.With(out, paramBool, ok)
		}

		return out
	}
}

// Position stores the positional argument index.
func Position(pos uint64) param.Option {
	return param.WithPostition(pos)
}

// Hidden marks the option as hidden.
func Hidden(v param.Params) param.Params {
	return param.With(v, paramHidden, true)
}

// Required marks the option as required.
func Required(v param.Params) param.Params {
	return param.With(v, paramRequired, true)
}

// Slice marks the option as a slice.
func Slice(v param.Params) param.Params {
	return param.With(v, paramSlice, true)
}

// Default stores the default value.
func Default(in any) param.Option {
	return param.WithDefault(in)
}

// Deprecated: use param.WithDescription.
func Description(in string) param.Option {
	return param.WithDescription(in)
}

// HasDefaut reports whether a default value is stored.
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

// IsSlice reports whether the option is a slice.
func IsSlice(fn param.Params) bool {
	data, ok := param.Bool(paramSlice, fn)

	return ok && data
}

// IsBool reports whether the option is a bool flag.
func IsBool(fn param.Params) bool {
	data, ok := param.Bool(paramBool, fn)

	return ok && data
}

// IsHidden reports whether the option is hidden.
func IsHidden(fn param.Params) bool {
	data, ok := param.Bool(paramHidden, fn)

	return ok && data
}

// IsRequired reports whether the option is required.
func IsRequired(fn param.Params) bool {
	data, ok := param.Bool(paramRequired, fn)

	return ok && data
}

// Deprecated: use param.Description.
func DataDescription(fn param.Params) string {
	return param.Description(fn)
}
