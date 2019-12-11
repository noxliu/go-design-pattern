package adapter

import "strconv"

type IntToStringAdapter struct {
	IntValue int
}

func (i *IntToStringAdapter) stringGenerator() string {
	return strconv.Itoa(i.IntValue)
}
