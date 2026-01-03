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
	params{{.}} := param.New(
	{{- if $.SkipContext }}view.WithSkipContext,{{ end }}
	view.WithStructName("{{$.Prefix}}_{{.}}_{{$.Suffix}}"),
	view.WithStructPrefix("{{$.Prefix}}"),
  view.WithStructSuffix("{{$.Suffix}}"),
	)

	def{{.}} := definition.New().With(params{{.}})
	if err := {{$.Pkg}}.{{.}}(ctx, def{{.}}); err != nil {
		return err
	}
	defs = append(defs,def{{.}})
{{ end }}

	if gerr := generate.Run(ctx,"{{.FullPkg}}",f, defs...);gerr != nil {
		return gerr
	}

  return nil
}
