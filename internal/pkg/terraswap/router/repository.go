package router

import (
	"github.com/delight-labs/terraswap-service/internal/pkg/cache"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap"
	"github.com/pkg/errors"
)

type terraswapCacheStore interface {
	GetSwapablePairs() []terraswap.Pair
	cache.Cache
}

type Repository interface {
	GetSwapablePairs() []terraswap.Pair
	CreateRouteInfo(pairs []terraswap.Pair) routeInfo
	GetRouteInfo() routeInfo
	SetRouteInfo(routeInfo) error
}

type repositoryImpl struct {
	cache terraswapCacheStore
}

var _ Repository = &repositoryImpl{}

func NewRepo(cache terraswapCacheStore) Repository {
	return &repositoryImpl{cache}
}

// CreateRouteInfo implements Repository
func (r *repositoryImpl) CreateRouteInfo(pairs []terraswap.Pair) routeInfo {
	return newRouteInfo(pairs)
}

// GetRouteInfo implements Repository
func (r *repositoryImpl) GetRouteInfo() routeInfo {
	v := r.cache.Get(routeInfoKey)
	ri, ok := v.(routeInfo)
	if !ok {
		return nil
	}
	return ri

}

// GetSwapablePairs implements Repository
func (r *repositoryImpl) GetSwapablePairs() []terraswap.Pair {
	return r.cache.GetSwapablePairs()
}

// SetRouteInfo implements Repository
func (r *repositoryImpl) SetRouteInfo(ri routeInfo) error {
	if err := r.cache.Set(routeInfoKey, ri); err != nil {
		return errors.Wrap(err, "SetRouteInfo")
	}
	return nil
}

const routeInfoKey = "route_info"
