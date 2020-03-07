package duedatecalc

import "duedatecalc/types"

type Issue struct {
	Submitted      types.DateTime
	TurnaroundTime types.Time
}

func (i Issue) CalculateDueDate() types.DateTime {

	return types.DateTime{}
}
