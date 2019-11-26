package visitor

import "fmt"

type VisitEngine struct {
}

func (*VisitEngine) VisitEngine(engine Engine) string {
	fmt.Println(engine.EngineType)
	return "visit engine"
}
