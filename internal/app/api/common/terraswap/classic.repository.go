package terraswap

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap/databases"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap/databases/rdb"
)

type classicRepositoryImpl struct {
	repository
	rdb.TerraswapRdb
}

type classicMapper struct {
	mapperImpl
}

var _ repository = &classicRepositoryImpl{}

func NewClassicRepo(chainId string, store databases.TerraswapDb, rdb rdb.TerraswapRdb) repository {
	repo := &repositoryImpl{chainId, store, &classicMapper{}}
	return &classicRepositoryImpl{repo, rdb}
}

// GetZeroPoolPairs implements repository
func (r *classicRepositoryImpl) getZeroPoolPairs(pairs []terraswap.Pair) (map[string]bool, error) {
	if r.TerraswapRdb == nil {
		return r.repository.getZeroPoolPairs(pairs)
	}
	zeroPoolPairsMap, err := r.GetZeroPoolPairs(pairs)
	if err != nil {
		return nil, errors.Wrap(err, "terraswap.ClassicRepository.GetZeroPoolPairs")
	}
	return zeroPoolPairsMap, nil
}

func (m *classicMapper) denomAddrToToken(denom string) terraswap.Token {
	icon := ""
	symbol := ""
	if denom == "uluna" {
		symbol = "LUNC"
		icon = fmt.Sprintf(`%s/%s.svg`, terraswap.DenomIconUrl, symbol)
	} else { // ex) uusd, ukrw ...
		symbol = strings.ToUpper(denom[1:3]) + "T"
		icon = fmt.Sprintf(`%s/%s.svg`, terraswap.ClassicDenomIconUrl, strings.ToUpper(symbol))
	}

	return terraswap.Token{
		Name:         denom,
		Symbol:       symbol,
		Decimals:     6,
		ContractAddr: denom,
		Icon:         icon,
		Verified:     true,
	}
}
