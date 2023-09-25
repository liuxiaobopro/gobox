package time

import (
	"time"
)

const (
	FormatDateTime = "2006-01-02 15:04:05"
	FormatDate     = "2006-01-02"
	FormatTime     = "15:04:05"

	TimeZoneSH = "Asia/Shanghai"
)

// IntToString 将int转换为时间字符串
// 例如: 1560000000 -> 2019-06-09 00:00:00
func IntToString(t int) string {
	return time.Unix(int64(t), 0).Format(FormatDateTime)
}

// IntToStringDate 将int转换为日期字符串
// 例如: 1560000000 -> 2019-06-09
func IntToStringDate(t int) string {
	return time.Unix(int64(t), 0).Format(FormatDate)
}

// IntToStringTime 将int转换为时间字符串
// 例如: 1560000000 -> 00:00:00
func IntToStringTime(t int) string {
	return time.Unix(int64(t), 0).Format(FormatTime)
}

// StringToInt 将时间字符串转换为int
// 例如: 2019-06-09 00:00:00 -> 1560000000
func StringToInt(t string) int {
	loc, err := time.LoadLocation("Local") // 获取时区(中国上海; TimeZoneSH)
	if err != nil {
		panic(err)
	}
	tt, err := time.ParseInLocation(FormatDateTime, t, loc)
	if err != nil {
		panic(err)
	}
	return int(tt.Unix())
}

// NowTimeStr 获取当前时间字符串
// 例如: 2019-06-09 00:00:00
func NowTimeStr() string {
	return time.Now().Format(FormatDateTime)
}

// NowTimeStamp 获取当前时间戳
// 例如: 1560000000
func NowTimeStamp() int {
	return int(time.Now().Unix())
}

// NowMilliTimeStamp 获取当前时间戳(毫秒)
// 例如: 1560000000000
func NowMilliTimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// CheckSETime 检查开始时间和结束时间
func CheckSETime(startTime, endTime string) bool {
	if startTime == "" || endTime == "" {
		return false
	}
	if StringToInt(startTime) > StringToInt(endTime) {
		return false
	}
	return true
}
