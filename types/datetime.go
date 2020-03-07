package types

import "fmt"

type DateTime struct {
	Date
	Time
}

func (dt DateTime) String() string {
	return fmt.Sprintf("%s %s", dt.Date, dt.Time)
}
