package visitorCarExample

type Visitor interface {
	VisitEngine(Engine Engine) (int, string)
	VisitWheel(While Wheel) string
	VisitBody(body Body) string
}
