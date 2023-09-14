package terraswap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCw20Allowlisted(t *testing.T) {
	assert := assert.New(t)
	r := repositoryImpl{chainId: "phoenix-1"}

	res, _ := r.getCw20Allowlist("https://assets.terra.dev/cw20/tokens.json")
	assert.Equal("terra14xsm2wzvu7xaf567r693vgfkhmvfs08l68h4tjj5wjgyn5ky8e2qvzyanh", res["terra14xsm2wzvu7xaf567r693vgfkhmvfs08l68h4tjj5wjgyn5ky8e2qvzyanh"].ContractAddr)
	assert.Equal("LunaX", res["terra14xsm2wzvu7xaf567r693vgfkhmvfs08l68h4tjj5wjgyn5ky8e2qvzyanh"].Symbol)
	assert.Equal("Stader", res["terra14xsm2wzvu7xaf567r693vgfkhmvfs08l68h4tjj5wjgyn5ky8e2qvzyanh"].Protocol)
	assert.Equal("https://raw.githubusercontent.com/stader-labs/assets/main/terra/LunaX_1.png", res["terra14xsm2wzvu7xaf567r693vgfkhmvfs08l68h4tjj5wjgyn5ky8e2qvzyanh"].Icon)
}

func TestGetIBCAllowlisted(t *testing.T) {
	assert := assert.New(t)
	r := repositoryImpl{"phoenix-1", nil, &mapperImpl{}}

	res, _ := r.getIbcAllowlist("https://assets.terra.dev/ibc/tokens.json")
	assert.NotNil(res["ibc/0471F1C4E7AFD3F07702BEF6DC365268D64570F7C1FDC98EA6098DD6DE59817B"])
	assert.Equal("OSMO", res["ibc/0471F1C4E7AFD3F07702BEF6DC365268D64570F7C1FDC98EA6098DD6DE59817B"].Symbol)
	assert.Equal("https://assets.terra.dev/icon/svg/ibc/OSMO.svg", res["ibc/0471F1C4E7AFD3F07702BEF6DC365268D64570F7C1FDC98EA6098DD6DE59817B"].Icon)
	assert.Equal("Osmosis", res["ibc/0471F1C4E7AFD3F07702BEF6DC365268D64570F7C1FDC98EA6098DD6DE59817B"].Name)
}
