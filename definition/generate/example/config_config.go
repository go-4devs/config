// Code generated gitoa.ru/go-4devs/config DO NOT EDIT.
package example

import (
	context "context"
	fmt "fmt"
	config "gitoa.ru/go-4devs/config"
)

func WithInputConfigLog(log func(context.Context, string, ...any)) func(*InputConfig) {
	return func(ci *InputConfig) {
		ci.log = log
	}
}

func NewInputConfig(prov config.Provider, opts ...func(*InputConfig)) InputConfig {
	i := InputConfig{
		Provider: prov,
		log: func(_ context.Context, format string, args ...any) {
			fmt.Printf(format, args...)
		},
	}

	for _, opt := range opts {
		opt(&i)
	}

	return i
}

type InputConfig struct {
	config.Provider
	log func(context.Context, string, ...any)
}

// readTest test string.
func (i InputConfig) readTest(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "test")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"test"}, err)

	}

	return val.ParseString()

}

// ReadTest test string.
func (i InputConfig) ReadTest() (string, error) {
	return i.readTest(context.Background())
}

// Test test string.
func (i InputConfig) Test() string {
	val, err := i.readTest(context.Background())
	if err != nil {
		i.log(context.Background(), "get [%v]: %v", []string{"test"}, err)
	}

	return val
}

type InputConfigUser struct {
	InputConfig
}

// User configure user.
func (i InputConfig) User() InputConfigUser {
	return InputConfigUser{i}
}

// readName name.
func (i InputConfigUser) readName(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "user", "name")
	if err != nil {
		i.log(context.Background(), "read [%v]: %v", []string{"user", "name"}, err)

		return "4devs", nil
	}

	return val.ParseString()

}

// ReadName name.
func (i InputConfigUser) ReadName(ctx context.Context) (string, error) {
	return i.readName(ctx)
}

// Name name.
func (i InputConfigUser) Name(ctx context.Context) string {
	val, err := i.readName(ctx)
	if err != nil {
		i.log(ctx, "get [%v]: %v", []string{"user", "name"}, err)
	}

	return val
}

// readPass password.
func (i InputConfigUser) readPass(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "user", "pass")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"user", "pass"}, err)

	}

	return val.ParseString()

}

// ReadPass password.
func (i InputConfigUser) ReadPass(ctx context.Context) (string, error) {
	return i.readPass(ctx)
}

// Pass password.
func (i InputConfigUser) Pass(ctx context.Context) string {
	val, err := i.readPass(ctx)
	if err != nil {
		i.log(ctx, "get [%v]: %v", []string{"user", "pass"}, err)
	}

	return val
}

type InputConfigLog struct {
	InputConfig
}

// Log configure logger.
func (i InputConfig) Log() InputConfigLog {
	return InputConfigLog{i}
}

// readLevel log level.
func (i InputConfigLog) readLevel(ctx context.Context) (v Level, e error) {
	val, err := i.Value(ctx, "log", "level")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"log", "level"}, err)

	}

	pval, perr := val.ParseString()
	if perr != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"log", "level"}, perr)
	}

	return v, v.UnmarshalText([]byte(pval))
}

// ReadLevel log level.
func (i InputConfigLog) ReadLevel(ctx context.Context) (Level, error) {
	return i.readLevel(ctx)
}

// Level log level.
func (i InputConfigLog) Level(ctx context.Context) Level {
	val, err := i.readLevel(ctx)
	if err != nil {
		i.log(ctx, "get [%v]: %v", []string{"log", "level"}, err)
	}

	return val
}

type InputConfigLogService struct {
	InputConfigLog
	service string
}

// Service servise logger.
func (i InputConfigLog) Service(key string) InputConfigLogService {
	return InputConfigLogService{i, key}
}

// readLevel log level.
func (i InputConfigLogService) readLevel(ctx context.Context) (v Level, e error) {
	val, err := i.Value(ctx, "log", i.service, "level")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"log", i.service, "level"}, err)

	}

	pval, perr := val.ParseString()
	if perr != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"log", i.service, "level"}, perr)
	}

	return v, v.UnmarshalText([]byte(pval))
}

// ReadLevel log level.
func (i InputConfigLogService) ReadLevel(ctx context.Context) (Level, error) {
	return i.readLevel(ctx)
}

// Level log level.
func (i InputConfigLogService) Level(ctx context.Context) Level {
	val, err := i.readLevel(ctx)
	if err != nil {
		i.log(ctx, "get [%v]: %v", []string{"log", i.service, "level"}, err)
	}

	return val
}
