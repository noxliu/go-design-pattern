package main

import (
	"fmt"
	"go-design-pattern/factoryMethod/simpleFactoryExample/simpleFactory"
)

func main() {
	car := simpleFactory.GetCar("bmw")
	fmt.Println(car.Drive())

	benz := simpleFactory.GetCar("benz")
	fmt.Println(benz.Drive())
}
