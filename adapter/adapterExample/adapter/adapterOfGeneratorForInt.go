package adapter

import "strconv"

type AdapterOfGeneratorForInt struct {
	IntValue int
}

func (i *AdapterOfGeneratorForInt) StringGenerator() string {
	return strconv.Itoa(i.IntValue)
}
