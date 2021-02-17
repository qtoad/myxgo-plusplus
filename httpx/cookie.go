package httpx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

// 设置Cookie
func SetCookie(c *gin.Context, domain, name, value string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		MaxAge:   -1,
		Path:     "/",
		Domain:   domain,
		Secure:   false,
		HttpOnly: true,
	})
}
