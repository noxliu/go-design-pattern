package visitor

type Car struct {
	Wheel    Wheel
	CarColor string
	Engine   Engine
}

func (c Car) Accept(visitor Visitor) {
	elements := []Element{
		&c.Engine,
		&c.Wheel,
	}

	for _, elem := range elements {
		elem.Accept(visitor)
	}
}
