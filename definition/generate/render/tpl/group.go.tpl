type {{.StructName}} struct {
    {{.ParentStruct}}
}

// {{.FuncName}} {{.Description}}.
func (i {{.ParentStruct}}) {{.FuncName}}() {{.StructName}} {
	return {{.StructName}}{i}
}
