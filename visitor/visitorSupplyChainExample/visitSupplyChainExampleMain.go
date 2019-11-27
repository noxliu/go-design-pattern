package main

import (
	"fmt"
	"go-design-pattern/visitor/visitorSupplyChainExample/visitor"
)

func main() {
	order := visitor.Order{
		Price:  100,
		Amount: 5,
	}

	visitOrder := visitor.VisitOrder{}
	totalPrice := visitOrder.VisitOrder(order)
	fmt.Println(totalPrice)
}
