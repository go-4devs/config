package arg

import (
	"sync/atomic"

	"github.com/go-4devs/config/param"
)

type keyParam int

const (
	paramArgument keyParam = iota + 1
	paramDumpReferenceView
)

//nolint:gochecknoglobals
var argNum atomic.Uint64

// Argument marks the option as the next positional argument.
func Argument(v param.Params) param.Params {
	return param.With(v, paramArgument, argNum.Add(1)-1)
}

// ParamArgument returns the positional argument index.
func ParamArgument(fn param.Params) (uint64, bool) {
	return param.Uint64(paramArgument, fn)
}

// PosArgument reports whether the option has that positional index.
func PosArgument(in uint64) param.Has {
	return func(p param.Params) bool {
		idx, ok := ParamArgument(p)

		return ok && idx == in
	}
}

// HasArgument reports whether the option is a positional argument.
func HasArgument(fn param.Params) bool {
	_, ok := ParamArgument(fn)

	return ok
}
