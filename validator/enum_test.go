package validator_test

import (
	"errors"
	"testing"

	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/test/require"
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
	if !errors.Is(iErr, validator.ErrEnum) {
		t.Errorf("expected err:%s, got: %s", validator.ErrEnum, iErr)
	}
}

type color int

const (
	colorRed color = iota
	colorBlue
)

var errColor = errors.New("unknown color")

func (c color) String() string {
	switch c {
	case colorRed:
		return "red"
	case colorBlue:
		return "blue"
	default:
		return "unknown"
	}
}

func (c *color) UnmarshalText(text []byte) error {
	switch string(text) {
	case colorRed.String():
		*c = colorRed
	case colorBlue.String():
		*c = colorBlue
	default:
		return errColor
	}

	return nil
}

func TestEnumStringer(t *testing.T) {
	t.Parallel()

	opt := option.New("color", "color", color(0))
	check := validator.Enum(colorRed, colorBlue)

	require.NoError(t, check(opt, value.New("red")))
	require.NoError(t, check(opt, value.JString("blue")))
	require.NoError(t, check(opt, value.New(colorBlue)))
	require.ErrorIs(t, check(opt, value.New("green")), validator.ErrEnum)
	require.ErrorIs(t, check(opt, value.New(color(9))), validator.ErrEnum)
	require.ErrorIs(t, validator.Enum(colorRed)(opt, value.New("blue")), validator.ErrEnum)
	require.ErrorIs(t, check(option.String("name", "name"), value.New("red")), validator.ErrInvalid)
	require.ErrorIs(t, check(opt, value.New(1)), validator.ErrInvalid)
}
