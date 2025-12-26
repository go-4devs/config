// read{{.FuncName}} {{.Description}}.
func (i {{.ParentName}}) read{{.FuncName}}(ctx context.Context) (v {{.Type}},e error) {
	val, err := i.Value(ctx, {{ .Keys "i" }})
	if err != nil {
        {{- if .HasDefault }}
        i.log({{ if not .SkipContext }}context.Background(){{else}}ctx{{ end }}, "read [%v]: %v",[]string{ {{- .Keys "i" -}} }, err)
        
        {{ .Default "val" -}}
        {{ else }}
        return v, fmt.Errorf("read [%v]:%w",[]string{ {{- .Keys "i" -}} }, err)
        {{ end }}
	}

	{{ .Value "val" "v" }}
}

// Read{{.FuncName}} {{.Description}}.
func (i {{.ParentName}}) Read{{.FuncName}}({{if not .SkipContext}} ctx context.Context {{end}}) ({{.Type}}, error) {
    return i.read{{.FuncName}}({{if .SkipContext}}context.Background(){{else}}ctx{{end}})
}

// {{.FuncName}} {{.Description}}.
func (i {{.ParentName}}) {{.FuncName}}({{if not .SkipContext}} ctx context.Context {{end}}) {{.Type}} {
    val, err := i.read{{.FuncName}}({{ if .SkipContext }}context.Background(){{else}}ctx{{ end }})
	if err != nil {
		i.log({{ if .SkipContext }}context.Background(){{else}}ctx{{ end }}, "get [%v]: %v",[]string{ {{- .Keys "i" -}} }, err)
	}

	return val
}
