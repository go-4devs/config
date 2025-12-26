package validator

import (
	"errors"
	"fmt"
)

var (
	ErrInvalid  = errors.New("invalid value")
	ErrNotBlank = errors.New("not blank")
)

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

func (e Error) Error() string {
	return fmt.Sprintf("%s: expext: %s, given: %s", e.err, e.expect, e.value)
}

func (e Error) Is(err error) bool {
	return errors.Is(e.err, err)
}

func (e Error) Unwrap() error {
	return e.err
}
