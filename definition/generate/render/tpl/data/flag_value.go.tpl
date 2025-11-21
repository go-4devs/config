{{ block "FlagValue" . -}}
    pval, perr := {{.ValName}}.ParseString()
    if perr != nil {
        return {{.Value}}, fmt.Errorf("read [%v]:%w",[]string{ {{- .Keys "i" -}} }, perr)
    }

    return {{.Value}}, {{.Value}}.Set(pval)
{{- end }}