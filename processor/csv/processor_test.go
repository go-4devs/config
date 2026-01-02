package csv_test

import (
	"context"
	"testing"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/processor/csv"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/value"
)

func TestCsv(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	val := value.String("test2,test3,other")

	data, derr := csv.Csv(ctx, val)
	require.NoError(t, derr)

	var resString []string

	require.NoError(t, data.Unmarshal(&resString))
	require.Equal(t, []string{"test2", "test3", "other"}, resString)
}

func TestCsv_int(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	val := value.String("42,0,1")

	data, derr := csv.Csv(ctx, val, csv.WithInt)
	require.NoError(t, derr)

	var resInt []int

	require.NoError(t, data.Unmarshal(&resInt))
	require.Equal(t, []int{42, 0, 1}, resInt)
}

func TestCsv_invalidValue(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	val := value.String("42,0.1,1")

	data, derr := csv.Csv(ctx, val, csv.WithInt)
	require.ErrorIs(t, derr, config.ErrInvalidValue)
	require.Equal(t, nil, data)
}
