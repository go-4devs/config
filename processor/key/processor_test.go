package key_test

import (
	"context"
	"encoding/json"
	"testing"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/processor/key"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/value"
)

func TestKey(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	res, rerr := key.Key(ctx, keyData(t), key.WithKey("key"))
	require.NoError(t, rerr)

	require.Equal(t, "value", res.String())
}

func TestKey_required(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	res, rerr := key.Key(ctx, keyData(t))
	require.ErrorIs(t, rerr, config.ErrRequired)

	require.Equal(t, res, nil)
}

func TestKey_notFound(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	res, rerr := key.Key(ctx, keyData(t), key.WithKey("wrong"))
	require.ErrorIs(t, rerr, config.ErrNotFound)

	require.Equal(t, res, nil)
}

func keyData(t *testing.T) config.Value {
	t.Helper()

	data := map[string]string{
		"key": "value",
	}

	jdata, err := json.Marshal(data)
	require.NoError(t, err)

	return value.JBytes(jdata)
}
