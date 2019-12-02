package factoryMethod

type Benz struct {
	CarType CarType
}

func (b Benz) Drive() string {
	return b.CarType.carName
}
