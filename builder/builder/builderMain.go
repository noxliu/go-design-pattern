package main

import (
	. "./example"
	"fmt"
)

func main() {
	carBuilder := CarBuilder{}
	car := carBuilder.SetSeats(4).SetStructure("SUV").SetWheels(4).GetVehicle()
	//上面的没有写对
	fmt.Println(car)
}
