package httpx

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

// 转发请求
func NewRequestForward(c *gin.Context, url string) (*http.Response, error) {
	requestBody, _ := ioutil.ReadAll(c.Request.Body)

	proxyRequest, _ := http.NewRequest(c.Request.Method, url, bytes.NewReader(requestBody))
	proxyRequest.Header = c.Request.Header

	return http.DefaultClient.Do(proxyRequest)
}

func ScanRequestBody(c *gin.Context, params interface{}) error {
	paramsErr := c.ShouldBindJSON(params)
	return paramsErr
}

func GetParamFromURI(c *gin.Context, key string) (int, error) {
	idStr := c.Param(key)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return 0, err
	}

	return id, err
}

func GetQueryNumber(c *gin.Context, key string, defaultValue int) int {
	numString := c.Query(key)
	numQuery, err := strconv.Atoi(numString)

	if err != nil {
		return defaultValue
	}
	return numQuery
}

func GetPage(c *gin.Context) (int, error) {
	page := GetQueryNumber(c, "page", 1)
	if page < 0 {
		return 0, errors.New("ERROR_INVALID_PARAMS")
	}

	return page, nil
}

func GetPageSize(c *gin.Context) (int, error) {
	pageSize := GetQueryNumber(c, "page_size", 20)
	if pageSize < 0 {
		return 0, errors.New("ERROR_INVALID_PARAMS")
	}

	return pageSize, nil
}
