package value

import (
	"fmt"
	"time"

	"github.com/go-4devs/config"
)

var _ config.Value = String("")

type String string

// ParseString returns the text.
func (s String) ParseString() (string, error) {
	return string(s), nil
}

// Unmarshal copies the text into *string.
func (s String) Unmarshal(in any) error {
	v, ok := in.(*string)
	if !ok {
		return fmt.Errorf("%w: expect *string", config.ErrWrongType)
	}

	*v = string(s)

	return nil
}

// Any returns the text.
func (s String) Any() any {
	return string(s)
}

// ParseInt parses the text as an int.
func (s String) ParseInt() (int, error) {
	return Atoi(s.String())
}

// Int64 returns the text parsed as an int64. A parse error becomes 0.
func (s String) Int64() int64 {
	out, _ := s.ParseInt64()

	return out
}

// ParseInt64 parses the text as an int64.
func (s String) ParseInt64() (int64, error) {
	return ParseInt64(s.String())
}

// ParseUint parses the text as a uint.
func (s String) ParseUint() (uint, error) {
	return ParseUint(s.String())
}

// ParseUint64 parses the text as a uint64.
func (s String) ParseUint64() (uint64, error) {
	return ParseUint64(s.String())
}

// ParseFloat64 parses the text as a float64.
func (s String) ParseFloat64() (float64, error) {
	return ParseFloat(s.String())
}

// ParseBool parses the text as a bool.
func (s String) ParseBool() (bool, error) {
	return ParseBool(s.String())
}

// ParseDuration parses the text as a time.Duration.
func (s String) ParseDuration() (time.Duration, error) {
	return ParseDuration(s.String())
}

// ParseTime parses the text as RFC3339 time.
func (s String) ParseTime() (time.Time, error) {
	return ParseTime(s.String())
}

// IsEquals reports whether val is the same text.
func (s String) IsEquals(val config.Value) bool {
	data, ok := val.(String)

	return ok && data == s
}

// String returns the text.
func (s String) String() string {
	data, _ := s.ParseString()

	return data
}

// Int returns the text parsed as an int. A parse error becomes 0.
func (s String) Int() int {
	data, _ := s.ParseInt()

	return data
}

// Uint returns the text parsed as a uint. A parse error becomes 0.
func (s String) Uint() uint {
	data, _ := s.ParseUint()

	return data
}

// Uint64 returns the text parsed as a uint64. A parse error becomes 0.
func (s String) Uint64() uint64 {
	data, _ := s.ParseUint64()

	return data
}

// Float64 returns the text parsed as a float64. A parse error becomes 0.
func (s String) Float64() float64 {
	data, _ := s.ParseFloat64()

	return data
}

// Bool returns the text parsed as a bool. A parse error becomes false.
func (s String) Bool() bool {
	data, _ := s.ParseBool()

	return data
}

// Duration returns the text parsed as a time.Duration. A parse error becomes 0.
func (s String) Duration() time.Duration {
	data, _ := s.ParseDuration()

	return data
}

// Time returns the text parsed as RFC3339 time. A parse error becomes the zero time.
func (s String) Time() time.Time {
	data, _ := s.ParseTime()

	return data
}
