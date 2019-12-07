package decorator

import "fmt"

type PlaceOrderImplment struct {
}

func (p *PlaceOrderImplment) PlaceOrder(order Order) {
	fmt.Println()
}
