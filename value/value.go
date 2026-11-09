package value

import (
	"fmt"
	"reflect"
	"time"

	"gitoa.ru/go-4devs/config"
)

var _ config.Value = (*Value)(nil)

// New returns data when it already implements config.Value. Otherwise it stores data in a Value.
func New(data any) config.Value {
	switch val := data.(type) {
	case config.Value:
		return val
	default:
		return Value{Val: data}
	}
}

type Value struct {
	Val any
}

// Int returns the stored int. Another type becomes 0.
func (s Value) Int() int {
	v, _ := s.ParseInt()

	return v
}

// Int64 returns the stored int64. Another type becomes 0.
func (s Value) Int64() int64 {
	v, _ := s.ParseInt64()

	return v
}

// Uint returns the stored uint. Another type becomes 0.
func (s Value) Uint() uint {
	v, _ := s.ParseUint()

	return v
}

// Uint64 returns the stored uint64. Another type becomes 0.
func (s Value) Uint64() uint64 {
	v, _ := s.ParseUint64()

	return v
}

// Float64 returns the stored float64. Another type becomes 0.
func (s Value) Float64() float64 {
	in, _ := s.ParseFloat64()

	return in
}

// String returns the stored string. Another type becomes an empty string.
func (s Value) String() string {
	v, _ := s.ParseString()

	return v
}

// Bool returns the stored bool. Another type becomes false.
func (s Value) Bool() bool {
	v, _ := s.ParseBool()

	return v
}

// Duration returns the stored time.Duration. Another type becomes 0.
func (s Value) Duration() time.Duration {
	v, _ := s.ParseDuration()

	return v
}

// Time returns the stored time.Time. Another type becomes the zero time.
func (s Value) Time() time.Time {
	v, _ := s.ParseTime()

	return v
}

// Unmarshal assigns the stored value to a pointer of the same kind.
func (s Value) Unmarshal(target any) error {
	return typeAssert(s.Val, target)
}

// ParseInt returns the stored value when it is an int.
func (s Value) ParseInt() (int, error) {
	if r, ok := s.Any().(int); ok {
		return r, nil
	}

	return 0, config.ErrInvalidValue
}

// ParseInt64 returns the stored value when it is an int64.
func (s Value) ParseInt64() (int64, error) {
	if r, ok := s.Any().(int64); ok {
		return r, nil
	}

	return 0, config.ErrInvalidValue
}

// ParseUint returns the stored value when it is a uint.
func (s Value) ParseUint() (uint, error) {
	if r, ok := s.Any().(uint); ok {
		return r, nil
	}

	return 0, config.ErrInvalidValue
}

// ParseUint64 returns the stored value when it is a uint64.
func (s Value) ParseUint64() (uint64, error) {
	if r, ok := s.Any().(uint64); ok {
		return r, nil
	}

	return 0, config.ErrInvalidValue
}

// ParseFloat64 returns the stored value when it is a float64.
func (s Value) ParseFloat64() (float64, error) {
	if r, ok := s.Any().(float64); ok {
		return r, nil
	}

	return 0, config.ErrInvalidValue
}

// ParseString returns the stored value when it is a string.
func (s Value) ParseString() (string, error) {
	if r, ok := s.Any().(string); ok {
		return r, nil
	}

	return "", config.ErrInvalidValue
}

// ParseBool returns the stored value when it is a bool.
func (s Value) ParseBool() (bool, error) {
	if b, ok := s.Any().(bool); ok {
		return b, nil
	}

	return false, config.ErrInvalidValue
}

// ParseDuration returns the stored value when it is a time.Duration.
func (s Value) ParseDuration() (time.Duration, error) {
	if b, ok := s.Any().(time.Duration); ok {
		return b, nil
	}

	return 0, config.ErrInvalidValue
}

// ParseTime returns the stored value when it is a time.Time.
func (s Value) ParseTime() (time.Time, error) {
	if b, ok := s.Any().(time.Time); ok {
		return b, nil
	}

	return time.Time{}, config.ErrInvalidValue
}

// IsEquals reports whether the stored values are equal.
func (s Value) IsEquals(in config.Value) bool {
	return s.Any() == in.Any()
}

// Any returns the stored value.
func (s Value) Any() any {
	return s.Val
}

func typeAssert(source, target any) error {
	if source == nil {
		return nil
	}

	if directTypeAssert(source, target) {
		return nil
	}

	valTarget := reflect.ValueOf(target)
	if !valTarget.IsValid() || valTarget.Kind() != reflect.Pointer {
		return fmt.Errorf("ptr target:%w", config.ErrInvalidValue)
	}

	valTarget = valTarget.Elem()

	if !valTarget.IsValid() {
		return fmt.Errorf("elem targer:%w", config.ErrInvalidValue)
	}

	valSource := reflect.ValueOf(source)
	if !valSource.IsValid() {
		return fmt.Errorf("source:%w", config.ErrInvalidValue)
	}

	valSource = deReference(valSource)
	if err := canSet(valSource, valTarget); err != nil {
		return fmt.Errorf("can set:%w", err)
	}

	valTarget.Set(valSource)

	return nil
}

func canSet(source, target reflect.Value) error {
	if source.Kind() != target.Kind() {
		return fmt.Errorf("source=%v target=%v:%w", source.Kind(), target.Kind(), config.ErrInvalidValue)
	}

	if source.Kind() == reflect.Slice && source.Type().Elem().Kind() != target.Type().Elem().Kind() {
		return fmt.Errorf("slice source=%v, slice target=%v:%w",
			source.Type().Elem().Kind(), target.Type().Elem().Kind(), config.ErrInvalidValue)
	}

	return nil
}

func directTypeAssert(source, target any) bool {
	var ok bool

	switch val := target.(type) {
	case *string:
		*val, ok = source.(string)
	case *[]byte:
		*val, ok = source.([]byte)
	case *int:
		*val, ok = source.(int)
	case *int64:
		*val, ok = source.(int64)
	case *uint:
		*val, ok = source.(uint)
	case *uint64:
		*val, ok = source.(uint64)
	case *bool:
		*val, ok = source.(bool)
	case *float64:
		*val, ok = source.(float64)
	case *time.Duration:
		*val, ok = source.(time.Duration)
	case *time.Time:
		*val, ok = source.(time.Time)
	case *[]string:
		*val, ok = source.([]string)
	case *map[string]string:
		*val, ok = source.(map[string]string)
	case *map[string]any:
		*val, ok = source.(map[string]any)
	}

	return ok
}

func deReference(v reflect.Value) reflect.Value {
	if (v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface) && !v.IsNil() {
		return v.Elem()
	}

	return v
}
