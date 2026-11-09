package validator_test

import (
	"testing"

	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func TestEmail(t *testing.T) {
	t.Parallel()

	opt := option.String("email", "email")

	require.NoError(t, validator.Email(opt, value.New("user@example.com")))
	require.NoError(t, validator.Email(opt, value.JString("user@example.com")))
	require.ErrorIs(t, validator.Email(opt, value.New("User <user@example.com>")), validator.ErrInvalid)
	require.ErrorIs(t, validator.Email(opt, value.New("user")), validator.ErrInvalid)
	require.ErrorIs(t, validator.Email(opt, value.New("")), validator.ErrInvalid)
	require.ErrorIs(t, validator.Email(option.Int("port", "port"), value.New(1)), validator.ErrInvalid)
}
