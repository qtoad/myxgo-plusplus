package http

import (
	"io/ioutil"
	"net/http"
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
