package config

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidValue   = errors.New("invalid value")
	ErrUnknowType     = errors.New("unknow type")
	ErrInitFactory    = errors.New("init factory")
	ErrStopWatch      = errors.New("stop watch")
	ErrNotFound       = errors.New("not found")
	ErrValueNotFound  = fmt.Errorf("value %w", ErrNotFound)
	ErrToManyArgs     = errors.New("to many args")
	ErrWrongType      = errors.New("wrong type")
	ErrInvalidName    = errors.New("ivalid name")
	ErrUnexpectedType = errors.New("unexpected type")
	ErrRequired       = errors.New("required")
)
