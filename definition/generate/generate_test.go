package generate_test

import (
	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/definition/generate/view"
	"gitoa.ru/go-4devs/config/definition/group"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/definition/proto"
)

type LogLevel string

func (l *LogLevel) UnmarshalText(in []byte) error {
	data := string(in)
	*l = LogLevel(data)

	return nil
}

func Configure(def *definition.Definition) {
	def.Add(
		option.String("test", "test string", view.WithSkipContext),
		group.New("user", "configure user",
			option.String("name", "name", option.Default("4devs")),
			option.String("pass", "password"),
		),

		group.New("log", "configure logger",
			option.New("level", "log level", LogLevel("")),
			proto.New("service", "servise logger", option.New("level", "log level", LogLevel(""))),
		),
	)
}
