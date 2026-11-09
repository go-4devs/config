package example

import (
	"context"

	configs "gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/generate/view"
	"gitoa.ru/go-4devs/config/definition/group"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/definition/proto"
)

//go:generate go run ../../../cmd/config/main.go config:generate

// User returns the user settings.
type User interface {
	// Name returns the user name.
	Name(ctx context.Context) string
	// ReadName reads the user name.
	ReadName(ctx context.Context) (string, error)
	// Pass returns the user password.
	Pass(ctx context.Context) string
	// ReadPass reads the user password.
	ReadPass(ctx context.Context) (string, error)
}

type Level string

// UnmarshalText sets the log level from text.
func (l *Level) UnmarshalText(in []byte) error {
	data := string(in)
	*l = Level(data)

	return nil
}

func Example(_ context.Context, def configs.Definition) error {
	def.Add(
		option.String("test", "the test string", view.WithSkipContext),
		group.New("user", "the user settings",
			option.String("name", "the user name", option.Default("4devs")),
			option.String("pass", "the user password"),
		).With(view.WithContext, view.WithInterface[User]()),

		group.New("log", "the logger settings",
			option.New("level", "the log level", Level(""), option.Default(Level("info"))),
			proto.New("service", "the logger service", option.New("level", "the log level", Level(""))),
		),
	)

	return nil
}
