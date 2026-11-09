package validator_test

import (
	"errors"
	"testing"

	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

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
