package validator_test

import (
	"context"
	"os"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/provider/memory"
	"github.com/go-4devs/config/validator"
)

func ExampleFileExists() {
	file, err := os.CreateTemp("", "config-file-*")
	if err != nil {
		panic(err)
	}

	name := file.Name()
	if err = file.Close(); err != nil {
		panic(err)
	}

	defer os.Remove(name)

	ctx := context.Background()
	store := memory.NewMap("app")
	store.SetOption(name, "path")

	vars := config.NewVars(
		option.String("path", "the file path", validator.Valid(validator.FileExists)),
	)
	show(memory.Valid(store).Bind(ctx, vars))

	store.SetOption(name+"-missing", "path")
	show(memory.Valid(store).Bind(ctx, vars))
	// Output:
	// ok
	// invalid
}
