package ini_test

import (
	"embed"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gitoa.ru/go-4devs/config/provider/ini"
	"gitoa.ru/go-4devs/config/test"
	lib "gopkg.in/ini.v1"
)

//go:embed fixture/*
var fixtures embed.FS

func TestProvider(t *testing.T) {
	t.Parallel()

	data, derr := fixtures.ReadFile("fixture/config.ini")
	require.NoError(t, derr)

	file, err := lib.Load(data)
	require.NoError(t, err)

	read := []test.Read{
		test.NewRead("To Do, In Progress, Done", "project", "PROJECT_BOARD_BASIC_KANBAN_TYPE"),
		test.NewRead("markdown", "repository.editor", "PREVIEWABLE_FILE_MODES"),
		test.NewRead("http://0.0.0.0:3000/", "server", "LOCAL_ROOT_URL"),
		test.NewRead(20*time.Minute, "server", "LFS_HTTP_AUTH_EXPIRY"),
		test.NewRead(5120, "repository.pull-request", "DEFAULT_MERGE_MESSAGE_SIZE"),
		test.NewRead(true, "ui", "SHOW_USER_EMAIL"),
		test.NewRead(false, "cors", "enabled"),
	}

	prov := ini.New(file)
	test.Run(t, prov, read)
}
