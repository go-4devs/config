package validator_test

import (
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func ExampleRegex() {
	name := option.String("name", "the user name")
	check := validator.Regex(`[a-z]+`)

	show(check(name, value.New("ada")))
	show(check(name, value.New("ada1")))
	// Output:
	// ok
	// invalid
}
