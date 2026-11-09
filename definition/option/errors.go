package option

import (
	"errors"
	"fmt"
)

type Error struct {
	Key []string
	Err error
}

// Error returns the key and the wrapped error.
func (o Error) Error() string {
	return fmt.Sprintf("%s: %s", o.Key, o.Err)
}

// Is reports whether err matches the wrapped error.
func (o Error) Is(err error) bool {
	return errors.Is(err, o.Err)
}

// Unwrap returns the wrapped error.
func (o Error) Unwrap() error {
	return o.Err
}

// Err returns an error tagged with key.
func Err(err error, key []string) Error {
	return Error{
		Key: key,
		Err: err,
	}
}
