package json_test

import (
	"embed"
	"testing"
	"time"

	"github.com/go-4devs/config/provider/json"
	"github.com/go-4devs/config/test"
	"github.com/go-4devs/config/test/require"
)

//go:embed fixture/*
var fixture embed.FS

func TestProvider(t *testing.T) {
	t.Parallel()

	js, err := fixture.ReadFile("fixture/config.json")
	require.NoError(t, err)

	prov := json.New(js)
	sl := []string{}
	read := []test.Read{
		test.NewRead("config title", "app.name.title"),
		test.NewRead(time.Minute, "app.name.timeout"),
		test.NewReadUnmarshal(&[]string{"name"}, &sl, "app.name.var"),
		test.NewReadConfig("cfg"),
		test.NewRead(true, "app", "name", "success"),
	}

	test.Run(t, prov, read)
}
