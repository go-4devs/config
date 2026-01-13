package handler

import (
	"context"
	"fmt"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/processor/env"
	"gitoa.ru/go-4devs/config/provider/memory"
	"gitoa.ru/go-4devs/config/value"
)

var (
	_ config.DumpProvider = Env(nil)
	_ config.BindProvider = Env(nil)
)

const (
	envPreffix = "%env("
	envSuffix  = ")%"
)

type EnvOption func(*EnvHandler)

func WithEnvProcessor(proc config.Processor) EnvOption {
	return func(eh *EnvHandler) {
		eh.Processor = proc
	}
}

func Env(parent config.Provider, opts ...EnvOption) EnvHandler {
	handler := EnvHandler{
		WrapProvider: memory.Wrap(parent),
		Processor:    config.ProcessFunc(env.Env),
	}

	for _, opt := range opts {
		opt(&handler)
	}

	return handler
}

type EnvHandler struct {
	memory.WrapProvider
	config.Processor
}

func (e EnvHandler) Value(ctx context.Context, key ...string) (config.Value, error) {
	val, err := e.Provider.Value(ctx, key...)
	if err != nil {
		return nil, fmt.Errorf("get %v:%w", e.Name(), err)
	}

	data, serr := val.ParseString()
	if serr != nil || !strings.HasPrefix(data, envPreffix) || !strings.HasSuffix(data, envSuffix) {
		return val, nil //nolint:nilerr
	}

	pval, perr := e.Process(ctx, value.String(data[len(envPreffix):len(data)-len(envSuffix)]))
	if perr != nil {
		return nil, fmt.Errorf("process[%v]:%w", e.Name(), perr)
	}

	return pval, nil
}
