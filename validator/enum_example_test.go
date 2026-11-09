package validator_test

import (
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func ExampleEnum() {
	level := option.String("level", "the log level")
	check := validator.Enum("info", "warn", "error")

	show(check(level, value.New("info")))
	show(check(level, value.New("debug")))
	// Output:
	// ok
	// invalid
}
