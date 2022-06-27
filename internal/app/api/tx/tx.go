package tx

import (
	"github.com/gin-gonic/gin"
	"github.com/terraswap/terraswap-service/internal/app/api/common/terraswap"
	"github.com/terraswap/terraswap-service/internal/pkg/logging"
)

var logger logging.Logger

func Init(db terraswap.DataHandler, e *gin.Engine, appLogger logging.Logger, isClassic bool) {
	logger = appLogger
	r := newRepo(db)
	s := newService(r)
	if isClassic {
		s = newClassicService(r)
	}
	newController(e, s)
}
