package visitor

import (
	"fmt"
)

type VisitWheel struct {
}

func (*VisitWheel) VisitWheel(wheel Wheel) string {
	fmt.Println(wheel.WheelSize)
	return "visit wheel"
}
