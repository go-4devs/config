func With{{.StructName}}Handle(fn func(context.Context, error)) func(*{{.StructName}}) {
	return func(ci *{{.StructName}}) {
		ci.handle = fn
	}
}

func New{{.StructName}}({{if or .SkipContext .ClildSkipContext }} ctx context.Context,{{end}}prov config.Provider, opts ...func(*{{.StructName}})) {{.StructName}} {
	i := {{.StructName}}{
		Provider: prov,
		handle: func(_ context.Context, err error) {
			fmt.Printf("{{.StructName}}:%v",err)
		},
		{{if or .SkipContext .ClildSkipContext }} ctx: ctx, {{end}}
	}

	for _, opt := range opts {
		opt(&i)
	}

	return i
}

type {{.StructName}} struct {
	config.Provider
	handle   func(context.Context, error)
	{{if or .SkipContext .ClildSkipContext}}ctx context.Context {{end}}
}
