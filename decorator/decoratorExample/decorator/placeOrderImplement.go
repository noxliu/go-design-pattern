package decorator

import "fmt"

type PlaceOrderImplement struct {
}

func (p *PlaceOrderImplement) PlaceOrder(order Order) {
	fmt.Println("Complete order...")
}
