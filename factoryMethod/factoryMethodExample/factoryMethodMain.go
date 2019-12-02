package main

import (
	"fmt"
	"go-design-pattern/factoryMethod/factoryMethodExample/factoryMethod"
)

func main() {
	bmwFactory := factoryMethod.BmwFactory{}
	bmw1 := bmwFactory.GetCar("bmw1")
	fmt.Println(bmw1.Drive())

	benzFactory := factoryMethod.BenzFactory{}
	benz1 := benzFactory.GetCar("my benz")
	fmt.Println(benz1.Drive())
}
