package validator_test

import (
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func ExampleJSON() {
	auth := option.String("auth", "the auth json")

	show(validator.JSON(auth, value.New(`{"a":1}`)))
	show(validator.JSON(auth, value.New("{")))
	// Output:
	// ok
	// invalid
}
