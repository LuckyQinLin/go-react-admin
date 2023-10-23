package utils

import (
	"fmt"
	"time"
)

type DateTime time.Time
type DateTime1 time.Time

func (d DateTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(d).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (d DateTime1) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(d).Format("15:04:05"))
	return []byte(stamp), nil
}
