package validator_test

import (
	"testing"

	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func TestMinMax(t *testing.T) {
	t.Parallel()

	opt := option.Int("port", "port")
	atLeast := validator.Min(1)
	atMost := validator.Max(10)

	require.NoError(t, atLeast(opt, value.New(1)))
	require.NoError(t, atMost(opt, value.New(10)))
	require.ErrorIs(t, atLeast(opt, value.New(0)), validator.ErrMin)
	require.ErrorIs(t, atMost(opt, value.New(11)), validator.ErrMax)

	flt := option.Float64("load", "load")
	require.NoError(t, validator.Max(1.5)(flt, value.New(1.5)))
	require.ErrorIs(t, validator.Min(0.5)(flt, value.New(0.1)), validator.ErrMin)
	require.ErrorIs(t, validator.Min(1)(option.String("name", "name"), value.New("a")), validator.ErrInvalid)
	require.ErrorIs(t, validator.Min(int64(1))(opt, value.New(1)), validator.ErrInvalid)

	wide := option.Uint64("id", "id")
	require.NoError(t, validator.Max(uint64(1<<54))(wide, value.New(uint64(1<<54))))
	require.ErrorIs(t, validator.Max(uint64(1<<54))(wide, value.New(uint64(1<<54)+1)), validator.ErrMax)
}
