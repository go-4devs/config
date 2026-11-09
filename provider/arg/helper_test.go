package arg_test

import (
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition"
	"github.com/go-4devs/config/definition/group"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/definition/proto"
	"github.com/go-4devs/config/provider/arg"
	"github.com/go-4devs/config/test"
)

func testOptions(t *testing.T) config.Options {
	t.Helper()

	def := definition.New()
	def.Add(
		option.Int("listen", "listen", option.Short('l'), option.Default(8080)),
		option.String("config", "config", arg.Argument, option.Default("config.hcl")),
		group.New("user", "configure user",
			option.String("name", "username", arg.Argument),
			option.String("password", "user pass", option.Short('p')),
		),
		option.String("url", "url", option.Short('u'), option.Slice),
		option.Duration("timeout", "timeout", option.Short('t'), option.Slice),
		option.Time("start-at", "start at", option.Default(test.Time("2010-01-02T15:04:05Z"))),
		group.New("end", "end at",
			option.Time("after", "after", option.Slice),
			proto.New("service", "service after",
				option.Time("after", "after"),
			),
		),
	)

	return def
}
