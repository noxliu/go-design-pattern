package decorator

import "fmt"

type OrderTicketImplement struct {
}

func (p *OrderTicketImplement) PlaceOrder(order Order) {
	fmt.Println("Complete order...")
}
