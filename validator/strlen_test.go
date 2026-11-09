package validator_test

import (
	"testing"

	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/test/require"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func TestStringLen(t *testing.T) {
	t.Parallel()

	opt := option.String("name", "name")
	atLeast := validator.MinLen(1)
	atMost := validator.MaxLen(2)

	require.NoError(t, atLeast(opt, value.New("a")))
	require.NoError(t, atMost(opt, value.New("йй")))
	require.ErrorIs(t, atLeast(opt, value.New("")), validator.ErrMinLen)
	require.ErrorIs(t, atMost(opt, value.New("abc")), validator.ErrMaxLen)
	require.ErrorIs(t, validator.MinLen(-1)(opt, value.New("a")), validator.ErrInvalid)
	require.ErrorIs(t, validator.MinLen(1)(option.Int("port", "port"), value.New(1)), validator.ErrInvalid)
}
