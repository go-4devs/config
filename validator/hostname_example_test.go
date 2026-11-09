package validator_test

import (
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func ExampleHostname() {
	host := option.String("host", "the server hostname")

	show(validator.Hostname(host, value.New("db.internal")))
	show(validator.Hostname(host, value.New("db.internal:5432")))
	// Output:
	// ok
	// invalid
}
