package duedatecalc

import (
	"duedatecalc/types"
	"math"
)

type Issue struct {
	Submitted      types.DateTime
	TurnaroundTime types.Duration
}

func (i Issue) CalculateDueDate() types.DateTime {
	deadline := types.DateTime{
		Date: types.Date{
			Year:        i.Submitted.Year,
			MonthOfYear: i.Submitted.MonthOfYear,
			DayOfMonth:  i.Submitted.DayOfMonth,
		},
		Time: types.Time{
			Hours:   i.Submitted.Hours,
			Minutes: i.Submitted.Minutes,
			Seconds: i.Submitted.Seconds,
		},
	}

	deadline.Minutes += i.TurnaroundTime.Minutes
	if deadline.Minutes > 59 {
		deadline.Minutes -= 60
		deadline.Hours += 1
	}

	firstWeekday := func(year int, month types.Month) (dayOfMonth int) {
		date := types.Date{Year: year, MonthOfYear: month}
		for i := 1; i < date.DaysInMonth(); i++ {
			date.DayOfMonth = i
			if !date.Weekday().IsWeekend() {
				return i
			}
		}

		return 1
	}

	if i.TurnaroundTime.Hours > 8 {
		deadline.DayOfMonth += int(math.Ceil(float64(i.TurnaroundTime.Hours) / 8.0))
		i.TurnaroundTime.Hours -= (i.TurnaroundTime.Hours / 8) * 8

		if deadline.Date.Weekday().IsWeekend() {
			deadline.DayOfMonth += 2
			if deadline.Date.DaysInMonth() < deadline.DayOfMonth {
				deadline.DayOfMonth = firstWeekday(deadline.Year, deadline.MonthOfYear+1)
				deadline.MonthOfYear += 1
			}
		}
	}

	deadline.Hours += i.TurnaroundTime.Hours
	if deadline.Hours >= 17 {
		deadline.Hours = deadline.Hours - 17 + 8
		deadline.DayOfMonth += 1

		if deadline.Date.Weekday().IsWeekend() {
			deadline.DayOfMonth += 2
			if deadline.Date.DaysInMonth() < deadline.DayOfMonth {
				deadline.DayOfMonth = firstWeekday(deadline.Year, deadline.MonthOfYear+1)
				deadline.MonthOfYear += 1
			}
		}
	}

	return deadline
}
