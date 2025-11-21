package param

func String(key any, fn Param) (string, bool) {
	val, ok := fn.Value(key)
	if !ok {
		return "", false
	}

	data, ok := val.(string)

	return data, ok
}

func Bool(key any, fn Param) (bool, bool) {
	val, ok := fn.Value(key)
	if !ok {
		return false, false
	}

	data, ok := val.(bool)

	return data, ok
}
