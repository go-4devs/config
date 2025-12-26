package validator_test

import (
	"errors"
	"testing"

	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func TestEnum(t *testing.T) {
	t.Parallel()

	vars := option.String("test", "test")
	validValue := value.New("valid")
	invalidValue := value.New("invalid")

	enum := validator.Enum("valid", "other", "three")

	err := enum(vars, validValue)
	if err != nil {
		t.Errorf("expected valid value got err:%s", err)
	}

	iErr := enum(vars, invalidValue)
	if !errors.Is(iErr, validator.ErrInvalid) {
		t.Errorf("expected err:%s, got: %s", validator.ErrInvalid, iErr)
	}
}
