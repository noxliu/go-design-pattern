package visitor

type VisitWheel struct {
}

func (*VisitWheel) VisitWheel(wheel Wheel) (int, string) {
	return wheel.WheelSize, wheel.WheelBrand
}
