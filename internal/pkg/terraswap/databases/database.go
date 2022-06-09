package databases

import (
	ibctype "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap"
)

type TerraswapDb interface {
	GetIBCDenom(ibcHash string) (*ibctype.DenomTrace, error)
	GetPairs(lastPair terraswap.Pair) (pairs []terraswap.Pair, err error)
	GetTokenInfo(tokenAddress string) (*terraswap.Token, error)
	GetZeroPoolPairs(pairs []terraswap.Pair) (map[string]bool, error)
}
