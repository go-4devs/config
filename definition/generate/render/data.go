package render

import (
	"gitoa.ru/go-4devs/config/definition/generate/pkg"
	"gitoa.ru/go-4devs/config/definition/generate/view"
)

func NewViewData(render Rendering, view view.View) ViewData {
	return ViewData{
		Rendering: render,
		View:      view,
	}
}

type ViewData struct {
	Rendering
	view.View
}

func (d ViewData) StructName() string {
	return d.Rendering.StructName(d.View.StructName())
}

func (d ViewData) FuncName() string {
	return d.Rendering.FuncName(d.View.FuncName())
}

func (d ViewData) ParentStruct() string {
	name := d.View.ParentStruct()
	if name == "" {
		name = d.Name()
	}

	return d.Rendering.StructName(name)
}

func (d ViewData) Name() string {
	return pkg.AliasName(d.View.Name())
}

func (d ViewData) Type() string {
	return Type(d)
}

func (d ViewData) Keys(parent string) string {
	return Keys(d.View.Keys(), parent)
}

func (d ViewData) Value(name, val string) string {
	return Value(name, val, d)
}

func (d ViewData) Default(name string) string {
	return Data(d.View.Default(), name, d)
}

type Rendering interface {
	StructName(name string) string
	FuncName(name string) string
	AddType(pkg string) (string, error)
}
