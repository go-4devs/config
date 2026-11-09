package value

import (
	"errors"
	"time"

	"gitoa.ru/go-4devs/config"
)

var emptyValue = Empty{Err: nil}

// EmptyValue returns a value that holds no data.
func EmptyValue() config.Value {
	return emptyValue
}

type Empty struct {
	Err error
}

// IsEmpty reports whether v is nil or the empty value.
func IsEmpty(v config.Value) bool {
	return v == nil || v == emptyValue
}

// Unmarshal returns the stored error.
func (e Empty) Unmarshal(_ any) error {
	return e.Err
}

// ParseString returns an empty string and the stored error.
func (e Empty) ParseString() (string, error) {
	return "", e.Err
}

// ParseInt returns 0 and the stored error.
func (e Empty) ParseInt() (int, error) {
	return 0, e.Err
}

// ParseInt64 returns 0 and the stored error.
func (e Empty) ParseInt64() (int64, error) {
	return 0, e.Err
}

// ParseUint returns 0 and the stored error.
func (e Empty) ParseUint() (uint, error) {
	return 0, e.Err
}

// ParseUint64 returns 0 and the stored error.
func (e Empty) ParseUint64() (uint64, error) {
	return 0, e.Err
}

// ParseFloat64 returns 0 and the stored error.
func (e Empty) ParseFloat64() (float64, error) {
	return 0, e.Err
}

// ParseBool returns false and the stored error.
func (e Empty) ParseBool() (bool, error) {
	return false, e.Err
}

// ParseDuration returns 0 and the stored error.
func (e Empty) ParseDuration() (time.Duration, error) {
	return 0, e.Err
}

// ParseTime returns the zero time and the stored error.
func (e Empty) ParseTime() (time.Time, error) {
	return time.Time{}, e.Err
}

// String returns an empty string.
func (e Empty) String() string {
	return ""
}

// Int returns 0.
func (e Empty) Int() int {
	return 0
}

// Int64 returns 0.
func (e Empty) Int64() int64 {
	return 0
}

// Uint returns 0.
func (e Empty) Uint() uint {
	return 0
}

// Uint64 returns 0.
func (e Empty) Uint64() uint64 {
	return 0
}

// Float64 returns 0.
func (e Empty) Float64() float64 {
	return 0
}

// Bool returns false.
func (e Empty) Bool() bool {
	return false
}

// Duration returns 0.
func (e Empty) Duration() time.Duration {
	return 0
}

// Time returns the zero time.
func (e Empty) Time() time.Time {
	return time.Time{}
}

// Any returns the stored error.
func (e Empty) Any() any {
	return e.Err
}

// IsEquals reports whether v is an Empty with the same error.
func (e Empty) IsEquals(v config.Value) bool {
	em, ok := v.(Empty)

	return ok && errors.Is(em.Err, e.Err)
}
