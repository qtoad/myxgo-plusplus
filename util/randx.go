package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var Rander = rand.New(rand.NewSource(time.Now().UnixNano()))

const letterString = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numLetterString = "0123456789"

// 随机生成字符串
func RandStr(n int, letter string) string {
	str := []byte(letter)
	res := ""
	for i := 0; i < n; i++ {
		res += fmt.Sprintf("%c", str[Rander.Intn(strings.Count(letter, "")-1)])
	}
	return res
}

func RandNumStr(n int) string {
	return RandStr(n, numLetterString)
}

func RandString(n int) string {
	return RandStr(n, letterString)
}

func RandOrder(n int) string {
	return time.Now().Format("20060102150405") + RandNumStr(n)
}

// 包含min, max
func RandNum(min, max int) int {
	return Rander.Intn(max-min+1) + min
}
