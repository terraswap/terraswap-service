package tx

import (
	"net/http"

	request_config "github.com/delight-labs/terraswap-service/internal/app/api/common/request-config"
	"github.com/delight-labs/terraswap-service/internal/app/api/utils/responser"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service Service
}

func (c *controller) GetSwapTxs(con *gin.Context) {
	dto := &SwapTxRequest{}
	dto.HopCount = request_config.DEFAULT_HOP_COUNT

	if err := con.ShouldBind(dto); err != nil {
		logger.Debug(err.Error())
		con.JSON(http.StatusBadRequest, responser.GetBadRequest("Bad Request", err.Error()))
		return
	}

	if dto.HopCount < 0 {
		con.JSON(http.StatusBadRequest, responser.GetBadRequest("Bad Request", "hop_count shouldn't be negative"))
		return
	}

	resBody, resErr := c.service.GetSwapTxs(dto.From, dto.To, dto.Amount, dto.Sender, dto.MaxSpread, dto.BeliefPrice, dto.HopCount)

	if resErr != nil {
		con.JSON(resErr.Code, resErr)
		return
	}

	con.JSON(http.StatusOK, resBody)
}

func (c *controller) GetProvideTx(con *gin.Context) {
	dto := &ProvideTxRequest{}
	if err := con.ShouldBind(dto); err != nil {
		logger.Debug(err.Error())
		con.JSON(http.StatusBadRequest, responser.GetBadRequest("Bad Request", err.Error()))
		return
	}

	unsignedtxs, resErr := c.service.GetProvideTx(dto.From, dto.To, dto.FromAmount, dto.ToAmount, dto.Slippage, dto.Sender)
	if resErr != nil {
		con.JSON(resErr.Code, resErr)
		return
	}

	con.JSON(http.StatusOK, unsignedtxs)
}

func (c *controller) GetWithdrawTx(con *gin.Context) {
	dto := &WithdrawTxRequest{}
	if err := con.ShouldBind(dto); err != nil {
		logger.Debug(err.Error())
		con.JSON(http.StatusBadRequest, responser.GetBadRequest("Bad Request", err.Error()))
		return
	}

	unsignedtxs, resErr := c.service.GetWithdrawTx(dto.LpAddr, dto.Amount, dto.Sender)
	if resErr != nil {
		con.JSON(resErr.Code, resErr)
		return
	}

	con.JSON(http.StatusOK, unsignedtxs)
}

func newController(e *gin.Engine, s Service) controller {

	c := controller{
		service: s,
	}

	e.GET("/tx/swap", c.GetSwapTxs)
	e.GET("/tx/provide", c.GetProvideTx)
	e.GET("/tx/withdraw", c.GetWithdrawTx)

	return c
}
