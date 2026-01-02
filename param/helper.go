package param

func String(fn Params, key any) (string, bool) {
	val, ok := fn.Param(key)
	if !ok {
		return "", false
	}

	data, ok := val.(string)

	return data, ok
}

func Rune(fn Params, key any) (rune, bool) {
	val, ok := fn.Param(key)
	if !ok {
		return '0', false
	}

	data, dok := val.(rune)

	return data, dok
}

func Bool(key any, fn Params) (bool, bool) {
	val, ok := fn.Param(key)
	if !ok {
		return false, false
	}

	data, ok := val.(bool)

	return data, ok
}

func Uint64(key any, fn Params) (uint64, bool) {
	data, ok := fn.Param(key)
	if !ok {
		return 0, false
	}

	res, ok := data.(uint64)

	return res, ok
}
