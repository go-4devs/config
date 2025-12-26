package value_test

import (
	"encoding/json"
	"testing"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/value"
)

func TestStringUnmarshal(t *testing.T) {
	t.Parallel()

	st := value.String("test")

	data, err := json.Marshal([]string{"test1", "test2"})
	require.NoError(t, err)

	sta := value.JBytes(data)

	ac := ""
	require.NoError(t, st.Unmarshal(&ac))
	require.Equal(t, "test", ac)

	aca := []string{}
	require.NoError(t, sta.Unmarshal(&aca))
	require.Equal(t, []string{"test1", "test2"}, aca)

	require.ErrorIs(t, sta.Unmarshal(ac), config.ErrWrongType)
	require.ErrorIs(t, sta.Unmarshal(&ac), config.ErrWrongType)
	require.ErrorIs(t, st.Unmarshal(aca), config.ErrWrongType)
	require.ErrorIs(t, st.Unmarshal(&aca), config.ErrWrongType)
}
