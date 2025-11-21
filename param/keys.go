package param

type key int

const (
	paramTimeFormat key = iota + 1
)

func WithTimeFormat(format string) Option {
	return func(p Params) Params {
		return With(p, paramTimeFormat, format)
	}
}

func TimeFormat(fn Params) (string, bool) {
	return String(paramTimeFormat, fn)
}
