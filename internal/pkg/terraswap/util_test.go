package terraswap

import (
	"encoding/base64"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSwapSendMsg(t *testing.T) {
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

		n, err := GetSwapSendMsg("", "", 0)

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

type getWithdrawSendMsgSuite struct {
	suite.Suite
}

func (s *getWithdrawSendMsgSuite) TestEmptyMsg() {
	assert := assert.New(s.T())

	expected := "eyJ3aXRoZHJhd19saXF1aWRpdHkiOnt9fQ==" // encoded `{"withdraw_liquidity":{}}`
	actual, err := GetWithdrawSendMsg(0)

	assert.NoError(err)
	assert.Equal(expected, actual)
}

func (s *getWithdrawSendMsgSuite) TestDeadline() {
	assert := assert.New(s.T())

	expected := "eyJ3aXRoZHJhd19saXF1aWRpdHkiOnsiZGVhZGxpbmUiOjEyM319" // encoded `{"withdraw_liquidity":{"deadline":123}}`
	actual, err := GetWithdrawSendMsg(123)

	assert.NoError(err)
	assert.Equal(expected, actual)
}

func TestGetWithdrawSendMsg(t *testing.T) {
	suite.Run(t, new(getWithdrawSendMsgSuite))
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
