package handler_test

import (
	"context"
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/provider/handler"
	"github.com/go-4devs/config/test/require"
	"github.com/go-4devs/config/value"
)

type provider struct {
	value config.Value
}

func (p provider) Value(context.Context, ...string) (config.Value, error) {
	return p.value, nil
}

func (p provider) Name() string {
	return "test"
}

func TestEnvValue(t *testing.T) {
	const except = "env value"
	t.Setenv("APP_ENV", except)

	ctx := context.Background()
	process := handler.Env(provider{value: value.String("%env(APP_ENV)%")})
	data, err := process.Value(ctx, "any")
	require.NoError(t, err)
	require.Equal(t, except, data.String())
}
