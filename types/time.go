package types

import "fmt"

type Time struct {
	Hours   int
	Minutes int
	Seconds int
}

func (t Time) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hours, t.Minutes, t.Seconds)
}
