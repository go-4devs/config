package value

import (
	"fmt"
	"time"

	"gitoa.ru/go-4devs/config"
)

var _ config.Value = (String)("")

type String string

func (s String) ParseString() (string, error) {
	return string(s), nil
}

func (s String) Unmarshal(in any) error {
	v, ok := in.(*string)
	if !ok {
		return fmt.Errorf("%w: expect *string", config.ErrWrongType)
	}

	*v = string(s)

	return nil
}

func (s String) Any() any {
	return string(s)
}

func (s String) ParseInt() (int, error) {
	return Atoi(s.String())
}

func (s String) Int64() int64 {
	out, _ := s.ParseInt64()

	return out
}

func (s String) ParseInt64() (int64, error) {
	return ParseInt64(s.String())
}

func (s String) ParseUint() (uint, error) {
	return ParseUint(s.String())
}

func (s String) ParseUint64() (uint64, error) {
	return ParseUint64(s.String())
}

func (s String) ParseFloat64() (float64, error) {
	return ParseFloat(s.String())
}

func (s String) ParseBool() (bool, error) {
	return ParseBool(s.String())
}

func (s String) ParseDuration() (time.Duration, error) {
	return ParseDuration(s.String())
}

func (s String) ParseTime() (time.Time, error) {
	return ParseTime(s.String())
}

func (s String) IsEquals(val config.Value) bool {
	data, ok := val.(String)

	return ok && data == s
}

func (s String) String() string {
	data, _ := s.ParseString()

	return data
}

func (s String) Int() int {
	data, _ := s.ParseInt()

	return data
}

func (s String) Uint() uint {
	data, _ := s.ParseUint()

	return data
}

func (s String) Uint64() uint64 {
	data, _ := s.ParseUint64()

	return data
}

func (s String) Float64() float64 {
	data, _ := s.ParseFloat64()

	return data
}

func (s String) Bool() bool {
	data, _ := s.ParseBool()

	return data
}

func (s String) Duration() time.Duration {
	data, _ := s.ParseDuration()

	return data
}

func (s String) Time() time.Time {
	data, _ := s.ParseTime()

	return data
}
