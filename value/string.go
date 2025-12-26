package value

import (
	"fmt"
	"strconv"
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
	v, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("string int:%w", err)
	}

	return v, nil
}

func (s String) Int64() int64 {
	out, _ := s.ParseInt64()

	return out
}

func (s String) ParseInt64() (int64, error) {
	v, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("string int64:%w", err)
	}

	return v, nil
}

func (s String) ParseUint() (uint, error) {
	uout, err := s.ParseUint64()

	return uint(uout), err
}

func (s String) ParseUint64() (uint64, error) {
	uout, err := strconv.ParseUint(string(s), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("string uint:%w", err)
	}

	return uout, nil
}

func (s String) ParseFloat64() (float64, error) {
	fout, err := strconv.ParseFloat(string(s), 64)
	if err != nil {
		return 0, fmt.Errorf("string float64:%w", err)
	}

	return fout, nil
}

func (s String) ParseBool() (bool, error) {
	v, err := strconv.ParseBool(string(s))
	if err != nil {
		return false, fmt.Errorf("string bool:%w", err)
	}

	return v, nil
}

func (s String) ParseDuration() (time.Duration, error) {
	v, err := time.ParseDuration(string(s))
	if err != nil {
		return 0, fmt.Errorf("string duration:%w", err)
	}

	return v, nil
}

func (s String) ParseTime() (time.Time, error) {
	v, err := time.Parse(time.RFC3339, string(s))
	if err != nil {
		return time.Time{}, fmt.Errorf("string time:%w", err)
	}

	return v, nil
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
