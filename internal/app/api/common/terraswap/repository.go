package terraswap

import (
	"fmt"
	"strings"

	ibctype "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap/allowlist"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap/databases"
	"github.com/pkg/errors"
)

type repository interface {
	getAllPairs() ([]terraswap.Pair, error)
	getZeroPoolPairs(pairs []terraswap.Pair) (map[string]bool, error)
	getCw20Allowlist(url string) terraswap.TokensMap
	getIbcAllowlist(url string) terraswap.TokensMap
	getIbcDenom(ibcHash string) (*terraswap.Token, error)
	getToken(addr string) (*terraswap.Token, error)
}

type repositoryImpl struct {
	chainId string
	store   databases.TerraswapDb
	mapper  mapper
}

type mapper struct{}

var _ repository = &repositoryImpl{}

func NewRepo(chainId string, store databases.TerraswapDb) repository {
	return &repositoryImpl{chainId, store, mapper{}}
}

// GetAllPairs implements repository
func (r *repositoryImpl) getAllPairs() ([]terraswap.Pair, error) {
	allPairs := []terraswap.Pair{}
	lastPair := terraswap.Pair{}

	for {
		pairs, err := r.store.GetPairs(lastPair)
		if err != nil {
			return nil, errors.Wrap(err, "cannot get all pairs")
		}
		if len(pairs) == 0 {
			break
		}
		lastPair = pairs[len(pairs)-1]
		allPairs = append(allPairs, pairs...)
	}

	return allPairs, nil
}

// getCw20Allowlist implements repository
func (r *repositoryImpl) getCw20Allowlist(url string) terraswap.TokensMap {
	res, err := allowlist.GetAllowlistMapResponse[allowlist.Cw20AllowlistResponse](url)
	if err != nil {
		return nil
	}
	var cw20TokenMap map[string]allowlist.Cw20Allowlist
	if terraswap.IsClassic(r.chainId) {
		cw20TokenMap = res.Classic
	} else if terraswap.IsMainnet(r.chainId) {
		cw20TokenMap = res.Mainnet
	} else if terraswap.IsTestnet(r.chainId) {
		cw20TokenMap = res.Testnet
	}

	tokensMap := terraswap.TokensMap{}

	for k, v := range cw20TokenMap {
		tokensMap[k] = terraswap.Token{
			Name:         v.Name,
			Symbol:       v.Symbol,
			ContractAddr: v.Token,
			Verified:     true,
			Icon:         v.Icon,
			Decimals:     int(v.Decimals),
			Protocol:     v.Protocol,
		}
	}

	return tokensMap
}

// getIbcAllowlist implements repository
func (r *repositoryImpl) getIbcAllowlist(url string) terraswap.TokensMap {
	res, err := allowlist.GetAllowlistMapResponse[allowlist.IbcAllowlistResponse](url)
	if err != nil {
		return nil
	}

	var ibcTokenMap map[string]allowlist.IbcTokenAllowlist
	if terraswap.IsClassic(r.chainId) {
		ibcTokenMap = res.Classic
	} else if terraswap.IsMainnet(r.chainId) {
		ibcTokenMap = res.Mainnet
	} else if terraswap.IsTestnet(r.chainId) {
		ibcTokenMap = res.Testnet
	}

	tokensMap := terraswap.TokensMap{}
	for k, allowlist := range ibcTokenMap {
		if !strings.Contains(k, string(terraswap.IbcTokenPrefix)) {
			k = fmt.Sprintf("%s%s", terraswap.IbcTokenPrefix, k)
		}
		tokensMap[k] = r.mapper.ibcAllowlistToToken(allowlist)
	}

	return tokensMap
}

// GetZeroPoolPairs implements repository
func (r *repositoryImpl) getZeroPoolPairs(pairs []terraswap.Pair) (map[string]bool, error) {
	var err error
	if len(pairs) == 0 {
		if pairs, err = r.getAllPairs(); err != nil {
			return nil, errors.Wrap(err, "terraswap.Repository.GetZeroPoolPairs")
		}
	}
	zeroPoolPairsMap, err := r.store.GetZeroPoolPairs(pairs)
	if err != nil {
		return nil, errors.Wrap(err, "terraswap.Repository.GetZeroPoolPairs")
	}
	return zeroPoolPairsMap, nil
}

// getIbcDenom implements repository
func (r *repositoryImpl) getIbcDenom(ibcHash string) (*terraswap.Token, error) {
	denomTrace, err := r.store.GetIBCDenom(ibcHash)
	if err != nil {
		return nil, errors.Wrap(err, "repository.getIbcDenom")
	}
	tokenInfo := r.mapper.ibcDenomTraceToToken(ibcHash, *denomTrace)

	return &tokenInfo, nil
}

// getToken implements repository
func (r *repositoryImpl) getToken(addr string) (*terraswap.Token, error) {
	if terraswap.IsIbcToken(addr) {
		return r.getIbcDenom(addr)
	}

	if terraswap.IsNativeToken(addr) {
		token := r.mapper.denomAddrToToken(addr)
		return &token, nil
	}

	token, err := r.store.GetTokenInfo(addr)
	if err != nil {
		return nil, errors.Wrap(err, "repository.getToken")
	}
	return token, nil
}

func (m *mapper) denomAddrToToken(denom string) terraswap.Token {
	symbol := terraswap.ToDenomSymbol(denom)
	return terraswap.Token{
		Name:         denom,
		Symbol:       symbol,
		Decimals:     6,
		ContractAddr: denom,
		Icon:         fmt.Sprintf(`%s/%s.png`, terraswap.DenomIconUrl, strings.ToUpper(symbol)),
		Verified:     true,
	}
}

func (m *mapper) ibcDenomTraceToToken(ibcHash string, trace ibctype.DenomTrace) terraswap.Token {
	symbol := terraswap.ToDenomSymbol(trace.BaseDenom)
	return terraswap.Token{
		Name:         ibcHash,
		Symbol:       symbol,
		Decimals:     6,
		ContractAddr: ibcHash,
		Icon:         fmt.Sprintf(`%s/%s.png`, terraswap.IbcIconUrl, strings.ToUpper(symbol)),
	}
}

func (m *mapper) ibcAllowlistToToken(v allowlist.IbcTokenAllowlist) terraswap.Token {
	decimals := v.Decimals
	if decimals == 0 {
		decimals = 6
	}
	return terraswap.Token{
		Name:         v.Name,
		Symbol:       v.Symbol,
		ContractAddr: v.Denom,
		Verified:     true,
		Icon:         v.Icon,
		Decimals:     decimals,
	}
}
