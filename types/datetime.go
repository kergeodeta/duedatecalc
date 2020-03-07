package types

import (
	"fmt"
)

type DateTime struct {
	Date
	Time
}

func CreateDateTime(year, month, day, hour, minute, second int) DateTime {
	return DateTime{
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

func (dt DateTime) Copy() DateTime {
	return DateTime{
		Date: Date{
			Year:        dt.Year,
			MonthOfYear: dt.MonthOfYear,
			DayOfMonth:  dt.DayOfMonth,
		},
		Time: Time{
			Hours:   dt.Hours,
			Minutes: dt.Minutes,
			Seconds: dt.Seconds,
		},
	}
}

func (dt DateTime) String() string {
	return fmt.Sprintf("%s %s", dt.Date, dt.Time)
}
