type {{.StructName}} struct {
    {{.ParentStruct}}
}

// {{.FuncName}} {{.Description}}.
func (i {{.ParentStruct}}) {{.FuncName}}() {{.ReturnType}} {
	return {{.StructName}}{i}
}
