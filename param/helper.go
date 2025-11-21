package param

import "gitoa.ru/go-4devs/config"

func String(key any, fn Params) (string, bool) {
	val, ok := fn.Param(key)
	if !ok {
		return "", false
	}

	data, ok := val.(string)

	return data, ok
}

func Bool(key any, fn Params) (bool, bool) {
	val, ok := fn.Param(key)
	if !ok {
		return false, false
	}

	data, ok := val.(bool)

	return data, ok
}

func Value(key any, fn Params) (config.Value, bool) {
	data, ok := fn.Param(key)
	if !ok {
		return nil, false
	}

	res, ok := data.(config.Value)

	return res, ok
}

func Uint64(key any, fn Params) (uint64, bool) {
	data, ok := fn.Param(key)
	if !ok {
		return 0, false
	}

	res, ok := data.(uint64)

	return res, ok
}
