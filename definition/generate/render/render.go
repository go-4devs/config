package render

import (
	"fmt"
	"io"
	"reflect"

	"github.com/go-4devs/config/definition"
	"github.com/go-4devs/config/definition/generate/view"
	"github.com/go-4devs/config/definition/group"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/definition/proto"
)

type Execute func(w io.Writer, vi view.View, rnd Rendering) error

var randders = map[reflect.Type]Execute{
	reflect.TypeFor[*definition.Definition](): Template(defTpl),
	reflect.TypeFor[*group.Group]():           Template(groupTpl),
	reflect.TypeFor[option.Option]():          Template(optTpl),
	reflect.TypeFor[proto.Proto]():            Template(protoTpl),
}

// Renders returns the templates by option type.
func Renders() map[reflect.Type]Execute {
	return randders
}

// Add registers a template for an option type.
func Add(rt reflect.Type, fn Execute) {
	randders[rt] = fn
}

// Render writes the view and its children.
func Render(w io.Writer, view view.View, data Rendering) error {
	rnd, ok := randders[view.Kind()]
	if !ok {
		return fmt.Errorf("%w:%v", ErrNotFound, view.Kind())
	}

	if err := rnd(w, view, data); err != nil {
		return fmt.Errorf("render:%v, err:%w", view.Kind(), err)
	}

	for _, ch := range view.Views() {
		if err := Render(w, ch, data); err != nil {
			return fmt.Errorf("render[%v]:%w", ch.Name(), err)
		}
	}

	return nil
}
