package value

import (
	"errors"
	"fmt"
	"time"

	"gitoa.ru/go-4devs/config"
)

var _ config.Value = (*JBytes)(nil)

type JBytes []byte

// Unmarshal JSON-decodes the bytes into v.
func (s JBytes) Unmarshal(v any) error {
	return JUnmarshal(s.Bytes(), v)
}

// ParseString JSON-decodes a string.
func (s JBytes) ParseString() (string, error) {
	var data string

	return data, s.Unmarshal(&data)
}

// ParseInt JSON-decodes an int.
func (s JBytes) ParseInt() (int, error) {
	return JParse[int](s.Bytes())
}

// ParseInt64 JSON-decodes an int64.
func (s JBytes) ParseInt64() (int64, error) {
	return JParse[int64](s.Bytes())
}

// ParseUint JSON-decodes a uint.
func (s JBytes) ParseUint() (uint, error) {
	return JParse[uint](s.Bytes())
}

// ParseUint64 JSON-decodes a uint64.
func (s JBytes) ParseUint64() (uint64, error) {
	return JParse[uint64](s.Bytes())
}

// ParseFloat64 JSON-decodes a float64.
func (s JBytes) ParseFloat64() (float64, error) {
	return JParse[float64](s.Bytes())
}

// ParseBool JSON-decodes a bool.
func (s JBytes) ParseBool() (bool, error) {
	return JParse[bool](s.Bytes())
}

// ParseDuration JSON-decodes a duration, then a JSON string passed to time.ParseDuration.
func (s JBytes) ParseDuration() (time.Duration, error) {
	jdata, jerr := JParse[time.Duration](s.Bytes())
	if jerr == nil {
		return jdata, nil
	}

	ustr, serr := s.ParseString()
	if serr != nil {
		return 0, errors.Join(jerr, fmt.Errorf("parse duration:%w", serr))
	}

	return ParseDuration(ustr)
}

// ParseTime JSON-decodes a time.Time.
func (s JBytes) ParseTime() (time.Time, error) {
	return JParse[time.Time](s.Bytes())
}

// Bytes returns the raw bytes.
func (s JBytes) Bytes() []byte {
	return []byte(s)
}

// String returns the JSON-decoded string. Decoding failure returns an empty string.
func (s JBytes) String() string {
	data, _ := s.ParseString()

	return data
}

// Int returns the JSON-decoded int. A decode error becomes 0.
func (s JBytes) Int() int {
	in, _ := s.ParseInt()

	return in
}

// Int64 returns the JSON-decoded int64. A decode error becomes 0.
func (s JBytes) Int64() int64 {
	in, _ := s.ParseInt64()

	return in
}

// Uint returns the JSON-decoded uint. A decode error becomes 0.
func (s JBytes) Uint() uint {
	in, _ := s.ParseUint()

	return in
}

// Uint64 returns the JSON-decoded uint64. A decode error becomes 0.
func (s JBytes) Uint64() uint64 {
	in, _ := s.ParseUint64()

	return in
}

// Float64 returns the JSON-decoded float64. A decode error becomes 0.
func (s JBytes) Float64() float64 {
	in, _ := s.ParseFloat64()

	return in
}

// Bool returns the JSON-decoded bool. A decode error becomes false.
func (s JBytes) Bool() bool {
	in, _ := s.ParseBool()

	return in
}

// Duration returns the decoded duration. A decode error becomes 0.
func (s JBytes) Duration() time.Duration {
	in, _ := s.ParseDuration()

	return in
}

// Time returns the JSON-decoded time. A decode error becomes the zero time.
func (s JBytes) Time() time.Time {
	in, _ := s.ParseTime()

	return in
}

// Any returns the raw bytes.
func (s JBytes) Any() any {
	return s.Bytes()
}

// IsEquals reports whether in.String equals the JSON-decoded string.
func (s JBytes) IsEquals(in config.Value) bool {
	return s.String() == in.String()
}
