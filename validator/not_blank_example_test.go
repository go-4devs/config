package validator_test

import (
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func ExampleNotBlank() {
	name := option.String("name", "the user name")

	show(validator.NotBlank(name, value.New("ada")))
	show(validator.NotBlank(name, value.New("")))
	// Output:
	// ok
	// invalid
}
