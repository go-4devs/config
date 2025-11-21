type {{.StructName}} struct {
    {{.ParentName}}
}

// {{.FuncName}} {{.Description}}.
func (i {{.ParentName}}) {{.FuncName}}() {{.StructName}} {
	return {{.StructName}}{i}
}
