package validator_test

import (
	"errors"
	"testing"

	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func TestSentinelsWrapInvalid(t *testing.T) {
	t.Parallel()

	sentinels := []error{
		validator.ErrNotBlank,
		validator.ErrMin,
		validator.ErrMax,
		validator.ErrEqual,
		validator.ErrMinLen,
		validator.ErrMaxLen,
		validator.ErrEmail,
		validator.ErrJSON,
		validator.ErrURL,
		validator.ErrURI,
		validator.ErrRegex,
		validator.ErrEnum,
	}

	for _, err := range sentinels {
		require.ErrorIs(t, err, validator.ErrInvalid)
	}
}

func TestErrorTypes(t *testing.T) {
	t.Parallel()

	port := option.Int("port", "port")
	err := validator.Max(10)(port, value.New(11))

	var bound validator.MaxError[int]
	if !errors.As(err, &bound) {
		t.Fatalf("type: %T", err)
	}

	if bound.Limit != 10 || bound.Got != 11 {
		t.Fatalf("limit %v, got %v", bound.Limit, bound.Got)
	}

	email := option.String("email", "email")

	err = validator.Email(email, value.New("user"))

	var addr validator.EmailError
	if !errors.As(err, &addr) {
		t.Fatalf("type: %T", err)
	}

	if addr.Got != "user" {
		t.Fatalf("got %v", addr.Got)
	}
}
