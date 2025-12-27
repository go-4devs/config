package value_test

import (
	"encoding/json"
	"testing"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/value"
)

func TestJBytes_String(t *testing.T) {
	t.Parallel()

	data := value.JBytes([]byte("\"data\""))
	res, err := data.ParseString()
	require.NoError(t, err)
	require.Equal(t, res, "data")

	dataErr := value.JBytes([]byte("data"))
	res2, err2 := dataErr.ParseString()
	require.ErrorIs(t, err2, config.ErrInvalidValue)
	require.Equal(t, "", res2)
}

func TestJBytes_Unmarshal(t *testing.T) {
	t.Parallel()

	data, err := json.Marshal([]string{"test1", "test2"})
	require.NoError(t, err)

	sta := value.JBytes(data)

	ac := ""

	aca := []string{}
	require.NoError(t, sta.Unmarshal(&aca))
	require.Equal(t, []string{"test1", "test2"}, aca)

	require.ErrorIs(t, sta.Unmarshal(ac), config.ErrInvalidValue)
	require.ErrorIs(t, sta.Unmarshal(&ac), config.ErrInvalidValue)
}
