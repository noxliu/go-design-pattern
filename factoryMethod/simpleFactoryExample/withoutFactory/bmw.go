package simpleFactory

type Bmw struct {
	CarType CarType
}

func (b Bmw) Drive() string {
	return b.CarType.carName
}
