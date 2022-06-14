package token

import (
	tsApp "github.com/terraswap/terraswap-service/internal/app/api/common/terraswap"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

type repository interface {
	GetAll() *terraswap.Tokens
	GetToken(contractAddr string) *terraswap.Token
	GetSwapableTokens(from string, hopCount int) []string
}

var _ repository = &repositoryImpl{}

type repositoryImpl struct {
	db tsApp.DataHandler
}

// GetSwapableTokens implements repository
func (r *repositoryImpl) GetSwapableTokens(from string, hopCount int) []string {
	return r.db.GetTokensFrom(from, hopCount)
}

//Fixme change db interface
func newRepo(db tsApp.DataHandler) repository {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) GetAll() *terraswap.Tokens {
	tokens := r.db.GetTokens()
	if len(tokens.Map()) == 0 {
		return nil
	}
	return &tokens
}

func (r *repositoryImpl) GetToken(contractAddr string) *terraswap.Token {
	return r.db.GetToken(contractAddr)
}
