package validator_test

import (
	"context"
	"fmt"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/provider/memory"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func ExampleValid() {
	ctx := context.Background()
	store := memory.NewMap("app")
	store.SetOption("user@example.com", "email")
	store.SetOption(8080, "port")
	store.SetOption("", "auth")

	vars := config.NewVars(
		option.String("email", "the user email", validator.Valid(validator.Email)),
		option.Int("port", "the listen port", validator.Valid(validator.Min(1), validator.Max(65535))),
		option.String("auth", "the auth json", validator.Valid(validator.Optional(validator.JSON))),
	)
	show(memory.BindValid(store).Bind(ctx, vars))

	store.SetOption("not-email", "email")
	show(memory.BindValid(store).Bind(ctx, vars))
	// Output:
	// ok
	// invalid
}

func ExampleChain() {
	port := option.Int("port", "the listen port")
	check := validator.Chain(validator.Min(1), validator.Max(65535))

	show(check(port, value.New(8080)))
	show(check(port, value.New(0)))
	// Output:
	// ok
	// invalid
}

func ExampleOptional() {
	auth := option.String("auth", "the auth json")
	check := validator.Optional(validator.JSON)

	show(check(auth, value.New("")))
	show(check(auth, value.New(`{"a":1}`)))
	show(check(auth, value.New("{")))
	// Output:
	// ok
	// ok
	// invalid
}

func ExampleWhen() {
	num := option.Int("n", "the number")
	check := validator.When(
		validator.Chain(validator.Min(1), validator.Max(10)),
		validator.NotEqual(5),
	)

	show(check(num, value.New(0)))
	show(check(num, value.New(4)))
	show(check(num, value.New(5)))
	// Output:
	// ok
	// ok
	// invalid
}

func show(err error) {
	if err != nil {
		fmt.Println("invalid")

		return
	}

	fmt.Println("ok")
}
