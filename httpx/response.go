package httpx

import (
	"github.com/gin-gonic/gin"
	_ "net/http"
)

func Response(c *gin.Context, httpCode int, code, msg string, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})

	return
}
