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

type serviceMixinImpl struct {
	repo repository
}

type terra2ServiceImpl struct {
	serviceMixinImpl
}

var _ service = &terra2ServiceImpl{}

func newService(r repository) service {
	return &terra2ServiceImpl{
		serviceMixinImpl{r},
	}
}
func (s *terra2ServiceImpl) GetSwapableTokens(from string, hopCount int) []string {
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

func (s *serviceMixinImpl) GetAllTokens() *terraswap.Tokens {
	return s.repo.GetAll()
}

func (s *serviceMixinImpl) GetToken(addr string) *terraswap.Token {
	return s.repo.GetToken(addr)
}
