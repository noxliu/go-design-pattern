package visitorCarExample

import "fmt"

type VisitCar struct {
}

func (*VisitCar) VisitCar(car Car) string {
	fmt.Println(car.Body.Size)
	fmt.Println(car.Body.BodyColor)
	return "visit body"
}
