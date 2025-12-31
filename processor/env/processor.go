package env

import (
	"context"
	"fmt"
	"os"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/value"
)

var _ config.ProcessFunc = Env

func Env(_ context.Context, in config.Value) (config.Value, error) {
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
