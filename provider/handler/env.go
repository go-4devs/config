package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/processor/env"
	"github.com/go-4devs/config/provider/memory"
	"github.com/go-4devs/config/value"
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

// WithEnvProcessor sets the function that reads an environment placeholder.
func WithEnvProcessor(proc config.Processor) EnvOption {
	return func(eh *EnvHandler) {
		eh.Processor = proc
	}
}

// Env returns a provider that resolves %env(NAME)% strings from the parent value.
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

// Value returns the parent value. A %env(NAME)% string is replaced with that environment variable.
func (e EnvHandler) Value(ctx context.Context, key ...string) (config.Value, error) {
	val, err := e.Provider.Value(ctx, key...)
	if err != nil {
		return nil, fmt.Errorf("get %v:%w", e.Name(), err)
	}

	data, ok := envRef(val)
	if !ok {
		return val, nil
	}

	pval, perr := e.Process(ctx, value.String(data[len(envPreffix):len(data)-len(envSuffix)]))
	if perr != nil {
		return nil, fmt.Errorf("process[%v]:%w", e.Name(), perr)
	}

	return pval, nil
}

func envRef(val config.Value) (string, bool) {
	data, err := val.ParseString()
	if err != nil || !strings.HasPrefix(data, envPreffix) || !strings.HasSuffix(data, envSuffix) {
		return "", false
	}

	return data, true
}
