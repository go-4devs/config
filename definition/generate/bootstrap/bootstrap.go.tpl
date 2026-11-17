//go:build ignore
// +build ignore

package main

import (
	{{range .Imports}}
    {{- .Alias }}"{{ .Package }}"
    {{end}}
)

func main() {
	if err := run(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(f io.Writer) error {
	ctx := context.Background()

	defs:=make([]config.Options,0)
{{ range .Configure }} 
	params{{.Name}} := param.New(
	{{- if $.SkipContext }}view.WithSkipContext,{{ end }}
	view.WithStructName("{{$.Prefix}}_{{if .StructName}}{{.StructName}}{{else}}{{.Name}}{{end}}_{{$.Suffix}}"),
	view.WithStructPrefix("{{$.Prefix}}"),
  view.WithStructSuffix("{{$.Suffix}}"),
	)

	def{{.Name}} := definition.New().With(params{{.Name}})
	if err := {{$.Pkg}}.{{.Name}}(ctx, def{{.Name}}); err != nil {
		return err
	}
	defs = append(defs,def{{.Name}})
{{ end }}

	if gerr := generate.Run(ctx,"{{.FullPkg}}",f, defs...);gerr != nil {
		return gerr
	}

  return nil
}
