package time

import (
	"fmt"
	"time"
)

var (
	TimeZero = Time(time.Time{})
	TimeNow  = Time(time.Now())
)

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format(FormatDateTime))
	return []byte(stamp), nil
}

type Time time.Time // 和JsonTime一样，只是为了方便使用

func (t Time) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format(FormatDateTime))
	return []byte(stamp), nil
}

func (t Time) IsZero() bool {
	return time.Time(t).IsZero()
}

func (t Time) Format(layout string) string {
	return time.Time(t).Format(layout)
}

// Raw 返回原始time.Time
func (t Time) Raw() time.Time {
	return time.Time(t)
}

func Now() Time {
	return Time(time.Now())
}
