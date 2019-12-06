package main

import (
	"fmt"
	"go-design-pattern/builder/builder/example"
)

func main() {
	carBuilder := example.CarBuilder{}
	car := carBuilder.SetSeats(4).SetStructure("SUV").SetWheels(4).GetVehicle()
	//上面的没有写对
	fmt.Println(car)
}
