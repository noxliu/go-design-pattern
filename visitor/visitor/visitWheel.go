package visitor

import (
	. "go-design-pattern/visitor/element"
)

type VisitWheel struct {
}

func (*VisitWheel) VisitWheel(wheel Wheel) string {
	return "wheel"
}
