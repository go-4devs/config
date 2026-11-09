package validator_test

import (
	"errors"
	"testing"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func TestChain(t *testing.T) {
	t.Parallel()

	vars := option.String("test", "test")
	validValue := value.New("one")
	invalidValue := value.New([]string{"one"})

	valid := validator.Chain(
		validator.NotBlank,
		validator.Enum("one", "two"),
	)

	err := valid(vars, validValue)
	if err != nil {
		t.Errorf("expected valid value, got: %s", err)
	}

	ierr := valid(vars, invalidValue)
	if !errors.Is(ierr, validator.ErrNotBlank) {
		t.Errorf("expected not blank, got:%s", ierr)
	}
}

func TestOptional(t *testing.T) {
	t.Parallel()

	auth := option.String("auth", "auth")
	check := validator.Optional(validator.JSON)

	require.NoError(t, check(auth, value.New("")))
	require.NoError(t, check(auth, value.New(`{"a":1}`)))
	require.ErrorIs(t, check(auth, value.New("{")), validator.ErrInvalid)

	port := option.Int("port", "port")
	atLeast := validator.Optional(validator.Min(1))

	require.NoError(t, atLeast(port, value.New(0)))
	require.ErrorIs(t, atLeast(port, value.New(-1)), validator.ErrMin)

	names := option.String("names", "names", option.Slice)
	list := validator.Optional(validator.JSON)

	require.NoError(t, list(names, value.New([]string{})))
	require.ErrorIs(t, list(names, value.New("bad")), config.ErrInvalidValue)
}

func TestWhen(t *testing.T) {
	t.Parallel()

	auth := option.String("auth", "auth")
	check := validator.When(validator.NotBlank, validator.JSON)

	require.NoError(t, check(auth, value.New("")))
	require.NoError(t, check(auth, value.New(`{"a":1}`)))
	require.ErrorIs(t, check(auth, value.New("{")), validator.ErrInvalid)

	num := option.Int("n", "n")
	forbid := validator.When(validator.Chain(validator.Min(1), validator.Max(10)), validator.NotEqual(5))

	require.NoError(t, forbid(num, value.New(0)))
	require.NoError(t, forbid(num, value.New(11)))
	require.NoError(t, forbid(num, value.New(4)))
	require.ErrorIs(t, forbid(num, value.New(5)), validator.ErrEqual)
}
