package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/types"
	ibc "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	"github.com/terraswap/terraswap-service/configs"
	"github.com/terraswap/terraswap-service/internal/pkg/logging"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
	"github.com/terraswap/terraswap-service/pkg/classic/grpc/oracle"
	"github.com/terraswap/terraswap-service/pkg/classic/util"
	"google.golang.org/grpc"
)

type terraswapClassicGrpcCon struct {
	logger  logging.Logger
	con     *grpc.ClientConn
	chainId string
	version string
}

func NewClassic(host, chainId, version string, insecure bool, log configs.LogConfig) TerraswapGrpcClient {
	logger := logging.New("TerraswapGrpcClient", log)
	config := types.NewConfig()

	config.SetCoinType(util.CoinType)
	config.SetBech32PrefixForAccount(util.Bech32PrefixAccAddr, util.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(util.Bech32PrefixValAddr, util.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(util.Bech32PrefixConsAddr, util.Bech32PrefixConsPub)
	config.Seal()
	con := connectGRPC(host, insecure)

	return &terraswapClassicGrpcCon{logger, con, chainId, version}
}

// GetZeroPoolPairs implements TerraswapGrpcClient
func (t *terraswapClassicGrpcCon) GetZeroPoolPairs(pairs []terraswap.Pair) (map[string]bool, error) {
	zeroPool := make(map[string]bool)

	for idx, pair := range pairs {
		poolInfo, err := t.getPoolInfo(pair.ContractAddr)
		if err != nil {
			t.logger.Debug(errors.Wrapf(err, "grpc.GetZeroPoolPairs(%s) %d", pair.ContractAddr, idx))
			continue
		}

		for _, asset := range poolInfo.Assets {
			if asset.Amount == "" || asset.Amount == "0" {
				zeroPool[pair.ContractAddr] = true
			}
		}

	}
	return zeroPool, nil
}

func (t *terraswapClassicGrpcCon) getPoolInfo(addr string) (*terraswap.PoolInfo, error) {
	client := wasm.NewQueryClient(t.con)
	res, err := client.SmartContractState(context.Background(), &wasm.QuerySmartContractStateRequest{
		Address:   addr,
		QueryData: []byte(`{"pool":{}}`),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "grpc.getPoolInfo(%s)", addr)
	}

	type poolInfoRes struct {
		Result terraswap.PoolInfo `json:"query_result"`
	}

	var poolInfo poolInfoRes
	err = json.Unmarshal(res.Data, &poolInfo.Result)
	if err != nil {
		return nil, errors.Wrapf(err, "unmarshal token(%s)", addr)
	}

	return &poolInfo.Result, nil
}
func (c *terraswapClassicGrpcCon) GetDenoms() (denoms []string, err error) {
	client := oracle.NewQueryClient(c.con)

	res, err := client.Actives(context.Background(), &oracle.QueryActivesRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "cannot get denoms from oracle")
	}

	res.Actives = append(res.Actives, "uluna")

	return res.Actives, nil
}

func (c *terraswapClassicGrpcCon) GetIbcDenom(ibcHash string) (*terraswap.IbcDenomTrace, error) {
	params := strings.Split(ibcHash, "/")
	if len(params) != 2 || params[PREFIX] != IBC {
		return nil, errors.Errorf(`format of the ibc does not match (ibc/HASH), but %v`, ibcHash)
	}

	client := ibc.NewQueryClient(c.con)
	res, err := client.DenomTrace(context.Background(), &ibc.QueryDenomTraceRequest{
		Hash: params[HASH],
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get hash(%v) from ibc", ibcHash)
	}

	denomTrace := terraswap.NewIbcDenomTrace(res.DenomTrace.Path, res.DenomTrace.BaseDenom)
	return &denomTrace, nil
}

func (c *terraswapClassicGrpcCon) GetPairs(lastPair terraswap.Pair) (pairs []terraswap.Pair, err error) {

	qmsg, err := terraswap.GetQueryMsg(lastPair)
	if err != nil {
		errors.Wrapf(err, "cannot getQueryMsg from last pair %v", lastPair)
	}

	client := wasm.NewQueryClient(c.con)
	res, err := client.SmartContractState(context.Background(), &wasm.QuerySmartContractStateRequest{
		Address:   terraswap.GetFactoryAddress(c.chainId, c.version),
		QueryData: qmsg,
	})
	if err != nil {
		println(err.Error())
		return nil, err
	}

	var pres terraswap.Pairs
	err = json.Unmarshal(res.Data, &pres)
	if err != nil {
		return nil, err
	}

	return pres.Pairs, nil
}

func (c *terraswapClassicGrpcCon) GetTokenInfo(tokenAddress string) (*terraswap.Token, error) {

	client := wasm.NewQueryClient(c.con)
	res, err := client.SmartContractState(context.Background(), &wasm.QuerySmartContractStateRequest{
		Address:   tokenAddress,
		QueryData: []byte(`{"token_info":{}}`),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "getTokenInfoGRPC: query fail for token(%s)", tokenAddress)
	}

	type tokenResponse struct {
		Height string          `json:"height"`
		Result terraswap.Token `json:"result"`
	}

	var token tokenResponse
	err = json.Unmarshal(res.Data, &token.Result)
	if err != nil {
		return nil, errors.Wrapf(err, "unmarshal token(%s)", tokenAddress)
	}

	if token.Result.Name == "" {
		msg := fmt.Sprintf("unknown token(%s)", tokenAddress)
		return nil, errors.New(msg)
	}

	token.Result.ContractAddr = tokenAddress

	return &token.Result, nil

}
