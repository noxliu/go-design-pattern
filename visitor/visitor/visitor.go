package visitor

import (
	. "go-design-pattern/visitor/element"
)

type Visitor interface {
	VisitWheel(wheel Wheel) string
}
