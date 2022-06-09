package terraswap

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {

	type testcase struct {
		input    SwapSendMsg
		expected string
		err      error
	}

	tcs := [...]testcase{
		{
			expected: "eyJzd2FwIjp7fX0=",
			err:      nil,
		},
	}

	assert := assert.New(t)
	for _, c := range tcs {
		_, _ = json.Marshal(&c.input)

		n, err := GetSwapSendMsg("", "")

		if c.err != nil {
			assert.Error(err)
			assert.Empty(n)
			continue
		}
		data, _ := base64.StdEncoding.DecodeString(c.expected)
		var expected SwapSendMsg
		json.Unmarshal(data, &expected)

		assert.Nil(err)
		assert.Exactly(expected, n, "must be same")
	}
}
func TestToDenomSymbol(t *testing.T) {
	assert := assert.New(t)

	res := ToDenomSymbol("uluna")
	assert.Equal("Luna", res)

	res = ToDenomSymbol("ukrw")
	assert.Equal("KRW", res)

	res = ToDenomSymbol("uusd")
	assert.Equal("USD", res)
}
