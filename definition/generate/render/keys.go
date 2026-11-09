package render

import (
	"strings"

	"github.com/go-4devs/config/definition/generate/pkg"
	"github.com/go-4devs/config/key"
)

// Keys returns the generated argument list for a value path.
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
