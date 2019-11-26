package visitor

import (
	"fmt"
)

type Wheel struct {
	WheelSize int
}

func (c Wheel) Accept(visitor Visitor) {
	//
	fmt.Println("wheel implement")
	visitor.VisitWheel(c)

}
