package util

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// golang 切割空字符串依然会得到包含一个空元素的数组 => [""]
func Split(str, sep string) []string {
	if str != "" {
		return strings.Split(str, sep)
	} else {
		return make([]string, 0)
	}
}

func ReplaceAll(s string, oldnew ...string) string {
	repl := strings.NewReplacer(oldnew...)
	return repl.Replace(s)
}

func ReplaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

func InArray(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func HtmlToPureText(html string) string {
	reHTML, _ := regexp.Compile("<[^>]*>")
	return reHTML.ReplaceAllString(html, "")
}

// 截取文本 支持中文
func SubString(src string, start, end int) string {
	runSrc := []rune(src)
	maxLen := len(runSrc)
	validEnd := Min(maxLen, end)
	return string(runSrc[start:validEnd])
}

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 时间戳转时间
func UnixToTime(e string) (datatime time.Time, err error) {
	data, err := strconv.ParseInt(e, 10, 64)
	datatime = time.Unix(data/1000, 0)
	return
}

// 时间转时间戳
func TimeToUnix(e time.Time) int64 {
	timeUnix, _ := time.Parse("2006-01-02 15:04:05", e.Format("2006-01-02 15:04:05"))
	return timeUnix.UnixNano() / 1e6
}

func GetCurrentTimeUnix() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetCurrentTime() time.Time {
	return time.Now()
}

// Copied from golint
var commonInitialisms = []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}
var commonInitialismsReplacer *strings.Replacer

func init() {
	var commonInitialismsForReplacer []string
	for _, initialism := range commonInitialisms {
		commonInitialismsForReplacer = append(commonInitialismsForReplacer, initialism, strings.Title(strings.ToLower(initialism)))
	}
	commonInitialismsReplacer = strings.NewReplacer(commonInitialismsForReplacer...)
}

// ToLowerUnderlinedNamer 转换为小写下划线命名
func ToLowerUnderlinedNamer(name string) string {
	const (
		lower = false
		upper = true
	)

	if name == "" {
		return ""
	}

	var (
		value                                    = commonInitialismsReplacer.Replace(name)
		buf                                      = bytes.NewBufferString("")
		lastCase, currCase, nextCase, nextNumber bool
	)

	for i, v := range value[:len(value)-1] {
		nextCase = bool(value[i+1] >= 'A' && value[i+1] <= 'Z')
		nextNumber = bool(value[i+1] >= '0' && value[i+1] <= '9')

		if i > 0 {
			if currCase == upper {
				if lastCase == upper && (nextCase == upper || nextNumber == upper) {
					buf.WriteRune(v)
				} else {
					if value[i-1] != '_' && value[i+1] != '_' {
						buf.WriteRune('_')
					}
					buf.WriteRune(v)
				}
			} else {
				buf.WriteRune(v)
				if i == len(value)-2 && (nextCase == upper && nextNumber == lower) {
					buf.WriteRune('_')
				}
			}
		} else {
			currCase = upper
			buf.WriteRune(v)
		}
		lastCase = currCase
		currCase = nextCase
	}

	buf.WriteByte(value[len(value)-1])

	s := strings.ToLower(buf.String())
	return s
}

/*
 * 获取当前唯一Token字符串
 *  */
func GetTokenString() string {
	timestamp := UnixNanoTimestamp()
	return Md5(strconv.FormatInt(timestamp, 10))
}

/*
 * 判断内容字符串头是否包含指定的字符串
 *  */
func HasPrefix(content, target string) bool {
	isPrefix := false

	if len(content) > 0 && len(target) > 0 {
		isPrefix = strings.HasPrefix(content, target)
	}

	return isPrefix
}

/*
 * 判断内容字符串尾是否包含指定的字符串
 *  */
func HasSuffix(content, target string) bool {
	isSuffix := false

	if len(content) > 0 && len(target) > 0 {
		isSuffix = strings.HasSuffix(content, target)
	}

	return isSuffix
}

/*
 * 判断内容字符串头尾是否包含指定的字符串
 *  */
func HasPrefixSuffix(content, target string) bool {
	isPrefixSuffix := false

	if len(content) > 0 && len(target) > 0 {
		isPrefix := strings.HasPrefix(content, target)
		isSuffix := strings.HasSuffix(content, target)

		isPrefixSuffix = isPrefix && isSuffix
	}

	return isPrefixSuffix
}

/*
 * 字符串替换
 * sourceString: 原始字符串
 * args[0...n-2]: 被替换字符串集合
 * args[n-1]: 替换字符串
 *  */
func StringReplace(sourceString string, args ...string) string {
	target := ""
	replaces := []string{}
	result := sourceString
	count := len(args)

	if len(result) > 0 && count > 0 {
		if count == 1 {
			replaces = append(replaces, args[0])
		} else if count > 1 {
			for index, value := range args {
				if index == count-1 {
					target = value
				} else {
					replaces = append(replaces, value)
				}
			}
		}

		for _, value := range replaces {
			result = strings.Replace(result, value, target, -1)
		}
	}

	return result
}

/*
 * 获取字符串个数（不是字节数）
 *  */
func GetStringCount(sourceString string) int {
	if sourceString == "" {
		return 0
	}
	return utf8.RuneCountInString(sourceString)
}

/*
 * 获取指定长度的字符串
 *  */
func GetSubString(sourceString string, count int, args ...string) string {
	if len(sourceString) <= count {
		return sourceString
	}

	newString, sourceStringRune := "", []rune(sourceString)
	sl, rl := 0, 0

	more := ""
	isRune := true

	if len(args) > 0 {
		more = args[0]
	}

	for _, r := range sourceStringRune {
		if isRune {
			rl = 1
		} else {
			if int(r) < 128 {
				rl = 1
			} else {
				rl = 2
			}
		}

		if sl+rl > count {
			break
		}

		sl += rl

		newString += string(r)
	}

	if sl < len(sourceStringRune) {
		if len(more) > 0 {
			newString += more
		}
	}

	return newString
}

/*
 * 过滤主机协议
 *  */
func FilterHostProtocol(path string) string {
	if len(path) > 0 {
		path = strings.Trim(path, " ")

		if paths := StringToStringSlice(path, ":"); len(paths) > 1 {
			path = paths[1]
		}

		path = strings.TrimPrefix(path, "//")
		path = strings.TrimPrefix(path, "/")
		path = strings.TrimSuffix(path, "/")
	}

	return path
}

/*
 * 获取指定个数的uint64slice
 *  */
func GetUint64SliceRange(int64Slice []uint64, count int) []uint64 {
	length := len(int64Slice)
	if length > count-1 {
		length = count
	}
	return int64Slice[0:length]
}

/*
 * 翻转字符串
 *  */
func ReverseString(sourceString string) string {
	sourceRunes := []rune(sourceString)
	for from, to := 0, len(sourceRunes)-1; from < to; from, to = from+1, to-1 {
		sourceRunes[from], sourceRunes[to] = sourceRunes[to], sourceRunes[from]
	}
	return string(sourceRunes)
}

/*
 * 翻转uint64 Slice
 *  */
func ReverseUint64Slice(int64Slice []uint64) {
	for from, to := 0, len(int64Slice)-1; from < to; from, to = from+1, to-1 {
		int64Slice[from], int64Slice[to] = int64Slice[to], int64Slice[from]
	}
}

/*
 * 字符串转换为bool
 *  */
func StringToBool(stringValue string) bool {
	var boolValue bool = true

	if len(stringValue) == 0 ||
		stringValue == "0" ||
		strings.ToLower(stringValue) == "f" ||
		strings.ToLower(stringValue) == "false" {

		boolValue = false
	}

	return boolValue
}

/*
 * 字符串转换为int32
 *  */
func StringToInt32(stringValue string) int32 {
	var intValue int64 = 0

	if len(stringValue) == 0 {
		return int32(intValue)
	}

	intValue, err := strconv.ParseInt(stringValue, 10, 32)
	if err != nil {
		return 0
	}

	return int32(intValue)
}

/*
 * 字符串转换为uint32
 *  */
func StringToUint32(stringValue string) uint32 {
	var uintValue uint64 = 0

	if len(stringValue) == 0 {
		return uint32(uintValue)
	}

	uintValue, err := strconv.ParseUint(stringValue, 10, 32)
	if err != nil {
		return 0
	}

	return uint32(uintValue)
}

/*
 * 字符串转换为uint64
 *  */
func StringToUint64(stringValue string) uint64 {
	var uintValue uint64 = 0

	if len(stringValue) == 0 {
		return uintValue
	}

	uintValue, err := strconv.ParseUint(stringValue, 10, 64)
	if err != nil {
		return 0
	}

	return uintValue
}

/*
 * 字符串转换为int64
 *  */
func StringToInt64(stringValue string) int64 {
	var intValue int64 = 0

	if len(stringValue) == 0 {
		return intValue
	}

	intValue, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		return 0
	}

	return intValue
}

/*
 * 字符串转换为float64
 *  */
func StringToFloat64(stringValue string) float64 {
	var floatValue float64 = 0.0

	if len(stringValue) == 0 {
		return floatValue
	}

	floatValue, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return 0.0
	}

	return floatValue
}

/*
 * 用指定的字符串把源字符串分隔为uint64切片
 *  */
func StringToUint64Slice(sourceString string, args ...string) []uint64 {
	result := make([]uint64, 0)

	if len(sourceString) == 0 {
		return result
	}

	splitString := ","
	if len(args) == 1 {
		splitString = args[0]
	}

	stringSlice := strings.Split(sourceString, splitString)
	for _, v := range stringSlice {
		if value, err := strconv.ParseUint(v, 10, 64); err == nil {
			result = append(result, value)
		}
	}

	return result
}

/*
 * 用指定的字符串把源字符串分隔为int64切片
 *  */
func StringToInt64Slice(sourceString string, args ...string) []int64 {
	result := make([]int64, 0)

	if len(sourceString) == 0 {
		return result
	}

	splitString := ","
	if len(args) == 1 {
		splitString = args[0]
	}

	stringSlice := strings.Split(sourceString, splitString)
	for _, v := range stringSlice {
		if value, err := strconv.ParseInt(v, 10, 64); err == nil {
			result = append(result, value)
		}
	}

	return result
}

/*
 * 用指定的字符串把源字符串分隔为int切片
 *  */
func StringToIntSlice(sourceString string, args ...string) []int {
	result := make([]int, 0)

	if len(sourceString) == 0 {
		return result
	}

	splitString := ","
	if len(args) == 1 {
		splitString = args[0]
	}

	stringSlice := strings.Split(sourceString, splitString)
	for _, v := range stringSlice {
		if value, err := strconv.Atoi(v); err == nil {
			result = append(result, value)
		}
	}

	return result
}

/*
 * 字符串切片转为整型64切片
 *  */
func StringSliceToInt64Slice(values []string) []int64 {
	results := make([]int64, 0)

	for _, value := range values {
		valueInt64 := StringToInt64(value)
		results = append(results, valueInt64)
	}

	return results
}

/*
 * 字符串切片转为无符号整型64切片
 *  */
func StringSliceToUint64Slice(values []string) []uint64 {
	results := make([]uint64, 0)

	for _, value := range values {
		valueInt64 := StringToUint64(value)
		results = append(results, valueInt64)
	}

	return results
}

/*
 * 整型64切片转为字符串切片
 *  */
func Int64SliceToStringSlice(values []int64) []string {
	results := make([]string, 0)

	for _, value := range values {
		valueString := Int64ToString(value)
		results = append(results, valueString)
	}

	return results
}

/*
 * 无符号整型64切片转为字符串切片
 *  */
func Uint64SliceToStringSlice(values []uint64) []string {
	results := make([]string, 0)

	for _, value := range values {
		valueString := Uint64ToString(value)
		results = append(results, valueString)
	}

	return results
}

/*
 * 无符号整型64转为字符串
 *  */
func Uint64ToString(value uint64) string {
	result := fmt.Sprintf("%d", value)
	return result
}

/*
 * 整型64转为字符串
 *  */
func Int64ToString(value int64) string {
	result := fmt.Sprintf("%d", value)
	return result
}

/*
 * 浮点64转为字符串
 *  */
func Float64ToString(value float64) string {
	result := fmt.Sprintf("%f", value)
	return result
}

/*
 * 用指定的字符串把uint64切片链接为字符串
 *  */
func Uint64SliceToString(uintSlice []uint64, args ...string) string {
	result := ""

	if len(uintSlice) == 0 {
		return result
	}

	joinString := ","
	if len(args) == 1 {
		joinString = args[0]
	}

	count := len(uintSlice)
	if count == 1 {
		result = fmt.Sprintf("%d", uintSlice[0])
	} else if count > 1 {
		for _, value := range uintSlice {
			valueString := fmt.Sprintf("%d", value)
			if len(result) == 0 {
				result = result + valueString
			} else {
				result = result + joinString + valueString
			}
		}
	}

	return result
}

/*
 * 用指定的字符串把int64切片链接为字符串
 *  */
func Int64SliceToString(intSlice []int64, args ...string) string {
	result := ""

	if len(intSlice) == 0 {
		return result
	}

	joinString := ","
	if len(args) == 1 {
		joinString = args[0]
	}

	count := len(intSlice)
	if count == 1 {
		result = fmt.Sprintf("%d", intSlice[0])
	} else if count > 1 {
		for _, value := range intSlice {
			valueString := fmt.Sprintf("%d", value)
			if len(result) == 0 {
				result = result + valueString
			} else {
				result = result + joinString + valueString
			}
		}
	}

	return result
}

/*
 * 用指定的字符串分隔源字符串为字符串切片
 *  */
func StringToStringSlice(sourceString string, args ...string) []string {
	result := make([]string, 0)

	if len(sourceString) == 0 {
		return result
	}

	splitString := ","
	if len(args) == 1 {
		splitString = args[0]
	}

	stringSlice := strings.Split(sourceString, splitString)
	for _, v := range stringSlice {
		if v != "" {
			result = append(result, v)
		}
	}

	return result
}

/*
 * 用指定的字符串链接字符串切片
 *  */
func StringSliceToString(stringSlice []string, args ...string) string {
	result := ""

	if len(stringSlice) == 0 {
		return result
	}

	joinString := ","
	if len(args) == 1 {
		joinString = args[0]
	}

	if len(stringSlice) == 1 {
		result = strings.Join(stringSlice, "")
	} else {
		for _, v := range stringSlice {
			if len(result) == 0 {
				result = result + v
			} else {
				result = result + joinString + v
			}
		}
	}

	return result
}

/*
 * 保留指定长度字符串切片，前面的数据移除
 *  */
func StringSliceLatest(srcSlice []string, maxCount int) []string {
	destSlice := srcSlice
	count := len(destSlice)
	if count > maxCount {
		offsetIndex := count - maxCount
		destSlice = destSlice[offsetIndex:count]
	}

	return destSlice
}

/*
 * 判断字符串切片及单个项的字符数是否匹配指定大小，
 *  */
func IsMatchStringSliceCount(srcSlice []string, maxCount, stringItemCount int) bool {
	isMatch := true
	srcSliceCount := len(srcSlice)

	if srcSliceCount > 0 {
		if srcSliceCount > maxCount {
			isMatch = false
		}

		for _, stringItem := range srcSlice {
			if GetStringCount(stringItem) > stringItemCount {
				isMatch = false
				break
			}
		}
	}

	return isMatch
}

/*
 * 过滤int64数组（从all中过滤所有other中的数据，返回未被过滤的数据集合）
 *  */
func FilterInt64Slice(all, other []int64) []int64 {
	allMaps := make(map[int64]bool, 0)
	sliceResult := make([]int64, 0)

	for _, v := range all {
		allMaps[v] = true
	}

	for _, v := range other {
		if _, isOk := allMaps[v]; isOk {
			allMaps[v] = false
		}
	}

	for _, v := range all {
		if allMaps[v] {
			sliceResult = append(sliceResult, v)
		}
	}

	return sliceResult
}

/*
 * 过滤uint64数组（从all中过滤所有other中的数据，返回未被过滤的数据集合）
 *  */
func FilterUint64Slice(all, other []uint64) []uint64 {
	allMaps := make(map[uint64]bool, 0)
	sliceResult := make([]uint64, 0)

	for _, v := range all {
		allMaps[v] = true
	}

	for _, v := range other {
		if _, isOk := allMaps[v]; isOk {
			allMaps[v] = false
		}
	}

	for _, v := range all {
		if allMaps[v] {
			sliceResult = append(sliceResult, v)
		}
	}

	return sliceResult
}

/*
 * 过滤字符数组（从all中过滤所有other中的数据，返回未被过滤的数据集合）
 *  */
func FilterStringSlice(all, other []string) []string {
	allMaps := make(map[string]bool, 0)
	sliceResult := make([]string, 0)

	for _, v := range all {
		allMaps[v] = true
	}

	for _, v := range other {
		if _, isOk := allMaps[v]; isOk {
			allMaps[v] = false
		}
	}

	for _, v := range all {
		if allMaps[v] {
			sliceResult = append(sliceResult, v)
		}
	}

	return sliceResult
}

/*
 * 字符集合的交集合
 *  */
func StringInter(one, two []string) []string {
	allMap := make(map[string]bool, 0)
	interSet := make([]string, 0)

	if len(one) == 0 && len(two) == 0 {
		return interSet
	}

	if len(one) == 0 {
		for _, v := range two {
			allMap[v] = true
		}
	} else if len(two) == 0 {
		for _, v := range one {
			allMap[v] = true
		}
	} else {
		for _, v := range one {
			allMap[v] = true
		}

		for _, v := range two {
			if _, isOk := allMap[v]; isOk {
				allMap[v] = false
			}
		}
	}

	for k, v := range allMap {
		if !v {
			interSet = append(interSet, k)
		}
	}

	return interSet
}

/*
 * 字符集合的并集合
 *  */
func StringUnion(one, two []string) []string {
	allMap := make(map[string]string, 0)
	union := make([]string, 0)

	for _, v := range one {
		allMap[v] = v
	}

	for _, v := range two {
		if _, isOk := allMap[v]; !isOk {
			allMap[v] = v
		}
	}

	for _, v := range allMap {
		union = append(union, v)
	}

	return union
}

/*
 * 字符集合的差集合
 *  */
func StringDiff(one, two []string) []string {
	//并集合
	union := StringUnion(one, two)

	//交集合
	inter := StringInter(one, two)

	//差集合
	diff := FilterStringSlice(union, inter)

	return diff
}

/*
 * int64交集合
 *  */
func Int64Inter(one, two []int64) []int64 {
	allMap := make(map[int64]bool, 0)
	interSet := make([]int64, 0)

	if len(one) == 0 && len(two) == 0 {
		return interSet
	}

	if len(one) == 0 {
		for _, v := range two {
			allMap[v] = true
		}
	} else if len(two) == 0 {
		for _, v := range one {
			allMap[v] = true
		}
	} else {
		for _, v := range one {
			allMap[v] = true
		}
		for _, v := range two {
			if _, isOk := allMap[v]; isOk {
				allMap[v] = false
			}
		}
	}

	for k, v := range allMap {
		if !v {
			interSet = append(interSet, k)
		}
	}

	return interSet
}

/*
 * int64并集合
 *  */
func Int64Union(one, two []int64) []int64 {
	allMap := make(map[int64]int64, 0)
	union := make([]int64, 0)

	for _, v := range one {
		allMap[v] = v
	}

	for _, v := range two {
		if _, isOk := allMap[v]; !isOk {
			allMap[v] = v
		}
	}

	for _, v := range allMap {
		union = append(union, v)
	}

	return union
}

/*
 * int64差集合
 *  */
func Int64Diff(one, two []int64) []int64 {
	//并集合
	union := Int64Union(one, two)

	//交集合
	inter := Int64Inter(one, two)

	//差集合
	diff := FilterInt64Slice(union, inter)

	return diff
}

/*
 * uint64交集合
 *  */
func Uint64Inter(one, two []uint64) []uint64 {
	allMap := make(map[uint64]bool, 0)
	interSet := make([]uint64, 0)

	if len(one) == 0 && len(two) == 0 {
		return interSet
	}

	if len(one) == 0 {
		for _, v := range two {
			allMap[v] = true
		}
	} else if len(two) == 0 {
		for _, v := range one {
			allMap[v] = true
		}
	} else {
		for _, v := range one {
			allMap[v] = true
		}
		for _, v := range two {
			if _, isOk := allMap[v]; isOk {
				allMap[v] = false
			}
		}
	}

	for k, v := range allMap {
		if !v {
			interSet = append(interSet, k)
		}
	}

	return interSet
}

/*
 * uint64并集合
 *  */
func Uint64Union(one, two []uint64) []uint64 {
	allMap := make(map[uint64]uint64, 0)
	union := make([]uint64, 0)

	for _, v := range one {
		allMap[v] = v
	}

	for _, v := range two {
		if _, isOk := allMap[v]; !isOk {
			allMap[v] = v
		}
	}

	for _, v := range allMap {
		union = append(union, v)
	}

	return union
}

/*
 * uint64差集合
 *  */
func Uint64Diff(one, two []uint64) []uint64 {
	//并集合
	union := Uint64Union(one, two)

	//交集合
	inter := Uint64Inter(one, two)

	//差集合
	diff := FilterUint64Slice(union, inter)

	return diff
}

/*
 * 保留浮点数指定长度的小数位数
 *  */
func ModFloat64(value float64, length int) float64 {
	format := fmt.Sprintf("0.%df", length)
	format = "%" + format
	valueString := fmt.Sprintf(format, value)
	floatValue, _ := strconv.ParseFloat(valueString, 64)
	return floatValue
}

/*
 * Base64字符编码
 *  */
func ToBase64(data string, args ...bool) string {
	resultString := ""
	isUrlEncoding := false

	if len(args) > 0 && args[0] {
		isUrlEncoding = true
	}

	if isUrlEncoding {
		resultString = base64.URLEncoding.EncodeToString([]byte(data))
	} else {
		resultString = base64.StdEncoding.EncodeToString([]byte(data))
	}

	return resultString
}

/*
 * Base64字符解码
 *  */
func FromBase64(data string, args ...bool) (string, error) {
	var bytesData []byte
	var err error
	isUrlDecoding := false

	if len(args) > 0 && args[0] {
		isUrlDecoding = true
	}

	if isUrlDecoding {
		bytesData, err = base64.URLEncoding.DecodeString(data)
	} else {
		bytesData, err = base64.StdEncoding.DecodeString(data)
	}

	if err != nil {
		return "", err
	}

	return string(bytesData), nil
}

/*
 * 对象转换成Json字符串
 *  */
func ToJson(object interface{}) (string, error) {
	v, err := json.Marshal(object)
	if err != nil {
		return "", err
	}

	return string(v), nil
}

/*
 * Json字符串转换成对象
 *  */
func FromJson(jsonString string, object interface{}) error {
	bytesData := []byte(jsonString)

	jsonDecoder := json.NewDecoder(bytes.NewBuffer(bytesData))
	jsonDecoder.UseNumber()

	return jsonDecoder.Decode(&object)
	//return json.Unmarshal(bytesData, object)
}

/*
 * 对象转换成Xml字符串
 *  */
func ToXml(object interface{}) (string, error) {
	v, err := xml.Marshal(object)
	if err != nil {
		return "", err
	}

	return string(v), nil
}

/*
 * Xml字符串转换成对象
 *  */
func FromXml(xmlString string, object interface{}) error {
	bytesData := []byte(xmlString)
	return xml.Unmarshal(bytesData, object)
}

/*
 * 字典参数升序排序，组成键＝值集合，然后把集合用&拼接成字符串
 *  */
func JoinMapToString(params map[string]string, filterKeys []string, isEscape bool) string {
	var keys []string = make([]string, 0)
	var values []string = make([]string, 0)

	filterKeyMaps := make(map[string]string, 0)
	if len(filterKeys) > 0 {
		for _, key := range filterKeys {
			filterKeyMaps[key] = key
		}
	}

	//请求参数排序（字母升序）
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	//拼接KeyValue字符串
	for _, key := range keys {
		//过滤空值
		if len(params[key]) > 0 {
			//过滤指定的key
			if _, isExists := filterKeyMaps[key]; isExists {
				continue
			}

			keyValue := params[key]
			if isEscape {
				keyValue = url.QueryEscape(keyValue)
			}

			//键＝值集合
			value := fmt.Sprintf("%s=%s", key, keyValue)
			values = append(values, value)
		}
	}

	//用&连接起来
	paramString := strings.Join(values, "&")

	return paramString
}

/*
 * 签名算法
 * params里的每个Value都需要进行url编码
 * fmt.Sprintf("%s=%s", key, url.QueryEscape(value))
 *  */
func MapDataSign(params map[string]string, secret string) (string, bool) {
	isExpired := false

	var timestamp int64
	var keys []string = make([]string, 0)
	var values []string = make([]string, 0)

	//请求参数排序（字母升序）
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	//拼接KeyValue字符串
	for _, key := range keys {
		if len(params[key]) > 0 {
			values = append(values, key)         //Key
			values = append(values, params[key]) //Value

			if key == "timestamp" {
				if _timestamp, err := strconv.ParseInt(params[key], 10, 64); err != nil {
					timestamp = _timestamp
				}
			}
		}
	}
	paramString := strings.Join(values, "")

	//是否已过期
	if timestamp > 0 {
		isExpired = time.Unix(timestamp, 0).Add(time.Minute * time.Duration(5)).Before(time.Now())
	}

	//Md5签名（在拼接的字符串头尾附加上api密匙，然后md5，md5串是大写）
	paramString = fmt.Sprintf("%s%s%s", secret, paramString, secret)
	sign := Md5(paramString)

	return strings.ToUpper(sign), isExpired
}
