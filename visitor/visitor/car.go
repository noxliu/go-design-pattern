package visitor

import (
	"fmt"
)

type Car struct {
	Wheel    Wheel
	CarColor string
	Engine   Engine
}

func (c Car) Accept(visitor Visitor) {
	//
	fmt.Println("wheel implement")
}
