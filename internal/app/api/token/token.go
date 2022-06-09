package token

import (
	"github.com/delight-labs/terraswap-service/internal/app/api/common/terraswap"

	"github.com/gin-gonic/gin"
)

func Init(d terraswap.DataHandler, e *gin.Engine) {
	r := newRepo(d)
	s := newService(r)
	newController(e, s)
}
