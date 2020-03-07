package duedatecalc

import (
	"duedatecalc/types"
	"math"
)

const (
	startOfDay = 9
	endOfDay   = 17
)

type Issue struct {
	Submitted      types.DateTime
	TurnaroundTime types.Duration
}

func (i Issue) CalculateDueDate() types.DateTime {
	deadline := i.Submitted.Copy()
	deadline.Minutes += i.TurnaroundTime.Minutes
	if deadline.Minutes > 59 {
		deadline.Minutes -= 60
		deadline.Hours += 1
	}

	// Első napon elkölthető időmennyiség
	startFraction := calculateFirstDayHours(i.Submitted)
	// Első nap után visszamaradó időmennyiség
	var remainingDuration types.Duration
	if i.TurnaroundTime.Minutes-startFraction.Minutes < 0 {
		remainingDuration.Minutes = 60 - (startFraction.Minutes - i.TurnaroundTime.Minutes)
		remainingDuration.Hours = i.TurnaroundTime.Hours - startFraction.Hours - 1
	} else {
		remainingDuration.Minutes = i.TurnaroundTime.Minutes - startFraction.Minutes
		remainingDuration.Hours = i.TurnaroundTime.Hours - startFraction.Hours
	}

	// Visszamaradó időmennyiség napokra szétosztva
	remainingDays := int(math.Ceil(float64(remainingDuration.Hours) / 8.0))
	if remainingDuration.Hours%8 == 0 && remainingDuration.Minutes > 0 {
		remainingDays += 1
	}

	// Ha a munkaintervallumba beleesik egy hétvége, akkor növeljük a munkával töltendő napok számát
	if hasWeekend(i.Submitted.Date, remainingDays) {
		remainingDays += 2
	}

	// Ha a feladat hosszabb mint 8 óra
	if i.TurnaroundTime.Hours > 8 {
		deadline.DayOfMonth += remainingDays
		i.TurnaroundTime.Hours -= (i.TurnaroundTime.Hours / 8) * 8
	}

	deadline.Hours += i.TurnaroundTime.Hours
	// Ha a feladat kilóg az adott napból
	if deadline.Hours >= 17 {
		deadline.Hours = deadline.Hours - 17 + 8
		deadline.DayOfMonth += 1
	}

	// Hétvége. Szánsájn, bícs
	handleWeekend(&deadline)

	// TODO: évváltás lekezelése, ha esetleg év végén abbamaradna a munka

	return deadline
}

// Kiszámítja, hogy az első nap mikor beérkezik a taszk mennyi idő tölthető el még vele az
// aktuálus munkanap végéig
func calculateFirstDayHours(dt types.DateTime) (startFraction types.Time) {
	if dt.Minutes > 0 {
		startFraction.Minutes = 60 - dt.Minutes
	}

	if dt.Minutes > 0 {
		startFraction.Hours = endOfDay - dt.Hours - 1
	} else {
		startFraction.Hours = endOfDay - dt.Hours
	}

	if dt.Hours < startOfDay {
		startFraction.Hours = 8
	}

	return
}

// Beérkezési dátumtól számítva van-e az időintervallumban hétvége
func hasWeekend(start types.Date, interval int) bool {
	act := start.Copy()
	for i := 0; i < interval; i++ {
		act.DayOfMonth += i
		if act.Weekday().IsWeekend() {
			return true
		}
	}
	return false
}

// Amennyiben a teljesítési időszakba belecsúszik egy hétvége úgy meghosszabbítja a teljesítési időszakot
// TODO: annak kezelése, ha több hétvége is belenyúlik a munkába
func handleWeekend(d *types.DateTime) {
	if d.Weekday().IsWeekend() {
		d.DayOfMonth += 2
		if d.Date.DaysInMonth() < d.DayOfMonth {
			d.DayOfMonth = firstWeekday(d.Year, d.MonthOfYear+1)
			d.MonthOfYear += 1
		}
	}
}

// Kiszámítja, hogy az adott hónapban hányadika az első munkanap
func firstWeekday(year int, month types.Month) (dayOfMonth int) {
	date := types.Date{Year: year, MonthOfYear: month}
	for i := 1; i < date.DaysInMonth(); i++ {
		date.DayOfMonth = i
		if !date.Weekday().IsWeekend() {
			return i
		}
	}

	return 1
}
