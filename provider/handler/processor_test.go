package handler_test

import (
	"context"
	"testing"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/group"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/definition/proto"
	"gitoa.ru/go-4devs/config/processor/csv"
	"gitoa.ru/go-4devs/config/provider/handler"
	"gitoa.ru/go-4devs/config/provider/memory"
	"gitoa.ru/go-4devs/config/test/require"
)

var (
	testKey    = []string{"test"}
	testBool   = []string{"group", "service", "bool"}
	testInt    = []string{"group", "int"}
	testTime   = []string{"group", "time"}
	testUint64 = []string{"uint64"}
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
	require.Equal(t, []string{"test1", "test2 data", "test3"}, tdata)

	bval, berr := prov.Value(ctx, testBool...)
	require.NoError(t, berr)

	var bdata []bool

	require.NoError(t, bval.Unmarshal(&bdata))
	require.Equal(t, []bool{true, false, true}, bdata)

	ival, ierr := prov.Value(ctx, testInt...)
	require.NoError(t, ierr)

	var idata []int

	require.NoError(t, ival.Unmarshal(&idata))
	require.Equal(t, []int{-42, 0, 42}, idata)

	tival, tierr := prov.Value(ctx, testTime...)
	require.NoError(t, tierr)

	var tidata []time.Time

	require.NoError(t, tival.Unmarshal(&tidata))
	require.Equal(t, []time.Time{time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)}, tidata)

	uval, uerr := prov.Value(ctx, testUint64...)
	require.NoError(t, uerr)

	var udata []uint64

	require.NoError(t, uval.Unmarshal(&udata))
	require.Equal(t, []uint64{42}, udata)
}

func testVariables(t *testing.T) config.Variables {
	t.Helper()

	vars := config.NewVars(
		option.String("test", "test",
			option.Slice,
			option.Default("test1,\"test2 data\",test3"),
			handler.Process(config.ProcessFunc(csv.Csv)),
		),
		group.New("group", "group",
			proto.New("proto", "proto",
				option.Bool("bool", "bool",
					option.Slice,
					option.Default("true|false|true"),
					handler.FormatFn(csv.Csv, csv.WithBool, csv.WithDelimiter('|')),
				),
			),
			option.Int("int", "int",
				option.Slice,
				option.Default("-42,0,42"),
				handler.FormatFn(csv.Csv, csv.WithInt),
			),
			option.Time("time", "time",
				option.Slice,
				option.Default("2006-01-02T15:04:05Z"),
				handler.FormatFn(csv.Csv, csv.WithTime),
			),
		),
		option.Uint64("uint64", "uint64",
			option.Slice,
			option.Default("42"),
			handler.FormatFn(csv.Csv, csv.WithUint64),
		),
	)

	return vars
}
