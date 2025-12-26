package param

import (
	"slices"
)

var emptyParam = empty{}

type (
	Option func(p Params) Params
	Has    func(Params) bool
	Params interface {
		Param(key any) (any, bool)
	}
)

func Chain(vals ...Params) Params {
	slices.Reverse(vals)

	return chain(vals)
}

func With(parent Params, key, val any) Params {
	return value{
		Params: parent,
		key:    key,
		val:    val,
	}
}

func New(opts ...Option) Params {
	var parms Params

	parms = emptyParam

	for _, opt := range opts {
		parms = opt(parms)
	}

	return parms
}

type empty struct{}

func (v empty) Param(_ any) (any, bool) {
	return nil, false
}

type value struct {
	Params

	key, val any
}

func (v value) Param(key any) (any, bool) {
	if v.key == key {
		return v.val, true
	}

	return v.Params.Param(key)
}

type chain []Params

func (c chain) Param(key any) (any, bool) {
	for _, p := range c {
		val, ok := p.Param(key)
		if ok {
			return val, ok
		}
	}

	return nil, false
}
