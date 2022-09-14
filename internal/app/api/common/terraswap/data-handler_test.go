package terraswap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/terraswap/terraswap-service/configs"
	"github.com/terraswap/terraswap-service/internal/pkg/cache"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"

	tscache "github.com/terraswap/terraswap-service/internal/pkg/terraswap/cache"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap/router"
)

func TestDataHandler_SwapablePairs(t *testing.T) {
	type testcase struct {
		cw20Allowlist     terraswap.TokensMap
		ibcAllowlist      terraswap.TokensMap
		zeroPoolPairCount int
		initialPairs      []terraswap.Pair
	}

	config := configs.Config{
		Terraswap: configs.TerraswapConfig{
			ChainId:          "pisco-1",
			FilterUnverified: false,
		}}
	var repo *repositoryMock
	var c cache.Cache
	var tsCache tscache.TerraswapCache
	var route router.Router

	beforeEach := func(t testcase) {
		repo = newRepositoryMock()
		c = cache.New(config.Cache)
		tsCache = tscache.New(c)
		route = router.NewRouterMock()
		addresses := []string{}
		for _, pair := range t.initialPairs {
			addresses = append(addresses, pair.LiquidityToken)
			for _, asset := range pair.AssetInfos {
				addresses = append(addresses, asset.GetKey())
			}
		}
		for _, addr := range addresses {
			token := terraswap.NewFakeToken()
			token.ContractAddr = addr
			tsCache.SetToken(token)
			repo.On("getToken", addr).Return(&token, nil)
		}
		repo.On("getAllPairs").Return(t.initialPairs, nil).Twice()
		repo.On("getCw20Allowlist", config.Terraswap.Cw20AllowlistUrl).Return(t.cw20Allowlist, nil)
		repo.On("getIbcAllowlist", config.Terraswap.IbcAllowlistUrl).Return(t.ibcAllowlist, nil)
		zeroPoolPairs := make(map[string]bool)
		for i := 0; i < t.zeroPoolPairCount; i++ {
			zeroPoolPairs[t.initialPairs[i].ContractAddr] = true
		}
		repo.On("getZeroPoolPairs", t.initialPairs).Return(zeroPoolPairs, nil).Once()
	}

	testCases := []testcase{
		{
			initialPairs:      terraswap.NewFakePairs(2),
			cw20Allowlist:     terraswap.TokensMap{},
			ibcAllowlist:      terraswap.TokensMap{},
			zeroPoolPairCount: 0,
		},
		{
			initialPairs:      terraswap.NewFakePairs(3),
			cw20Allowlist:     terraswap.TokensMap{},
			ibcAllowlist:      terraswap.TokensMap{},
			zeroPoolPairCount: 1,
		},
	}

	for _, tc := range testCases {
		beforeEach(tc)
		dh := NewDataHandler(repo, route, tsCache, config)
		dh.Run()
		expectedPairs := []terraswap.Pair{}
		for i := tc.zeroPoolPairCount; i < len(tc.initialPairs); i++ {
			expectedPairs = append(expectedPairs, tc.initialPairs[i])
		}
		swapablePairs := dh.GetSwapablePairs()
		assert := assert.New(t)
		assert.Equal(expectedPairs, swapablePairs, "pairs must be the same")
	}

}

func TestDataHandler_SwapablePairsMustChange(t *testing.T) {
	type testcase struct {
		cw20Allowlist     terraswap.TokensMap
		ibcAllowlist      terraswap.TokensMap
		zeroPoolPairCount int
		zeroPoolPairs     map[string]bool
		initialPairs      []terraswap.Pair
		newPairCount      int
	}

	config := configs.Config{
		Terraswap: configs.TerraswapConfig{
			ChainId:          "pisco-1",
			FilterUnverified: false,
		}}
	var repo *repositoryMock
	var c cache.Cache
	var tsCache tscache.TerraswapCache
	var route router.Router
	var dh DataHandler

	mockTokens := func(pairs []terraswap.Pair) {
		addresses := []string{}
		for _, pair := range pairs {
			addresses = append(addresses, pair.LiquidityToken)
			for _, asset := range pair.AssetInfos {
				addresses = append(addresses, asset.GetKey())
			}
		}
		for _, addr := range addresses {
			token := terraswap.NewFakeToken()
			token.ContractAddr = addr
			tsCache.SetToken(token)
			repo.On("getToken", addr).Return(&token, nil)
		}
	}

	beforeEach := func(t *testcase) {
		repo = newRepositoryMock()
		c = cache.New(config.Cache)
		tsCache = tscache.New(c)
		route = router.NewRouterMock()

		mockTokens(t.initialPairs)
		repo.On("getAllPairs").Return(t.initialPairs, nil).Twice()
		repo.On("getCw20Allowlist", config.Terraswap.Cw20AllowlistUrl).Return(t.cw20Allowlist, nil)
		repo.On("getIbcAllowlist", config.Terraswap.IbcAllowlistUrl).Return(t.ibcAllowlist, nil)
		zeroPoolPairs := make(map[string]bool)
		for i := 0; i < t.zeroPoolPairCount; i++ {
			zeroPoolPairs[t.initialPairs[i].ContractAddr] = true
		}
		t.zeroPoolPairs = zeroPoolPairs
		repo.On("getZeroPoolPairs", t.initialPairs).Return(zeroPoolPairs, nil)
		dh = NewDataHandler(repo, route, tsCache, config)
		dh.Run()

	}

	testCases := []testcase{
		{
			initialPairs:      terraswap.NewFakePairs(2),
			cw20Allowlist:     terraswap.TokensMap{},
			ibcAllowlist:      terraswap.TokensMap{},
			zeroPoolPairCount: 1,
			newPairCount:      1,
		},
		{
			initialPairs:      terraswap.NewFakePairs(3),
			cw20Allowlist:     terraswap.TokensMap{},
			ibcAllowlist:      terraswap.TokensMap{},
			zeroPoolPairCount: 1,
			newPairCount:      2,
		},
	}

	for _, tc := range testCases {
		beforeEach(&tc)
		newPairs := terraswap.NewFakePairs(uint(tc.newPairCount))
		mockTokens(newPairs)
		allPairs := append(tc.initialPairs, newPairs...)
		repo.On("getAllPairs").Return(allPairs, nil).Twice()
		repo.On("getZeroPoolPairs", allPairs).Return(tc.zeroPoolPairs, nil)
		expectedPairs := []terraswap.Pair{}
		for i := tc.zeroPoolPairCount; i < len(allPairs); i++ {
			expectedPairs = append(expectedPairs, allPairs[i])
		}

		dh.Run()
		swapablePairs := dh.GetSwapablePairs()
		assert.Equal(t, expectedPairs, swapablePairs, "pairs must be the same")
	}

}
