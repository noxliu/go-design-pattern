package simpleFactory

import "fmt"

type Benz struct {
	CarType CarType
}

func (b *Benz) Drive(carName string) string {
	return b.CarType.carName
}

func (c *Benz) ShowMyCar() {
	fmt.Println(c.CarType)
}
