package validator

import (
	"fmt"
	"unicode/utf8"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/param"
)

var (
	ErrMinLen = fmt.Errorf("%w: shorter than minimum", ErrInvalid)
	ErrMaxLen = fmt.Errorf("%w: longer than maximum", ErrInvalid)
)

// MinLen returns a validator that rejects strings shorter than limit.
// The option type must be string. The bound is inclusive and counts Unicode code points.
func MinLen(limit int) Validator {
	return lenBound(limit, func(count int) bool { return count < limit }, func(count int) error {
		return MinLenError{Limit: limit, Got: count}
	})
}

// MaxLen returns a validator that rejects strings longer than limit.
// The option type must be string. The bound is inclusive and counts Unicode code points.
func MaxLen(limit int) Validator {
	return lenBound(limit, func(count int) bool { return count > limit }, func(count int) error {
		return MaxLenError{Limit: limit, Got: count}
	})
}

// MinLenError is returned when the text is shorter than Limit characters.
type MinLenError struct {
	Limit int
	Got   int
}

// Error returns the minimum length and the received length.
func (e MinLenError) Error() string {
	return fmt.Sprintf("%s: want %v, given %v", ErrMinLen, e.Limit, e.Got)
}

// Unwrap returns ErrMinLen.
func (e MinLenError) Unwrap() error {
	return ErrMinLen
}

// MaxLenError is returned when the text is longer than Limit characters.
type MaxLenError struct {
	Limit int
	Got   int
}

// Error returns the maximum length and the received length.
func (e MaxLenError) Error() string {
	return fmt.Sprintf("%s: want %v, given %v", ErrMaxLen, e.Limit, e.Got)
}

// Unwrap returns ErrMaxLen.
func (e MaxLenError) Unwrap() error {
	return ErrMaxLen
}

func lenBound(limit int, tooFar func(int) bool, fail func(int) error) Validator {
	return func(fn param.Params, in config.Value) error {
		if limit < 0 {
			return ErrInvalid
		}

		count, err := strLen(fn, in)
		if err != nil {
			return err
		}

		if tooFar(count) {
			return fail(count)
		}

		return nil
	}
}

func strLen(fn param.Params, in config.Value) (int, error) {
	if _, ok := param.Type(fn).(string); !ok {
		return 0, ErrInvalid
	}

	text, err := in.ParseString()
	if err != nil {
		return 0, fmt.Errorf("string length:%w", err)
	}

	return utf8.RuneCountInString(text), nil
}
