package json //nolint:revive

import (
	"context"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/value"
)

//nolint:revive
func Json(_ context.Context, in config.Value, _ ...param.Option) (config.Value, error) {
	data, err := in.ParseString()
	if err != nil {
		return in, nil //nolint:nilerr
	}

	return value.JString(data), nil
}
