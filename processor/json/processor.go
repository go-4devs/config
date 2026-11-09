package json

import (
	"context"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/param"
	"github.com/go-4devs/config/value"
)

//nolint:revive
func Json(_ context.Context, in config.Value, _ ...param.Option) (config.Value, error) {
	data, err := in.ParseString()
	if err != nil {
		return in, nil //nolint:nilerr
	}

	return value.JString(data), nil
}
