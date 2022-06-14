package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/terraswap/terraswap-service/configs"
	"github.com/terraswap/terraswap-service/internal/pkg/cache"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

func TestTerraswapCache(t *testing.T) {
	config := configs.CacheConfig{}
	cache := cache.New(config)
	tsCache := New(cache)

	token := terraswap.NewFakeToken()
	tsCache.SetToken(token)

	cached := tsCache.GetToken(token.ContractAddr)
	assert.Equal(t, token, *cached)
}
