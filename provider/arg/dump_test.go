package arg_test

import (
	"bytes"
	"testing"

	"gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/config/test/require"
)

func TestDumpReference(t *testing.T) {
	t.Parallel()

	//nolint:dupword
	const expect = `
Arguments:
  config                                             config [default:config.hcl]
  user-name                                          username

Options:
      --end-after[=END-AFTER]                        after  (multiple values allowed)
      --end-{service}-after[=END-{SERVICE}-AFTER]    after
  -l, --listen[=LISTEN]                              listen [default:8080]
      --start-at[=START-AT]                          start at [default:2010-01-02T15:04:05Z]
  -t, --timeout[=TIMEOUT]                            timeout  (multiple values allowed)
  -u, --url[=URL]                                    url  (multiple values allowed)
  -p, --user-password[=USER-PASSWORD]                user pass
`

	dump := arg.NewDump()

	var buff bytes.Buffer
	require.NoError(t, dump.Reference(&buff, testOptions(t)))
	require.Equal(t, expect, buff.String())
}
