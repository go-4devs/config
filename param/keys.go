package param

type key int

const (
	paramTimeFormat key = iota + 1
	paramType
	paramDescription
	paramDefault
	paramPosition
)

// WithTimeFormat stores a time layout.
func WithTimeFormat(format string) Option {
	return func(p Params) Params {
		return With(p, paramTimeFormat, format)
	}
}

// TimeFormat returns the stored time layout.
func TimeFormat(fn Params) (string, bool) {
	return String(fn, paramTimeFormat)
}

// WithType stores the value type.
func WithType(in any) Option {
	return func(v Params) Params {
		return With(v, paramType, in)
	}
}

// Type returns the stored value type.
func Type(fn Params) any {
	param, _ := fn.Param(paramType)

	return param
}

// WithDescription stores the description.
func WithDescription(in string) Option {
	return func(p Params) Params {
		return With(p, paramDescription, in)
	}
}

// Description returns the stored description.
func Description(fn Params) string {
	data, _ := String(fn, paramDescription)

	return data
}

// WithDefault stores the default value.
func WithDefault(in any) Option {
	return func(p Params) Params {
		return With(p, paramDefault, in)
	}
}

// Default returns the stored default value.
func Default(p Params) (any, bool) {
	data, ok := p.Param(paramDefault)

	return data, ok
}

// WithPostition stores the positional argument index.
func WithPostition(in uint64) Option {
	return func(p Params) Params {
		return With(p, paramPosition, in)
	}
}

// Position returns the stored positional argument index.
func Position(in Params) uint64 {
	pos, _ := Uint64(paramPosition, in)

	return pos
}
