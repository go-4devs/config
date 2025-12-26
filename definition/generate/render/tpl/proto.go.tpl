type {{.StructName}} struct {
    {{.ParentName}}
    {{ .Name }} string
}

// {{.FuncName}} {{.Description}}.
func (i {{.ParentName}}) {{.FuncName}}(key string) {{.StructName}} {
	return {{.StructName}}{i,key}
}
