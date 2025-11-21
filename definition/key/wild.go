package key

const minWildCount = 3

func IsWild(name string) bool {
	if len(name) < minWildCount {
		return false
	}

	return name[0] == '{' && name[len(name)-1] == '}'
}

func Wild(name string) string {
	return "{" + name + "}"
}
