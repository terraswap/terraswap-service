package token

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	request_config "github.com/delight-labs/terraswap-service/internal/app/api/common/request-config"
	"github.com/delight-labs/terraswap-service/internal/app/api/utils/responser"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetAllTokens(c *gin.Context)
	GetToken(c *gin.Context)
	GetSwapableTokens(con *gin.Context)
}

type controller struct {
	service service
}

var _ Controller = &controller{}

func (c *controller) GetAllTokens(con *gin.Context) {

	tokens := c.service.GetAllTokens()
	if tokens == nil {
		con.JSON(http.StatusNotFound, responser.NotFound("Not Found", "Not Found"))
		return
	}
	con.JSON(http.StatusOK, tokens.Slice())
}

func (c *controller) GetToken(con *gin.Context) {
	addr := con.Param("contract-address")

	if addr == "" {
		con.JSON(http.StatusBadRequest, errors.New("contract-address param required"))
		return
	}

	token := c.service.GetToken(addr)
	if token == nil {
		msg := fmt.Sprintf("%s not found", addr)
		con.JSON(http.StatusNotFound, responser.NotFound(msg, "Not Found"))
		return
	}

	con.JSON(http.StatusOK, token)
}

func (c *controller) GetSwapableTokens(con *gin.Context) {
	from, ok := con.GetQuery("from")
	if !ok || from == "" {
		con.JSON(http.StatusBadRequest, errors.New("from query param required"))
		return
	}
	hopCount := request_config.DEFAULT_HOP_COUNT
	hopCountStr, _ := con.GetQuery("hopCount")

	if regexp.MustCompile(`\d`).MatchString(hopCountStr) {
		val, err := strconv.ParseUint(hopCountStr, 10, 32)
		if err != nil || val > request_config.MAX_HOP_COUNT {
			msg := fmt.Sprintf("hopCount(%s) is invalid. It must be between 0 to %d", hopCountStr, request_config.MAX_HOP_COUNT)
			con.JSON(http.StatusBadRequest, responser.GetBadRequest(msg, msg))
			return
		}
		hopCount = int(val)
	}

	tokens := c.service.GetSwapableTokens(from, hopCount)
	if tokens == nil {
		tokens = []string{}
	}

	con.JSON(http.StatusOK, tokens)
}

func newController(e *gin.Engine, s service) controller {

	c := controller{
		service: s,
	}
	e.GET("/tokens/swap", c.GetSwapableTokens)

	e.GET("/tokens/:contract-address", c.GetToken)
	e.GET("/token/:contract-address", c.GetToken)

	e.GET("/tokens", c.GetAllTokens)

	return c
}
