package pkg

import "errors"

var (
	ErrWrongFormat = errors.New("wrong format")
	ErrNotFound    = errors.New("not found")
)
