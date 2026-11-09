package validator_test

import (
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func ExampleNotEqual() {
	id := option.Int("id", "the id")

	show(validator.NotEqual(0)(id, value.New(1)))
	show(validator.NotEqual(0)(id, value.New(0)))
	// Output:
	// ok
	// invalid
}
