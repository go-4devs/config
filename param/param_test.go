package param_test

import (
	"testing"

	"gitoa.ru/go-4devs/config/param"
)

func TestChainReplace(t *testing.T) {
	t.Parallel()

	const (
		replaceParam = "param1"
		replaceValue = "replace"
	)

	params1 := param.With(param.New(), replaceParam, "param1")
	params2 := param.With(param.New(), replaceParam, replaceValue)

	data, ok := param.String(param.Chain(params1, params2), replaceParam)
	if !ok {
		t.Errorf("param %v: not found", replaceParam)
	}

	if data != replaceValue {
		t.Errorf("got:%v, expect:%v", data, replaceValue)
	}
}

func TestChainExtend(t *testing.T) {
	t.Parallel()

	const (
		extendParam = "param1"
		extendValue = "replace"
	)

	params1 := param.With(param.New(), extendParam, extendValue)
	params2 := param.With(param.New(), "new_value", "param2")

	data1, ok := param.String(param.Chain(params1, params2), extendParam)
	if !ok {
		t.Errorf("param %v: not found", extendParam)
	}

	if data1 != extendValue {
		t.Errorf("got:%v, expect:%v", data1, extendParam)
	}
}
