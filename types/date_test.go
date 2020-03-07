package types

import (
	"fmt"
	"testing"
)

func TestDate_IsLeapYear(t *testing.T) {
	tests := [...]struct {
		name string
		Date
		want bool
	}{
		{"Test the year 2000", Date{Year: 2000}, true},
		{"Test the year 2020", Date{Year: 2020}, true},
		{"Test the year 2400", Date{Year: 2400}, true},
		{"Test the year 1800", Date{Year: 1800}, false},
		{"Test the year 1900", Date{Year: 1900}, false},
		{"Test the year 2100", Date{Year: 2100}, false},
		{"Test the year 2200", Date{Year: 2200}, false},
		{"Test the year 2300", Date{Year: 2300}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			have := test.Date.IsLeapYear()
			if have != test.want {
				t.Error(fmt.Sprintf("expected result: %v, got %v. Year: %d", test.want, have, test.Date.Year))
			}
		})
	}
}

func TestDate_Weekday(t *testing.T) {
	tests := [...]struct {
		name string
		Date
		want Weekday
	}{
		{"Test date: 1970-01-01", Date{1970, 1, 1}, Thursday},
		{"Test date: 2020-01-01", Date{2020, 1, 1}, Wednesday},
		{"Test date: 2020-02-29", Date{2020, 2, 29}, Saturday},
		{"Test date: 2020-12-31", Date{2020, 12, 31}, Thursday},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			have := test.Date.Weekday()
			if have != test.want {
				t.Error(fmt.Sprintf("expected result: %v, got %v (%d). Date: %s", test.want, have, have, test.Date))
			}
		})
	}
}
