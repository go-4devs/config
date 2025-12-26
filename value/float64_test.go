package value_test

import (
	"encoding/json"
	"math"
	"testing"

	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/value"
)

func TestFloat64_Unmarshal(t *testing.T) {
	t.Parallel()

	f := value.Float64(math.Pi)

	var out float64

	require.NoError(t, f.Unmarshal(&out))
	require.Equal(t, math.Pi, out)
}

func TestFloat64_Any(t *testing.T) {
	t.Parallel()

	f := value.Float64(math.Pi)

	require.Equal(t, math.Pi, f.Any(), 0)
}

func TestFloat64s_Unmarshal(t *testing.T) {
	t.Parallel()

	data, err := json.Marshal([]float64{math.Pi, math.Sqrt2})
	require.NoError(t, err)

	f := value.JBytes(data)

	var out []float64

	require.NoError(t, f.Unmarshal(&out))
	require.Equal(t, []float64{math.Pi, math.Sqrt2}, out)
}

func TestFloat64s_Any(t *testing.T) {
	t.Parallel()

	f := value.New([]float64{math.Pi, math.Sqrt2})

	require.Equal(t, []float64{math.Pi, math.Sqrt2}, f.Any())
}
