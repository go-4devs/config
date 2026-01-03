package view

import (
	"fmt"
	"reflect"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
)

type key int

const (
	viewParamFunctName key = iota + 1
	viewParamSkipContext
	viewParamStructName
	viewParamStructPrefix
	viewParamStructSuffix
)

func WithStructName(name string) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, viewParamStructName, name)
	}
}

func WithStructPrefix(prefix string) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, viewParamStructPrefix, prefix)
	}
}

func WithStructSuffix(suffix string) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, viewParamStructSuffix, suffix)
	}
}

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

func WithParent(in *View) Option {
	return func(v *View) {
		v.parent = in
	}
}

func NewViews(option config.Group, opts ...Option) View {
	view := newView(option, opts...)

	for _, op := range option.Options() {
		view.children = append(view.children, NewView(op, WithParent(&view)))
	}

	return view
}

type IOption any

func newView(option config.Option, opts ...Option) View {
	vi := View{
		Option:   option,
		parent:   nil,
		children: nil,
	}

	for _, opt := range opts {
		opt(&vi)
	}

	return vi
}

func NewView(opt config.Option, opts ...Option) View {
	vi := newView(opt, opts...)

	if data, ok := opt.(config.Group); ok {
		for _, chi := range data.Options() {
			vi.children = append(vi.children, NewView(chi, WithParent(&vi)))
		}
	}

	return vi
}

type View struct {
	config.Option

	children []View
	parent   *View
}

func (v View) Types() []any {
	types := make([]any, 0)
	if v.Type() != "" {
		types = append(types, v.Type())
	}

	for _, child := range v.children {
		types = append(types, child.Types()...)
	}

	return types
}

func (v View) Kind() reflect.Type {
	return reflect.TypeOf(v.Option)
}

func (v View) Views() []View {
	return v.children
}

func (v View) Param(key any) string {
	data, has := v.Option.Param(key)
	if has {
		return fmt.Sprintf("%v", data)
	}

	if v.parent != nil {
		return v.parent.Param(key)
	}

	return ""
}

func (v View) ClildSkipContext() bool {
	for _, child := range v.children {
		if child.SkipContext() {
			return true
		}
	}

	return false
}

func (v View) SkipContext() bool {
	if IsSkipContext(v.Option) {
		return true
	}

	if v.parent != nil {
		return v.parent.SkipContext()
	}

	return false
}

func (v View) Name() string {
	return v.Option.Name()
}

func (v View) Keys() []string {
	keys := make([]string, 0, 1)
	if v.parent != nil {
		keys = append(keys, v.parent.Keys()...)
	}

	if name := v.Option.Name(); name != "" {
		keys = append(keys, name)
	}

	return keys
}

func (v View) Type() any {
	return param.Type(v.Option)
}

func (v View) FuncName() string {
	data, ok := v.Option.Param(viewParamFunctName)
	name, valid := data.(string)

	if !ok || !valid {
		return v.Name()
	}

	return name
}

func (v View) StructName() string {
	name, ok := param.String(v.Option, viewParamStructName)
	if ok {
		return name
	}

	keys := make([]string, 0, len(v.Keys())+2) //nolint:mnd

	prefix := v.Param(viewParamStructPrefix)
	if prefix != "" {
		keys = append(keys, prefix)
	}

	keys = append(keys, v.Keys()...)

	suffix := v.Param(viewParamStructSuffix)
	if suffix != "" {
		keys = append(keys, suffix)
	}

	return strings.Join(keys, "_")
}

func (v View) ParentStruct() string {
	if v.parent == nil {
		return ""
	}

	return v.parent.StructName()
}

func (v View) Description() string {
	return param.Description(v.Option)
}

func (v View) Default() any {
	data, ok := option.DataDefaut(v.Option)
	if !ok {
		return nil
	}

	return data
}

func (v View) HasDefault() bool {
	return option.HasDefaut(v.Option)
}
