package value

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-4devs/config"
)

// ParseDuration parses raw with time.ParseDuration.
func ParseDuration(raw string) (time.Duration, error) {
	d, err := time.ParseDuration(raw)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", config.ErrInvalidValue, err)
	}

	return d, nil
}

// ParseInt64 parses s as a base-10 int64.
func ParseInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", config.ErrInvalidValue, err)
	}

	return i, nil
}

// ParseUint parses s as a base-10 uint.
func ParseUint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", config.ErrInvalidValue, err)
	}

	return uint(i), nil
}

// Atoi parses s as an int.
func Atoi(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", config.ErrInvalidValue, err)
	}

	return i, nil
}

// ParseTime parses s as RFC3339 time.
func ParseTime(s string) (time.Time, error) {
	i, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: %w", config.ErrInvalidValue, err)
	}

	return i, nil
}

// ParseFloat parses s as a float64.
func ParseFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", config.ErrInvalidValue, err)
	}

	return f, nil
}

// ParseBool parses s as a bool.
func ParseBool(s string) (bool, error) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, fmt.Errorf("%w: %w", config.ErrInvalidValue, err)
	}

	return b, nil
}

// ParseUint64 parses s as a base-10 uint64.
func ParseUint64(s string) (uint64, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", config.ErrInvalidValue, err)
	}

	return i, nil
}

// JUnmarshal JSON-decodes b into v.
func JUnmarshal(b []byte, v any) error {
	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("%w: %w", config.ErrInvalidValue, err)
	}

	return nil
}

// JParse JSON-decodes b into T.
func JParse[T any](b []byte) (T, error) {
	var data T

	return data, JUnmarshal(b, &data)
}

// ParseSlice parses every string with fn and returns the slice.
func ParseSlice[T any](in []string, fn func(string) (T, error)) (config.Value, error) {
	out := make([]T, 0, len(in))
	for idx, one := range in {
		data, err := fn(one)
		if err != nil {
			return nil, fmt.Errorf("parse[%T][%d]:%w", out, idx, err)
		}

		out = append(out, data)
	}

	return New(out), nil
}
