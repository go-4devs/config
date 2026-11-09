package validator_test

import (
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
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
