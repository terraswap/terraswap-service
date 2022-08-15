package terraswap

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumber_NewPositiveNumberFromString(t *testing.T) {

	type testcase struct {
		input    string
		expected string
		decimals int
		err      error
	}

	tcs := [...]testcase{
		{
			input:    "0",
			decimals: 10,
			expected: "0",
			err:      nil,
		},
		{
			input:    "0.1",
			decimals: 6,
			expected: "100000",
			err:      nil,
		},
		{
			input:    "1",
			expected: "1000000",
			decimals: 6,
			err:      nil,
		},
		{
			input:    "-0",
			decimals: 6,
			expected: "0",
			err:      errors.New("should throw error due to negative number"),
		},

		{
			input:    "-0.1",
			decimals: 10,
			expected: "0",
			err:      errors.New("should throw error due to negative number"),
		},

		{
			input:    "-1",
			expected: "0",
			decimals: 10,
			err:      errors.New("should throw error due to negative number"),
		},
		{
			input:    "1",
			expected: "10000000000",
			decimals: 10,
			err:      nil,
		},
		{
			input:    "1",
			expected: "1",
			decimals: 0,
			err:      nil,
		},
		{
			input:    "1.0",
			expected: "",
			decimals: 0,
			err:      errors.New("should throw error due to negative number"),
		},
	}

	assert := assert.New(t)
	for _, c := range tcs {
		n, err := ToTerraAmount(c.input, c.decimals)

		if c.err != nil {
			assert.Error(err)
			assert.Empty(n)
			continue
		}
		assert.Nil(err)
		assert.Exactly(c.expected, n, "must be same")
	}
}
