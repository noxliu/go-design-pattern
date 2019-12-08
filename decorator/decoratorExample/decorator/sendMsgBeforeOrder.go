package decorator

import "fmt"

type SendMsgBeforeOrder struct {
	Decorator OrderTicketI
}

func (s *SendMsgBeforeOrder) PlaceOrder(order Order) {
	sendMsg()
	s.Decorator.PlaceOrder(order)
}

func sendMsg() {
	fmt.Println("发送短信")
}
