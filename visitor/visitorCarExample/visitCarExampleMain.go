package main

import (
	"fmt"
	"go-design-pattern/visitor/visitorCarExample/visitor"
)

func main() {
	wheel := visitor.Wheel{
		WheelSize:  225,
		WheelBrand: "MICHELIN",
	}
	engine := visitor.Engine{
		EngineType: "turbocharging",
	}
	carBody := visitor.Body{
		BodyColor: "green",
		Size:      5,
	}
	car := visitor.Car{
		Wheel:  wheel,
		Engine: engine,
		Body:   carBody,
	}

	//模拟对汽车对象的访问，并进行额外的业务逻辑操作，这些操作不会侵入原有业务逻辑
	visitorWheel := visitor.VisitWheel{}
	visitEngine := visitor.VisitEngine{}
	visitBody := visitor.VisitBody{}

	engineType := visitEngine.VisitEngine(car.Engine)
	wheelSize, wheelBrand := visitorWheel.VisitWheel(car.Wheel)
	bodyColor := visitBody.VisitBody(car.Body)

	fmt.Printf("Engine type: %s, Wheel size: %d, wheel brand: %s, body color: %s", engineType, wheelSize, wheelBrand, bodyColor)
}
