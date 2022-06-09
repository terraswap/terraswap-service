package cache

import (
	"testing"

	"github.com/delight-labs/terraswap-service/configs"
	"github.com/delight-labs/terraswap-service/internal/pkg/cache"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap"
	"github.com/stretchr/testify/assert"
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
