package validator_test

import (
	"testing"

	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func TestJSON(t *testing.T) {
	t.Parallel()

	opt := option.String("auth", "auth")

	require.NoError(t, validator.JSON(opt, value.New(`{"a":1}`)))
	require.NoError(t, validator.JSON(opt, value.JString(`[1]`)))
	require.NoError(t, validator.JSON(opt, value.New("null")))
	require.ErrorIs(t, validator.JSON(opt, value.New("{")), validator.ErrJSON)
	require.ErrorIs(t, validator.JSON(opt, value.New("")), validator.ErrJSON)
	require.ErrorIs(t, validator.JSON(option.Int("port", "port"), value.New(1)), validator.ErrInvalid)
}
