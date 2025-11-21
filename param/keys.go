package param

type key int

const (
	paramTimeFormat key = iota + 1
)

func WithTimeFormat(parent Param, format string) Param {
	return With(parent, paramTimeFormat, format)
}

func TimeFormat(fn Param) (string, bool) {
	return String(paramTimeFormat, fn)
}
