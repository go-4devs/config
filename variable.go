package config

import (
	"fmt"

	"gitoa.ru/go-4devs/config/key"
	"gitoa.ru/go-4devs/config/param"
)

func NewVar(opt Option) Variable {
	return Variable{
		param: opt,
		names: []string{opt.Name()},
	}
}

type Variable struct {
	names []string
	param param.Params
}

func (v Variable) Param(key any) (any, bool) {
	return v.param.Param(key)
}

func (v Variable) Key() []string {
	return v.names
}

func (v Variable) Type() any {
	return param.Type(v)
}

func (v Variable) With(opt Option) Variable {
	return Variable{
		names: append(v.Key(), opt.Name()),
		param: param.Chain(v.param, opt),
	}
}

func NewVars(opts ...Option) Vars {
	vars := newVars(opts...)
	pos := key.Map{}

	for idx, one := range vars {
		pos.Add(idx, one.Key())
	}

	return Vars{
		vars: vars,
		pos:  pos,
	}
}

func newVars(opts ...Option) []Variable {
	vars := make([]Variable, 0, len(opts))
	for _, opt := range opts {
		one := NewVar(opt)
		switch data := opt.(type) {
		case Group:
			vars = append(vars, groupVars(one, data.Options()...)...)
		default:
			vars = append(vars, one)
		}
	}

	return vars
}

func groupVars(parent Variable, opts ...Option) []Variable {
	vars := make([]Variable, 0, len(opts))
	for _, opt := range opts {
		switch data := opt.(type) {
		case Group:
			vars = append(vars, groupVars(parent.With(opt), data.Options()...)...)
		default:
			vars = append(vars, parent.With(opt))
		}
	}

	return vars
}

type Vars struct {
	vars []Variable
	pos  key.Map
}

func (d Vars) Variables() []Variable {
	return d.vars
}

func (d Vars) ByName(path ...string) (Variable, error) {
	if idx, ok := d.pos.Index(path); ok {
		return d.vars[idx], nil
	}

	return Variable{}, ErrNotFound
}

func (d Vars) ByParam(fn param.Has) (Variable, error) {
	opts := d.filter(fn)
	switch {
	case len(opts) == 0:
		return Variable{}, fmt.Errorf("%w:%T", ErrNotFound, fn)
	case len(opts) > 1:
		return Variable{}, fmt.Errorf("%w:%v", ErrToManyArgs, opts)
	default:
		return opts[0], nil
	}
}

func (d Vars) filter(fn param.Has) []Variable {
	opts := make([]Variable, 0, len(d.vars))
	for idx := range d.vars {
		if fn(d.vars[idx]) {
			opts = append(opts, d.vars[idx])
		}
	}

	return opts
}
