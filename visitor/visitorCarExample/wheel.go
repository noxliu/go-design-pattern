package visitorCarExample

type Wheel struct {
	WheelSize  int
	WheelBrand string
}

func (e Wheel) Accept(visitor Visitor) {
	visitor.VisitWheel(e)
}
