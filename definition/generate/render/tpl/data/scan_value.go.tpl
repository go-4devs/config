{{ block "ScanValue" . -}}
    return {{.Value}}, {{.Value}}.Scan({{.ValName}}.Any())
{{- end }}