package main

import (
	. "go-design-pattern/visitor/visitor"
)

func main() {
	wheel := Wheel{WheelSize: 100}
	engine := Engine{EngineType: "turbo"}
	carBody := Body{BodyColor: "green", Size: 5}
	car := Car{Wheel: wheel, Engine: engine, Body: carBody}

	visitorWheel := VisitWheel{}
	visitEngine := VisitEngine{}
	visitBody := VisitBody{}

	visitEngine.VisitEngine(car.Engine)
	visitorWheel.VisitWheel(car.Wheel)
	visitBody.VisitBody(car.Body)
}
