package decorator

import "fmt"

type SendMsgBeforeOrder struct {
	Function OrderTicketI
}

func (s *SendMsgBeforeOrder) PlaceOrder(order Order) {
	sendMsg()
	s.Function.PlaceOrder(order)
}

func sendMsg() {
	fmt.Println("send message...")
}
