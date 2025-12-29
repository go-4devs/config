package group_test

import (
	"testing"

	"gitoa.ru/go-4devs/config/definition/group"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/test/require"
)

func TestGroupAdd(t *testing.T) {
	t.Parallel()

	var gr group.Group

	gr.Add(group.New("test", "test"))
	require.Truef(t, len(gr.Options()) == 1, "len(%v) != 1", len(gr.Options()))
}

func TestGroupWith(t *testing.T) {
	t.Parallel()

	const descrition = "group description"

	gr := group.New("test", "test desc")
	gr = gr.With(option.Description(descrition))

	require.Equal(t, descrition, option.DataDescription(gr))
}
