package xhttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(code int, msg string, data interface{}) {
	g.C.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	return
}

// 公共结构体返回
func Response(c *gin.Context, httpCode int, code int, msg string, data interface{}) {
	g := Gin{C: c}
	g.Response(code, msg, data)
	return
}

//成功
func RespSUCESS(c *gin.Context, data interface{}) {
	code := http.StatusOK
	msg := http.StatusText(code)
	g := Gin{C: c}
	g.Response(code, msg, data)
	return
}

//失败
func RespFAIL(c *gin.Context, code int, msg string) {
	if code == 0 {
		code = http.StatusBadRequest
	}
	if msg == "" {
		msg = http.StatusText(http.StatusBadRequest)
	}
	g := Gin{C: c}
	g.Response(code, msg, nil)
	return
}
