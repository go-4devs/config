package validator

import (
	"fmt"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
)

func NotBlank(fn param.Params, in config.Value) error {
	dataType := param.Type(fn)
	if option.IsSlice(fn) {
		return sliceByType(dataType, in)
	}

	if notBlank(dataType, in) {
		return nil
	}

	return ErrNotBlank
}

func sliceByType(vType any, in config.Value) error {
	switch vType.(type) {
	case string:
		return sliceBy[string](in)
	case int:
		return sliceBy[int](in)
	case int64:
		return sliceBy[int64](in)
	case uint:
		return sliceBy[uint](in)
	case uint64:
		return sliceBy[uint64](in)
	case float64:
		return sliceBy[float64](in)
	case bool:
		return sliceBy[bool](in)
	case time.Duration:
		return sliceBy[time.Duration](in)
	case time.Time:
		return sliceBy[time.Time](in)
	default:
		return sliceBy[any](in)
	}
}
func sliceBy[T any](in config.Value) error {
	var data []T
	if err := in.Unmarshal(&data); err != nil {
		return fmt.Errorf("%w:%w", ErrNotBlank, err)
	}

	if len(data) > 0 {
		return nil
	}

	return fmt.Errorf("%w", ErrNotBlank)
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
