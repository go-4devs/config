package param_test

import (
	"testing"

	"github.com/go-4devs/config/param"
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

func TestChainLeavesInputOrder(t *testing.T) {
	t.Parallel()

	first := param.With(param.New(), "k", "first")
	second := param.With(param.New(), "k", "second")
	items := []param.Params{first, second}

	data, ok := param.String(param.Chain(items...), "k")
	if !ok || data != "second" {
		t.Fatalf("got %q ok %v", data, ok)
	}

	got, ok := param.String(items[0], "k")
	if !ok || got != "first" {
		t.Fatalf("input changed: %q ok %v", got, ok)
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
