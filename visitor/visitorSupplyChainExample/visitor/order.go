package visitor

type Order struct {
	Amount int
	Price  int
}

func (o Order) Accept(visitor Visitor) {
	visitor.VisitOrder(o)
}
