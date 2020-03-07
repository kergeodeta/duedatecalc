package types

import "testing"

func TestWeekday_IsWeekend(t *testing.T) {
	tests := []struct {
		name string
		d    Weekday
		want bool
	}{
		{"Saturday", Saturday, true},
		{"Monday", Monday, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.IsWeekend(); got != tt.want {
				t.Errorf("IsWeekend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeekday_String(t *testing.T) {
	tests := []struct {
		name string
		d    Weekday
		want string
	}{
		{"Before Monday", Weekday(0), "Unknown"},
		{"After Sunday", Weekday(8), "Unknown"},
		{"Wednesday", Weekday(3), "Wednesday"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
