package xhttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 公共结构体返回
func Response(c *gin.Context, httpCode int, code int, msg string, data interface{}) {
	c.JSON(httpCode, Resp{
		Code: code,
		Msg:  msg,
		Data: data,
	})

	return
}

//成功
func RespOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Resp{
		Code: 200,
		Msg:  "OK",
		Data: data,
	})
}

//失败
func RespFAIL(c *gin.Context, code int, msg string) {
	if code == 0 {
		code = 400
	}
	if msg == "" {
		msg = "unknown"
	}
	c.JSON(http.StatusBadRequest, Resp{
		Code: code,
		Msg:  msg,
	})
}
