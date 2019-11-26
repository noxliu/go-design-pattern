package visitorCarExample

type Body struct {
	BodyColor string
	Size      int
}

func (b Body) Accept(visitor Visitor) {
	visitor.VisitBody(b)
}
