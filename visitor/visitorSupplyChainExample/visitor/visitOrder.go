package visitor

type VisitOrder struct {
}

func (*VisitOrder) VisitOrder(order Order) int {
	return order.Amount * order.Price
}
