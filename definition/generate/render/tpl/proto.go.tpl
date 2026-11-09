type {{.StructName}} struct {
    {{.ParentStruct}}
    {{ .Name }} string
}

// {{.FuncName}} returns {{.Description}}.
func (i {{.ParentStruct}}) {{.FuncName}}(key string) {{.StructName}} {
	return {{.StructName}}{i,key}
}
