package env

import (
	"context"
	"fmt"
	"os"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/value"
)

const (
	prefix = "%env("
	suffix = ")%"
)

func New(prov config.Provider) Env {
	env := Env{
		Provider: prov,
		name:     "",
		prefix:   prefix,
		suffix:   suffix,
	}

	return env
}

type Env struct {
	config.Provider

	name   string
	prefix string
	suffix string
}

func (e Env) Value(ctx context.Context, key ...string) (config.Value, error) {
	val, err := e.Provider.Value(ctx, key...)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	pval, perr := e.Process(ctx, val)
	if perr != nil {
		return nil, fmt.Errorf("%w", perr)
	}

	return pval, nil
}

func (e Env) Name() string {
	if e.name != "" {
		return e.name
	}

	return e.Provider.Name()
}

func (e Env) Process(_ context.Context, in config.Value) (config.Value, error) {
	data, err := in.ParseString()
	if err != nil || !strings.HasPrefix(data, e.prefix) || !strings.HasSuffix(data, e.suffix) {
		return in, nil //nolint:nilerr
	}

	key := data[len(e.prefix) : len(data)-len(e.suffix)]

	res, ok := os.LookupEnv(key)
	if !ok {
		return nil, fmt.Errorf("%v:%w", e.Name(), config.ErrNotFound)
	}

	return value.String(res), nil
}
