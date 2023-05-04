package cache

import (
	"fmt"

	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

const (
	allPairsKey      = "all_pairs"
	allTokensKey     = "all_tokens"
	swapablePairsKey = "swapable_pairs"
	allDenomsKey     = "all_denoms"
	cw20AllowlistKey = "cw20_allowlist"
	ibcAllowlistKey  = "ibc_allowlist"
	zeroPoolPairKey  = "zero_pool_pairs"
)

func getAllPairsKey() string {
	return allPairsKey
}

func getSwapablePairsKey() string {
	return swapablePairsKey
}

func getAllTokensKey() string {
	return allTokensKey
}

func getCw20AllowlistKey() string {
	return cw20AllowlistKey
}

func getIbcAllowlistKey() string {
	return ibcAllowlistKey
}

func getPairKeyByAssets(addr1, addr2 string) string {
	return terraswap.GetPairKeyByAssets(addr1, addr2)
}

func getPairKey(key string) string {
	return fmt.Sprintf("%s:%s", "pair", key)
}

func getTokenKey(key string) string {
	return fmt.Sprintf("%s:%s", "token", key)
}

func getZeroPoolPairsKey() string {
	return zeroPoolPairKey
}
