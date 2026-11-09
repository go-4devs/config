// Package value decode value.
//
//nolint:nonamedreturns
package value

import (
	"time"

	"github.com/go-4devs/config"
)

var _ config.Value = (*Decode)(nil)

type Decode func(v any) error

// Unmarshal calls the decode function.
func (s Decode) Unmarshal(v any) error {
	return s(v)
}

// ParseInt decodes an int.
func (s Decode) ParseInt() (v int, err error) {
	return v, s.Unmarshal(&v)
}

// ParseInt64 decodes an int64.
func (s Decode) ParseInt64() (v int64, err error) {
	return v, s.Unmarshal(&v)
}

// ParseUint decodes a uint.
func (s Decode) ParseUint() (v uint, err error) {
	return v, s.Unmarshal(&v)
}

// ParseUint64 decodes a uint64.
func (s Decode) ParseUint64() (v uint64, err error) {
	return v, s.Unmarshal(&v)
}

// ParseFloat64 decodes a float64.
func (s Decode) ParseFloat64() (v float64, err error) {
	return v, s.Unmarshal(&v)
}

// ParseString decodes a string.
func (s Decode) ParseString() (v string, err error) {
	return v, s.Unmarshal(&v)
}

// ParseDuration decodes a time.Duration.
func (s Decode) ParseDuration() (v time.Duration, err error) {
	return v, s.Unmarshal(&v)
}

// ParseTime decodes a time.Time.
func (s Decode) ParseTime() (v time.Time, err error) {
	return v, s.Unmarshal(&v)
}

// Int returns the decoded int. A decode error becomes 0.
func (s Decode) Int() int {
	in, _ := s.ParseInt()

	return in
}

// Int64 returns the decoded int64. A decode error becomes 0.
func (s Decode) Int64() int64 {
	in, _ := s.ParseInt64()

	return in
}

// Uint returns the decoded uint. A decode error becomes 0.
func (s Decode) Uint() uint {
	in, _ := s.ParseUint()

	return in
}

// Uint64 returns the decoded uint64. A decode error becomes 0.
func (s Decode) Uint64() uint64 {
	in, _ := s.ParseUint64()

	return in
}

// Float64 returns the decoded float64. A decode error becomes 0.
func (s Decode) Float64() float64 {
	in, _ := s.ParseFloat64()

	return in
}

// Bytes returns the decoded string as bytes.
func (s Decode) Bytes() []byte {
	return []byte(s.String())
}

// String returns the decoded string. A decode error becomes an empty string.
func (s Decode) String() string {
	v, _ := s.ParseString()

	return v
}

// Bool returns the decoded bool. A decode error becomes false.
func (s Decode) Bool() bool {
	in, _ := s.ParseBool()

	return in
}

// ParseBool decodes a bool.
func (s Decode) ParseBool() (v bool, err error) {
	return v, s.Unmarshal(&v)
}

// Duration returns the decoded time.Duration. A decode error becomes 0.
func (s Decode) Duration() time.Duration {
	in, _ := s.ParseDuration()

	return in
}

// Time returns the decoded time. A decode error becomes the zero time.
func (s Decode) Time() time.Time {
	in, _ := s.ParseTime()

	return in
}

// Any returns the decode function.
func (s Decode) Any() any {
	return s
}

// IsEquals reports whether the decoded strings are equal.
func (s Decode) IsEquals(in config.Value) bool {
	return s.String() == in.String()
}
