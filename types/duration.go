package types

import "fmt"

type Duration struct {
	Hours   int
	Minutes int
}

func CreateDuration(h, m int) *Duration {
	return &Duration{
		Hours:   h,
		Minutes: m,
	}
}

func (d Duration) String() string {
	return fmt.Sprintf("%02d:%02d", d.Hours, d.Minutes)
}
