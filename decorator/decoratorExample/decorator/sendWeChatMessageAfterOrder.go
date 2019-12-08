package decorator

import "fmt"

type SendWeChatMessageAfterOrder struct {
	Decorator OrderTicketI
}

func (s *SendWeChatMessageAfterOrder) PlaceOrder(order Order) {
	s.Decorator.PlaceOrder(order)
	sendWeChatMessage()
}

func sendWeChatMessage() {
	fmt.Println("发送微信信息")
}
