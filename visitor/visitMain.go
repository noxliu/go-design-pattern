package main

import (
	. "go-design-pattern/visitor/visitor"
)

func main() {
	wheel := Wheel{WheelSize: 100}
	visitor := VisitWheel{}
	visitor.VisitWheel(wheel)
}
