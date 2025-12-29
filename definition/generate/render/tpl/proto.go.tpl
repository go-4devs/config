type {{.StructName}} struct {
    {{.ParentStruct}}
    {{ .Name }} string
}

// {{.FuncName}} {{.Description}}.
func (i {{.ParentStruct}}) {{.FuncName}}(key string) {{.StructName}} {
	return {{.StructName}}{i,key}
}
