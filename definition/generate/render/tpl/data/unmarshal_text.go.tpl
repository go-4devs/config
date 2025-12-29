{{ block "UnmarshalText" . -}}
    pval, perr := {{.ValName}}.ParseString()
    if perr != nil {
        return {{.Value}}, fmt.Errorf("parse [%v]:%w", []string{ {{- .Keys "i" -}} }, perr)
    }

    return {{.Value}}, {{.Value}}.UnmarshalText([]byte(pval))
{{- end -}}
