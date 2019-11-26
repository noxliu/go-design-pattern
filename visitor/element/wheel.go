package element

import (
	"fmt"
	"go-design-pattern/visitor/visitor"
)

type Wheel struct {
}

func (c *Wheel) Accept(visitor visitor.Visitor) {
	//
	fmt.Println("wheel implement")
	visitor.VisitWheel(c)
}
