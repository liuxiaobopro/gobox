package time

import (
	"fmt"
	"time"
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
