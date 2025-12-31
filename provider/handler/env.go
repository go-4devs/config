package handler

import (
	"context"
	"fmt"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/processor/env"
	"gitoa.ru/go-4devs/config/value"
)

const (
	envPreffix = "%env("
	envSuffix  = ")%"
)

func Env(parent config.Provider) EnvHandler {
	handler := EnvHandler{
		Provider:  parent,
		Processor: config.ProcessFunc(env.Env),
		name:      "env:" + parent.Name(),
	}

	return handler
}

type EnvHandler struct {
	config.Provider
	config.Processor

	name string
}

func (e EnvHandler) Name() string {
	return e.name
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
