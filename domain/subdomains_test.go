package domain

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"testing"
)

func TestSubdomains(t *testing.T) {
	r := gin.Default()
	r2 := gin.Default()
	hs := make(Subdomains)
	hs["admin"] = r
	hs["analytics"] = r2
	r.GET("/ping", adminHandlerOne)
	r2.GET("/ping", adminHandlerOne)
	http.ListenAndServe(":9090", hs)
}


func adminHandlerOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func TestServeHTTP(t *testing.T) {
	r := gin.Default()
    r.GET(
		"/:username",
		func(c *gin.Context) {
    	uri, ok := c.Get("location")
	    if !ok {
	    	c.JSON(500, gin.H{
	    		"reason": "Location unknown",
		    })
	    }
	    domain := "awesome.io"
	    if uri.(*url.URL).Host == domain {
		     // Want to send client to "https://auth.awesome.io/:username"
		     s := fmt.Sprintf("https://auth.%s/%s", domain, c.Param("username"))
		     uri, err := url.Parse(s)
		     if err != nil {
			    c.JSON(500, gin.H{
				"reason": "Subdomain is wrong",
			 })
		}
		rp := new(httputil.ReverseProxy)
		rp.Director = func(req *http.Request) {
			req.URL = uri
		}
		rp.ServeHTTP(c.Writer, c.Request)
	}
}
