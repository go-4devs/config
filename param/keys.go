package param

type key int

const (
	paramTimeFormat key = iota + 1
)

func WithTimeFormat(format string) Option {
	return func(p Param) Param {
		return With(p, paramTimeFormat, format)
	}
}

func TimeFormat(fn Param) (string, bool) {
	return String(paramTimeFormat, fn)
}
