package validator_test

import (
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func ExampleJSON() {
	auth := option.String("auth", "the auth json")

	show(validator.JSON(auth, value.New(`{"a":1}`)))
	show(validator.JSON(auth, value.New("{")))
	// Output:
	// ok
	// invalid
}
