package value

import (
	"fmt"
	"time"

	"gitoa.ru/go-4devs/config"
)

var _ config.Value = Float64(0)

type Float64 float64

func (f Float64) Any() any {
	return float64(f)
}

func (f Float64) ParseString() (string, error) {
	return fmt.Sprint(float64(f)), nil
}

func (f Float64) ParseInt() (int, error) {
	return int(f), nil
}

func (f Float64) ParseInt64() (int64, error) {
	return int64(f), nil
}

func (f Float64) ParseUint() (uint, error) {
	return uint(f), nil
}

func (f Float64) ParseUint64() (uint64, error) {
	return uint64(f), nil
}

func (f Float64) ParseFloat64() (float64, error) {
	return float64(f), nil
}

func (f Float64) ParseBool() (bool, error) {
	return false, fmt.Errorf("float64:%w", config.ErrWrongType)
}

func (f Float64) ParseDuration() (time.Duration, error) {
	return time.Duration(f), nil
}

func (f Float64) ParseTime() (time.Time, error) {
	return time.Unix(0, int64(f*Float64(time.Second))), nil
}

func (f Float64) Unmarshal(in any) error {
	v, ok := in.(*float64)
	if !ok {
		return fmt.Errorf("%w: expect *float64", config.ErrWrongType)
	}

	*v = float64(f)

	return nil
}

func (f Float64) IsEquals(val config.Value) bool {
	data, ok := val.(Float64)

	return ok && data == f
}

func (f Float64) String() string {
	data, _ := f.ParseString()

	return data
}

func (f Float64) Int() int {
	data, _ := f.ParseInt()

	return data
}

func (f Float64) Int64() int64 {
	data, _ := f.ParseInt64()

	return data
}

func (f Float64) Uint() uint {
	data, _ := f.ParseUint()

	return data
}

func (f Float64) Uint64() uint64 {
	data, _ := f.ParseUint64()

	return data
}

func (f Float64) Float64() float64 {
	data, _ := f.ParseFloat64()

	return data
}

func (f Float64) Bool() bool {
	data, _ := f.ParseBool()

	return data
}

func (f Float64) Duration() time.Duration {
	data, _ := f.ParseDuration()

	return data
}

func (f Float64) Time() time.Time {
	data, _ := f.ParseTime()

	return data
}
