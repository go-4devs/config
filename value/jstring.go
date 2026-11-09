package value

import (
	"time"

	"github.com/go-4devs/config"
)

var _ config.Value = (*JString)(nil)

type JString string

// Unmarshal JSON-decodes the raw text into v.
func (s JString) Unmarshal(v any) error {
	return JUnmarshal(s.Bytes(), v)
}

// ParseString returns the raw text.
func (s JString) ParseString() (string, error) {
	return s.String(), nil
}

// ParseInt parses the raw text as an int.
func (s JString) ParseInt() (int, error) {
	return Atoi(s.String())
}

// ParseInt64 parses the raw text as an int64.
func (s JString) ParseInt64() (int64, error) {
	return ParseInt64(s.String())
}

// ParseUint parses the raw text as a uint.
func (s JString) ParseUint() (uint, error) {
	return ParseUint(s.String())
}

// ParseUint64 parses the raw text as a uint64.
func (s JString) ParseUint64() (uint64, error) {
	return ParseUint64(s.String())
}

// ParseFloat64 parses the raw text as a float64.
func (s JString) ParseFloat64() (float64, error) {
	return ParseFloat(s.String())
}

// ParseBool parses the raw text as a bool.
func (s JString) ParseBool() (bool, error) {
	return ParseBool(s.String())
}

// ParseDuration parses the raw text as a time.Duration.
func (s JString) ParseDuration() (time.Duration, error) {
	return ParseDuration(s.String())
}

// ParseTime parses the raw text as RFC3339 time.
func (s JString) ParseTime() (time.Time, error) {
	return ParseTime(s.String())
}

// Bytes returns the raw text as bytes.
func (s JString) Bytes() []byte {
	return []byte(s)
}

// String returns the raw text.
func (s JString) String() string {
	return string(s)
}

// Int returns the raw text parsed as an int. A parse error becomes 0.
func (s JString) Int() int {
	in, _ := s.ParseInt()

	return in
}

// Int64 returns the raw text parsed as an int64. A parse error becomes 0.
func (s JString) Int64() int64 {
	in, _ := s.ParseInt64()

	return in
}

// Uint returns the raw text parsed as a uint. A parse error becomes 0.
func (s JString) Uint() uint {
	in, _ := s.ParseUint()

	return in
}

// Uint64 returns the raw text parsed as a uint64. A parse error becomes 0.
func (s JString) Uint64() uint64 {
	in, _ := s.ParseUint64()

	return in
}

// Float64 returns the raw text parsed as a float64. A parse error becomes 0.
func (s JString) Float64() float64 {
	in, _ := s.ParseFloat64()

	return in
}

// Bool returns the raw text parsed as a bool. A parse error becomes false.
func (s JString) Bool() bool {
	in, _ := s.ParseBool()

	return in
}

// Duration returns the raw text parsed as a time.Duration. A parse error becomes 0.
func (s JString) Duration() time.Duration {
	in, _ := s.ParseDuration()

	return in
}

// Time returns the raw text parsed as RFC3339 time. A parse error becomes the zero time.
func (s JString) Time() time.Time {
	in, _ := s.ParseTime()

	return in
}

// Any returns the JString.
func (s JString) Any() any {
	return s
}

// IsEquals reports whether in.String equals the raw text.
func (s JString) IsEquals(in config.Value) bool {
	return s.String() == in.String()
}
