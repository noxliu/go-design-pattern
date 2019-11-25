package element

import "go-design-pattern/visitor/visitor"

type Wheel struct {
}

func (*Wheel) Accept(visitor visitor.Visitor) {
	//
}
