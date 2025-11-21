package render

import (
	"embed"
	"fmt"
	"io"
	"strings"
	"text/template"

	"gitoa.ru/go-4devs/config/definition/generate/view"
)

//go:embed tpl/*
var tplFS embed.FS

var (
	tpls = template.Must(
		template.New("tpls").
			Funcs(template.FuncMap{
				"trim": strings.Trim,
			}).
			ParseFS(tplFS, "tpl/*.go.tpl"),
	)
	defTpl   = tpls.Lookup("definition.go.tpl")
	groupTpl = tpls.Lookup("group.go.tpl")
	protoTpl = tpls.Lookup("proto.go.tpl")
	optTpl   = template.Must(
		template.New("opt").ParseFS(tplFS, "tpl/option/option.go.tpl"),
	).Lookup("option.go.tpl")
	parceTpls = template.Must(template.New("data").ParseFS(tplFS, "tpl/data/*.go.tpl"))
)

func Template(tpl *template.Template) Execute {
	return func(w io.Writer, v view.View, rnd Rendering) error {
		if err := tpl.Execute(w, NewViewData(rnd, v)); err != nil {
			return fmt.Errorf("template[%v]:%w", tpl.Name(), err)
		}

		return nil
	}
}
