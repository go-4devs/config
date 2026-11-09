package validator_test

import (
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func ExampleNotBlank() {
	name := option.String("name", "the user name")

	show(validator.NotBlank(name, value.New("ada")))
	show(validator.NotBlank(name, value.New("")))
	// Output:
	// ok
	// invalid
}
