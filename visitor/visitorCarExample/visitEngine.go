package visitorCarExample

type VisitEngine struct {
}

func (*VisitEngine) VisitEngine(engine Engine) string {
	return engine.EngineType
}
