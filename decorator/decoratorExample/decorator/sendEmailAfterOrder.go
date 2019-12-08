package decorator

import "fmt"

type SendEmailAfterOrder struct {
	Function OrderTicketI
}

func (s *SendEmailAfterOrder) PlaceOrder(order Order) {
	sendEmail()
	s.Function.PlaceOrder(order)
}

func sendEmail() {
	fmt.Println("发送邮件")
}
