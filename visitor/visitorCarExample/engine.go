package visitorCarExample

type Engine struct {
	EngineType string
}

func (e Engine) Accept(visitor Visitor) {
	visitor.VisitEngine(e)
}
