package cache

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/terraswap/terraswap-service/internal/pkg/cache"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

type TerraswapCache interface {
	GetAllPairs() []terraswap.Pair
	SetAllPairs(pairs []terraswap.Pair) error

	GetAllTokens() terraswap.Tokens
	SetAllTokens(tokens terraswap.Tokens) error

	GetPairByAddr(addr string) *terraswap.Pair
	GetPairByAssets(addr0, addr1 string) *terraswap.Pair
	SetPair(pair terraswap.Pair) error

	GetToken(addr string) *terraswap.Token
	SetToken(token terraswap.Token) error

	GetSwapablePairs() []terraswap.Pair
	SetSwapablePairs(pairs []terraswap.Pair) error

	GetIbcAllowlistMap() terraswap.TokensMap
	SetIbcAllowlistMap(m terraswap.TokensMap) error

	GetCw20AllowlistMap() terraswap.TokensMap
	SetCw20AllowlistMap(m terraswap.TokensMap) error

	GetZeroPoolPairs() map[string]bool
	SetZeroPoolPairs(pairs map[string]bool) error
	cache.Cache
}

type cacheImpl struct {
	cache.Cache
}

var _ TerraswapCache = &cacheImpl{}

func New(cache cache.Cache) TerraswapCache {
	return &cacheImpl{cache}
}

// GetAllPairs implements terraswapCache
func (c *cacheImpl) GetAllPairs() []terraswap.Pair {
	a := c.Get(getAllPairsKey())
	pairs, ok := a.([]terraswap.Pair)
	if !ok {
		return nil
	}
	return pairs
}

// SetAllPairs implements terraswapCache
func (c *cacheImpl) SetAllPairs(pairs []terraswap.Pair) error {
	if err := c.Set(getAllPairsKey(), pairs); err != nil {
		return errors.Wrap(err, "terraswap.SetAllPairs")
	}
	return nil
}

// GetAllTokens implements terraswapCache
func (c *cacheImpl) GetAllTokens() terraswap.Tokens {
	v := c.Get(getAllTokensKey())
	tokens, ok := v.(terraswap.Tokens)
	if !ok {
		return *terraswap.NewTokens(nil)
	}
	return tokens
}

// SetAllTokens implements terraswapCache
func (c *cacheImpl) SetAllTokens(tokens terraswap.Tokens) error {
	if err := c.Set(getAllTokensKey(), tokens); err != nil {
		return errors.Wrap(err, "terraswap.SetAllTokens")
	}
	return nil
}

// GetCw20 terraswap.TokensMap implements terraswapCache
func (c *cacheImpl) GetCw20AllowlistMap() terraswap.TokensMap {
	v := c.Get(getCw20AllowlistKey())
	allowlistMap, ok := v.(terraswap.TokensMap)
	if !ok {
		return terraswap.TokensMap{}
	}
	return allowlistMap
}

// SetCw20 terraswap.TokensMap implements terraswapCache
func (c *cacheImpl) SetCw20AllowlistMap(m terraswap.TokensMap) error {
	if err := c.Set(getCw20AllowlistKey(), m); err != nil {
		return errors.Wrap(err, "terraswap.SetCw20AllowlistMap")
	}
	return nil
}

// GetIbc terraswap.TokensMap implements terraswapCache
func (c *cacheImpl) GetIbcAllowlistMap() terraswap.TokensMap {
	v := c.Get(getIbcAllowlistKey())
	allowlistMap, ok := v.(terraswap.TokensMap)
	if !ok {
		return terraswap.TokensMap{}
	}
	return allowlistMap
}

func (c *cacheImpl) SetIbcAllowlistMap(m terraswap.TokensMap) error {
	if err := c.Set(getIbcAllowlistKey(), m); err != nil {
		return errors.Wrap(err, "terraswap.SetIbcAllowlistMap")
	}
	return nil
}

func (c *cacheImpl) GetPairByAddr(addr string) *terraswap.Pair {
	v := c.Get(getPairKey(addr))
	pair, ok := v.(terraswap.Pair)
	if !ok {
		return nil
	}
	return &pair
}

func (c *cacheImpl) GetPairByAssets(addr0, addr1 string) *terraswap.Pair {
	v := c.Get(getPairKeyByAssets(addr0, addr1))
	pair, ok := v.(terraswap.Pair)
	if !ok {
		return nil
	}
	return &pair
}

// GetSwapablePairs implements terraswapCache
func (c *cacheImpl) GetSwapablePairs() []terraswap.Pair {
	a := c.Get(getSwapablePairsKey())
	pairs, ok := a.([]terraswap.Pair)
	if !ok {
		return nil
	}
	return pairs
}

// GetToken implements terraswapCache
func (c *cacheImpl) GetToken(addr string) *terraswap.Token {
	v := c.Get(getTokenKey(addr))
	token, ok := v.(terraswap.Token)
	if !ok {
		return nil
	}
	return &token
}

// SetPair implements terraswapCache
func (c *cacheImpl) SetPair(pair terraswap.Pair) error {

	if err := c.Set(getPairKey(pair.ContractAddr), pair); err != nil {
		return errors.Wrapf(err, "terraswap.SetPair(%s)", pair.ContractAddr)
	}
	if err := c.Set(getPairKey(pair.LiquidityToken), pair); err != nil {
		return errors.Wrapf(err, "terraswap.SetPair(%s)", pair.ContractAddr)
	}
	assets := pair.AssetInfos

	if err := c.Set(getPairKeyByAssets(assets[0].GetKey(), assets[1].GetKey()), pair); err != nil {
		return errors.Wrapf(err, "terraswap.SetPair(%s)", pair.ContractAddr)
	}

	return nil
}

// SetSwapablePairs implements terraswapCache
func (c *cacheImpl) SetSwapablePairs(pairs []terraswap.Pair) error {

	if err := c.Set(getSwapablePairsKey(), pairs); err != nil {
		return errors.Wrapf(err, "terraswap.SetSwapablePairs(%v)", pairs)
	}

	return nil
}

// SetToken implements terraswapCache
func (c *cacheImpl) SetToken(token terraswap.Token) error {
	if err := c.Set(getTokenKey(token.ContractAddr), token); err != nil {
		return errors.Wrapf(err, "terraswap.SetToken(%s)", token.ContractAddr)
	}

	return nil
}

// SetZeroPoolPairs implements terraswapCache
func (c *cacheImpl) GetZeroPoolPairs() map[string]bool {
	a := c.Get(getZeroPoolPairsKey())
	pairs, ok := a.(map[string]bool)
	if !ok {
		return nil
	}
	return pairs
}

func (c *cacheImpl) SetZeroPoolPairs(pairs map[string]bool) error {
	if err := c.Set(getZeroPoolPairsKey(), pairs); err != nil {
		return errors.Wrap(err, "terraswap.SetZeroPoolPairs")
	}

	return nil
}
