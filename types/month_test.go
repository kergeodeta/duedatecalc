package types

import "testing"

func TestMonth_String(t *testing.T) {
	tests := []struct {
		name string
		m    Month
		want string
	}{
		{"Before January", Month(-1), "Unknown"},
		{"After December", Month(13), "Unknown"},
		{"May", Month(5), "May"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
