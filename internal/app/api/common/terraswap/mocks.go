package terraswap

import (
	"github.com/delight-labs/terraswap-service/internal/pkg/logging"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap"
	"github.com/stretchr/testify/mock"
)

type dataHandlerMock struct {
	mock.Mock
}

var _ DataHandler = &dataHandlerMock{}

// GetAllPairs implements DataHandler
func (d *dataHandlerMock) GetAllPairs() []terraswap.Pair {
	args := d.Mock.MethodCalled("GetAllPairs")
	return args.Get(0).([]terraswap.Pair)
}

func NewDataHandlerMock() *dataHandlerMock {
	return &dataHandlerMock{}
}

// Run implements DataHandler
func (*dataHandlerMock) Run() {}

// GetLogger implements DataHandler
func (d *dataHandlerMock) GetLogger() logging.Logger {
	return nil
}

// GetRouterAddress implements DataHandler
func (d *dataHandlerMock) GetRouterAddress() string {
	args := d.Mock.MethodCalled("GetRouterAddress")
	return args.Get(0).(string)
}

// GetRoutes implements DataHandler
func (d *dataHandlerMock) GetRoutes(from string, to string) [][]string {
	args := d.Mock.MethodCalled("GetRoutes", from, to)
	return args.Get(0).([][]string)
}

// GetTokensFrom implements DataHandler
func (d *dataHandlerMock) GetTokensFrom(from string, hopCount int) []string {
	args := d.Mock.MethodCalled("GetTokensFrom", from)
	return args.Get(0).([]string)
}

// GetPair implements DataHandler
func (d *dataHandlerMock) GetPair(addr string) *terraswap.Pair {
	args := d.Mock.MethodCalled("GetPair", addr)
	return args.Get(0).(*terraswap.Pair)
}

// GetPairByAssets implements DataHandler
func (d *dataHandlerMock) GetPairByAssets(addr0 string, addr1 string) *terraswap.Pair {
	args := d.Mock.MethodCalled("GetPairByAssets", addr0)
	return args.Get(0).(*terraswap.Pair)
}

// GetSwapablePairs implements DataHandler
func (d *dataHandlerMock) GetSwapablePairs() []terraswap.Pair {
	args := d.Mock.MethodCalled("GetSwapablePairs")
	return args.Get(0).([]terraswap.Pair)
}

// GetToken implements DataHandler
func (d *dataHandlerMock) GetToken(addr string) *terraswap.Token {
	args := d.Mock.MethodCalled("GetToken", addr)
	return args.Get(0).(*terraswap.Token)
}

// GetTokens implements DataHandler
func (d *dataHandlerMock) GetTokens() terraswap.Tokens {
	args := d.Mock.MethodCalled("GetTokens")
	return args.Get(0).(terraswap.Tokens)
}

type repositoryMock struct {
	mock.Mock
}

var _ repository = &repositoryMock{}

func newRepositoryMock() *repositoryMock {
	return &repositoryMock{}
}

// getAllPairs implements repository
func (r *repositoryMock) getAllPairs() ([]terraswap.Pair, error) {
	args := r.Mock.MethodCalled("getAllPairs")
	return args.Get(0).([]terraswap.Pair), args.Error(1)
}

// getCw20Allowlist implements repository
func (r *repositoryMock) getCw20Allowlist(url string) terraswap.TokensMap {
	args := r.Mock.MethodCalled("getCw20Allowlist", url)
	return args.Get(0).(terraswap.TokensMap)
}

// getIbcAllowlist implements repository
func (r *repositoryMock) getIbcAllowlist(url string) terraswap.TokensMap {
	args := r.Mock.MethodCalled("getIbcAllowlist", url)
	return args.Get(0).(terraswap.TokensMap)
}

// getIbcDenom implements repository
func (r *repositoryMock) getIbcDenom(ibcHash string) (*terraswap.Token, error) {
	args := r.Mock.MethodCalled("getIbcDenom", ibcHash)
	return args.Get(0).(*terraswap.Token), args.Error(1)
}

// getToken implements repository
func (r *repositoryMock) getToken(addr string) (*terraswap.Token, error) {
	args := r.Mock.MethodCalled("getToken", addr)
	return args.Get(0).(*terraswap.Token), args.Error(1)
}

// getZeroPoolPairs implements repository
func (r *repositoryMock) getZeroPoolPairs(pairs []terraswap.Pair) (map[string]bool, error) {
	args := r.Mock.MethodCalled("getZeroPoolPairs", pairs)
	return args.Get(0).(map[string]bool), args.Error(1)
}
