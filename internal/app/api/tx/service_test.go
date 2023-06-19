package tx

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
	"testing"
)

var (
	assetInfoLuna = terraswap.AssetInfo{
		NativeToken: &terraswap.AssetNativeToken{
			Denom: "uluna",
		},
	}
	assetInfoIbc = terraswap.AssetInfo{
		NativeToken: &terraswap.AssetNativeToken{
			Denom: "ibc/3A14BA2D5493AF5632260E354E46004562C46AB7EC0DA4D1DA14E9E20E2545A8",
		},
	}
	assetInfoCw20 = terraswap.AssetInfo{
		Token: &terraswap.AssetCWToken{
			ContractAddr: "terra123",
		},
	}
	pairLunaCw20 = &terraswap.Pair{
		AssetInfos: []terraswap.AssetInfo{assetInfoLuna, assetInfoCw20},
	}
	pairLunaIbc = &terraswap.Pair{
		AssetInfos: []terraswap.AssetInfo{assetInfoLuna, assetInfoIbc},
	}
)

type mixinImplSuite struct {
	suite.Suite

	Service mixinImpl
}

func (s *mixinImplSuite) SetupSuite() {
	s.Service = mixinImpl{}
}

func (s *mixinImplSuite) TestParseMinAssetsWithEmptyStr() {
	assert := assert.New(s.T())

	var expected []terraswap.OfferAsset
	actual, err := s.Service.parseMinAssets(*pairLunaCw20, "")

	assert.NoError(err)
	assert.Equal(expected, actual)
}

func (s *mixinImplSuite) TestParseMinAssets() {
	assert := assert.New(s.T())

	expected := []terraswap.OfferAsset{
		{
			Amount: "100",
			Info:   assetInfoLuna,
		},
		{
			Amount: "100",
			Info:   assetInfoCw20,
		},
	}
	actual, err := s.Service.parseMinAssets(*pairLunaCw20, "100uluna,100terra123")

	assert.NoError(err)
	assert.Equal(expected, actual)
}

func (s *mixinImplSuite) TestParseMinAssetsIbc1() {
	assert := assert.New(s.T())

	expected := []terraswap.OfferAsset{
		{
			Amount: "100",
			Info:   assetInfoLuna,
		},
		{
			Amount: "100",
			Info:   assetInfoIbc,
		},
	}
	actual, err := s.Service.parseMinAssets(*pairLunaIbc, "100uluna,100ibc/3A14BA2D5493AF5632260E354E46004562C46AB7EC0DA4D1DA14E9E20E2545A8")

	assert.NoError(err)
	assert.Equal(expected, actual)
}

func (s *mixinImplSuite) TestParseMinAssetsIbc2() {
	assert := assert.New(s.T())

	expected := []terraswap.OfferAsset{
		{
			Amount: "100",
			Info:   assetInfoIbc,
		},
		{
			Amount: "100",
			Info:   assetInfoLuna,
		},
	}
	actual, err := s.Service.parseMinAssets(*pairLunaIbc, "100ibc/3A14BA2D5493AF5632260E354E46004562C46AB7EC0DA4D1DA14E9E20E2545A8,100uluna")

	assert.NoError(err)
	assert.Equal(expected, actual)
}

func (s *mixinImplSuite) TestParseMinAssetsReverseOrder() {
	assert := assert.New(s.T())

	expected := []terraswap.OfferAsset{
		{
			Amount: "100",
			Info:   assetInfoCw20,
		},
		{
			Amount: "100",
			Info:   assetInfoLuna,
		},
	}
	actual, err := s.Service.parseMinAssets(*pairLunaCw20, "100terra123,100uluna")

	assert.NoError(err)
	assert.Equal(expected, actual)
}

func (s *mixinImplSuite) TestParseMinAssetsUnmatchedPair1() {
	assert := assert.New(s.T())

	_, err := s.Service.parseMinAssets(*pairLunaCw20, "100terra123,100uusd")

	assert.Error(err)
	assert.ErrorContains(err, "unmatched pair")
}

func (s *mixinImplSuite) TestParseMinAssetsUnmatchedPair2() {
	assert := assert.New(s.T())

	_, err := s.Service.parseMinAssets(*pairLunaCw20, "100terra123,100terra123")

	assert.Error(err)
	assert.ErrorContains(err, "minAssets is invalid")
}

func (s *mixinImplSuite) TestParseMinAssetsInvalidValue1() {
	assert := assert.New(s.T())

	_, err := s.Service.parseMinAssets(*pairLunaCw20, "abcd")

	assert.Error(err)
	assert.ErrorContains(err, "minAssets is invalid")
}

func (s *mixinImplSuite) TestParseMinAssetsInvalidValue2() {
	assert := assert.New(s.T())

	_, err := s.Service.parseMinAssets(*pairLunaCw20, "100terra123,")

	assert.Error(err)
	assert.ErrorContains(err, "minAssets is invalid")
}

func (s *mixinImplSuite) TestParseMinAssetsInvalidValue3() {
	assert := assert.New(s.T())

	_, err := s.Service.parseMinAssets(*pairLunaCw20, "100terra123,100")

	assert.Error(err)
	assert.ErrorContains(err, "minAssets is invalid")
}

func TestMixinImplSuite(t *testing.T) {
	suite.Run(t, new(mixinImplSuite))
}
