package visitor

type Visitor interface {
	VisitOrder(orderer Order)
}
