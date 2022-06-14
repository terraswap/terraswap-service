package pair

import (
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

type service interface {
	GetAllPairs() *terraswap.Pairs
	GetSwapablePairs() *terraswap.Pairs
	GetPair(addr string) *terraswap.Pair
}

var _ service = &swapServiceImpl{}

type swapServiceImpl struct {
	repo repository
}

func newService(r repository) service {
	return &swapServiceImpl{
		repo: r,
	}
}

func (s *swapServiceImpl) GetAllPairs() *terraswap.Pairs {
	return s.repo.GetAllPairs()
}

func (s *swapServiceImpl) GetSwapablePairs() *terraswap.Pairs {
	return s.repo.GetSwapableAll()
}

func (s *swapServiceImpl) GetPair(addr string) *terraswap.Pair {
	return s.repo.GetPair(addr)
}
