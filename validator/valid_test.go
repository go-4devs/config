package validator_test

import (
	"errors"
	"testing"

	"gitoa.ru/go-4devs/config/definition/option"
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
