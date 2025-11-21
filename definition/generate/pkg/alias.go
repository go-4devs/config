package pkg

import (
	"strings"
	"unicode"
)

func AliasName(name string) string {
	data := strings.Builder{}
	toUp := false

	for _, char := range name {
		isLeter := unicode.IsLetter(char)
		isAllowed := isLeter || unicode.IsDigit(char)

		switch {
		case isAllowed && !toUp:
			data.WriteRune(char)
		case !isAllowed && data.Len() > 0:
			toUp = true
		case toUp:
			data.WriteString(strings.ToUpper(string(char)))

			toUp = false
		}
	}

	return data.String()
}
