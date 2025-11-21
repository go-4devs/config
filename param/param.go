package param

import (
	"slices"
)

var (
	emptyParam = empty{}
)

type Param interface {
	Value(key any) (any, bool)
}

func Chain(vals ...Param) Param {
	slices.Reverse(vals)

	return chain(vals)
}

func With(parent Param, key, val any) Param {
	return value{
		Param: parent,
		key:   key,
		val:   val,
	}
}

func New() Param {
	return emptyParam
}

type empty struct{}

func (v empty) Value(_ any) (any, bool) {
	return nil, false
}

type value struct {
	Param

	key, val any
}

func (v value) Value(key any) (any, bool) {
	if v.key == key {
		return v.val, true
	}

	return v.Param.Value(key)
}

type chain []Param

func (c chain) Value(key any) (any, bool) {
	for _, p := range c {
		val, ok := p.Value(key)
		if ok {
			return val, ok
		}
	}

	return nil, false
}
