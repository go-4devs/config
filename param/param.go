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

// Chain returns params searched from the last argument to the first. The first stored key wins.
func Chain(vals ...Params) Params {
	ordered := slices.Clone(vals)
	slices.Reverse(ordered)

	return chain(ordered)
}

// With returns params that store val under key and fall back to parent.
func With(parent Params, key, val any) Params {
	return value{
		Params: parent,
		key:    key,
		val:    val,
	}
}

// New applies opts to an empty parameter set.
func New(opts ...Option) Params {
	var parms Params

	parms = emptyParam

	for _, opt := range opts {
		parms = opt(parms)
	}

	return parms
}

type empty struct{}

// Param reports that the set is empty.
func (v empty) Param(_ any) (any, bool) {
	return nil, false
}

type value struct {
	Params

	key, val any
}

// Param returns val when key matches. Otherwise it asks the parent.
func (v value) Param(key any) (any, bool) {
	if v.key == key {
		return v.val, true
	}

	return v.Params.Param(key)
}

type chain []Params

// Param returns the first value found in the chain.
func (c chain) Param(key any) (any, bool) {
	for _, p := range c {
		val, ok := p.Param(key)
		if ok {
			return val, ok
		}
	}

	return nil, false
}
