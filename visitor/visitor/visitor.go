package visitor

type Visitor interface {
	VisitWheel(wheel Wheel) string
}
