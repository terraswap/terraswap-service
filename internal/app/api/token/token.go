package token

import (
	"github.com/terraswap/terraswap-service/internal/app/api/common/terraswap"

	"github.com/gin-gonic/gin"
)

func Init(d terraswap.DataHandler, e *gin.Engine, isClassic bool) {
	r := newRepo(d)
	s := newService(r)
	if isClassic {
		s = newClassicService(r)
	}
	newController(e, s)
}
