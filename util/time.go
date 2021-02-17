package util

import "time"

// WeekDay 获取当前日的周名称
func WeekDay(t time.Time) string {
	w := ""

	switch t.Weekday() {
	case time.Sunday:
		w = "周日"
	case time.Monday:
		w = "周一"
	case time.Tuesday:
		w = "周二"
	case time.Wednesday:
		w = "周三"
	case time.Thursday:
		w = "周四"
	case time.Friday:
		w = "周五"
	case time.Saturday:
		w = "周六"
	}

	return w
}
