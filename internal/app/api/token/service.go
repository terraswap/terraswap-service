package token

import (
	"sort"

	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

type service interface {
	GetAllTokens() *terraswap.Tokens
	GetToken(addr string) *terraswap.Token
	GetSwapableTokens(from string, hopCount int) []string
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

func (s *swapServiceImpl) GetAllTokens() *terraswap.Tokens {
	return s.repo.GetAll()
}

func (s *swapServiceImpl) GetToken(addr string) *terraswap.Token {
	return s.repo.GetToken(addr)
}

func (s *swapServiceImpl) GetSwapableTokens(from string, hopCount int) []string {
	tokens := s.repo.GetSwapableTokens(from, hopCount)
	if terraswap.IsCw20Token(from) {
		return tokens
	}

	uniqueToken := make(map[string]bool)

	for _, v := range tokens {
		uniqueToken[v] = true
	}

	keys := make([]string, 0, len(uniqueToken))
	for k := range uniqueToken {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}
