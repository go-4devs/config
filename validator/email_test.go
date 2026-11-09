package validator_test

import (
	"testing"

	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/test/require"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func TestEmail(t *testing.T) {
	t.Parallel()

	opt := option.String("email", "email")

	require.NoError(t, validator.Email(opt, value.New("user@example.com")))
	require.NoError(t, validator.Email(opt, value.JString("user@example.com")))
	require.ErrorIs(t, validator.Email(opt, value.New("User <user@example.com>")), validator.ErrEmail)
	require.ErrorIs(t, validator.Email(opt, value.New("user")), validator.ErrEmail)
	require.ErrorIs(t, validator.Email(opt, value.New("")), validator.ErrEmail)
	require.ErrorIs(t, validator.Email(option.Int("port", "port"), value.New(1)), validator.ErrInvalid)
}
