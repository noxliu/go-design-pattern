package main

import (
	. "./builder"
	"fmt"
)

func main() {
	c := Test{Name: "aa"}
	fmt.Println(c)
	fmt.Println("-------")
	vehi := VehicleProduct{}
	//builder := builder.ManufacturingDirector{}
	carbuilder := CarBuilder{}
	carbuilder.VehicleProduct = vehi
	car := carbuilder.SetSeats(4).SetStructure("ss").SetWheels(4).GetVehicle()
	fmt.Println(car)
}
