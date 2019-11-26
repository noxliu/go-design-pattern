package visitor

type Wheel struct {
	WheelSize int
}

func (e Wheel) Accept(visitor Visitor) {
	visitor.VisitWheel(e)
}
