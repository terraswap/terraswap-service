package pair

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/terraswap/terraswap-service/internal/app/api/utils/responser"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

type controller struct {
	service service
}

func (c *controller) GetPairs(con *gin.Context) {

	unverified := con.Query("unverified")
	var pairs *terraswap.Pairs

	if unverified == "true" {
		pairs = c.service.GetAllPairs()
	} else {
		pairs = c.service.GetSwapablePairs()
	}

	if pairs == nil {
		pairs = &terraswap.Pairs{
			Pairs: []terraswap.Pair{},
		}
	}

	con.JSON(http.StatusOK, pairs)
}

func (s *controller) GetPair(c *gin.Context) {
	addr := c.Param("contract-address")
	pair := s.service.GetPair(addr)
	if pair == nil {
		c.JSON(http.StatusNotFound, responser.NotFound(`Not Found`, ""))
	}
	c.JSON(http.StatusOK, pair)
}

func newController(e *gin.Engine, s service) controller {

	c := controller{
		service: s,
	}
	e.GET("/pairs", c.GetPairs)
	e.GET("/pairs/:contract-address", c.GetPair)

	return c
}
