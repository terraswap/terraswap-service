package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/terraswap/terraswap-service/configs"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

func TestClassicGetPairs(t *testing.T) {
	classicCon := NewClassic("terra.delightlabs.io:9090", "columbus-5", "v1", configs.LogConfig{})
	res, err := classicCon.GetPairs(terraswap.Pair{})
	assert.NoError(t, err)
	assert.Greater(t, len(res), 1)
}

func TestClassicActiveDenoms(t *testing.T) {
	var tsCon = NewClassic("terra.delightlabs.io:9090", "columbus-5", "v1", configs.LogConfig{})
	res, err := tsCon.GetDenoms()
	assert.NoError(t, err)
	assert.NotZero(t, res)

}
func TestClassicGetIbcDenom(t *testing.T) {
	var tsCon = NewClassic("terra.delightlabs.io:9090", "columbus-5", "v1", configs.LogConfig{})
	res, err := tsCon.GetIbcDenom("ibc/18ABA66B791918D51D33415DA173632735D830E2E77E63C91C11D3008CFD5262")
	assert.NoError(t, err)

	assert.Equal(t, &terraswap.IbcDenomTrace{
		Path:      "transfer/channel-41",
		BaseDenom: "uatom",
	}, res)
}

func TestClassicGetTokenInfo(t *testing.T) {
	var tsCon = NewClassic("terra.delightlabs.io:9090", "columbus-5", "v1", configs.LogConfig{})
	res, err := tsCon.GetTokenInfo("terra1kc87mu460fwkqte29rquh4hc20m54fxwtsx7gp")
	assert.NoError(t, err)
	assert.Equal(t, 6, res.Decimals)
	assert.Equal(t, "BLUNA", res.Symbol)
	assert.Equal(t, 6, res.Decimals)
	assert.Greater(t, res.TotalSupply, "0")
}
