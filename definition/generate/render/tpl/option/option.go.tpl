// read{{.FuncName}} {{.Description}}.
func (i {{.ParentStruct}}) read{{.FuncName}}(ctx context.Context) (v {{.Type}},e error) {
	val, err := i.Value(ctx, {{ .Keys "i" }})
	if err != nil {
        {{- if .HasDefault }}
        i.handle(ctx, err)
        
        {{ .Default "v" -}}
        {{ else }}
        return v, fmt.Errorf("read [%v]:%w",[]string{ {{- .Keys "i" -}} }, err)
        {{ end }}
	}

	{{ .Value "val" "v" }}
}

// Read{{.FuncName}} {{.Description}}.
func (i {{.ParentStruct}}) Read{{.FuncName}}({{if not .SkipContext}} ctx context.Context {{end}}) ({{.Type}}, error) {
    return i.read{{.FuncName}}({{if .SkipContext}}i.ctx{{else}}ctx{{end}})
}

// {{.FuncName}} {{.Description}}.
func (i {{.ParentStruct}}) {{.FuncName}}({{if not .SkipContext}} ctx context.Context {{end}}) {{.Type}} {
    val, err := i.read{{.FuncName}}({{ if .SkipContext }}i.ctx{{else}}ctx{{ end }})
	if err != nil {
		i.handle({{ if .SkipContext }}i.ctx{{else}}ctx{{ end }}, err)
	}

	return val
}
