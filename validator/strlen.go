package validator

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

var (
	ErrMinLen = errors.New("shorter than minimum")
	ErrMaxLen = errors.New("longer than maximum")
)

// MinLen returns a validator that rejects strings shorter than limit.
// The option type must be string. The bound is inclusive and counts Unicode code points.
func MinLen(limit int) Validator {
	return lenBound(limit, func(count int) bool { return count < limit }, ErrMinLen)
}

// MaxLen returns a validator that rejects strings longer than limit.
// The option type must be string. The bound is inclusive and counts Unicode code points.
func MaxLen(limit int) Validator {
	return lenBound(limit, func(count int) bool { return count > limit }, ErrMaxLen)
}

func lenBound(limit int, tooFar func(int) bool, sentinel error) Validator {
	return func(fn param.Params, in config.Value) error {
		if limit < 0 {
			return ErrInvalid
		}

		count, err := strLen(fn, in)
		if err != nil {
			return err
		}

		if tooFar(count) {
			return NewError(sentinel, count, limit)
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
