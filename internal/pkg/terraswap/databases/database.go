package databases

import (
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

type TerraswapDb interface {
	GetIbcDenom(ibcHash string) (*terraswap.IbcDenomTrace, error)
	GetPairs(lastPair terraswap.Pair) (pairs []terraswap.Pair, err error)
	GetTokenInfo(tokenAddress string) (*terraswap.Token, error)
	GetZeroPoolPairs(pairs []terraswap.Pair) (map[string]bool, error)
}
