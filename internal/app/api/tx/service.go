package tx

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/terraswap/terraswap-service/internal/app/api/utils/responser"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

type Service interface {
	getSwapTx(from, to, amount, sender, max_spread, belief_price string) []*terraswap.UnsignedTx
	GetSwapTxs(from, to, amount, sender, max_spread, belief_price string, hopCount int) ([][]*terraswap.UnsignedTx, *responser.ErrorResponse)
	GetProvideTx(from, to, fromAmount, toAmount, slippage, sender string) ([]*terraswap.UnsignedTx, *responser.ErrorResponse)
	GetWithdrawTx(ldAddr, amount, sender string) ([]*terraswap.UnsignedTx, *responser.ErrorResponse)
}

type mixinImpl struct {
	repo Repository
}

func (s *mixinImpl) getRouteSwapTx(from, amount, sender string, path []string) ([]*terraswap.UnsignedTx, error) {
	routerAddr := s.repo.GetRouteContractAddress()
	if routerAddr == "" {
		return nil, errors.New("api.tx.service.getRouteSwapTx(): there is no router")
	}
	addr := from
	if from[0:1] == "u" {
		addr = routerAddr
	}
	txs := make([]*terraswap.UnsignedTx, 0)
	utx := terraswap.BaseUnsignedTx(addr, sender)

	swapMsg, err := s.repo.GetSwapRouteExecuteMsg(from, path)
	if err != nil {
		return nil, err
	}

	if terraswap.IsCw20Token(from) {
		utx.Value.ExecuteMsg = terraswap.ExecuteMsg{
			Send: &terraswap.SendMsg{
				Amount:   amount,
				Contract: routerAddr,
				Msg:      swapMsg,
			},
		}
	} else {
		utx.Value.ExecuteMsg = *swapMsg
	}

	return append(txs, utx), nil
}

func (s *mixinImpl) getSwapTx(from, to, amount, sender, max_spread, belief_price string) []*terraswap.UnsignedTx {
	pair := s.repo.GetPairByAssets(from, to)
	if pair == nil {
		return nil
	}

	fromAsset := pair.AssetInfos[0]
	if fromAsset.GetKey() != from {
		fromAsset = pair.AssetInfos[1]
	}
	addr := pair.ContractAddr
	if fromAsset.GetTokenType() == terraswap.Cw20TokenType {
		addr = fromAsset.GetKey()
	}

	txs := make([]*terraswap.UnsignedTx, 0)
	tx := terraswap.BaseUnsignedTx(addr, sender)
	tx.Value.ExecuteMsg = *(s.repo.GetSwapExecuteMsg(fromAsset, pair.ContractAddr, amount, max_spread, belief_price))

	if fromAsset.GetTokenType() == terraswap.NativeTokenType {
		tx.Value.Coins = append(tx.Value.Coins, terraswap.NewCoin(amount, fromAsset.GetKey()))
	}

	txs = append(txs, tx)
	return txs
}

func (s *mixinImpl) GetProvideTx(from, to, fromAmount, toAmount, slippage, sender string) ([]*terraswap.UnsignedTx, *responser.ErrorResponse) {

	fromAmount, err := s.convertToTerraAmount(fromAmount, from)
	if err != nil {
		msg := fmt.Sprintf("cannot convert amount(%s) for %s", fromAmount, from)
		res := responser.GetBadRequest(msg, "")
		return nil, &res
	}

	toAmount, err = s.convertToTerraAmount(toAmount, to)
	if err != nil {
		msg := fmt.Sprintf("cannot convert amount(%s) for %s", toAmount, to)
		res := responser.GetBadRequest(msg, "")
		return nil, &res
	}

	pair := s.repo.GetPairByAssets(from, to)
	if pair == nil {
		msg := fmt.Sprintf("cannot find a pair(%s, %s)", from, to)
		res := responser.NotFound(msg, "")
		return nil, &res
	}

	txs := make([]*terraswap.UnsignedTx, 0)
	tx := terraswap.BaseUnsignedTx(pair.ContractAddr, sender)

	for _, a := range pair.AssetInfos {
		amount := fromAmount
		if to == a.GetKey() {
			amount = toAmount
		}

		if a.GetTokenType() == terraswap.Cw20TokenType {
			increaseAllowanceTx := terraswap.BaseUnsignedTx(a.GetKey(), sender)
			increaseAllowanceTx.Value.ExecuteMsg = *s.repo.GetIncreaseAllowance(amount, pair.ContractAddr)
			txs = append(txs, increaseAllowanceTx)
		}

		if a.GetTokenType() == terraswap.NativeTokenType {
			tx.Value.Coins = append(tx.Value.Coins, terraswap.NewCoin(amount, a.GetKey()))
		}
	}

	tx.Value.ExecuteMsg = *(s.repo.GetProvideLiquidityExecuteMsg(from, fromAmount, toAmount, slippage, *pair))

	txs = append(txs, tx)
	return txs, nil

}

func (s *mixinImpl) GetWithdrawTx(lpAddr, amount, sender string) ([]*terraswap.UnsignedTx, *responser.ErrorResponse) {
	pair := s.repo.GetPair(lpAddr)
	if pair == nil {
		msg := fmt.Sprintf("cannot find a pair by lpAddr(%s)", lpAddr)
		res := responser.NotFound(msg, "")
		return nil, &res
	}

	amount, err := s.convertToTerraAmount(amount, lpAddr)
	if err != nil {
		msg := fmt.Sprintf("cannot convert amount(%s) for %s", amount, lpAddr)
		res := responser.GetBadRequest(msg, "")
		return nil, &res
	}

	tx := terraswap.BaseUnsignedTx(pair.LiquidityToken, sender)
	tx.Value.ExecuteMsg = *(s.repo.GetWithdrawExecuteMsg(*pair, amount))

	return append(make([]*terraswap.UnsignedTx, 0), tx), nil

}

func (s mixinImpl) convertToTerraAmount(amount string, tokenAddr string) (string, error) {
	decimals := 6

	if tokenAddr[0:1] == "t" {
		decimals = s.repo.GetTokenDecimals(tokenAddr)
	}

	if decimals == -1 {
		return "", errors.New("cannot find a token")
	}

	amount, err := terraswap.ToTerraAmount(amount, decimals)
	if err != nil {
		err = errors.Wrap(err, "convert amount fail")
		return "", err
	}

	return amount, nil
}
