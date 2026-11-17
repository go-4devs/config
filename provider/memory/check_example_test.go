package memory_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-4devs/config/definition"
	"github.com/go-4devs/config/definition/group"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/provider/memory"
)

func ExampleCheckPrefix() {
	err := memory.CheckPrefix(func(_ context.Context, def *definition.Definition) {
		def.Add(
			option.String("port", "port").With(memory.Prefix("extra", "node")),
			group.New("db", "db", option.String("port", "port")).With(memory.Prefix("extra")),
		)
	}, memory.ParamPrefix)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ok")
	// Output:
	// ok
}

func ExampleCheckPrefix_conflict() {
	err := memory.CheckPrefix(func(_ context.Context, def *definition.Definition) {
		def.Add(
			option.String("port", "port").With(memory.Prefix("db")),
			group.New("db", "db", option.String("port", "port")),
		)
	}, memory.ParamPrefix)
	if errors.Is(err, memory.ErrConflict) {
		fmt.Println(err)
	}
	// Output:
	// prefix conflict: port and db/port -> db/port
}
