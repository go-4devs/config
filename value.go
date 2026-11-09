package config

import (
	"time"
)

type Value interface {
	ReadValue
	ParseValue
	UnmarshalValue
	// IsEquals reports whether the value is equal to in.
	IsEquals(in Value) bool
}

type UnmarshalValue interface {
	// Unmarshal decodes the value into val.
	Unmarshal(val any) error
}

type ReadValue interface {
	// String returns the value as a string.
	String() string
	// Int returns the value as an int.
	Int() int
	// Int64 returns the value as an int64.
	Int64() int64
	// Uint returns the value as a uint.
	Uint() uint
	// Uint64 returns the value as a uint64.
	Uint64() uint64
	// Float64 returns the value as a float64.
	Float64() float64
	// Bool returns the value as a bool.
	Bool() bool
	// Duration returns the value as a time.Duration.
	Duration() time.Duration
	// Time returns the value as a time.Time.
	Time() time.Time
	// Any returns the underlying value.
	Any() any
}

type ParseValue interface {
	// ParseString parses the value as a string.
	ParseString() (string, error)
	// ParseInt parses the value as an int.
	ParseInt() (int, error)
	// ParseInt64 parses the value as an int64.
	ParseInt64() (int64, error)
	// ParseUint parses the value as a uint.
	ParseUint() (uint, error)
	// ParseUint64 parses the value as a uint64.
	ParseUint64() (uint64, error)
	// ParseFloat64 parses the value as a float64.
	ParseFloat64() (float64, error)
	// ParseBool parses the value as a bool.
	ParseBool() (bool, error)
	// ParseDuration parses the value as a time.Duration.
	ParseDuration() (time.Duration, error)
	// ParseTime parses the value as a time.Time.
	ParseTime() (time.Time, error)
}
