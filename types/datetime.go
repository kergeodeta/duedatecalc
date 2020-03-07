package types

import (
	"fmt"
)

type DateTime struct {
	Date
	Time
}

func CreateDateTime(year, month, day, hour, minute, second int) *DateTime {
	return &DateTime{
		Date: Date{
			Year:        year,
			MonthOfYear: Month(month),
			DayOfMonth:  day,
		},
		Time: Time{
			Hours:   hour,
			Minutes: minute,
			Seconds: second,
		},
	}
}

func (dt DateTime) String() string {
	return fmt.Sprintf("%s %s", dt.Date, dt.Time)
}
