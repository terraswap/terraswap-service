package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type statusError struct {
	status int
	Err    error
}

func codedErrorHandle(c *gin.Context, err interface{}) {
	if e, ok := err.(statusError); ok {
		c.JSON(e.status, e.Err.Error())
	}
	c.JSON(http.StatusInternalServerError, "Internal Server Error")
}
