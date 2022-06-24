package token

import (
	"sort"

	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

var _ service = &classicServiceImpl{}

type classicServiceImpl struct {
	serviceMixinImpl
}

func newClassicService(r repository) service {
	return &classicServiceImpl{
		serviceMixinImpl{r},
	}
}

func (s *classicServiceImpl) GetSwapableTokens(from string, hopCount int) []string {
	tokens := s.repo.GetSwapableTokens(from, hopCount)
	if terraswap.IsCw20Token(from) {
		return tokens
	}

	tokens = append(tokens, s.repo.GetActiveDenoms()...)
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
