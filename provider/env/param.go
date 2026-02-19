package env

import (
	"gitoa.ru/go-4devs/config/param"
)

type ekey int

const (
	keyDeprecated ekey = iota + 1
	keyAlias
)

func Deprecated(name string) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, keyDeprecated, name)
	}
}

func ParamDeprecated(in param.Params) (any, bool) {
	return in.Param(keyDeprecated)
}

func Alias(name string) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, keyAlias, name)
	}
}

func ParamAlias(in param.Params) (any, bool) {
	return in.Param(keyAlias)
}
