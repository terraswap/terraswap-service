package terraswap

import (
	"fmt"
	"sort"
	"sync"

	"github.com/delight-labs/terraswap-service/configs"
	"github.com/delight-labs/terraswap-service/internal/pkg/logging"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap/cache"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap/router"
	"github.com/pkg/errors"
)

type DataHandler interface {
	GetAllPairs() []terraswap.Pair
	GetSwapablePairs() []terraswap.Pair
	GetPair(addr string) *terraswap.Pair
	GetPairByAssets(addr0, addr1 string) *terraswap.Pair
	GetTokens() terraswap.Tokens
	GetTokensFrom(addr string, hopCount int) []string
	GetToken(addr string) *terraswap.Token
	GetLogger() logging.Logger
	// Routes handler
	router.Router
}

type dataHandlerImpl struct {
	repo   repository
	routes router.Router
	cache  cache.TerraswapCache
	config configs.TerraswapConfig
	logger logging.Logger
	mutex  *sync.Mutex
}

var _ DataHandler = &dataHandlerImpl{}

func NewDataHandler(repo repository, router router.Router, cache cache.TerraswapCache, c configs.Config) DataHandler {
	name := "TerraswapService"
	logger := logging.New(name, c.Log)
	if c.Sentry.DSN != "" {
		logging.ConfigureReporter(logger, c.Sentry.DSN)
	}

	service := dataHandlerImpl{
		repo,
		router,
		cache,
		c.Terraswap,
		logger,
		&sync.Mutex{},
	}

	return &service
}

// GetPair implements DataHandler
func (s *dataHandlerImpl) GetPair(addr string) *terraswap.Pair {
	return s.cache.GetPairByAddr(addr)
}

// GetPairByAssets implements DataHandler
func (s *dataHandlerImpl) GetPairByAssets(addr0 string, addr1 string) *terraswap.Pair {
	return s.cache.GetPairByAssets(addr0, addr1)
}

// GetSwapablePairs implements DataHandler
func (s *dataHandlerImpl) GetAllPairs() []terraswap.Pair {
	pairs := s.cache.GetAllPairs()
	if pairs == nil {
		pairs = []terraswap.Pair{}
	}
	return pairs
}

// GetSwapablePairs implements DataHandler
func (s *dataHandlerImpl) GetSwapablePairs() []terraswap.Pair {
	pairs := s.cache.GetSwapablePairs()
	if pairs == nil {
		pairs = []terraswap.Pair{}
	}
	return pairs
}

// GetToken implements DataHandler
func (s *dataHandlerImpl) GetToken(addr string) *terraswap.Token {

	token := s.cache.GetToken(addr)
	if token != nil {
		return token
	}

	t, err := s.repo.getToken(addr)
	if err != nil {
		return nil
	}

	return t
}

// GetTokens implements DataHandler
func (s *dataHandlerImpl) GetTokens() terraswap.Tokens {
	return s.cache.GetAllTokens()
}

func (s *dataHandlerImpl) Run() {
	cw20Allowlist := s.repo.getCw20Allowlist(s.config.Cw20AllowlistUrl)
	ibcAllowlist := s.repo.getIbcAllowlist(s.config.IbcAllowlistUrl)

	allPairs, err := s.repo.getAllPairs()
	if err != nil {
		err := errors.Wrap(err, "terraswap.Service.Run")
		panic(err)
	}

	zeroPoolPairs, err := s.repo.getZeroPoolPairs(allPairs)
	if err != nil {
		err := errors.Wrap(err, "terraswap.Service.Run")
		panic(err)
	}

	s.mutex.Lock()

	if s.shouldUpdate(cw20Allowlist, ibcAllowlist, allPairs, zeroPoolPairs) {
		tokens := s.getTokensFromPairs(allPairs)
		markedTokens := s.markVerified(*tokens, ibcAllowlist, cw20Allowlist)
		allTokens := markedTokens
		filteredPairs, err := s.filterPairs(allPairs, zeroPoolPairs, *allTokens)
		if err != nil {
			err = errors.Wrap(err, "terraswap.Service.Run")
			panic(err)
		}

		s.cache.SetCw20AllowlistMap(cw20Allowlist)
		s.cache.SetIbcAllowlistMap(ibcAllowlist)
		s.cache.SetZeroPoolPairs(zeroPoolPairs)
		s.cache.SetSwapablePairs(filteredPairs)

		s.cacheAllTokens(*allTokens)
		s.cacheAllPairs(allPairs)
	}

	s.mutex.Unlock()
}

func (s *dataHandlerImpl) shouldUpdate(cw20WhiteList terraswap.TokensMap, ibcWhiteList terraswap.TokensMap, allPairs []terraswap.Pair, zeroPoolPairMap map[string]bool) bool {

	cachedCw20WhiteList := s.cache.GetCw20AllowlistMap()
	if !cw20WhiteList.Equal(cachedCw20WhiteList) {
		return true
	}

	cachedIbcWhiteList := s.cache.GetIbcAllowlistMap()
	if !ibcWhiteList.Equal(cachedIbcWhiteList) {
		return true
	}

	cachedZeroPoolPairMap := s.cache.GetZeroPoolPairs()
	if cachedZeroPoolPairMap == nil || len(cachedZeroPoolPairMap) != len(zeroPoolPairMap) {
		return true
	}

	for k := range zeroPoolPairMap {
		if !cachedZeroPoolPairMap[k] {
			return true
		}
	}

	cachedAllPairs := s.cache.GetAllPairs()
	if cachedAllPairs == nil {
		return true
	}

	return len(allPairs) != len(cachedAllPairs)
}

func (s *dataHandlerImpl) filterPairs(allPairs []terraswap.Pair, zeroPoolPairs map[string]bool, tokens terraswap.Tokens) (filteredPairs []terraswap.Pair, err error) {

	filteredPairs = []terraswap.Pair{}
	pairs, err := s.repo.getAllPairs()
	if err != nil {
		return nil, errors.Wrap(err, "terraswap.service.getSwapablePairs")
	}

	tokenMap := tokens.Map()

	for _, pair := range pairs {
		token0, ok0 := tokenMap[pair.AssetInfos[0].GetKey()]
		token1, ok1 := tokenMap[pair.AssetInfos[1].GetKey()]

		if !ok0 || !ok1 {
			continue
		}

		if s.config.FilterUnverified && (!token0.Verified || !token1.Verified) {
			continue
		}

		if zeroPoolPairs[pair.ContractAddr] {
			continue
		}

		filteredPairs = append(filteredPairs, pair)
	}

	return filteredPairs, nil
}

func (s *dataHandlerImpl) getTokensFromPairs(pairs []terraswap.Pair) *terraswap.Tokens {
	addresses := make(map[string]bool)
	for _, pair := range pairs {
		addresses[pair.LiquidityToken] = true
		for _, asset := range pair.AssetInfos {
			addresses[asset.GetKey()] = true
		}
	}

	tokenSlice := []terraswap.Token{}
	for addr := range addresses {
		var err error
		token := s.cache.GetToken(addr)

		if token != nil {
			tokenSlice = append(tokenSlice, *token)
			continue
		}

		token, err = s.repo.getToken(addr)
		if err != nil {
			msg := fmt.Sprintf("getTokensFromPairs: token(%s)", addr)
			s.logger.Debug(msg)
			continue
		}

		tokenSlice = append(tokenSlice, *token)
	}

	return terraswap.NewTokens(tokenSlice)
}

func (s *dataHandlerImpl) markVerified(tokens terraswap.Tokens, ibcAllowlist terraswap.TokensMap, cw20Allowlist terraswap.TokensMap) *terraswap.Tokens {
	tokensMap := terraswap.TokensMap{}
	cachedIbc := s.cache.GetIbcAllowlistMap()
	cachedCw20 := s.cache.GetCw20AllowlistMap()

	_, ibcRemoved := cachedIbc.GetDiffMap(ibcAllowlist)
	_, cw20Removed := cachedCw20.GetDiffMap(cw20Allowlist)

	for _, token := range tokens.Slice() {
		if t, ok := ibcAllowlist[token.ContractAddr]; ok {
			token = t
		}
		if t, ok := cw20Allowlist[token.ContractAddr]; ok {
			token = t
		}
		if ibcRemoved.Has(token.ContractAddr) || cw20Removed.Has(token.ContractAddr) {
			token.Verified = false
		}
		tokensMap[token.ContractAddr] = token
	}
	return tokensMap.ToTokens()
}

func (s *dataHandlerImpl) cacheAllTokens(tokens terraswap.Tokens) {
	tokenSlice := tokens.Slice()
	for _, token := range tokenSlice {
		s.cache.SetToken(token)
	}
	sort.Slice(tokenSlice, func(i, j int) bool {
		return tokenSlice[i].ContractAddr > tokenSlice[j].ContractAddr
	})
	s.cache.SetAllTokens(tokens)
}

func (s *dataHandlerImpl) cacheAllPairs(pairs []terraswap.Pair) {
	for _, p := range pairs {
		s.cache.SetPair(p)
	}
	s.cache.SetAllPairs(pairs)
}

// GetRouterAddress implements DataHandler
func (s *dataHandlerImpl) GetRouterAddress() string {
	return s.routes.GetRouterAddress()
}

// GetRoutes implements DataHandler
func (s *dataHandlerImpl) GetRoutes(from string, to string) [][]string {
	return s.routes.GetRoutes(from, to)
}

// GetTokensFrom implements DataHandler
func (s *dataHandlerImpl) GetTokensFrom(from string, hopCount int) []string {
	return s.routes.GetTokensFrom(from, hopCount)
}

func (s *dataHandlerImpl) GetLogger() logging.Logger {
	return s.logger
}
