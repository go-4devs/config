package arg

import (
	"sync/atomic"

	"gitoa.ru/go-4devs/config/param"
)

type keyParam int

const (
	paramArgument keyParam = iota + 1
	paramDumpReferenceView
)

//nolint:gochecknoglobals
var argNum uint64

func Argument(v param.Params) param.Params {
	return param.With(v, paramArgument, atomic.AddUint64(&argNum, 1)-1)
}

func ParamArgument(fn param.Params) (uint64, bool) {
	return param.Uint64(paramArgument, fn)
}

func PosArgument(in uint64) param.Has {
	return func(p param.Params) bool {
		idx, ok := ParamArgument(p)

		return ok && idx == in
	}
}

func HasArgument(fn param.Params) bool {
	_, ok := ParamArgument(fn)

	return ok
}
