package validator_test

import (
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func ExampleNotEqual() {
	id := option.Int("id", "the id")

	show(validator.NotEqual(0)(id, value.New(1)))
	show(validator.NotEqual(0)(id, value.New(0)))
	// Output:
	// ok
	// invalid
}
