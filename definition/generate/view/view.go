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
	viewParamInterface
)

// WithFuncName sets the generated function name.
func WithFuncName(name string) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, viewParamFunctName, name)
	}
}

// WithStructName sets the generated struct name.
func WithStructName(name string) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, viewParamStructName, name)
	}
}

// WithStructPrefix sets the prefix of a generated struct name.
func WithStructPrefix(prefix string) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, viewParamStructPrefix, prefix)
	}
}

// WithStructSuffix sets the suffix of a generated struct name.
func WithStructSuffix(suffix string) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, viewParamStructSuffix, suffix)
	}
}

// WithInterface sets the type returned by the generated group method.
func WithInterface[T any]() param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, viewParamInterface, reflect.TypeFor[T]())
	}
}

// WithSkipContext marks generated methods that do not take context.
func WithSkipContext(p param.Params) param.Params {
	return param.With(p, viewParamSkipContext, true)
}

// WithContext marks generated methods that take context.
func WithContext(p param.Params) param.Params {
	return param.With(p, viewParamSkipContext, false)
}

// IsSkipContext reports whether generated methods omit context.
func IsSkipContext(p param.Params) bool {
	data, has := p.Param(viewParamSkipContext)

	if has {
		skip, ok := data.(bool)

		return ok && skip
	}

	return false
}

type Option func(*View)

// WithParent sets the parent view.
func WithParent(in *View) Option {
	return func(v *View) {
		v.parent = in
	}
}

// NewViews returns a view tree for the options.
func NewViews(option config.Options, opts ...Option) View {
	view := newView(option, opts...)

	for _, op := range option.Options() {
		view.children = append(view.children, NewView(op, WithParent(&view)))
	}

	return view
}

func newView(params param.Params, opts ...Option) View {
	vi := View{
		Params:   params,
		parent:   nil,
		children: nil,
	}

	for _, opt := range opts {
		opt(&vi)
	}

	return vi
}

// NewView returns a view for one option. A group includes its children.
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
	param.Params

	children []View
	parent   *View
}

// Types returns the option value types in this view and its children.
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

// Kind returns the dynamic type stored in the view.
func (v View) Kind() reflect.Type {
	return reflect.TypeOf(v.Params)
}

// Views returns the child views.
func (v View) Views() []View {
	return v.children
}

// Param returns the parameter as text, walking to the parent when it is missing.
func (v View) Param(key any) string {
	data, has := v.Params.Param(key)
	if has {
		return fmt.Sprintf("%v", data)
	}

	if v.parent != nil {
		return v.parent.Param(key)
	}

	return ""
}

// ClildSkipContext reports whether any child omits context.
func (v View) ClildSkipContext() bool {
	for _, child := range v.children {
		if child.SkipContext() {
			return true
		}
	}

	return false
}

// SkipContext reports whether this view or a parent omits context.
func (v View) SkipContext() bool {
	if IsSkipContext(v.Params) {
		return true
	}

	if v.parent != nil {
		return v.parent.SkipContext()
	}

	return false
}

// Name returns the option name.
func (v View) Name() string {
	if data, ok := v.Params.(interface{ Name() string }); ok {
		return data.Name()
	}

	return ""
}

// Keys returns the path of option names.
func (v View) Keys() []string {
	keys := make([]string, 0, 1)
	if v.parent != nil {
		keys = append(keys, v.parent.Keys()...)
	}

	if name := v.Name(); name != "" {
		keys = append(keys, name)
	}

	return keys
}

// Type returns the option value type.
func (v View) Type() any {
	return param.Type(v.Params)
}

// FuncName returns the configured function name, or the option name.
func (v View) FuncName() string {
	data, ok := v.Params.Param(viewParamFunctName)
	name, valid := data.(string)

	if !ok || !valid {
		return v.Name()
	}

	return name
}

// StructName returns the configured struct name, or the path joined with the prefix and suffix.
func (v View) StructName() string {
	name, ok := param.String(v.Params, viewParamStructName)
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

// InterfaceType returns the type set by WithInterface.
func (v View) InterfaceType() reflect.Type {
	data, ok := v.Params.Param(viewParamInterface)
	if !ok {
		return nil
	}

	typ, _ := data.(reflect.Type)

	return typ
}

// ParentStruct returns the parent struct name.
func (v View) ParentStruct() string {
	if v.parent == nil {
		return ""
	}

	return v.parent.StructName()
}

// Description returns the option description.
func (v View) Description() string {
	return param.Description(v.Params)
}

// Default returns the default value.
func (v View) Default() any {
	data, ok := param.Default(v.Params)
	if !ok {
		return nil
	}

	return data
}

// HasDefault reports whether a default value is set.
func (v View) HasDefault() bool {
	return option.HasDefaut(v.Params)
}
