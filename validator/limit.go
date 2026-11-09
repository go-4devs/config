package validator

import (
	"fmt"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

type number interface {
	int | int64 | uint | uint64 | float64 | time.Duration
}

var (
	ErrMin = fmt.Errorf("%w: less than minimum", ErrInvalid)
	ErrMax = fmt.Errorf("%w: greater than maximum", ErrInvalid)
)

// Min returns a validator that rejects values below limit.
// The option type must be the same type as limit: int, int64, uint, uint64, float64, or time.Duration.
// The bound is inclusive.
func Min[T number](limit T) Validator {
	return bound(func(num T) bool { return num < limit }, func(num T) error {
		return MinError[T]{Limit: limit, Got: num}
	})
}

// Max returns a validator that rejects values above limit.
// The option type must be the same type as limit: int, int64, uint, uint64, float64, or time.Duration.
// The bound is inclusive.
func Max[T number](limit T) Validator {
	return bound(func(num T) bool { return num > limit }, func(num T) error {
		return MaxError[T]{Limit: limit, Got: num}
	})
}

// MinError is returned when the value is below Limit.
type MinError[T number] struct {
	Limit T
	Got   T
}

// Error returns the minimum and the received value.
func (e MinError[T]) Error() string {
	return fmt.Sprintf("%s: want %v, given %v", ErrMin, e.Limit, e.Got)
}

// Unwrap returns ErrMin.
func (e MinError[T]) Unwrap() error {
	return ErrMin
}

// MaxError is returned when the value is above Limit.
type MaxError[T number] struct {
	Limit T
	Got   T
}

// Error returns the maximum and the received value.
func (e MaxError[T]) Error() string {
	return fmt.Sprintf("%s: want %v, given %v", ErrMax, e.Limit, e.Got)
}

// Unwrap returns ErrMax.
func (e MaxError[T]) Unwrap() error {
	return ErrMax
}

func bound[T number](tooFar func(T) bool, fail func(T) error) Validator {
	return func(fn param.Params, in config.Value) error {
		num, err := numberOf[T](fn, in)
		if err != nil {
			return err
		}

		if tooFar(num) {
			return fail(num)
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
