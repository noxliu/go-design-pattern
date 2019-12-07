package decorator

import "fmt"

type SendEmailAfterOrder struct {
	Function PlaceOrderI
}

func (s *SendEmailAfterOrder) PlaceOrder(order Order) {
	sendEmail()
	s.Function.PlaceOrder(order)
}

func sendEmail() {
	fmt.Println("send email...")
}
