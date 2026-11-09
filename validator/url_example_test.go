package validator_test

import (
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func ExampleURL() {
	page := option.String("page", "the page url")

	show(validator.URL(page, value.New("https://example.com/path")))
	show(validator.URL(page, value.New("/path")))
	// Output:
	// ok
	// invalid
}

func ExampleURI() {
	ref := option.String("ref", "the reference")

	show(validator.URI(ref, value.New("/path")))
	show(validator.URI(ref, value.New("example.com")))
	// Output:
	// ok
	// invalid
}
