package types

import (
	"fmt"
	"math"
)

const yearOfUnix int = 1970

type Date struct {
	Year        int
	MonthOfYear Month
	DayOfMonth  int
}

func (d Date) String() string {
	return fmt.Sprintf("%d-%02d-%02d", d.Year, d.MonthOfYear, d.DayOfMonth)
}

func (d Date) IsLeapYear() bool {
	if d.Year%400 == 0 {
		return true
	} else if d.Year%100 == 0 {
		return false
	} else if d.Year%4 == 0 {
		return true
	} else {
		return false
	}
}

func (d Date) DaysInMonth() int {
	switch d.MonthOfYear {
	case January, March, May, July, August, October, December:
		return 31
	case February:
		if d.IsLeapYear() {
			return 29
		} else {
			return 28
		}
	default:
		return 30
	}
}

func (d Date) Weekday() Weekday {
	var elapsedDays int
	// Eltelt évek napjainak összegzése
	for i := yearOfUnix; i < d.Year; i++ {
		if (Date{Year: i}.IsLeapYear()) {
			elapsedDays += 366
		} else {
			elapsedDays += 365
		}
	}

	// Aktuális évben eltelt hónapok napjainak összegzése
	for i := January; i < d.MonthOfYear; i++ {
		elapsedDays += Date{Year: d.Year, MonthOfYear: i}.DaysInMonth()
	}

	// Aktuális hónapban eltelt napok hozzáaádása.
	elapsedDays += d.DayOfMonth - 1

	return Weekday(int(math.Mod(float64(elapsedDays+4), 7)))
}
