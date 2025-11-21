//go:build ignore
// +build ignore

package main

import (
	{{range .Imports}}
    {{- .Alias }}"{{ .Package }}"
    {{end}}
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	f, err := os.Create("{{.OutName}}")
	if err != nil {
		return err
	}

	defs:=make([]generate.Input,0)
{{ range .Configure }} 
	def{{.}} := definition.New()
	if err := {{$.Pkg}}.{{.}}(ctx, def{{.}}); err != nil {
		return err
	}
	defs = append(defs,generate.NewInput("{{.}}",def{{.}}))
{{ end }}

	opts := make([]generate.Option,0)
	{{ if .SkipContext }}opts = append(opts, generate.WithSkipContext){{ end }}
	opts = append(opts, 
		generate.WithPrefix("{{.Prefix}}"),
		generate.WithSuffix("{{.Suffix}}"),
		generate.WithFullPkg("{{.FullPkg}}"),
	)

	if gerr := generate.Run(ctx,generate.NewConfig(opts...),f, defs...);gerr != nil {
		return gerr
	}

	in, err := os.ReadFile(f.Name())
	if err != nil {
		return err
	}

	out, err := format.Source(in)
	if err != nil {
		return err
	}

	return os.WriteFile(f.Name(), out, 0644)
}
