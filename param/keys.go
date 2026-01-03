package param

type key int

const (
	paramTimeFormat key = iota + 1
	paramType
	paramDescription
)

func WithTimeFormat(format string) Option {
	return func(p Params) Params {
		return With(p, paramTimeFormat, format)
	}
}

func TimeFormat(fn Params) (string, bool) {
	return String(fn, paramTimeFormat)
}

func WithType(in any) Option {
	return func(v Params) Params {
		return With(v, paramType, in)
	}
}

func Type(fn Params) any {
	param, _ := fn.Param(paramType)

	return param
}

func WithDescription(in string) Option {
	return func(p Params) Params {
		return With(p, paramDescription, in)
	}
}

func Description(fn Params) string {
	data, _ := String(fn, paramDescription)

	return data
}
