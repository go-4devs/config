package validator

import (
	"errors"
	"fmt"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

type number interface {
	int | int64 | uint | uint64 | float64 | time.Duration
}

var (
	ErrMin = errors.New("less than minimum")
	ErrMax = errors.New("greater than maximum")
)

// Min returns a validator that rejects values below limit.
// The option type must be the same type as limit: int, int64, uint, uint64, float64, or time.Duration.
// The bound is inclusive.
func Min[T number](limit T) Validator {
	return bound(limit, func(num T) bool { return num < limit }, ErrMin)
}

// Max returns a validator that rejects values above limit.
// The option type must be the same type as limit: int, int64, uint, uint64, float64, or time.Duration.
// The bound is inclusive.
func Max[T number](limit T) Validator {
	return bound(limit, func(num T) bool { return num > limit }, ErrMax)
}

func bound[T number](limit T, tooFar func(T) bool, sentinel error) Validator {
	return func(fn param.Params, in config.Value) error {
		num, err := numberOf[T](fn, in)
		if err != nil {
			return err
		}

		if tooFar(num) {
			return NewError(sentinel, num, limit)
		}

		return nil
	}
}

func numberOf[T number](fn param.Params, in config.Value) (T, error) {
	var zero T
	if _, ok := param.Type(fn).(T); !ok {
		return zero, ErrInvalid
	}

	var (
		num any
		err error
	)

	switch any(zero).(type) {
	case int:
		num, err = in.ParseInt()
	case int64:
		num, err = in.ParseInt64()
	case uint:
		num, err = in.ParseUint()
	case uint64:
		num, err = in.ParseUint64()
	case float64:
		num, err = in.ParseFloat64()
	case time.Duration:
		num, err = in.ParseDuration()
	default:
		return zero, ErrInvalid
	}

	if err != nil {
		return zero, fmt.Errorf("number:%w", err)
	}

	typed, ok := num.(T)
	if !ok {
		return zero, ErrInvalid
	}

	return typed, nil
}
