package validator

import (
	"fmt"
	"time"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/param"
)

// NotBlankError is returned when Got is empty.
type NotBlankError struct {
	Got any
}

// Error returns the received value.
func (e NotBlankError) Error() string {
	return fmt.Sprintf("%s: want non-empty, given %v", ErrNotBlank, e.Got)
}

// Unwrap returns ErrNotBlank.
func (e NotBlankError) Unwrap() error {
	return ErrNotBlank
}

// NotBlank rejects a zero value. A slice must be non-empty.
func NotBlank(fn param.Params, in config.Value) error {
	dataType := param.Type(fn)
	if option.IsSlice(fn) {
		return sliceByType(dataType, in)
	}

	if notBlank(dataType, in) {
		return nil
	}

	return NotBlankError{Got: in.Any()}
}

func filled(fn param.Params, in config.Value) (bool, error) {
	dataType := param.Type(fn)
	if !option.IsSlice(fn) {
		return notBlank(dataType, in), nil
	}

	size, err := sliceLen(dataType, in)
	if err != nil {
		return false, err
	}

	return size > 0, nil
}

func sliceByType(vType any, in config.Value) error {
	size, err := sliceLen(vType, in)
	if err != nil {
		return err
	}

	if size > 0 {
		return nil
	}

	return NotBlankError{Got: in.Any()}
}

func sliceLen(vType any, in config.Value) (int, error) {
	switch vType.(type) {
	case string:
		return sliceCount[string](in)
	case int:
		return sliceCount[int](in)
	case int64:
		return sliceCount[int64](in)
	case uint:
		return sliceCount[uint](in)
	case uint64:
		return sliceCount[uint64](in)
	case float64:
		return sliceCount[float64](in)
	case bool:
		return sliceCount[bool](in)
	case time.Duration:
		return sliceCount[time.Duration](in)
	case time.Time:
		return sliceCount[time.Time](in)
	default:
		return sliceCount[any](in)
	}
}

func sliceCount[T any](in config.Value) (int, error) {
	var data []T
	if err := in.Unmarshal(&data); err != nil {
		return 0, fmt.Errorf("%w:%w", ErrNotBlank, err)
	}

	return len(data), nil
}

func notBlank(vType any, in config.Value) bool {
	switch vType.(type) {
	case int:
		return in.Int() != 0
	case int64:
		return in.Int64() != 0
	case uint:
		return in.Uint() != 0
	case uint64:
		return in.Uint64() != 0
	case float64:
		return in.Float64() != 0
	case time.Duration:
		return in.Duration() != 0
	case time.Time:
		return !in.Time().IsZero()
	case string:
		return len(in.String()) > 0
	default:
		return in.Any() != nil
	}
}
