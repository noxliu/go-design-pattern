package decorator

import "fmt"

type SendEmailAfterOrder struct {
	Decorator OrderTicketI
}

func (s *SendEmailAfterOrder) PlaceOrder(order Order) {
	s.Decorator.PlaceOrder(order)
	sendEmail()
}

func sendEmail() {
	fmt.Println("发送邮件")
}
