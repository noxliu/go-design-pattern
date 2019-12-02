package main

import (
	"fmt"
	"go-design-pattern/factoryMethod/factoryMethodExample/factoryMethod"
)

func main() {
	bmwFactory := factoryMethod.BmwFactory{}
	bmw1 := bmwFactory.GetCar("bmw1")

	fmt.Println("The car brand is %s, name is %s ", bmw1)

}
