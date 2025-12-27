package value_test

import (
	"encoding/json"
	"testing"
	"time"

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

func TestJBytes_Duration(t *testing.T) {
	t.Parallel()

	ex := 42 * time.Minute
	data, err := json.Marshal("42m")
	require.NoError(t, err)

	valueDuration := value.JBytes(data)

	var exp time.Duration

	require.ErrorIs(t, valueDuration.Unmarshal(&exp), config.ErrInvalidValue)

	pdur, perr := valueDuration.ParseDuration()
	require.NoError(t, perr)
	require.Equal(t, pdur, ex)

	jdata, jerr := json.Marshal(ex)
	require.NoError(t, jerr)

	jDuration := value.JBytes(jdata)

	var jexp time.Duration
	require.NoError(t, jDuration.Unmarshal(&jexp))
	require.Equal(t, jexp, ex)

	jpdur, jperr := jDuration.ParseDuration()
	require.NoError(t, jperr)
	require.Equal(t, jpdur, ex)
}
