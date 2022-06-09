package tx

import (
	"github.com/delight-labs/terraswap-service/internal/app/api/common/terraswap"
	"github.com/delight-labs/terraswap-service/internal/pkg/logging"
	"github.com/gin-gonic/gin"
)

var logger logging.Logger

func Init(db terraswap.DataHandler, e *gin.Engine, appLogger logging.Logger) {
	logger = appLogger
	r := newRepo(db)
	s := newService(r)
	newController(e, s)
}
