package value

import (
	"encoding/json"
	"fmt"
	"time"

	"gitoa.ru/go-4devs/config"
)

type Strings []string

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

func (s Strings) ParseString() (string, error) {
	return s.String(), nil
}

func (s Strings) ParseInt() (int, error) {
	data, err := Atoi(s.String())
	if err != nil {
		return 0, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

func (s Strings) ParseInt64() (int64, error) {
	data, err := ParseInt64(s.String())
	if err != nil {
		return 0, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

func (s Strings) ParseUint() (uint, error) {
	data, err := ParseUint(s.String())
	if err != nil {
		return 0, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

func (s Strings) ParseUint64() (uint64, error) {
	data, err := ParseUint64(s.String())
	if err != nil {
		return 0, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

func (s Strings) ParseFloat64() (float64, error) {
	data, err := ParseFloat(s.String())
	if err != nil {
		return 0, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

func (s Strings) ParseBool() (bool, error) {
	data, err := ParseBool(s.String())
	if err != nil {
		return false, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

func (s Strings) ParseDuration() (time.Duration, error) {
	data, err := ParseDuration(s.String())
	if err != nil {
		return 0, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

func (s Strings) ParseTime() (time.Time, error) {
	data, err := ParseTime(s.String())
	if err != nil {
		return time.Time{}, fmt.Errorf("strings: %w", err)
	}

	return data, nil
}

func (s Strings) String() string {
	if len(s) == 1 {
		return s[0]
	}

	return fmt.Sprintf("%v", []string(s))
}

func (s Strings) Int() int {
	val, _ := s.ParseInt()

	return val
}

func (s Strings) Int64() int64 {
	val, _ := s.ParseInt64()

	return val
}

func (s Strings) Uint() uint {
	val, _ := s.ParseUint()

	return val
}

func (s Strings) Uint64() uint64 {
	val, _ := s.ParseUint64()

	return val
}

func (s Strings) Float64() float64 {
	val, _ := s.ParseFloat64()

	return val
}

func (s Strings) Bool() bool {
	val, _ := s.ParseBool()

	return val
}

func (s Strings) Duration() time.Duration {
	val, _ := s.ParseDuration()

	return val
}

func (s Strings) Time() time.Time {
	val, _ := s.ParseTime()

	return val
}

func (s Strings) Any() any {
	return s
}

func (s Strings) Append(data string) Strings {
	return append(s, data)
}

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
