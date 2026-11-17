package env_test

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/provider/env"
)

func ExamplePrefixProvider_Value() {
	const key = "FDEVS_CONFIG_EXTRA_PORT"

	_ = os.Setenv(key, "8080")
	defer func() { _ = os.Unsetenv(key) }()

	ctx := context.Background()
	provider := env.NewPrefixProvider(env.New("fdevs", "config"))

	err := provider.Bind(ctx, config.NewVars(
		option.Int("port", "port").With(env.Prefix("extra")),
	))
	if err != nil {
		fmt.Println(err)

		return
	}

	val, err := provider.Value(ctx, "port")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(val.String())
	// Output:
	// 8080
}

func ExamplePrefixProvider_DumpReference() {
	provider := env.NewPrefixProvider(env.New("fdevs", "config"))
	buf := bytes.NewBuffer(nil)

	err := provider.DumpReference(context.Background(), buf, definition.New(
		option.Int("port", "port").With(env.Prefix("extra")),
		option.String("dsn", "dsn"),
	))
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Print(buf.String())
	// Output:
	// # port.
	// #FDEVS_CONFIG_EXTRA_PORT=
	// # dsn.
	// #FDEVS_CONFIG_DSN=
}
