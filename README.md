# config

![Build Status](https://github.com/go-4devs/config/actions/workflows/goaction.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-4devs/config)](https://goreportcard.com/report/github.com/go-4devs/config)
[![GoDoc](https://godoc.org/github.com/go-4devs/config?status.svg)](http://godoc.org/github.com/go-4devs/config)

`config` reads one value from several providers. `New` keeps them in the given order. The next provider is used when the current one returns not found.

Providers include command-line flags, environment variables, JSON, YAML, TOML, INI, dasel, etcd and Vault.

## Example

`listen` comes from the environment. `--host` replaces `FDEVS_CONFIG_HOST`.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/provider/arg"
	"github.com/go-4devs/config/provider/env"
)

func main() {
	ctx := context.Background()
	_ = os.Setenv("FDEVS_CONFIG_LISTEN", "8080")
	_ = os.Setenv("FDEVS_CONFIG_HOST", "localhost")

	cfg, err := config.New(
		arg.New(arg.WithArgs([]string{"--host=gitoa.ru"})),
		env.New("fdevs", "config"),
	)
	if err != nil {
		log.Fatal(err)
	}

	port, err := cfg.Value(ctx, "listen")
	if err != nil {
		log.Fatal(err)
	}

	host, err := cfg.Value(ctx, "host")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("listen from env: %d\n", port.Int())
	fmt.Printf("replace env host by args: %v\n", host.String())
}
```

```
listen from env: 8080
replace env host by args: gitoa.ru
```
