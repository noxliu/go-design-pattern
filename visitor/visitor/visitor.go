package visitor

type Visitor interface {
	visitWheel(wheel Wheel) string
	visitEngine(engine Engine) string
	visitBody(body Body) string
	visitCar(car Car) string
}
