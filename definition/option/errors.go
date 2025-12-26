package option

import (
	"errors"
	"fmt"
)

type Error struct {
	Key []string
	Err error
}

func (o Error) Error() string {
	return fmt.Sprintf("%s: %s", o.Key, o.Err)
}

func (o Error) Is(err error) bool {
	return errors.Is(err, o.Err)
}

func (o Error) Unwrap() error {
	return o.Err
}

func Err(err error, key []string) Error {
	return Error{
		Key: key,
		Err: err,
	}
}
