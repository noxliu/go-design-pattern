package main

import (
	. "./builder"
	"fmt"
)

func main() {
	carBuilder := CarBuilder{}
	car := carBuilder.SetSeats(4).SetStructure("SUV").SetWheels(4).GetVehicle()
	fmt.Println(car)
}
