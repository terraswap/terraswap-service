package router

import (
	"sort"

	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

type routeInfo interface {
	getAddressFromIndex(idx int) string
	getIndexFromAddress(from string) (v int, ok bool)
	getRoutesMap(fromIdx int) map[int]string
	getRoutePathsMap(fromIdx int) (destMap map[int][][]int, ok bool)
	getRoutePaths(fromIdx, toIdx int) [][]int
}

type routeInfoImpl struct {
	// from, to, pair address
	routes map[int]map[int]string
	// from, to, path indexes
	routesPath   map[int]map[int][][]int
	indexToAsset map[int]string
	assetToIndex map[string]int
}

var _ routeInfo = &routeInfoImpl{}

// newRouteInfo implements cache
func newRouteInfo(pairs []terraswap.Pair) routeInfo {
	ri := routeInfoImpl{}

	ri.setIndex(pairs)
	ri.setRoutes(pairs)
	ri.updateRoutesPath()

	return &ri
}

// getRoutePaths implements routeInfo
func (ri *routeInfoImpl) getRoutePaths(fromIdx int, toIdx int) [][]int {
	return ri.routesPath[fromIdx][toIdx]
}

// getRoutePathsMap implements routeInfo
func (ri *routeInfoImpl) getRoutePathsMap(fromIdx int) (destMap map[int][][]int, ok bool) {
	destMap, ok = ri.routesPath[fromIdx]
	return destMap, ok
}

// getRoutesMap implements routeInfo
func (ri *routeInfoImpl) getRoutesMap(fromIdx int) map[int]string {
	return ri.routes[fromIdx]
}

func (ri *routeInfoImpl) getAddressFromIndex(idx int) string {
	return ri.indexToAsset[idx]
}

func (ri *routeInfoImpl) getIndexFromAddress(from string) (v int, ok bool) {
	v, ok = ri.assetToIndex[from]
	return v, ok
}

func (ri *routeInfoImpl) setIndex(pairs []terraswap.Pair) {
	ri.indexToAsset = make(map[int]string)
	ri.assetToIndex = make(map[string]int)

	idx := 0
	for _, pair := range pairs {
		for _, asset := range pair.AssetInfos {
			token := asset.GetKey()
			if _, ok := ri.assetToIndex[token]; ok {
				continue
			}

			ri.assetToIndex[token] = idx
			ri.indexToAsset[idx] = token
			idx++
		}
	}
}

func (ri *routeInfoImpl) setRoutes(pairs []terraswap.Pair) {
	ri.routes = make(map[int]map[int]string)
	for _, p := range pairs {
		first := p.AssetInfos[0].GetKey()
		second := p.AssetInfos[1].GetKey()

		from := ri.assetToIndex[first]
		if ri.routes[from] == nil {
			ri.routes[from] = make(map[int]string)
		}
		to := ri.assetToIndex[second]
		if ri.routes[to] == nil {
			ri.routes[to] = make(map[int]string)
		}
		ri.routes[from][to] = p.ContractAddr
		ri.routes[to][from] = p.ContractAddr
	}
}

func (ri *routeInfoImpl) updateRoutesPath() {
	ri.routesPath = make(map[int]map[int][][]int)
	keys := []int{}
	visited := make(map[int]bool)
	for k := range ri.routes {
		keys = append(keys, k)
		visited[k] = false
	}

	for _, key := range keys {
		ri.routesPath[key] = make(map[int][][]int)
		visited[key] = true
		ri.updatePaths(key, key, []int{}, visited, 0)
		visited[key] = false
	}

	for _, pathMap := range ri.routesPath {
		for _, paths := range pathMap {
			sort.Slice(paths, func(i, j int) bool {
				if len(paths[i]) != len(paths[j]) {
					return len(paths[i]) < len(paths[j])
				}
				lmt := len(paths[i])
				for idx := 0; idx < lmt; idx++ {
					if paths[i][idx] == paths[j][idx] {
						continue
					}
					return paths[i][idx] < paths[j][idx]
				}
				return true
			})
		}
	}
}

func (ri *routeInfoImpl) updatePaths(start, current int, path []int, visited map[int]bool, pathLen int) {
	if pathLen > MAX_ROUTE_PATH_LEN {
		return
	}

	nodes := ri.routes[current]

	for to := range nodes {
		if v, ok := visited[to]; ok && v {
			continue
		}
		if ri.routesPath[start] == nil {
			ri.routesPath[start] = make(map[int][][]int)
		}
		if ri.routesPath[start][to] == nil {
			ri.routesPath[start][to] = make([][]int, 0)
		}

		copyPath := make([]int, len(path))
		copy(copyPath, path)
		newPath := append(copyPath, to)
		paths := ri.routesPath[start][to]
		ri.routesPath[start][to] = append(paths, newPath)
		visited[to] = true
		ri.updatePaths(start, to, newPath, visited, pathLen+1)
		visited[to] = false
	}
}
