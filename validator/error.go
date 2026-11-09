package validator

import (
	"errors"
	"fmt"
)

var (
	ErrInvalid  = errors.New("invalid value")
	ErrNotBlank = errors.New("not blank")
)

// NewError returns an error that records value and the expected value.
func NewError(err error, value, expect any) Error {
	return Error{
		err:    err,
		value:  value,
		expect: expect,
	}
}

type Error struct {
	err    error
	value  any
	expect any
}

// Error returns the error text with the expected and given values.
func (e Error) Error() string {
	return fmt.Sprintf("%s: expext: %s, given: %s", e.err, e.expect, e.value)
}

// Is reports whether err matches the wrapped error.
func (e Error) Is(err error) bool {
	return errors.Is(e.err, err)
}

// Unwrap returns the wrapped error.
func (e Error) Unwrap() error {
	return e.err
}
