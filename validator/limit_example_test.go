package validator_test

import (
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func ExampleMin() {
	port := option.Int("port", "the listen port")

	show(validator.Min(1)(port, value.New(1)))
	show(validator.Min(1)(port, value.New(0)))
	// Output:
	// ok
	// invalid
}

func ExampleMax() {
	port := option.Int("port", "the listen port")

	show(validator.Max(65535)(port, value.New(8080)))
	show(validator.Max(65535)(port, value.New(70000)))
	// Output:
	// ok
	// invalid
}
