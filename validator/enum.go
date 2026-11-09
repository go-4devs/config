package validator

import (
	"encoding"
	"fmt"
	"slices"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

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

		return NewError(ErrInvalid, val, names)
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
			return NewError(ErrInvalid, text, err)
		}

		got, ok := any(parsed).(fmt.Stringer)
		if !ok || !slices.Contains(names, got.String()) {
			return NewError(ErrInvalid, text, names)
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
