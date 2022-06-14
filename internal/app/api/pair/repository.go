package pair

import (
	tsApp "github.com/terraswap/terraswap-service/internal/app/api/common/terraswap"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

type repository interface {
	GetSwapableAll() *terraswap.Pairs
	GetAllPairs() *terraswap.Pairs
	GetPair(contractAddr string) *terraswap.Pair
}

var _ repository = &repositoryImpl{}

type repositoryImpl struct {
	db tsApp.DataHandler
}

//Fixme change db interface
func newRepo(tsRepo tsApp.DataHandler) repository {
	return &repositoryImpl{db: tsRepo}
}

func (r *repositoryImpl) GetSwapableAll() *terraswap.Pairs {
	pairs := r.db.GetSwapablePairs()
	return &terraswap.Pairs{
		Pairs: pairs,
	}
}

func (r *repositoryImpl) GetAllPairs() *terraswap.Pairs {
	pairs := r.db.GetAllPairs()
	return &terraswap.Pairs{
		Pairs: pairs,
	}
}

func (r *repositoryImpl) GetPair(contractAddr string) *terraswap.Pair {
	return r.db.GetPair(contractAddr)
}
