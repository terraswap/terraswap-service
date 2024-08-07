package grpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/pkg/errors"

	wasmtype "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/types"
	ibctype "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	"github.com/terraswap/terraswap-service/configs"
	"github.com/terraswap/terraswap-service/internal/pkg/logging"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap/databases"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type TerraswapGrpcClient interface {
	GetIbcDenom(ibcHash string) (*terraswap.IbcDenomTrace, error)
	GetPairs(lastPair terraswap.Pair) (pairs []terraswap.Pair, err error)
	GetTokenInfo(tokenAddress string) (*terraswap.Token, error)
	GetZeroPoolPairs(pairs []terraswap.Pair) (map[string]bool, error)
	GetDenoms() ([]string, error)
}

var _ databases.TerraswapDb = &terraswapGrpcCon{}

type terraswapGrpcCon struct {
	logger  logging.Logger
	con     *grpc.ClientConn
	chainId string
}

func New(host, chainId string, insecureCon bool, log configs.LogConfig) TerraswapGrpcClient {
	logger := logging.New("TerraswapGrpcClient", log)
	config := types.NewConfig()

	config.SetCoinType(types.CoinType)
	config.SetBech32PrefixForAccount(types.Bech32PrefixAccAddr, types.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(types.Bech32PrefixValAddr, types.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(types.Bech32PrefixConsAddr, types.Bech32PrefixConsPub)
	config.Seal()
	con := connectGRPC(host, insecureCon)

	return &terraswapGrpcCon{logger, con, chainId}
}

func connectGRPC(host string, isInsecure bool) *grpc.ClientConn {
	var opts []grpc.DialOption

	if isInsecure {
		cred := insecure.NewCredentials()
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		dialer := &net.Dialer{
			Timeout: 5 * time.Second,
		}
		// Establish a TLS connection to fetch server certificates with a timeout
		conn, err := tls.DialWithDialer(dialer, "tcp", host, &tls.Config{
			InsecureSkipVerify: true,
		})
		if err != nil {
			panic(err.Error())
		}

		certs := conn.ConnectionState().PeerCertificates
		conn.Close()

		pool := x509.NewCertPool()
		pool.AddCert(certs[0])

		clientCert := credentials.NewClientTLSFromCert(pool, "")
		opts = append(opts, grpc.WithTransportCredentials(clientCert))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts = append(opts, grpc.WithBlock())

	rpcConn, err := grpc.DialContext(ctx, host, opts...)
	if err != nil {
		panic(err.Error())
	}

	return rpcConn
}

const (
	PREFIX = iota
	HASH

	IBC = "ibc"
)

func (t *terraswapGrpcCon) GetIbcDenom(ibcHash string) (*terraswap.IbcDenomTrace, error) {
	params := strings.Split(ibcHash, "/")
	if len(params) == 2 && params[PREFIX] != IBC {
		return nil, errors.Errorf(`format of the ibc does not match (ibc/HASH), but %v`, ibcHash)
	}
	hashIdx := HASH
	if len(params) == 1 {
		hashIdx = 0
	}

	client := ibctype.NewQueryClient(t.con)
	res, err := client.DenomTrace(context.Background(), &ibctype.QueryDenomTraceRequest{
		Hash: params[hashIdx],
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get hash(%v) from ibc", ibcHash)
	}
	denomTrace := terraswap.NewIbcDenomTrace(res.DenomTrace.Path, res.DenomTrace.BaseDenom)
	return &denomTrace, nil
}

func (t *terraswapGrpcCon) GetPairs(lastPair terraswap.Pair) (pairs []terraswap.Pair, err error) {

	qmsg, err := terraswap.GetQueryMsg(lastPair)
	if err != nil {
		errors.Wrapf(err, "cannot getQueryMsg from last pair %v", lastPair)
	}

	client := wasmtype.NewQueryClient(t.con)
	res, err := client.SmartContractState(context.Background(), &wasmtype.QuerySmartContractStateRequest{
		Address:   terraswap.GetFactoryAddress(t.chainId, ""),
		QueryData: qmsg,
	})
	if err != nil {
		err := errors.Wrap(err, "terraswap.GrpcCon.GetPairs()")
		return nil, err
	}

	var pres terraswap.Pairs
	err = json.Unmarshal(res.Data, &pres)
	if err != nil {
		return nil, err
	}

	return pres.Pairs, nil
}

// GetZeroPoolPairs implements TerraswapGrpcClient
func (t *terraswapGrpcCon) GetZeroPoolPairs(pairs []terraswap.Pair) (map[string]bool, error) {
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

func (t *terraswapGrpcCon) getPoolInfo(addr string) (*terraswap.PoolInfo, error) {
	client := wasmtype.NewQueryClient(t.con)
	res, err := client.SmartContractState(context.Background(), &wasmtype.QuerySmartContractStateRequest{
		Address:   addr,
		QueryData: []byte(`{"pool":{}}`),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "grpc.getPairInfo(%s)", addr)
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

func (t *terraswapGrpcCon) GetTokenInfo(tokenAddress string) (*terraswap.Token, error) {
	client := wasmtype.NewQueryClient(t.con)
	res, err := client.SmartContractState(context.Background(), &wasmtype.QuerySmartContractStateRequest{
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

// GetDenoms implements TerraswapGrpcClient
func (*terraswapGrpcCon) GetDenoms() ([]string, error) {
	return []string{"uluna"}, nil
}
