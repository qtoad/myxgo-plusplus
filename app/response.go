package app

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/qtoad/myxgo-plusplus/xhttp"
)

// 公共结构体返回
func Response(c *gin.Context, httpCode int, code int, msg string, data interface{}) {
	xhttp.Response(c, httpCode, code, msg, data)
	return
}

//成功
func RespSUCESS(c *gin.Context, data interface{}) {
	xhttp.RespSUCESS(c, data)
	return
}

//失败
func RespFAIL(c *gin.Context, code int, msg string) {
	xhttp.RespFAIL(c, code, msg)
	return
}
