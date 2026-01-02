package json_test

import (
	"context"
	"testing"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/processor/json"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/value"
)

func TestJson(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	val, err := json.Json(ctx, value.String("42"))
	require.NoError(t, err)

	var res int
	require.NoError(t, val.Unmarshal(&res))
	require.Equal(t, 42, res)

	sval, serr := json.Json(ctx, value.String("\"test data\""))
	require.NoError(t, serr)

	var sres string
	require.NoError(t, sval.Unmarshal(&sres))
	require.Equal(t, "test data", sres)

	slval, slerr := json.Json(ctx, value.String("[\"test\",\"test2 data\",\"test3\"]"))
	require.NoError(t, slerr)

	var slres []string
	require.NoError(t, slval.Unmarshal(&slres))
	require.Equal(t, []string{"test", "test2 data", "test3"}, slres)
}

func TestJson_invalidValue(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	val, err := json.Json(ctx, value.New("42"))
	require.NoError(t, err)

	var data string
	require.ErrorIs(t, val.Unmarshal(&data), config.ErrInvalidValue)
}
