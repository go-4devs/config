package validator

import (
	"encoding"
	"fmt"
	"slices"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/param"
)

// ErrEnum is returned when the value is outside the allowed set.
var ErrEnum = fmt.Errorf("%w: enum", ErrInvalid)

// EnumError is returned when Got is outside Allowed.
type EnumError[T any] struct {
	Allowed []T
	Got     string
}

// Error returns the allowed values and the received text.
func (e EnumError[T]) Error() string {
	return fmt.Sprintf("%s: want %v, given %v", ErrEnum, e.Allowed, e.Got)
}

// Unwrap returns ErrEnum.
func (e EnumError[T]) Unwrap() error {
	return ErrEnum
}

// Enum returns a validator that accepts only allowed values.
// A string is compared as text. A fmt.Stringer is accepted when the option has the same type,
// UnmarshalText accepts the text, and the parsed value is one of allowed.
func Enum[T any](allowed ...T) Validator {
	var zero T
	switch any(zero).(type) {
	case string:
		return enumText(allowed)
	default:
		if _, ok := any(zero).(fmt.Stringer); ok {
			return enumStringer(allowed)
		}

		return invalidValue
	}
}

func enumText[T any](allowed []T) Validator {
	names := make([]string, len(allowed))
	for idx, item := range allowed {
		text, ok := any(item).(string)
		if !ok {
			return invalidValue
		}

		names[idx] = text
	}

	return func(_ param.Params, in config.Value) error {
		val := in.String()
		if slices.Contains(names, val) {
			return nil
		}

		return EnumError[T]{Allowed: allowed, Got: val}
	}
}

func enumStringer[T any](allowed []T) Validator {
	names := make([]string, len(allowed))
	for idx, item := range allowed {
		text, ok := any(item).(fmt.Stringer)
		if !ok {
			return invalidValue
		}

		names[idx] = text.String()
	}

	return func(fn param.Params, in config.Value) error {
		if _, ok := param.Type(fn).(T); !ok {
			return ErrInvalid
		}

		text, err := stringerText[T](in)
		if err != nil {
			return err
		}

		var parsed T

		parser, ok := any(&parsed).(encoding.TextUnmarshaler)
		if !ok {
			return ErrInvalid
		}

		if err = parser.UnmarshalText([]byte(text)); err != nil {
			return EnumError[T]{Allowed: allowed, Got: text}
		}

		got, ok := any(parsed).(fmt.Stringer)
		if !ok || !slices.Contains(names, got.String()) {
			return EnumError[T]{Allowed: allowed, Got: text}
		}

		return nil
	}
}

func stringerText[T any](in config.Value) (string, error) {
	if text, err := in.ParseString(); err == nil {
		return text, nil
	}

	typed, ok := in.Any().(T)
	if !ok {
		return "", ErrInvalid
	}

	text, ok := any(typed).(fmt.Stringer)
	if !ok {
		return "", ErrInvalid
	}

	return text.String(), nil
}

func invalidValue(param.Params, config.Value) error {
	return ErrInvalid
}
