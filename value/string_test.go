package value_test

import (
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/test/require"
	"github.com/go-4devs/config/value"
)

func TestStringUnmarshal(t *testing.T) {
	t.Parallel()

	st := value.String("test")

	ac := ""
	require.NoError(t, st.Unmarshal(&ac))
	require.Equal(t, "test", ac)

	aca := []string{}

	require.ErrorIs(t, st.Unmarshal(aca), config.ErrWrongType)
	require.ErrorIs(t, st.Unmarshal(&aca), config.ErrWrongType)
}
