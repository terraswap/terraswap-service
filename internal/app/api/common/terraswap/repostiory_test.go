package terraswap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCw20Allowlisted(t *testing.T) {
	assert := assert.New(t)
	r := repositoryImpl{chainId: "phoenix-1"}

	res := r.getCw20Allowlist("https://raw.githubusercontent.com/terra-money/assets/master/cw20/tokens.js")
	assert.Equal("terra14xsm2wzvu7xaf567r693vgfkhmvfs08l68h4tjj5wjgyn5ky8e2qvzyanh", res["terra14xsm2wzvu7xaf567r693vgfkhmvfs08l68h4tjj5wjgyn5ky8e2qvzyanh"].ContractAddr)
	assert.Equal("LunaX", res["terra14xsm2wzvu7xaf567r693vgfkhmvfs08l68h4tjj5wjgyn5ky8e2qvzyanh"].Symbol)
	assert.Equal("Stader", res["terra14xsm2wzvu7xaf567r693vgfkhmvfs08l68h4tjj5wjgyn5ky8e2qvzyanh"].Protocol)
	assert.Equal("https://raw.githubusercontent.com/stader-labs/assets/main/terra/LunaX_1.png", res["terra14xsm2wzvu7xaf567r693vgfkhmvfs08l68h4tjj5wjgyn5ky8e2qvzyanh"].Icon)
}

func TestGetIBCAllowlisted(t *testing.T) {
	assert := assert.New(t)
	r := repositoryImpl{chainId: "phoenix-1"}

	res := r.getIbcAllowlist("https://raw.githubusercontent.com/terra-money/assets/master/ibc/tokens.js")
	assert.Equal("ibc/0471F1C4E7AFD3F07702BEF6DC365268D64570F7C1FDC98EA6098DD6DE59817B", res["ibc/0471F1C4E7AFD3F07702BEF6DC365268D64570F7C1FDC98EA6098DD6DE59817B"].ContractAddr)
	assert.Equal("OSMO", res["ibc/0471F1C4E7AFD3F07702BEF6DC365268D64570F7C1FDC98EA6098DD6DE59817B"].Symbol)
	assert.Equal("https://assets.terra.money/icon/svg/ibc/OSMO.svg", res["ibc/0471F1C4E7AFD3F07702BEF6DC365268D64570F7C1FDC98EA6098DD6DE59817B"].Icon)
	assert.Equal("Osmosis", res["ibc/0471F1C4E7AFD3F07702BEF6DC365268D64570F7C1FDC98EA6098DD6DE59817B"].Name)
}
