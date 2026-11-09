package validator_test

import (
	"testing"

	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func TestRegex(t *testing.T) {
	t.Parallel()

	opt := option.String("name", "name")
	check := validator.Regex(`[a-z]+`)

	require.NoError(t, check(opt, value.New("ada")))
	require.NoError(t, check(opt, value.JString("ada")))
	require.ErrorIs(t, check(opt, value.New("ada1")), validator.ErrInvalid)
	require.ErrorIs(t, check(opt, value.New("")), validator.ErrInvalid)
	require.ErrorIs(t, validator.Regex(`(`)(opt, value.New("ada")), validator.ErrInvalid)
	require.ErrorIs(t, check(option.Int("port", "port"), value.New(1)), validator.ErrInvalid)
}
