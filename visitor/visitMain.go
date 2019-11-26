package main

import (
	. "go-design-pattern/visitor/visitor"
)

func main() {
	wheel := Wheel{WheelSize: 100}
	engine := Engine{EngineType: "turbo"}
	car := Car{Wheel: wheel, Engine: engine}
	visitorWheel := VisitWheel{}
	VisitEngine := VisitEngine{}
	VisitEngine.VisitEngine(car.Engine)
	visitorWheel.VisitWheel(car.Wheel)
}
