package validator_test

import (
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func ExampleHostname() {
	host := option.String("host", "the server hostname")

	show(validator.Hostname(host, value.New("db.internal")))
	show(validator.Hostname(host, value.New("db.internal:5432")))
	// Output:
	// ok
	// invalid
}
