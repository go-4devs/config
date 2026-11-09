package validator_test

import (
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func ExampleEmail() {
	email := option.String("email", "the user email")

	show(validator.Email(email, value.New("user@example.com")))
	show(validator.Email(email, value.New("user")))
	// Output:
	// ok
	// invalid
}
