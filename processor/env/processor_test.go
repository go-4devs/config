package env_test

import (
	"context"
	"testing"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/processor/env"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/value"
)

type provider struct {
	value config.Value
}

func (p provider) Value(_ context.Context, _ ...string) (config.Value, error) {
	return p.value, nil
}

func (p provider) Name() string {
	return "test"
}

func TestEnvValue(t *testing.T) {
	const except = "env value"
	t.Setenv("APP_ENV", except)

	ctx := context.Background()
	process := env.New(provider{value: value.String("%env(APP_ENV)%")})
	data, err := process.Value(ctx, "any")
	require.NoError(t, err)
	require.Equal(t, except, data.String())
}
