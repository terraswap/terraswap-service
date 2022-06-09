package cache

import (
	"sync"

	"github.com/delight-labs/terraswap-service/configs"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap"
)

type Cacheable interface {
	terraswap.Pair | []terraswap.Pair | terraswap.Token
}

type Cache interface {
	Get(key string) interface{}
	Set(key string, v interface{}) error
	Has(key string) bool
	Delete(key string)
}

var _ Cache = &cacheImpl{}

type cacheImpl struct {
	cache sync.Map
}

func (c *cacheImpl) Get(key string) interface{} {
	v, ok := c.cache.Load(key)
	if ok {
		return v
	}
	return nil

}

func (c *cacheImpl) Set(key string, v interface{}) error {
	c.cache.Store(key, v)
	return nil
}

func (d *cacheImpl) Has(key string) bool {
	_, ok := d.cache.Load(key)
	return ok
}

func (c *cacheImpl) Delete(key string) {
	c.cache.Delete(key)
}

func New(c configs.CacheConfig) Cache {
	if c.Host == "" {
		return &cacheImpl{}
	}

	//TODO implement redis or cache service
	return nil
}
