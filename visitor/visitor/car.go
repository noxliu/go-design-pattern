package visitor

type Car struct {
	Wheel  Wheel
	Body   Body
	Engine Engine
}

func (c Car) Accept(visitor Visitor) {
	elements := []Element{
		&c.Engine,
		&c.Wheel,
		&c.Body,
	}

	for _, elem := range elements {
		elem.Accept(visitor)
	}
}
