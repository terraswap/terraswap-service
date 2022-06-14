package terraswap

import (
	"github.com/pkg/errors"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap/databases"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap/databases/rdb"
)

type classicRepositoryImpl struct {
	repository
	rdb.TerraswapRdb
}

var _ repository = &classicRepositoryImpl{}

func NewClassicRepo(chainId string, store databases.TerraswapDb, rdb rdb.TerraswapRdb) repository {
	repo := &repositoryImpl{chainId, store, mapper{}}
	return &classicRepositoryImpl{repo, rdb}
}

// GetZeroPoolPairs implements repository
func (r *classicRepositoryImpl) getZeroPoolPairs(pairs []terraswap.Pair) (map[string]bool, error) {
	zeroPoolPairsMap, err := r.GetZeroPoolPairs(pairs)
	if err != nil {
		return nil, errors.Wrap(err, "terraswap.ClassicRepository.GetZeroPoolPairs")
	}
	return zeroPoolPairsMap, nil
}
