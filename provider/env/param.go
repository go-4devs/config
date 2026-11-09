package env

import (
	"github.com/go-4devs/config/param"
)

type ekey int

const (
	keyDeprecated ekey = iota + 1
	keyAlias
)

// Deprecated stores the deprecated environment variable name.
func Deprecated(name string) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, keyDeprecated, name)
	}
}

// ParamDeprecated returns the deprecated environment variable name.
func ParamDeprecated(in param.Params) (any, bool) {
	return in.Param(keyDeprecated)
}

// Alias stores an environment variable name to read instead.
func Alias(name string) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, keyAlias, name)
	}
}

// ParamAlias returns the alias environment variable name.
func ParamAlias(in param.Params) (any, bool) {
	return in.Param(keyAlias)
}
