package types

type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (d Weekday) String() string {
	names := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	if d < Monday || d > Sunday {
		return "Unknown"
	}

	return names[d-1]
}

// Megadja, hogy a hét adott napja hétvégére esik-e
func (d Weekday) IsWeekend() bool {
	switch d {
	case Saturday, Sunday:
		return true
	default:
		return false
	}
}
