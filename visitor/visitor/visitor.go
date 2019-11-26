package visitor

type Visitor interface {
	VisitEngine(Engine Engine) string
	VisitWheel(While Wheel) string
}
