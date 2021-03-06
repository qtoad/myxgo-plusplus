package xhttp

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func POST(url string, data string) (body string, err error) {
	resp, err1 := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(data))
	defer resp.Body.Close()
	if err1 != nil {
		err = err1
	}
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		err = err2
	}
	body = string(b)
	return
}

func GET(url string) (body string, err error) {
	resp, err1 := http.Get(url)
	defer resp.Body.Close()
	if err1 != nil {
		err = err1
	}
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		err = err2
	}
	body = string(b)
	return
}

/*
 * 上传文件
 *  */
func PostFile(url, filename, fileTag string, params map[string]string) (int, map[string][]string, string, error) {
	if !filepath.IsAbs(filename) {
		filename, _ = filepath.Abs(filename)
	}
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, "", err
	}
	defer file.Close()

	if fileTag == "" {
		fileTag = "img"
	}

	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	formFile, err := bodyWriter.CreateFormFile(fileTag, filepath.Base(filename))
	if err != nil {
		return 0, nil, "", err
	}

	//写入file数据
	_, err = io.Copy(formFile, file)
	if err != nil {
		return 0, nil, "", err
	}

	//写入参数
	for key, val := range params {
		_ = bodyWriter.WriteField(key, val)
	}

	contentType := bodyWriter.FormDataContentType()
	err = bodyWriter.Close()
	if err != nil {
		return 0, nil, "", err
	}

	//http post
	resp, err := http.Post(url, contentType, bodyBuffer)
	if err != nil {
		return 0, nil, "", err
	}
	defer resp.Body.Close()

	//状态码
	statusCode, _ := strconv.Atoi(resp.Status)

	header := resp.Header

	//获取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, "", err
	}

	return statusCode, header, string(body), err
}

/*
 * 获取url里指定参数的值
 *  */
func GetUrlParam(sourceUrl, paramName string) string {
	paramValue := ""
	if urlParser, err := url.Parse(sourceUrl); err == nil {
		values := urlParser.Query()
		for k, v := range values {
			if k == paramName {
				paramValue = v[0]
				break
			}
		}
	}

	return paramValue
}

/*
 * 字典转URL查询字符串
 *  */
func ToQueryString(values map[string]interface{}, args ...bool) string {
	isEncode := false
	queryString := ""

	if len(args) > 0 {
		isEncode = args[0]
	}

	for k, v := range values {
		if isEncode {
			v = QueryEncode(fmt.Sprintf("%v", v))
		}
		queryString = queryString + fmt.Sprintf("%s=%v&", k, v)
	}
	queryString = queryString[0 : len(queryString)-1]

	return queryString
}

/*
 * Query编码
 *  */
func QueryEncode(value string) string {
	encodeValue := value
	if value != "" {
		encodeValue = url.QueryEscape(value)
	}

	return encodeValue
}

/*
 * Url编码
 *  */
func UrlEncode(sourceUrl string) string {
	encodeUrl := sourceUrl
	if sourceUrl != "" {
		if urlParser, err := url.Parse(sourceUrl); err == nil {
			encodeUrl = urlParser.EscapedPath()
		}
	}

	return encodeUrl
}

/*
 * Url解码
 *  */
func UrlDecode(sourceUrl string) string {
	decodeUrl := sourceUrl
	if sourceUrl != "" {
		if urlParser, err := url.Parse(sourceUrl); err == nil {
			decodeUrl = urlParser.Path
		}
	}

	return decodeUrl
}

/*
 * 判断是否本地Ip
 *  */
func IsLocalIp(ip string) (bool, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return false, err
	}

	for i := range addrs {
		intf, _, err := net.ParseCIDR(addrs[i].String())
		if err != nil {
			return false, err
		}

		if net.ParseIP(ip).Equal(intf) {
			return true, nil
		}
	}
	return false, nil
}
