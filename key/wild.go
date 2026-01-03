package key

import "slices"

const minWildCount = 3

func IsWild(keys ...string) bool {
	return slices.ContainsFunc(keys, isWild)
}

func Wild(name string) string {
	return "{" + name + "}"
}

func isWild(name string) bool {
	if len(name) < minWildCount {
		return false
	}

	return name[0] == '{' && name[len(name)-1] == '}'
}
