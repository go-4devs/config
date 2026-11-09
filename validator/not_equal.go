package validator

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

// ErrEqual is returned when the value equals the excluded one.
var ErrEqual = errors.New("equal to excluded value")

// NotEqual returns a validator that rejects want.
// The option type must be the same type as want.
func NotEqual[T any](want T) Validator {
	return func(fn param.Params, in config.Value) error {
		got, err := readAs[T](fn, in)
		if err != nil {
			return err
		}

		if reflect.DeepEqual(got, want) {
			return NewError(ErrEqual, got, want)
		}

		return nil
	}
}

func readAs[T any](fn param.Params, in config.Value) (T, error) {
	var zero T
	if _, ok := param.Type(fn).(T); !ok {
		return zero, ErrInvalid
	}

	if got, ok := in.Any().(T); ok {
		return got, nil
	}

	if parsed, ok, err := readBuiltin[T](in); ok {
		return parsed, err
	}

	if reflect.TypeFor[T]().Kind() == reflect.String {
		raw, err := in.ParseString()
		if err == nil {
			reflect.ValueOf(&zero).Elem().SetString(raw)

			return zero, nil
		}
	}

	var parsed T
	if err := in.Unmarshal(&parsed); err != nil {
		return zero, ErrInvalid
	}

	return parsed, nil
}

func readBuiltin[T any](in config.Value) (T, bool, error) {
	var zero T

	var (
		parsed any
		err    error
		match  bool
	)

	switch reflect.TypeFor[T]() {
	case reflect.TypeFor[string]():
		parsed, err = in.ParseString()
		match = true
	case reflect.TypeFor[bool]():
		parsed, err = in.ParseBool()
		match = true
	case reflect.TypeFor[int]():
		parsed, err = in.ParseInt()
		match = true
	case reflect.TypeFor[int64]():
		parsed, err = in.ParseInt64()
		match = true
	case reflect.TypeFor[uint]():
		parsed, err = in.ParseUint()
		match = true
	case reflect.TypeFor[uint64]():
		parsed, err = in.ParseUint64()
		match = true
	case reflect.TypeFor[float64]():
		parsed, err = in.ParseFloat64()
		match = true
	case reflect.TypeFor[time.Duration]():
		parsed, err = in.ParseDuration()
		match = true
	case reflect.TypeFor[time.Time]():
		parsed, err = in.ParseTime()
		match = true
	}

	if !match {
		return zero, false, nil
	}

	if err != nil {
		return zero, true, fmt.Errorf("value:%w", err)
	}

	got, ok := parsed.(T)
	if !ok {
		return zero, true, ErrInvalid
	}

	return got, true, nil
}
