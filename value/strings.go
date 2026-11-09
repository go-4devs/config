package value

import (
	"encoding/json"
	"fmt"
	"time"

	"gitoa.ru/go-4devs/config"
)

type Strings []string

// Unmarshal JSON-decodes the slice into in.
func (s Strings) Unmarshal(in any) error {
	data, jerr := json.Marshal([]string(s))
	if jerr != nil {
		return fmt.Errorf("failed load data:%w", jerr)
	}

	if err := json.Unmarshal(data, in); err != nil {
		return fmt.Errorf("unmarshal:%w", err)
	}

	return nil
}

// ParseString returns the only element, or the slice formatted with %v.
func (s Strings) ParseString() (string, error) {
	return s.String(), nil
}

// ParseInt parses String as an int.
func (s Strings) ParseInt() (int, error) {
	data, err := Atoi(s.String())
	if err != nil {
		return 0, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

// ParseInt64 parses String as an int64.
func (s Strings) ParseInt64() (int64, error) {
	data, err := ParseInt64(s.String())
	if err != nil {
		return 0, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

// ParseUint parses String as a uint.
func (s Strings) ParseUint() (uint, error) {
	data, err := ParseUint(s.String())
	if err != nil {
		return 0, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

// ParseUint64 parses String as a uint64.
func (s Strings) ParseUint64() (uint64, error) {
	data, err := ParseUint64(s.String())
	if err != nil {
		return 0, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

// ParseFloat64 parses String as a float64.
func (s Strings) ParseFloat64() (float64, error) {
	data, err := ParseFloat(s.String())
	if err != nil {
		return 0, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

// ParseBool parses String as a bool.
func (s Strings) ParseBool() (bool, error) {
	data, err := ParseBool(s.String())
	if err != nil {
		return false, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

// ParseDuration parses String as a time.Duration.
func (s Strings) ParseDuration() (time.Duration, error) {
	data, err := ParseDuration(s.String())
	if err != nil {
		return 0, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

// ParseTime parses String as RFC3339 time.
func (s Strings) ParseTime() (time.Time, error) {
	data, err := ParseTime(s.String())
	if err != nil {
		return time.Time{}, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

// String returns the only element, or the slice formatted with %v.
func (s Strings) String() string {
	if len(s) == 1 {
		return s[0]
	}

	return fmt.Sprintf("%v", []string(s))
}

// Int returns String parsed as an int. A parse error becomes 0.
func (s Strings) Int() int {
	val, _ := s.ParseInt()

	return val
}

// Int64 returns String parsed as an int64. A parse error becomes 0.
func (s Strings) Int64() int64 {
	val, _ := s.ParseInt64()

	return val
}

// Uint returns String parsed as a uint. A parse error becomes 0.
func (s Strings) Uint() uint {
	val, _ := s.ParseUint()

	return val
}

// Uint64 returns String parsed as a uint64. A parse error becomes 0.
func (s Strings) Uint64() uint64 {
	val, _ := s.ParseUint64()

	return val
}

// Float64 returns String parsed as a float64. A parse error becomes 0.
func (s Strings) Float64() float64 {
	val, _ := s.ParseFloat64()

	return val
}

// Bool returns String parsed as a bool. A parse error becomes false.
func (s Strings) Bool() bool {
	val, _ := s.ParseBool()

	return val
}

// Duration returns String parsed as a time.Duration. A parse error becomes 0.
func (s Strings) Duration() time.Duration {
	val, _ := s.ParseDuration()

	return val
}

// Time returns String parsed as RFC3339 time. A parse error becomes the zero time.
func (s Strings) Time() time.Time {
	val, _ := s.ParseTime()

	return val
}

// Any returns the slice.
func (s Strings) Any() any {
	return s
}

// Append returns the slice with data added.
func (s Strings) Append(data string) Strings {
	return append(s, data)
}

// IsEquals reports whether in is a Strings value with the same elements.
func (s Strings) IsEquals(in config.Value) bool {
	val, ok := in.(Strings)
	if !ok || len(s) != len(val) {
		return false
	}

	for idx := range val {
		if s[idx] != val[idx] {
			return false
		}
	}

	return true
}
