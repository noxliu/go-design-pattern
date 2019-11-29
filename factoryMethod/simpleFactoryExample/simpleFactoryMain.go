package main

import (
	"fmt"
	"go-design-pattern/factoryMethod/simpleFactoryExample/simpleFactory"
)

func main() {
	car1 := simpleFactory.GetCar("bmw")
	fmt.Println(car1.Drive())

	car2 := simpleFactory.GetCar("benz")
	fmt.Println(car2.Drive())
}
