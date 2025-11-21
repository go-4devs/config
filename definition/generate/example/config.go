package example

import (
	"context"

	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/definition/generate/view"
	"gitoa.ru/go-4devs/config/definition/group"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/definition/proto"
)

type Level string

func (l *Level) UnmarshalText(in []byte) error {
	data := string(in)
	*l = Level(data)

	return nil
}

func Config(_ context.Context, def *definition.Definition) error {
	def.Add(
		option.String("test", "test string", view.WithSkipContext),
		group.New("user", "configure user",
			option.String("name", "name", option.Default("4devs")),
			option.String("pass", "password"),
		),

		group.New("log", "configure logger",
			option.New("level", "log level", Level("")),
			proto.New("service", "servise logger", option.New("level", "log level", Level(""))),
		),
	)

	return nil
}
