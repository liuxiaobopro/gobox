package time

import (
	"fmt"
	"time"

	definex "github.com/liuxiaobopro/golib/define"
)

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format(definex.FormatTime))
	return []byte(stamp), nil
}

// IntToString 将int转换为时间字符串
func IntToString(t int) string {
	return time.Unix(int64(t), 0).Format(definex.FormatTime)
}

// StringToInt 将时间字符串转换为int
func StringToInt(t string) int {
	loc, err := time.LoadLocation("Local") // 获取时区(中国上海; typesx.TimeZoneSH)
	if err != nil {
		panic(err)
	}
	tt, err := time.ParseInLocation(definex.FormatTime, t, loc)
	if err != nil {
		panic(err)
	}
	return int(tt.Unix())
}
