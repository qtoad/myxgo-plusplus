package chn

import (
	"encoding/hex"
	"strconv"
)

// 获取中文字符串第一个首字母
func GetChineseFirstLetter(hanzi string) string {
	// 获取中文字符串第一个字符
	firstChar := string([]rune(hanzi)[:1])
	// Utf8 转 GBK2312
	firstCharGbk, err := Utf8ToGbk([]byte(firstChar))
	if err != nil {
		return ""
	}
	// 获取第一个字符的16进制
	firstCharHex := hex.EncodeToString(firstCharGbk)
	// 16进制转十进制
	firstCharDec, err := strconv.ParseInt(firstCharHex, 16, 0)
	if err != nil {
		return ""
	}
	// 十进制落在GB 2312的某个拼音区间即为某个字母
	firstCharDecRelative := firstCharDec - 65536
	if firstCharDecRelative >= -20319 && firstCharDecRelative <= -20284 {
		return "A"
	}
	if firstCharDecRelative >= -20283 && firstCharDecRelative <= -19776 {
		return "B"
	}
	if firstCharDecRelative >= -19775 && firstCharDecRelative <= -19219 {
		return "C"
	}
	if firstCharDecRelative >= -19218 && firstCharDecRelative <= -18711 {
		return "D"
	}
	if firstCharDecRelative >= -18710 && firstCharDecRelative <= -18527 {
		return "E"
	}
	if firstCharDecRelative >= -18526 && firstCharDecRelative <= -18240 {
		return "F"
	}
	if firstCharDecRelative >= -18239 && firstCharDecRelative <= -17923 {
		return "G"
	}
	if firstCharDecRelative >= -17922 && firstCharDecRelative <= -17418 {
		return "H"
	}
	if firstCharDecRelative >= -17417 && firstCharDecRelative <= -16475 {
		return "J"
	}
	if firstCharDecRelative >= -16474 && firstCharDecRelative <= -16213 {
		return "K"
	}
	if firstCharDecRelative >= -16212 && firstCharDecRelative <= -15641 {
		return "L"
	}
	if firstCharDecRelative >= -15640 && firstCharDecRelative <= -15166 {
		return "M"
	}
	if firstCharDecRelative >= -15165 && firstCharDecRelative <= -14923 {
		return "N"
	}
	if firstCharDecRelative >= -14922 && firstCharDecRelative <= -14915 {
		return "O"
	}
	if firstCharDecRelative >= -14914 && firstCharDecRelative <= -14631 {
		return "P"
	}
	if firstCharDecRelative >= -14630 && firstCharDecRelative <= -14150 {
		return "Q"
	}
	if firstCharDecRelative >= -14149 && firstCharDecRelative <= -14091 {
		return "R"
	}
	if firstCharDecRelative >= -14090 && firstCharDecRelative <= -13319 {
		return "S"
	}
	if firstCharDecRelative >= -13318 && firstCharDecRelative <= -12839 {
		return "T"
	}
	if firstCharDecRelative >= -12838 && firstCharDecRelative <= -12557 {
		return "W"
	}
	if firstCharDecRelative >= -12556 && firstCharDecRelative <= -11848 {
		return "X"
	}
	if firstCharDecRelative >= -11847 && firstCharDecRelative <= -11056 {
		return "Y"
	}
	if firstCharDecRelative >= -11055 && firstCharDecRelative <= -10247 {
		return "Z"
	}
	return ""
}
