package pair

import (
	"github.com/gin-gonic/gin"
	"github.com/terraswap/terraswap-service/internal/app/api/common/terraswap"
)

func Init(d terraswap.DataHandler, e *gin.Engine) {
	r := newRepo(d)
	s := newService(r)
	newController(e, s)
}
