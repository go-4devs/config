package render

import (
	"strings"

	"gitoa.ru/go-4devs/config/definition/generate/pkg"
	"gitoa.ru/go-4devs/config/key"
)

func Keys(keys []string, val string) string {
	if len(keys) == 0 {
		return ""
	}

	var out strings.Builder

	for idx, one := range keys {
		if key.IsWild(one) {
			out.WriteString(val)
			out.WriteString(".")
			out.WriteString(pkg.AliasName(one))
		} else {
			out.WriteString("\"")
			out.WriteString(one)
			out.WriteString("\"")
		}

		if len(keys)-1 != idx {
			out.WriteString(", ")
		}
	}

	return out.String()
}
