package validator_test

import (
	"testing"

	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func TestNotEqual(t *testing.T) {
	t.Parallel()

	num := option.Int("n", "n")
	require.ErrorIs(t, validator.NotEqual(5)(num, value.New(5)), validator.ErrEqual)
	require.NoError(t, validator.NotEqual(5)(num, value.New(4)))
	require.ErrorIs(t, validator.NotEqual("root")(num, value.New(1)), validator.ErrInvalid)

	name := option.String("name", "name")
	require.ErrorIs(t, validator.NotEqual("root")(name, value.New("root")), validator.ErrEqual)
	require.NoError(t, validator.NotEqual("root")(name, value.JString("admin")))

	type pair struct{ A int }

	item := option.New("pair", "pair", pair{})
	require.ErrorIs(t, validator.NotEqual(pair{A: 1})(item, value.New(pair{A: 1})), validator.ErrEqual)
	require.NoError(t, validator.NotEqual(pair{A: 1})(item, value.New(pair{A: 2})))

	list := option.New("list", "list", []string(nil))
	require.ErrorIs(t, validator.NotEqual([]string{"a"})(list, value.New([]string{"a"})), validator.ErrEqual)
	require.NoError(t, validator.NotEqual([]string{"a"})(list, value.New([]string{"b"})))
}
