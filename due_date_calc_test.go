package duedatecalc

import (
	"duedatecalc/types"
	"reflect"
	"testing"
)

func TestIssue_CalculateDueDate(t *testing.T) {
	tests := []struct {
		name      string
		submitted types.DateTime
		duration  types.Duration
		want      types.DateTime
	}{
		{
			"5 perc hozzáadása úgy, hogy órán belül maradunk",
			types.CreateDateTime(2020, 2, 28, 12, 0, 0),
			*types.CreateDuration(0, 5),
			types.CreateDateTime(2020, 2, 28, 12, 5, 0),
		},
		{
			"30 perc hozzáadása úgy, hogy az óra értéke is növekszik",
			types.CreateDateTime(2020, 2, 28, 12, 45, 0),
			*types.CreateDuration(0, 30),
			types.CreateDateTime(2020, 2, 28, 13, 15, 0),
		},
		{
			"1 teljes óra hozzáadása",
			types.CreateDateTime(2020, 2, 28, 12, 0, 0),
			*types.CreateDuration(1, 5),
			types.CreateDateTime(2020, 2, 28, 13, 5, 0),
		},
		{
			"1 teljes óra hozzáadása, hogy napváltás is legyen",
			types.CreateDateTime(2020, 2, 27, 16, 30, 0),
			*types.CreateDuration(1, 5),
			types.CreateDateTime(2020, 2, 28, 8, 35, 0),
		},
		{
			"1 teljes óra hozzáadása, hogy napváltás is legyen hétvége figyelembevételével",
			types.CreateDateTime(2020, 2, 28, 16, 30, 0),
			*types.CreateDuration(1, 5),
			types.CreateDateTime(2020, 3, 2, 8, 35, 0),
		},
		{
			"16 teljes óra hozzáadása ahogy a kiírásban szereplő példa is megadja",
			types.CreateDateTime(2020, 3, 3, 14, 12, 0),
			*types.CreateDuration(16, 0),
			types.CreateDateTime(2020, 3, 5, 14, 12, 0),
		},
		{
			"18 teljes óra hozzáadása, hogy hétvége is legyen benne ",
			types.CreateDateTime(2020, 3, 13, 14, 12, 0),
			*types.CreateDuration(18, 0),
			types.CreateDateTime(2020, 3, 17, 16, 12, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Issue{
				Submitted:      tt.submitted,
				TurnaroundTime: tt.duration,
			}
			if got := i.CalculateDueDate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateDueDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateFirstDayHours(t *testing.T) {
	tests := []struct {
		name              string
		dt                types.DateTime
		wantStartFraction types.Time
	}{
		{
			"Ha reggel 8 órától nézzük",
			types.CreateDateTime(2020, 1, 1, 8, 0, 0),
			types.Time{8, 0, 0},
		},
		{
			"Ha reggel 9 órától nézzük",
			types.CreateDateTime(2020, 1, 1, 9, 0, 0),
			types.Time{8, 0, 0},
		},
		{
			"Ha nap közben nézzük",
			types.CreateDateTime(2020, 1, 1, 12, 30, 0),
			types.Time{4, 30, 0},
		},
		{
			"Ha nap közben nézzük",
			types.CreateDateTime(2020, 1, 1, 16, 15, 0),
			types.Time{0, 45, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStartFraction := calculateFirstDayHours(tt.dt); !reflect.DeepEqual(gotStartFraction, tt.wantStartFraction) {
				t.Errorf("calculateFirstDayHours() = %v, want %v", gotStartFraction, tt.wantStartFraction)
			}
		})
	}
}
