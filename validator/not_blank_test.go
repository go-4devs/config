package validator_test

import (
	"errors"
	"testing"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func TestNotBlank(t *testing.T) {
	t.Parallel()

	for name, ca := range casesNotBlank() {
		valid := validator.NotBlank

		if err := valid(ca.vars, ca.value); err != nil {
			t.Errorf("case: %s, expected error <nil>, got: %s", name, err)
		}

		if ca.empty == nil {
			ca.empty = value.EmptyValue()
		}

		emptErr := valid(ca.vars, ca.empty)
		if emptErr == nil || !errors.Is(emptErr, validator.ErrNotBlank) {
			t.Errorf("case empty: %s, expect: %s, got:%v", name, validator.ErrNotBlank, emptErr)
		}
	}
}

type cases struct {
	vars  config.Option
	value config.Value
	empty config.Value
}

func casesNotBlank() map[string]cases {
	return map[string]cases{
		"any": {
			vars:  option.New("any", "any", nil),
			value: value.New(float32(1)),
		},
		"array int": {
			vars:  option.Int("int", "array int", option.Slice),
			value: value.New([]int{1}),
			empty: value.New([]int{}),
		},
		"array int64": {
			vars:  option.Int64("int64", "array int64", option.Slice),
			value: value.New([]int64{1}),
			empty: value.New([]int64{}),
		},
		"array uint": {
			vars:  option.Uint("uint", "array uint", option.Slice),
			value: value.New([]uint{1}),
			empty: value.New([]uint{}),
		},
		"array uint64": {
			vars:  option.Uint64("uint64", "array uint64", option.Slice),
			value: value.New([]uint64{1}),
			empty: value.New([]uint64{}),
		},
		"array float64": {
			vars:  option.Float64("float64", "array float64", option.Slice),
			value: value.New([]float64{0.2}),
			empty: value.New([]float64{}),
		},
		"array bool": {
			vars:  option.Bool("bool", "array bool", option.Slice),
			value: value.New([]bool{true, false}),
			empty: value.New([]bool{}),
		},
		"array duration": {
			vars:  option.Duration("duration", "array duration", option.Slice),
			value: value.New([]time.Duration{time.Second}),
			empty: value.New([]time.Duration{}),
		},
		"array time": {
			vars:  option.Time("time", "array time", option.Slice),
			value: value.New([]time.Time{time.Now()}),
			empty: value.New([]time.Time{}),
		},
		"array string": {
			vars:  option.String("string", "array string", option.Slice),
			value: value.New([]string{"value"}),
			empty: value.New([]string{}),
		},
		"int": {
			vars:  option.Int("int", "int"),
			value: value.New(int(1)),
		},
		"int64": {
			vars:  option.Int64("int64", "int64"),
			value: value.New(int64(2)),
		},
		"uint": {
			vars:  option.Uint("uint", "uint"),
			value: value.New(uint(1)),
			empty: value.New([]uint{1}),
		},
		"uint64": {
			vars:  option.Uint64("uint64", "uint64"),
			value: value.New(uint64(10)),
		},
		"float64": {
			vars:  option.Float64("float64", "float64"),
			value: value.New(float64(.00001)),
		},
		"duration": {
			vars:  option.Duration("duration", "duration"),
			value: value.New(time.Minute),
			empty: value.New("same string"),
		},
		"time": {
			vars:  option.Time("time", "time"),
			value: value.New(time.Now()),
		},
		"string": {
			vars:  option.String("string", "string"),
			value: value.New("string"),
			empty: value.New(""),
		},
	}
}
