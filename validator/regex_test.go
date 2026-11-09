package validator_test

import (
	"testing"

	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/test/require"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func TestRegex(t *testing.T) {
	t.Parallel()

	opt := option.String("name", "name")
	check := validator.Regex(`[a-z]+`)

	require.NoError(t, check(opt, value.New("ada")))
	require.NoError(t, check(opt, value.JString("ada")))
	require.ErrorIs(t, check(opt, value.New("ada1")), validator.ErrRegex)
	require.ErrorIs(t, check(opt, value.New("")), validator.ErrRegex)
	require.ErrorIs(t, validator.Regex(`(`)(opt, value.New("ada")), validator.ErrRegex)
	require.ErrorIs(t, check(option.Int("port", "port"), value.New(1)), validator.ErrInvalid)
}
