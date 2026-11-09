type {{.StructName}} struct {
    {{.ParentStruct}}
}

// {{.FuncName}} returns {{.Description}}.
func (i {{.ParentStruct}}) {{.FuncName}}() {{.ReturnType}} {
	return {{.StructName}}{i}
}
