package decorator

import "fmt"

type SendEmailAfterOrder struct {
	Function OrderTicketI
}

func (s *SendEmailAfterOrder) PlaceOrder(order Order) {
	s.Function.PlaceOrder(order)
	sendEmail()
}

func sendEmail() {
	fmt.Println("发送邮件")
}
