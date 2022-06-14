package router

import (
	"github.com/stretchr/testify/mock"
	"github.com/terraswap/terraswap-service/internal/pkg/logging"
)

type routerMock struct {
	mock.Mock
}

var _ Router = &routerMock{}

func NewRouterMock() Router {
	return &routerMock{}
}

// Run implements Router
func (*routerMock) Run() {}

// GetLogger implements Router
func (*routerMock) GetLogger() logging.Logger {
	return nil
}

// GetRouterAddress implements Router
func (r *routerMock) GetRouterAddress() string {
	args := r.Mock.MethodCalled("GetRouterAddress")
	return args.Get(0).(string)
}

// GetRoutes implements Router
func (r *routerMock) GetRoutes(from string, to string) [][]string {
	args := r.Mock.MethodCalled("GetRoutes", from, to)
	return args.Get(0).([][]string)
}

// GetTokensFrom implements Router
func (r *routerMock) GetTokensFrom(from string, hopCount int) []string {
	args := r.Mock.MethodCalled("GetTokensFrom", from, hopCount)
	return args.Get(0).([]string)
}
