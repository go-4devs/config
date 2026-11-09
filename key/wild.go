package key

import "slices"

const minWildCount = 3

// IsWild reports whether any key is wrapped in braces.
func IsWild(keys ...string) bool {
	return slices.ContainsFunc(keys, isWild)
}

// Wild returns name wrapped as {name}.
func Wild(name string) string {
	return "{" + name + "}"
}

func isWild(name string) bool {
	if len(name) < minWildCount {
		return false
	}

	return name[0] == '{' && name[len(name)-1] == '}'
}
