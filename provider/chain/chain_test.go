package chain_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/provider/chain"
	"github.com/go-4devs/config/test/require"
	"github.com/go-4devs/config/value"
)

var errDown = errors.New("down")

type stub struct {
	err error
	val config.Value
}

func (s stub) Name() string { return "stub" }

func (s stub) Value(context.Context, ...string) (config.Value, error) {
	return s.val, s.err
}

func TestValueStopsOnRealError(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	got := value.New("ok")
	prov := chain.New(
		stub{err: errDown},
		stub{val: got},
	)

	_, err := prov.Value(ctx, "host")
	require.ErrorIs(t, err, errDown)

	prov = chain.New(
		stub{err: config.ErrNotFound},
		stub{val: got},
	)

	val, err := prov.Value(ctx, "host")
	require.NoError(t, err)
	require.Equal(t, "ok", val.String())

	prov = chain.New(
		stub{err: config.ErrInitFactory},
		stub{val: got},
	)

	_, err = prov.Value(ctx, "host")
	require.ErrorIs(t, err, config.ErrInitFactory)
}
