package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"gitoa.ru/go-4devs/config/definition/generate/command"
	"gitoa.ru/go-4devs/console"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan os.Signal, 1)
	defer close(ch)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-ch
		cancel()
	}()

	console.
		New().
		Add(
			command.Command(),
		).
		Execute(ctx)
}
