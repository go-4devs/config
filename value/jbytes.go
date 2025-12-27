package value

import (
	"time"

	"gitoa.ru/go-4devs/config"
)

var _ config.Value = (*JBytes)(nil)

type JBytes []byte

func (s JBytes) Unmarshal(v any) error {
	return JUnmarshal(s.Bytes(), v)
}

func (s JBytes) ParseString() (string, error) {
	var data string

	return data, s.Unmarshal(&data)
}

func (s JBytes) ParseInt() (int, error) {
	return JParce[int](s.Bytes())
}

func (s JBytes) ParseInt64() (int64, error) {
	return JParce[int64](s.Bytes())
}

func (s JBytes) ParseUint() (uint, error) {
	return JParce[uint](s.Bytes())
}

func (s JBytes) ParseUint64() (uint64, error) {
	return JParce[uint64](s.Bytes())
}

func (s JBytes) ParseFloat64() (float64, error) {
	return JParce[float64](s.Bytes())
}

func (s JBytes) ParseBool() (bool, error) {
	return JParce[bool](s.Bytes())
}

func (s JBytes) ParseDuration() (time.Duration, error) {
	return JParce[time.Duration](s.Bytes())
}

func (s JBytes) ParseTime() (time.Time, error) {
	return JParce[time.Time](s.Bytes())
}

func (s JBytes) Bytes() []byte {
	return []byte(s)
}

func (s JBytes) String() string {
	data, _ := s.ParseString()

	return data
}

func (s JBytes) Int() int {
	in, _ := s.ParseInt()

	return in
}

func (s JBytes) Int64() int64 {
	in, _ := s.ParseInt64()

	return in
}

func (s JBytes) Uint() uint {
	in, _ := s.ParseUint()

	return in
}

func (s JBytes) Uint64() uint64 {
	in, _ := s.ParseUint64()

	return in
}

func (s JBytes) Float64() float64 {
	in, _ := s.ParseFloat64()

	return in
}

func (s JBytes) Bool() bool {
	in, _ := s.ParseBool()

	return in
}

func (s JBytes) Duration() time.Duration {
	in, _ := s.ParseDuration()

	return in
}

func (s JBytes) Time() time.Time {
	in, _ := s.ParseTime()

	return in
}

func (s JBytes) Any() any {
	return s.Bytes()
}

func (s JBytes) IsEquals(in config.Value) bool {
	return s.String() == in.String()
}
