package generate

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"io"
	"strings"
	"text/template"
	"unicode"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/generate/pkg"
	"gitoa.ru/go-4devs/config/definition/generate/render"
	"gitoa.ru/go-4devs/config/definition/generate/view"
)

//go:embed tpl/*
var tpls embed.FS

var initTpl = template.Must(template.New("tpls").ParseFS(tpls, "tpl/*.tpl")).Lookup("init.go.tpl")

func Run(_ context.Context, fullPkg string, w io.Writer, defs ...config.Options) error {
	data := Data{
		Packages: pkg.NewImports(fullPkg).
			Adds("fmt", "context", "gitoa.ru/go-4devs/config"),
	}

	var buff bytes.Buffer

	for _, in := range defs {
		vi := view.NewViews(in)

		if err := render.Render(&buff, vi, data); err != nil {
			return fmt.Errorf("render:%w", err)
		}
	}

	if err := initTpl.Execute(w, data); err != nil {
		return fmt.Errorf("render base:%w", err)
	}

	if _, err := io.Copy(w, &buff); err != nil {
		return fmt.Errorf("copy:%w", err)
	}

	return nil
}

type Data struct {
	*pkg.Packages
}

func (f Data) StructName(name string) string {
	return FuncName(name)
}

func (f Data) FuncName(in string) string {
	return FuncName(in)
}

func FuncName(name string) string {
	data := strings.Builder{}
	toUp := true

	for _, char := range name {
		isLeter := unicode.IsLetter(char)
		isAllowed := isLeter || unicode.IsDigit(char)

		switch {
		case isAllowed && !toUp:
			data.WriteRune(char)
		case !isAllowed:
			toUp = true
		case toUp:
			data.WriteString(strings.ToUpper(string(char)))

			toUp = false
		}
	}

	return data.String()
}
