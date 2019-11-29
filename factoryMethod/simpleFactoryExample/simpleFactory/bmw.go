package simpleFactory

import "fmt"

type Bmw struct {
	CarType CarType
}

func (b Bmw) Drive() string {
	return b.CarType.carName
}

func (c *Bmw) ShowMyCar() {
	fmt.Println(c.CarType)
}
