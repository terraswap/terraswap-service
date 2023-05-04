package router

import (
	"sort"
	"sync"

	"github.com/terraswap/terraswap-service/configs"
	"github.com/terraswap/terraswap-service/internal/pkg/logging"
	"github.com/terraswap/terraswap-service/internal/pkg/repeater"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

type Router interface {
	GetTokensFrom(from string, hopCount int) []string
	GetRoutes(from, to string) [][]string
	GetRouterAddress() string
	GetLogger() logging.Logger
	repeater.Runner
}

var _ Router = &routerImpl{}

type routerImpl struct {
	Repository
	currentPairs  []terraswap.Pair
	routerAddress string
	logger        logging.Logger
	mutex         *sync.Mutex
}

func New(repo Repository, c configs.Config) Router {
	name := "TerraswapRouter"
	logger := logging.New(name, c.Log)
	r := &routerImpl{
		logger:        logger,
		Repository:    repo,
		routerAddress: terraswap.GetRouterAddress(c.Terraswap.ChainId, c.Terraswap.Version),
		mutex:         &sync.Mutex{},
	}
	return r
}
func (r *routerImpl) GetName() string {
	return "routes"
}

func (r *routerImpl) GetRouterAddress() string {
	return r.routerAddress
}

func (r *routerImpl) GetTokensFrom(from string, hopCount int) []string {
	routeInfo := r.GetRouteInfo()
	if routeInfo == nil {
		return nil
	}

	fromIdx, fromOk := routeInfo.getIndexFromAddress(from)
	pathLen := hopCount + 1

	destMap, ok := routeInfo.getRoutePathsMap(fromIdx)
	if !fromOk || !ok {
		return nil
	}
	tokenMap := make(map[string]bool)

	for k, paths := range destMap {
		for _, path := range paths {
			if len(path) > pathLen {
				continue
			}
			tokenMap[routeInfo.getAddressFromIndex(k)] = true
		}
	}

	tokens := make([]string, 0, len(tokenMap))
	for k := range tokenMap {
		tokens = append(tokens, k)
	}
	sort.Strings(tokens)
	return tokens
}

func (r *routerImpl) GetRoutes(from, to string) [][]string {
	currentRoute := r.GetRouteInfo()
	if currentRoute == nil {
		return nil
	}
	fromIdx, fromOk := currentRoute.getIndexFromAddress(from)
	toIdx, toOk := currentRoute.getIndexFromAddress(to)
	if !fromOk || !toOk || currentRoute.getRoutesMap(fromIdx) == nil || currentRoute.getRoutePaths(fromIdx, toIdx) == nil {
		return nil
	}

	pathsIndexes := currentRoute.getRoutePaths(fromIdx, toIdx)
	pathsArr := [][]string{}
	for _, pathIdx := range pathsIndexes {
		paths := []string{}
		for _, pathIdx := range pathIdx {
			paths = append(paths, currentRoute.getAddressFromIndex(pathIdx))
		}
		pathsArr = append(pathsArr, paths)
	}
	return pathsArr
}

func (r *routerImpl) getSwapablePairs() []terraswap.Pair {
	return r.GetSwapablePairs()
}

func (r *routerImpl) Run() {

	r.mutex.Lock()

	if r.shouldUpdate() {
		pairs := r.getSwapablePairs()
		ri := r.CreateRouteInfo(pairs)
		r.SetRouteInfo(ri)
		r.currentPairs = pairs
	}

	r.mutex.Unlock()
}

func (r *routerImpl) shouldUpdate() bool {
	pairs := r.getSwapablePairs()
	if len(r.currentPairs) != len(pairs) {
		return true
	}
	lmt := len(pairs)
	for idx := 0; idx < lmt; idx++ {
		if pairs[idx].ContractAddr != r.currentPairs[idx].ContractAddr {
			return true
		}
	}
	return false
}

func (r *routerImpl) GetLogger() logging.Logger {
	return r.logger
}
