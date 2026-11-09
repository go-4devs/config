package validator_test

import (
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func ExampleMinLen() {
	name := option.String("name", "the user name")

	show(validator.MinLen(3)(name, value.New("ada")))
	show(validator.MinLen(3)(name, value.New("al")))
	// Output:
	// ok
	// invalid
}

func ExampleMaxLen() {
	code := option.String("code", "the code")

	show(validator.MaxLen(2)(code, value.New("йй")))
	show(validator.MaxLen(2)(code, value.New("abc")))
	// Output:
	// ok
	// invalid
}
