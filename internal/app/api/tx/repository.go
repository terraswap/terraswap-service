package tx

import (
	tsApp "github.com/delight-labs/terraswap-service/internal/app/api/common/terraswap"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap"

	"github.com/pkg/errors"
)

type Repository interface {
	GetPairByAssets(nameOrAddr, nameOrAddr2 string) *terraswap.Pair
	GetPair(contractAddr string) *terraswap.Pair
	GetIncreaseAllowance(amount, pairAddress string) *terraswap.ExecuteMsg
	GetSwapExecuteMsg(fromAsset terraswap.AssetInfo, pairAddress, amount, max_spread, belief_price string) *terraswap.ExecuteMsg
	GetProvideLiquidityExecuteMsg(from, fromAmount, toAmount, slippage string, p terraswap.Pair) *terraswap.ExecuteMsg
	GetWithdrawExecuteMsg(p terraswap.Pair, amount string) *terraswap.ExecuteMsg
	GetSwapRouteExecuteMsg(from string, routes []string) (*terraswap.ExecuteMsg, error)
	GetTokenDecimals(tokenId string) int
	GetSwapableTokensFrom(from string, hopCount int) []string
	GetRouteContractAddress() string
	GetRoutes(from, to string) [][]string
}

var _ Repository = &repositoryImpl{}

type repositoryImpl struct {
	db tsApp.DataHandler
}

//Fixme change db interface
func newRepo(db tsApp.DataHandler) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) GetTokenDecimals(tokenAddr string) int {
	token := r.db.GetToken(tokenAddr)
	if token == nil {
		return -1
	}
	return token.Decimals
}

func (r *repositoryImpl) GetSwapableTokensFrom(from string, hopCount int) []string {
	return r.db.GetTokensFrom(from, hopCount)
}

func (r *repositoryImpl) GetRoutes(from, to string) [][]string {
	return r.db.GetRoutes(from, to)
}

func (r *repositoryImpl) GetPair(contractAddr string) *terraswap.Pair {
	return r.db.GetPair(contractAddr)

}

func (r *repositoryImpl) GetPairByAssets(nameOrAddr, nameOrAddr2 string) *terraswap.Pair {
	return r.db.GetPairByAssets(nameOrAddr, nameOrAddr2)
}

func (r *repositoryImpl) GetIncreaseAllowance(amount, pairAddress string) *terraswap.ExecuteMsg {
	return &terraswap.ExecuteMsg{
		IncreaseAllowance: &terraswap.IncreaseAllowanceMsg{
			Amount:  amount,
			Spender: pairAddress,
		},
	}
}

func (r *repositoryImpl) GetRouteContractAddress() string {
	return r.db.GetRouterAddress()
}

func (r *repositoryImpl) GetSwapExecuteMsg(fromAsset terraswap.AssetInfo, pairAddress, amount, max_spread, belief_price string) *terraswap.ExecuteMsg {
	if fromAsset.GetTokenType() == terraswap.Cw20TokenType {
		msg, err := terraswap.GetSwapSendMsg(max_spread, belief_price)
		if err != nil {
			panic(err)
		}
		return &terraswap.ExecuteMsg{
			Send: &terraswap.SendMsg{
				Amount:   amount,
				Contract: pairAddress,
				Msg:      msg,
			},
		}
	}

	return &terraswap.ExecuteMsg{
		Swap: &terraswap.SwapMsg{
			OfferAsset: terraswap.OfferAsset{
				Amount: amount,
				Info:   fromAsset,
			},
			MaxSpread:   max_spread,
			BeliefPrice: belief_price,
		},
	}
}

func (r *repositoryImpl) GetProvideLiquidityExecuteMsg(from, fromAmount, toAmount, slippage string, p terraswap.Pair) *terraswap.ExecuteMsg {
	fromIdx := 0
	toIdx := 1

	if p.AssetInfos[0].GetKey() != from {
		fromIdx = 1
		toIdx = 0
	}
	return &terraswap.ExecuteMsg{
		Provide: &terraswap.ProvideMsg{
			Assets: []terraswap.OfferAsset{
				{
					Info:   p.AssetInfos[fromIdx],
					Amount: fromAmount,
				}, {
					Info:   p.AssetInfos[toIdx],
					Amount: toAmount,
				},
			},
			SlippageTolerance: slippage,
		},
	}
}

func (r *repositoryImpl) GetWithdrawExecuteMsg(p terraswap.Pair, amount string) *terraswap.ExecuteMsg {

	msg, err := terraswap.GetWithdrawSendMsg()
	if err != nil {
		panic(err)
	}

	return &terraswap.ExecuteMsg{
		Send: &terraswap.SendMsg{
			Amount:   amount,
			Contract: p.ContractAddr,
			Msg:      msg,
		},
	}
}

func (r *repositoryImpl) GetSwapRouteExecuteMsg(from string, routes []string) (*terraswap.ExecuteMsg, error) {

	ops := []terraswap.RouteSwapOperation{}
	for _, to := range routes {
		offerAssetInfo, err := terraswap.AddressToAssetInfo(from)
		if err != nil {
			err := errors.Wrap(err, "invalid routes info")
			return nil, err
		}
		askAssetInfo, err := terraswap.AddressToAssetInfo(to)
		if err != nil {
			err := errors.Wrap(err, "invalid routes info")
			return nil, err
		}
		op := &terraswap.TerraSwapOperation{
			OfferAssetInfo: *offerAssetInfo,
			AskAssetInfo:   *askAssetInfo,
		}
		ops = append(ops, terraswap.RouteSwapOperation{TerraSwapOperation: op})

		from = to
	}

	return &terraswap.ExecuteMsg{
		RouteSwapOperation: &terraswap.RouteSwapOperationMsg{
			Operations:     ops,
			MinimumReceive: "",
		},
	}, nil
}
