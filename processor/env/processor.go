package env

import (
	"context"
	"fmt"
	"os"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/param"
	"github.com/go-4devs/config/value"
)

var _ config.ProcessFunc = Env

// Env reads the environment variable named by the string value.
func Env(_ context.Context, in config.Value, _ ...param.Option) (config.Value, error) {
	key, err := in.ParseString()
	if err != nil {
		return in, fmt.Errorf("process[env]:%w", err)
	}

	res, ok := os.LookupEnv(key)
	if !ok {
		return nil, fmt.Errorf("%w", config.ErrNotFound)
	}

	return value.String(res), nil
}
