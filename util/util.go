package util

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func Uuid(length int64) string {
	ele := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "v", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "Driver", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	ele, _ = Random(ele)
	uuid := ""
	var i int64
	for i = 0; i < length; i++ {
		uuid += ele[rand.Intn(59)]
	}
	return uuid
}

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

func WrapURL(u string) string {
	uarr := strings.Split(u, "?")
	if len(uarr) < 2 {
		return url.QueryEscape(strings.ReplaceAll(u, "/", "_"))
	}
	v, err := url.ParseQuery(uarr[1])
	if err != nil {
		return url.QueryEscape(strings.ReplaceAll(u, "/", "_"))
	}
	return url.QueryEscape(strings.ReplaceAll(uarr[0], "/", "_")) + "?" +
		strings.ReplaceAll(v.Encode(), "%7B%7B.Id%7D%7D", "{{.Id}}")
}

func ReplaceAll(s string, oldnew ...string) string {
	repl := strings.NewReplacer(oldnew...)
	return repl.Replace(s)
}

func PackageName(v interface{}) string {
	if v == nil {
		return ""
	}

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		return val.Elem().Type().PkgPath()
	}
	return val.Type().PkgPath()
}

func SetDefault(value, condition, def string) string {
	if value == condition {
		return def
	}
	return value
}

func AorB(condition bool, a, b string) string {
	if condition {
		return a
	}
	return b
}

func CopyMap(m map[string]string) map[string]string {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		panic(err)
	}
	var cm map[string]string
	err = dec.Decode(&cm)
	if err != nil {
		panic(err)
	}
	return cm
}

// TimeSincePro calculates the time interval and generate full user-friendly string.
func TimeSincePro(then time.Time, m map[string]string) string {
	now := time.Now()
	diff := now.Unix() - then.Unix()

	if then.After(now) {
		return "future"
	}

	var timeStr, diffStr string
	for {
		if diff == 0 {
			break
		}

		diff, diffStr = computeTimeDiff(diff, m)
		timeStr += ", " + diffStr
	}
	return strings.TrimPrefix(timeStr, ", ")
}

// Seconds-based time units
const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month
)

func computeTimeDiff(diff int64, m map[string]string) (int64, string) {
	diffStr := ""
	switch {
	case diff <= 0:
		diff = 0
		diffStr = "now"
	case diff < 2:
		diff = 0
		diffStr = "1 " + m["second"]
	case diff < 1*Minute:
		diffStr = fmt.Sprintf("%d "+m["seconds"], diff)
		diff = 0

	case diff < 2*Minute:
		diff -= 1 * Minute
		diffStr = "1 " + m["minute"]
	case diff < 1*Hour:
		diffStr = fmt.Sprintf("%d "+m["minutes"], diff/Minute)
		diff -= diff / Minute * Minute

	case diff < 2*Hour:
		diff -= 1 * Hour
		diffStr = "1 " + m["hour"]
	case diff < 1*Day:
		diffStr = fmt.Sprintf("%d "+m["hours"], diff/Hour)
		diff -= diff / Hour * Hour

	case diff < 2*Day:
		diff -= 1 * Day
		diffStr = "1 " + m["day"]
	case diff < 1*Week:
		diffStr = fmt.Sprintf("%d "+m["days"], diff/Day)
		diff -= diff / Day * Day

	case diff < 2*Week:
		diff -= 1 * Week
		diffStr = "1 " + m["week"]
	case diff < 1*Month:
		diffStr = fmt.Sprintf("%d "+m["weeks"], diff/Week)
		diff -= diff / Week * Week

	case diff < 2*Month:
		diff -= 1 * Month
		diffStr = "1 " + m["month"]
	case diff < 1*Year:
		diffStr = fmt.Sprintf("%d "+m["months"], diff/Month)
		diff -= diff / Month * Month

	case diff < 2*Year:
		diff -= 1 * Year
		diffStr = "1 " + m["year"]
	default:
		diffStr = fmt.Sprintf("%d "+m["years"], diff/Year)
		diff = 0
	}
	return diff, diffStr
}
