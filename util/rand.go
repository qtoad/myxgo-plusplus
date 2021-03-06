package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var Rander = rand.New(rand.NewSource(time.Now().UnixNano()))

func Random(strings []string) ([]string, error) {
	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := make([]string, 0)
	for i := 0; i < len(strings); i++ {
		str = append(str, strings[i])
	}
	return str, nil
}

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

func RandOrd(n int) string {
	return time.Now().Format("20060102150405") + RandNumStr(n)
}

// 包含min, max
func RandNum(min, max int) int {
	return Rander.Intn(max-min+1) + min
}
