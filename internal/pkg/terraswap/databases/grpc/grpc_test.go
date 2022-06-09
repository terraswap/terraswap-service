package grpc

import (
	"testing"

	ibctype "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"github.com/delight-labs/terraswap-service/configs"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap"
	"github.com/stretchr/testify/assert"
)

var tsCon = New("terra-phoenix.delightlabs.io:9090", "phoenix-1", configs.LogConfig{})

func TestGetPairs(t *testing.T) {
	res, err := tsCon.GetPairs(terraswap.Pair{})
	assert.NoError(t, err)
	assert.Greater(t, len(res), 1)
}

func TestGetIBCDenom(t *testing.T) {
	res, err := tsCon.GetIBCDenom("ibc/0471F1C4E7AFD3F07702BEF6DC365268D64570F7C1FDC98EA6098DD6DE59817B")
	assert.NoError(t, err)

	assert.Equal(t, &ibctype.DenomTrace{
		Path:      "transfer/channel-1",
		BaseDenom: "uosmo",
	}, res)
}

func TestGetTokenInfo(t *testing.T) {
	res, err := tsCon.GetTokenInfo("terra14xsm2wzvu7xaf567r693vgfkhmvfs08l68h4tjj5wjgyn5ky8e2qvzyanh")
	assert.NoError(t, err)
	assert.Equal(t, 6, res.Decimals)
	assert.Equal(t, "Stader LunaX Token", res.Name)
	assert.Equal(t, "LunaX", res.Symbol)
	assert.Equal(t, 6, res.Decimals)
	assert.Greater(t, res.TotalSupply, "0")
}

func TestGetPairInfo(t *testing.T) {
	pair := terraswap.Pair{}
	pair.ContractAddr = "terra1mecfcj3fkmsgxqa4eaq5w285u6cn0wqtwzkp9gfhpm3dyt8e3cesrpg5hq"

	res, err := tsCon.GetZeroPoolPairs([]terraswap.Pair{pair})
	assert := assert.New(t)
	assert.NoError(err)
	assert.Len(res, 1, "returned map should be 1")
}
