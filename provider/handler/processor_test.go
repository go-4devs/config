package handler_test

import (
	"context"
	"testing"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/group"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/definition/proto"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/processor/csv"
	"gitoa.ru/go-4devs/config/provider/handler"
	"gitoa.ru/go-4devs/config/provider/memory"
	"gitoa.ru/go-4devs/config/test/require"
)

var (
	testKey  = []string{"test"}
	testBool = []string{"group", "service", "bool"}
	testInt  = []string{"group", "int"}
)

func TestProcessor(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	prov := handler.Processor(&memory.Default{})
	require.NoError(t, prov.Bind(ctx, testVariables(t)))

	tval, terr := prov.Value(ctx, testKey...)
	require.NoError(t, terr)

	var tdata []string

	require.NoError(t, tval.Unmarshal(&tdata))
	require.Equal(t, []string{"test1", "test2"}, tdata)

	bval, berr := prov.Value(ctx, testBool...)
	require.NoError(t, berr)

	var bdata []bool

	require.NoError(t, bval.Unmarshal(&bdata))
	require.Equal(t, []bool{true, false, true}, bdata)

	ival, ierr := prov.Value(ctx, testInt...)
	require.NoError(t, ierr)

	var idata []int

	require.NoError(t, ival.Unmarshal(&idata))
	require.Equal(t, []int{42, 0, 1}, idata)
}

func testVariables(t *testing.T) config.Variables {
	t.Helper()

	vars := config.NewVars(
		option.String("test", "test",
			option.Slice,
			option.Default("test1,test2"),
			handler.Process(config.ProcessFunc(csv.Csv)),
		),
		group.New("group", "group",
			proto.New("proto", "proto",
				option.Bool("bool", "bool",
					option.Slice,
					option.Default("true|false|true"),
					handler.Process(config.ProcessFunc(func(ctx context.Context, in config.Value, _ ...param.Option) (config.Value, error) {
						return csv.Csv(ctx, in, csv.WithBool, csv.WithDelimiter('|'))
					})),
				),
			),
			option.Int("int", "int",
				option.Slice,
				option.Default("42,0,1"),
				handler.Process(config.ProcessFunc(func(ctx context.Context, in config.Value, _ ...param.Option) (config.Value, error) {
					return csv.Csv(ctx, in, csv.WithInt)
				})),
			),
		),
	)

	return vars
}
