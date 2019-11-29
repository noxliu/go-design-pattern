package simpleFactory

import "fmt"

type Benz struct {
	CarType CarType
}

func (*Benz) Drive(carName string) CarType {
	car := CarType{
		carName: carName,
		size:    10,
	}

	return car
}

func (c *Benz) ShowMyCar() {
	fmt.Println(c.CarType)
}
