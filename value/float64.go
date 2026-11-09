package value

import (
	"fmt"
	"time"

	"gitoa.ru/go-4devs/config"
)

var _ config.Value = Float64(0)

type Float64 float64

// Any returns the number.
func (f Float64) Any() any {
	return float64(f)
}

// ParseString formats the number.
func (f Float64) ParseString() (string, error) {
	return fmt.Sprint(float64(f)), nil
}

// ParseInt converts the number to int.
func (f Float64) ParseInt() (int, error) {
	return int(f), nil
}

// ParseInt64 converts the number to int64.
func (f Float64) ParseInt64() (int64, error) {
	return int64(f), nil
}

// ParseUint converts the number to uint.
func (f Float64) ParseUint() (uint, error) {
	return uint(f), nil
}

// ParseUint64 converts the number to uint64.
func (f Float64) ParseUint64() (uint64, error) {
	return uint64(f), nil
}

// ParseFloat64 returns the number.
func (f Float64) ParseFloat64() (float64, error) {
	return float64(f), nil
}

// ParseBool returns ErrWrongType.
func (f Float64) ParseBool() (bool, error) {
	return false, fmt.Errorf("float64:%w", config.ErrWrongType)
}

// ParseDuration converts the number to a time.Duration.
func (f Float64) ParseDuration() (time.Duration, error) {
	return time.Duration(f), nil
}

// ParseTime treats the number as seconds since the Unix epoch.
func (f Float64) ParseTime() (time.Time, error) {
	return time.Unix(0, int64(f*Float64(time.Second))), nil
}

// Unmarshal copies the number into *float64.
func (f Float64) Unmarshal(in any) error {
	v, ok := in.(*float64)
	if !ok {
		return fmt.Errorf("%w: expect *float64", config.ErrWrongType)
	}

	*v = float64(f)

	return nil
}

// IsEquals reports whether val is the same number.
func (f Float64) IsEquals(val config.Value) bool {
	data, ok := val.(Float64)

	return ok && data == f
}

// String returns the formatted number.
func (f Float64) String() string {
	data, _ := f.ParseString()

	return data
}

// Int returns the number converted to int.
func (f Float64) Int() int {
	data, _ := f.ParseInt()

	return data
}

// Int64 returns the number converted to int64.
func (f Float64) Int64() int64 {
	data, _ := f.ParseInt64()

	return data
}

// Uint returns the number converted to uint.
func (f Float64) Uint() uint {
	data, _ := f.ParseUint()

	return data
}

// Uint64 returns the number converted to uint64.
func (f Float64) Uint64() uint64 {
	data, _ := f.ParseUint64()

	return data
}

// Float64 returns the number.
func (f Float64) Float64() float64 {
	data, _ := f.ParseFloat64()

	return data
}

// Bool returns false.
func (f Float64) Bool() bool {
	data, _ := f.ParseBool()

	return data
}

// Duration returns the number converted to a time.Duration.
func (f Float64) Duration() time.Duration {
	data, _ := f.ParseDuration()

	return data
}

// Time returns the number as seconds since the Unix epoch.
func (f Float64) Time() time.Time {
	data, _ := f.ParseTime()

	return data
}
