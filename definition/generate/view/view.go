package view

import (
	"fmt"
	"reflect"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
)

type key int

const (
	viewParamFunctName key = iota + 1
	viewParamSkipContext
)

func WithSkipContext(p param.Params) param.Params {
	return param.With(p, viewParamSkipContext, true)
}

func WithContext(p param.Params) param.Params {
	return param.With(p, viewParamSkipContext, false)
}

func IsSkipContext(p param.Params) bool {
	data, has := p.Param(viewParamSkipContext)

	if has {
		skip, ok := data.(bool)

		return ok && skip
	}

	return false
}

type Option func(*View)

func WithParent(name string) Option {
	return func(v *View) {
		v.parent = name
	}
}

func WithKeys(keys ...string) Option {
	return func(v *View) {
		v.keys = keys
	}
}

func NewViews(name string, get param.Params, option config.Options, opts ...Option) View {
	view := newView(name, get, option, opts...)

	for _, op := range option.Options() {
		view.children = append(view.children, NewView(op, get, WithParent(name)))
	}

	return view
}

type IOption any

func newView(name string, get param.Params, in any, opts ...Option) View {
	vi := View{
		kind:     reflect.TypeOf(in),
		name:     name,
		Params:   get,
		dtype:    param.Type(get),
		children: nil,
		keys:     nil,
		parent:   "",
	}

	for _, opt := range opts {
		opt(&vi)
	}

	return vi
}

func NewView(opt config.Option, get param.Params, opts ...Option) View {
	vi := newView(opt.Name(), param.Chain(get, opt), opt, opts...)

	if data, ok := opt.(config.Group); ok {
		for _, chi := range data.Options() {
			vi.children = append(vi.children, NewView(
				chi,
				param.Chain(vi.Params, chi),
				WithParent(vi.ParentName()+"_"+opt.Name()),
				WithKeys(append(vi.keys, opt.Name())...),
			))
		}
	}

	return vi
}

type View struct {
	param.Params

	children []View
	keys     []string
	kind     reflect.Type
	name     string
	parent   string
	dtype    any
}

func (v View) Types() []any {
	types := make([]any, 0)
	if v.dtype != nil {
		types = append(types, v.dtype)
	}

	for _, child := range v.children {
		types = append(types, child.Types()...)
	}

	return types
}

func (v View) Kind() reflect.Type {
	return v.kind
}

func (v View) Views() []View {
	return v.children
}

func (v View) Param(key any) string {
	data, ok := v.Params.Param(key)
	if !ok {
		return ""
	}

	if res, ok := data.(string); ok {
		return res
	}

	return fmt.Sprintf("%v", data)
}

func (v View) SkipContext() bool {
	return IsSkipContext(v.Params)
}

func (v View) Name() string {
	return v.name
}

func (v View) Keys() []string {
	return v.keys
}

func (v View) Type() any {
	return v.dtype
}

func (v View) FuncName() string {
	data, ok := v.Params.Param(viewParamFunctName)
	name, valid := data.(string)

	if !ok || !valid {
		return v.name
	}

	return name
}

func (v View) ParentName() string {
	return v.parent
}

func (v View) Description() string {
	return option.DataDescription(v.Params)
}

func (v View) Default() any {
	data, ok := option.DataDefaut(v.Params)
	if !ok {
		return nil
	}

	return data
}

func (v View) HasDefault() bool {
	return option.HasDefaut(v.Params)
}
