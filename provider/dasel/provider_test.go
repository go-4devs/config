package dasel_test

import (
	"context"
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/provider/chain"
	"github.com/go-4devs/config/provider/dasel"
	"github.com/go-4devs/config/test/require"
	"github.com/go-4devs/config/value"
	"github.com/tomwright/dasel/v3/parsing/json"
)

type fallback struct{}

func (fallback) Name() string { return "fallback" }

func (fallback) Value(context.Context, ...string) (config.Value, error) {
	return value.New("from-next"), nil
}

func TestMissingKeyFallsThrough(t *testing.T) {
	t.Parallel()

	prov, err := dasel.New([]byte(`{"app":{"name":"demo"}}`), json.JSON)
	require.NoError(t, err)

	ctx := context.Background()

	_, err = prov.Value(ctx, "app", "missing")
	require.ErrorIs(t, err, config.ErrValueNotFound)

	got, err := chain.New(prov, fallback{}).Value(ctx, "app", "missing")
	require.NoError(t, err)
	require.Equal(t, "from-next", got.String())
}
