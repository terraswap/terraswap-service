package terraswap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tokensMap TokensMap = map[string]Token{
	"a": {
		Protocol:     "protocolA",
		Symbol:       "symbolA",
		ContractAddr: "tokenA",
		Icon:         "iconA",
	},
	"b": {
		Protocol:     "protocolB",
		Symbol:       "symbolB",
		ContractAddr: "tokenB",
		Icon:         "iconB",
	},
	"c": {
		Protocol:     "protocolC",
		Symbol:       "symbolC",
		ContractAddr: "tokenC",
		Icon:         "iconC",
	},
}

func Test_TokenMapEqual(t *testing.T) {
	next := map[string]Token{
		"a": {
			Protocol:     "protocolA",
			Symbol:       "symbolA",
			ContractAddr: "tokenA",
			Icon:         "iconA",
		},
		"b": {
			Protocol:     "protocolB",
			Symbol:       "symbolB",
			ContractAddr: "tokenB",
			Icon:         "iconB",
		},
		"c": {
			Protocol:     "protocolC",
			Symbol:       "symbolC",
			ContractAddr: "tokenC",
			Icon:         "iconC",
		},
		"d": {
			Protocol:     "protocolD",
			Symbol:       "symbolD",
			ContractAddr: "tokenD",
			Icon:         "iconD",
		},
	}
	assert := assert.New(t)
	assert.False(tokensMap.Equal(next))
	assert.True(tokensMap.Equal(tokensMap))
}

func TestGetDiffSame(t *testing.T) {
	assert := assert.New(t)

	next := map[string]Token{
		"a": {
			Protocol:     "protocolA",
			Symbol:       "symbolA",
			ContractAddr: "tokenA",
			Icon:         "iconA",
		},
		"b": {
			Protocol:     "protocolB",
			Symbol:       "symbolB",
			ContractAddr: "tokenB",
			Icon:         "iconB",
		},
		"c": {
			Protocol:     "protocolC",
			Symbol:       "symbolC",
			ContractAddr: "tokenC",
			Icon:         "iconC",
		},
	}

	addWhiteList, removeWhiteList := tokensMap.GetDiffMap(next)
	assert.Equal(0, len(addWhiteList))
	assert.Equal(0, len(removeWhiteList))
}

func TestGetDiffAdd(t *testing.T) {
	assert := assert.New(t)

	var next = TokensMap{
		"a": {
			Protocol:     "protocolA",
			Symbol:       "symbolA",
			ContractAddr: "tokenA",
			Icon:         "iconA",
		},
		"b": {
			Protocol:     "protocolB",
			Symbol:       "symbolB",
			ContractAddr: "tokenB",
			Icon:         "iconB",
		},
		"c": {
			Protocol:     "protocolC",
			Symbol:       "symbolC",
			ContractAddr: "tokenC",
			Icon:         "iconC",
		},
		"f": {
			Protocol:     "protocolF",
			Symbol:       "symbolF",
			ContractAddr: "tokenF",
			Icon:         "iconF",
		},
	}

	addWhiteList, removeWhiteList := tokensMap.GetDiffMap(next)
	assert.Equal(1, len(addWhiteList))
	assert.Equal(0, len(removeWhiteList))
	assert.Equal("protocolF", addWhiteList["f"].Protocol)
	assert.Equal("symbolF", addWhiteList["f"].Symbol)
	assert.Equal("tokenF", addWhiteList["f"].ContractAddr)
	assert.Equal("iconF", addWhiteList["f"].Icon)
}

func TestGetDiffRemove(t *testing.T) {
	assert := assert.New(t)
	next := TokensMap{
		"a": {
			Protocol:     "protocolA",
			Symbol:       "symbolA",
			ContractAddr: "tokenA",
			Icon:         "iconA",
		},
		"b": {
			Protocol:     "protocolB",
			Symbol:       "symbolB",
			ContractAddr: "tokenB",
			Icon:         "iconB",
		},
	}

	addWhiteList, removeWhiteList := tokensMap.GetDiffMap(next)
	assert.Equal(0, len(addWhiteList))
	assert.Equal(1, len(removeWhiteList))

	assert.Equal("protocolC", removeWhiteList["c"].Protocol)
	assert.Equal("symbolC", removeWhiteList["c"].Symbol)
	assert.Equal("tokenC", removeWhiteList["c"].ContractAddr)
	assert.Equal("iconC", removeWhiteList["c"].Icon)
}

func TestGetDiffChangeAll(t *testing.T) {
	assert := assert.New(t)

	next := TokensMap{
		"a": {
			Protocol:     "protocolA",
			Symbol:       "symbolA",
			ContractAddr: "tokenA",
			Icon:         "iconA",
		},
		"b": {
			Protocol:     "protocol_b",
			Symbol:       "symbol_b",
			ContractAddr: "token_b",
			Icon:         "icon_b",
		},
		"f": {
			Protocol:     "protocolF",
			Symbol:       "symbolF",
			ContractAddr: "tokenF",
			Icon:         "iconF",
		},
	}

	addWhiteList, removeWhiteList := tokensMap.GetDiffMap(next)
	assert.Equal(2, len(addWhiteList))
	assert.Equal(2, len(removeWhiteList))

	assert.Equal("protocolF", addWhiteList["f"].Protocol)
	assert.Equal("symbolF", addWhiteList["f"].Symbol)
	assert.Equal("tokenF", addWhiteList["f"].ContractAddr)
	assert.Equal("iconF", addWhiteList["f"].Icon)

	assert.Equal("protocol_b", addWhiteList["b"].Protocol)
	assert.Equal("symbol_b", addWhiteList["b"].Symbol)
	assert.Equal("token_b", addWhiteList["b"].ContractAddr)
	assert.Equal("icon_b", addWhiteList["b"].Icon)

	_, exist := removeWhiteList["b"]
	assert.True(exist)
	_, exist = removeWhiteList["c"]
	assert.True(exist)
}
