package adapter

import "strconv"

type AdapterOfGeneratorForInt struct {
	IntValue int
}

func (i *AdapterOfGeneratorForInt) stringGenerator() string {
	return strconv.Itoa(i.IntValue)
}
