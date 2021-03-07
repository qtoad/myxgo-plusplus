package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Format unix time int64 to string
func Date(ti int64, format string) string {
	t := time.Unix(int64(ti), 0)
	return DateT(t, format)
}

// Format unix time string to string
func DateS(ts string, format string) string {
	i, _ := strconv.ParseInt(ts, 10, 64)
	return Date(i, format)
}

// Format time.Time struct to string
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5
func DateT(t time.Time, format string) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}

// DateFormat pattern rules.
var datePatterns = []string{
	// year
	"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
	"y", "06", //A two digit representation of a year   Examples: 99 or 03

	// month
	"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
	"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
	"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
	"F", "January", // A full textual representation of a month, such as January or March   January through December

	// day
	"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
	"j", "2", // Day of the month without leading zeros 1 to 31

	// week
	"D", "Mon", // A textual representation of a day, three letters Mon through Sun
	"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

	// time
	"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
	"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
	"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
	"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

	"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
	"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

	"i", "04", // Minutes with leading zeros    00 to 59
	"s", "05", // Seconds, with leading zeros   00 through 59

	// time zone
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",

	// RFC 2822
	"r", time.RFC1123Z,
}

// Parse Date use PHP time format.
func DateParse(dateString, format string) (time.Time, error) {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	return time.ParseInLocation(format, dateString, time.Local)
}

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

/*
 * 判断firstDatetime是否在secondDatetime的后面，即firstDatetime比secondDatetime日期大
 *  */
func IsDateGreaterThan(firstDatetime, secondDatetime time.Time) bool {
	return firstDatetime.After(secondDatetime)
}

/*
 * 判断firstDatetime是否在secondDatetime的前面，即firstDatetime比secondDatetime日期小
 *  */
func IsDateLessThan(firstDatetime, secondDatetime time.Time) bool {
	return firstDatetime.Before(secondDatetime)
}

/*
 * 获取当前Unix时间戳
 * 当前日期距离197011000的秒数
 *  */
func UnixTimestamp() int64 {
	return GetNow().Unix()
}

/*
 * 获取当前Unix纳秒时间戳
 * 当前日期距离197011000的纳秒数
 *  */
func UnixNanoTimestamp() int64 {
	return GetNow().UnixNano()
}

/*
 * Unix时间戳日期（1970-01-01 00:00:00）
 *  */
func UnixTimestampDate() time.Time {
	dtTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	return dtTime
}

/*
 * Golang零时间日期（0001-01-01 00:00:00）
 * -62135596800
 *  */
func UnixDate() time.Time {
	dtTime, _ := StringToTime(time.UnixDate, true)
	return dtTime
}

/*
 * 获取指定日期的Unix秒时间戳
 *  */
func DateToUnixTimestamp(date time.Time) int64 {
	var unixValue int64

	if !date.IsZero() {
		unixValue = date.Unix()
	}

	return unixValue
}

/*
 * 获取指定日期的Unix纳秒时间戳
 *  */
func DateToUnixNanoTimestamp(date time.Time) int64 {
	var unixValue int64

	if !date.IsZero() {
		unixValue = date.UnixNano()
	}

	return unixValue
}

/*
 * 根据Unix时间戳返回日期
 *  */
func UnixTimestampToDate(unixTimestamp int64) time.Time {
	return time.Unix(unixTimestamp, 0).UTC()
}

/*
 * 根据Unix纳秒时间戳返回日期
 *  */
func UnixNanoTimestampToDate(unixNanoTimestamp int64) time.Time {
	return time.Unix(0, unixNanoTimestamp).UTC()
}

/*
 * 获取当前Local日期时间
 *  */
func GetNow() time.Time {
	return time.Now()
}

/*
 * 获取当前Utc日期时间
 *  */
func GetUtcNow() time.Time {
	return time.Now().UTC()
}

/*
 * 获取当前年月日的整型数字值
 *  */
func GetDateYearMonthDay(args ...time.Time) int {
	var yearMonthDay int

	format := "20060102"
	date := time.Now()

	if len(args) > 0 {
		date = args[0]
	}

	yearMonthDayString := TimeToString(date, format)
	if _yearMonthDay, err := strconv.Atoi(yearMonthDayString); err == nil {
		yearMonthDay = _yearMonthDay
	}

	return yearMonthDay
}

/*
 * 获取当前年份
 *  */
func GetCurrentYear() int32 {
	year, _, _ := GetNow().Date()
	return int32(year)
}

/*
 * 获取当前月份
 *  */
func GetCurrentMonth() int32 {
	_, month, _ := GetNow().Date()
	return int32(month)
}

/*
 * 获取当前日
 *  */
func GetCurrentDay() int32 {
	_, _, day := GetNow().Date()
	return int32(day)
}

/*
 * 获取当前小时
 *  */
func GetCurrentHour() int32 {
	hour, _, _ := GetNow().Clock()
	return int32(hour)
}

/*
 * 获取当前分钟
 *  */
func GetCurrentMinute() int32 {
	_, minute, _ := GetNow().Clock()
	return int32(minute)
}

/*
 * 获取当前秒数
 *  */
func GetCurrentSecond() int32 {
	_, _, second := GetNow().Clock()
	return int32(second)
}

/*
 * 获取日期时间的年份
 *  */
func GetDateYear(datetime time.Time) int32 {
	year, _, _ := datetime.Date()
	return int32(year)
}

/*
 * 获取日期时间的月份
 *  */
func GetDateMonth(datetime time.Time) int32 {
	_, month, _ := datetime.Date()
	return int32(month)
}

/*
 * 获取日期时间的日部分
 *  */
func GetDateDay(datetime time.Time) int32 {
	_, _, day := datetime.Date()
	return int32(day)
}

/*
 * 获取日期时间的小时部分
 *  */
func GetDateHour(datetime time.Time) int32 {
	hour, _, _ := datetime.Clock()
	return int32(hour)
}

/*
 * 获取日期时间的分钟部分
 *  */
func GetDateMinute(datetime time.Time) int32 {
	_, minute, _ := datetime.Clock()
	return int32(minute)
}

/*
 * 获取日期时间的秒部分
 *  */
func GetDateSecond(datetime time.Time) int32 {
	_, _, second := datetime.Clock()
	return int32(second)
}

/*
 * 返回日期的最小日期时间（2016-01-02 00:00:00）
 *  */
func GetMinDate(dtTime time.Time) time.Time {
	year, month, day := dtTime.UTC().Date()
	return time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)
}

/*
 * 返回日期的最大日期时间（2016-01-02 23:59:59 999999999）
 *  */
func GetMaxDate(dtTime time.Time) time.Time {
	year, month, day := dtTime.UTC().Date()
	return time.Date(int(year), time.Month(month), int(day), 23, 59, 59, 999999999, time.UTC)
}

/*
 * 获取日期的最小日期时间戳，单位秒
 *  */
func GetMinDateTimestamp(dtTime time.Time) int64 {
	minTime := GetMinDate(dtTime)
	return DateToUnixTimestamp(minTime)
}

/*
 * 获取日期的最大日期时间戳，单位秒
 *  */
func GetMaxDateTimestamp(dtTime time.Time) int64 {
	maxTime := GetMaxDate(dtTime)
	return DateToUnixTimestamp(maxTime)
}

/*
 * 获取日期的最小日期时间戳，单位纳秒
 *  */
func GetMinDateNanoTimestamp(dtTime time.Time) int64 {
	minTime := GetMinDate(dtTime)
	return DateToUnixNanoTimestamp(minTime)
}

/*
 * 获取日期的最大日期时间戳，单位纳秒
 *  */
func GetMaxDateNanoTimestamp(dtTime time.Time) int64 {
	maxTime := GetMaxDate(dtTime)
	return DateToUnixNanoTimestamp(maxTime)
}

/*
 * 获取当月里最小日期时间（2016-01-01 0:0:0 0）
 *  */
func GetCurrentMonthMinDate(args ...time.Time) time.Time {
	date := time.Now()
	if len(args) > 0 {
		date = args[0]
	}

	year, month, _ := date.Date()

	return time.Date(int(year), time.Month(month), int(1), 0, 0, 0, 0, time.Local)
}

/*
 * 获取当月里最大日期时间（2016-01-02 23:59:59 999999999）
 *  */
func GetCurrentMonthMaxDate(args ...time.Time) time.Time {
	date := time.Now()
	if len(args) > 0 {
		date = args[0]
	}

	daysForMonth := GetCurrentDayCount(args...)

	year, month, _ := date.Date()
	return time.Date(int(year), time.Month(month), int(daysForMonth), 23, 59, 59, 999999999, time.Local)
}

/*
 * 获取当月里最小日期时间戳(当月第一天最小时间)，单位秒
 *  */
func GetCurrentMonthMinTimestamp(args ...time.Time) int64 {
	minTimeForMonthFirstDay := GetCurrentMonthMinDate(args...)
	return DateToUnixTimestamp(minTimeForMonthFirstDay)
}

/*
 * 获取当月里最大日期时间戳(当月最后一天最大时间)，单位秒
 *  */
func GetCurrentMonthMaxTimestamp(args ...time.Time) int64 {
	maxTimeForMonthLastDay := GetCurrentMonthMaxDate(args...)
	return DateToUnixTimestamp(maxTimeForMonthLastDay)
}

/*
 * 获取当月里最小日期时间戳(当月第一天最小时间)，单位纳秒
 *  */
func GetCurrentMonthMinNanoTimestamp(args ...time.Time) int64 {
	minTimeForMonthFirstDay := GetCurrentMonthMinDate(args...)
	return DateToUnixNanoTimestamp(minTimeForMonthFirstDay)
}

/*
 * 获取当月里最大日期时间戳(当月最后一天最大时间)，单位纳秒
 *  */
func GetCurrentMonthMaxNanoTimestamp(args ...time.Time) int64 {
	maxTimeForMonthLastDay := GetCurrentMonthMaxDate(args...)
	return DateToUnixNanoTimestamp(maxTimeForMonthLastDay)
}

/*
 * 获取日期时间的日期和星期字符串
 *  */
func GetDatetimeWeekString(datetime time.Time) string {
	_, month, day := datetime.Date()
	hour, minute, _ := datetime.Clock()
	weekday := GetWeek(datetime)

	weekdays := make(map[int]string, 0)
	weekdays[1] = "星期一"
	weekdays[2] = "星期二"
	weekdays[3] = "星期三"
	weekdays[4] = "星期四"
	weekdays[5] = "星期五"
	weekdays[6] = "星期六"
	weekdays[7] = "星期日"
	weekdayString := weekdays[weekday]

	dateString := fmt.Sprintf("%d月%d日%s%d:%d", int(month), day, weekdayString, hour, minute)

	return dateString
}

/*
 * 月份数值集合转换成季节名集合
 *  */
func MonthsToSeasons(months []int32) []string {
	seasonDatas := make(map[string]bool, 0)
	seasonNames := make([]string, 0)

	seasons := map[string][]int32{
		"spring": []int32{1, 2, 3},
		"summer": []int32{4, 5, 6},
		"autumn": []int32{7, 8, 9},
		"winter": []int32{10, 11, 12},
	}

	for _, month := range months {
		seasonName := ""

		for k, v1 := range seasons {
			isFound := false
			for _, v2 := range v1 {
				if v2 == month {
					seasonName = k
					isFound = true
					break
				}
			}

			if isFound {
				break
			}
		}

		if _, isOk := seasonDatas[seasonName]; !isOk {
			seasonDatas[seasonName] = true
		}
	}

	for seasonName, _ := range seasonDatas {
		seasonNames = append(seasonNames, seasonName)
	}

	return seasonNames
}

/*
 * 日期时间转换成友好的显示字符串
 * isUtc:bool | format:string
 *  */
func TimeToFriendString(datetime time.Time, args ...interface{}) string {
	format := "2006-01-02 15:04:05"
	isUtc := false

	if len(args) > 0 {
		if v, ok := args[0].(bool); ok {
			isUtc = v
		}
	}

	if len(args) > 1 {
		if v, ok := args[1].(string); ok {
			format = v
		}
	}

	result := TimeToString(datetime.Local(), format)
	if isUtc {
		result = TimeToString(datetime.UTC(), format)
	}

	currentDate := time.Now().UTC()
	duration := currentDate.Sub(datetime.UTC())
	hourCount := int(duration.Hours())
	minuteCount := int(duration.Minutes())
	dayCount := int(hourCount / 24)
	monthCount := int(dayCount / 30)

	if monthCount > 24 {
		if monthCount < 24 {
			if monthCount == 1 {
				result = fmt.Sprintf("%s", "上个月")
			} else {
				if monthCount == 6 {
					result = fmt.Sprintf("%s", "半年前")
				} else {
					result = fmt.Sprintf("%d个月前", monthCount)
				}
			}
		} else {
			yearCount := int(monthCount / 12)
			if yearCount < 5 {
				result = fmt.Sprintf("%d年前", yearCount)
			}
		}
	} else if dayCount > 0 {
		if dayCount > 14 {
			result = "半个月前"
		} else if dayCount > 6 {
			result = "一周前"
		} else {
			if dayCount == 1 {
				result = fmt.Sprintf("%s", "昨天")
			} else {
				result = fmt.Sprintf("%d天前", dayCount)
			}
		}
	} else if hourCount > 0 {
		result = fmt.Sprintf("%d小时前", hourCount)
	} else {
		if minuteCount > 0 {
			result = fmt.Sprintf("%d分钟前", minuteCount)
		} else {
			result = "刚刚"
		}
	}

	return result
}

/*
 * 日期字符串切片转成日期
 *  */
func StringSliceToDate(dateStringSlice []string) (time.Time, error) {
	dateString := strings.Join(dateStringSlice, "-")

	dtTime, err := StringToTime(dateString)
	if err != nil {
		return UnixDate(), err
	}

	return GetMinDate(dtTime), nil
}

/*
 * 日期转成日期字符串切片
 *  */
func DateToStringSlice(date time.Time) []string {
	if date.IsZero() {
		date = time.Now()
	}

	dateStringSlice := make([]string, 0)
	year, month, day := date.Date()

	dateStringSlice = append(dateStringSlice, fmt.Sprintf("%d", year))
	dateStringSlice = append(dateStringSlice, fmt.Sprintf("%d", int(month)))
	dateStringSlice = append(dateStringSlice, fmt.Sprintf("%d", day))

	return dateStringSlice
}

/*
 * int切片转成日期
 *  */
func IntSliceToDate(intSlice []int) (time.Time, error) {
	dateStringSlice := make([]string, 3)
	for _, v := range intSlice {
		dateStringSlice = append(dateStringSlice, fmt.Sprintf("%d", v))
	}

	dtTime, err := StringToTime(strings.Join(dateStringSlice, "-"))
	if err != nil {
		return UnixDate(), err
	}

	return GetMinDate(dtTime), nil
}

/*
 * 日期转成int切片
 *  */
func DateToIntSlice(date time.Time) []int {
	intSlice := make([]int, 3)

	dateStringSlice := DateToStringSlice(date)
	intSlice[0], _ = strconv.Atoi(dateStringSlice[0])
	intSlice[1], _ = strconv.Atoi(dateStringSlice[1])
	intSlice[2], _ = strconv.Atoi(dateStringSlice[2])

	return intSlice
}

/*
 * firstDatetime加上时间间隔duration，返回日期时间
 *  */
func DatetimeAdd(firstDatetime time.Time, duration time.Duration) time.Time {
	return firstDatetime.Add(duration)
}

/*
 * firstDatetime加上指定的天数，返回日期时间
 *  */
func DatetimeAddDay(firstDatetime time.Time, dayValue int) time.Time {
	return DatetimeAddHour(firstDatetime, 24*dayValue)
}

/*
 * firstDatetime加上指定的小时数，返回日期时间
 *  */
func DatetimeAddHour(firstDatetime time.Time, hourValue int) time.Time {
	return DatetimeAdd(firstDatetime, time.Duration(hourValue)*time.Hour)
}

/*
 * firstDatetime加上指定的分钟数，返回日期时间
 *  */
func DatetimeAddMinute(firstDatetime time.Time, minuteValue int) time.Time {
	return DatetimeAdd(firstDatetime, time.Duration(minuteValue)*time.Minute)
}

/*
 * firstDatetime加上指定的秒数，返回日期时间
 *  */
func DatetimeAddSecond(firstDatetime time.Time, secondValue int) time.Time {
	return DatetimeAdd(firstDatetime, time.Duration(secondValue)*time.Second)
}

/*
 * firstDatetime减去secondDatetime，返回时间间隔
 *  */
func DatetimeSub(firstDatetime, secondDatetime time.Time) time.Duration {
	return firstDatetime.Sub(secondDatetime)
}

/*
 * 在当前的日期时间增加指定的分钟数，返回日期时间
 *  */
func AddMinutesForCurrent(minutes int) time.Time {
	return time.Now().Add(time.Duration(minutes) * time.Minute)
}

/*
 * 在当前的日期时间增加指定的小时数，返回日期时间
 *  */
func AddHoursForCurrent(hours int) time.Time {
	return time.Now().Add(time.Duration(hours) * time.Hour)
}

/*
 * 在当前的日期时间增加指定的天数，返回日期时间
 *  */
func AddDaysForCurrent(days int) time.Time {
	return AddHoursForCurrent(days * 24)
}

/*
 * 时间字符串加指定的分钟数，返回时间字符串
 *  */
func TimeStringAddMinutes(timeString string, minutes int) string {
	format := "15:04:05"

	var timeValue time.Time
	if time, err := time.ParseInLocation(format, timeString, time.UTC); err == nil {
		timeValue = time
	}

	timeValue = timeValue.Add(time.Duration(minutes) * time.Minute)
	return timeValue.Format(format)
}

/*
 * 日期时间的日期部分和时间字符串连接，返回日期时间
 *  */
func GetDatetimeForDateAndTimeString(date time.Time, timeString string) time.Time {
	format := "15:04:05"

	var timeValue time.Time
	if time, err := time.ParseInLocation(format, timeString, time.UTC); err == nil {
		timeValue = time
	}

	year, month, day := date.UTC().Date()
	hour, minute, second := timeValue.Clock()

	return time.Date(year, month, day, hour, minute, second, 0, time.UTC)
}

/*
 * 获取当前日期月份对应的天数
 *  */
func GetCurrentDayCount(args ...time.Time) int {
	date := time.Now()
	if len(args) > 0 {
		date = args[0]
	}

	return GetDayCount(date)
}

/*
 * 获取指定日期月份对应的天数
 *  */
func GetDayCount(datetime time.Time) int {
	year, month, _ := datetime.Date()
	dayCount := 31
	if month == 4 || month == 6 || month == 9 || month == 11 {
		dayCount = 30
	} else if month == 2 {
		if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
			dayCount = 29
		} else {
			dayCount = 28
		}
	}

	return dayCount
}

/*
 * 获取当前日期是周几（1:周一｜2:周二｜...|7:周日）
 *  */
func GetCurrentWeek() int {
	return GetWeek(time.Now())
}

/*
 * 获取指定的日期是周几（1:周一｜2:周二｜...|7:周日）
 *  */
func GetWeek(datetime time.Time) int {
	nowDate := datetime
	days := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
		0: 7,
	}
	weekday := nowDate.Weekday() //0：周日 | 1：周一 | .. ｜6：周六
	weekdayValue := days[int(weekday)]

	return weekdayValue
}

/*
 * 获取当前周对应的月份里的日期范围（minDay in month, maxDay in month）
 *  */
func GetCurrentWeekDayRange() (int, int) {
	nowDate := time.Now()
	_, _, day := nowDate.Date()
	weekdayValue := GetCurrentWeek()
	minDay := day - weekdayValue + 1
	maxDay := day + 7 - weekdayValue

	return minDay, maxDay
}

/*
 * 获取日期范围内的所属周几的日期集合
 * week：从1开始，1表示周一，依次类推
 *  */
func GetDateRangeForWeekInDateRange(startDate, endDate time.Time, week int) []time.Time {
	dateList := make([]time.Time, 0)
	date := startDate

	for date.Before(endDate) || date.Equal(endDate) {
		weekValue := GetWeek(date)
		if weekValue == week {
			dateList = append(dateList, date)
		}

		date = date.AddDate(0, 0, 1)
	}

	return dateList
}

/*
 * 获取一段时间范围内指定间隔的时间段集合
 *  */
func GetTimeIntervalStringSlice(startDate, endDate time.Time, minutes int64) []string {
	timeStringList := make([]string, 0)

	date := startDate
	for date.Before(endDate) || date.Equal(endDate) {
		timeString := TimeToString(date, "15:04")
		timeStringList = append(timeStringList, timeString)

		date = date.Add(time.Duration(minutes) * time.Minute)
	}

	return timeStringList
}

/*
 * 分钟数转时间字符串（HH:mm:ss）
 *  */
func MinutesToTimeString(minutes int64) string {
	hoursPart := minutes / 60
	minutesPart := minutes % 60

	timeString := fmt.Sprintf("%02d:%02d:00", hoursPart, minutesPart)

	return timeString
}

/*
 * 时间转字符串
 *  */
func CurrentTimeToString(args ...interface{}) string {
	return TimeToString(time.Now(), args...)
}

/*
 * 时间转字符串
 *  */
func TimeToString(datetime time.Time, args ...interface{}) string {
	format := "2006-01-02 15:04:05"
	if len(args) == 1 {
		if v, ok := args[0].(string); ok {
			format = v
		}
	}

	timeStrng := datetime.Format(format)
	return timeStrng
}

/*
 * 字符串转时间
 * isUtc参数决定传入的日期字符串是否是utc日期
 * isUtc:bool | format:string
 * 返回UTC日期
 *  */
func StringToTime(datetimeString string, args ...interface{}) (time.Time, error) {
	var retErr error
	var result time.Time

	format := "2006-01-02 15:04:05"
	isUtc := false

	if len(args) > 0 {
		if v, ok := args[0].(bool); ok {
			isUtc = v
		}
	}

	if len(args) > 1 {
		if v, ok := args[1].(string); ok {
			format = v
		}
	}

	if isUtc {
		result, retErr = time.ParseInLocation(format, datetimeString, time.UTC)
	} else {
		result, retErr = time.ParseInLocation(format, datetimeString, time.Local)
	}

	if retErr != nil {
		return result, retErr
	}

	return result.UTC(), retErr
}
