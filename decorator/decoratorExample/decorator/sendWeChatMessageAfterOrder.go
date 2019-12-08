package decorator

import "fmt"

type SendWeChatMessageAfterOrder struct {
	Function OrderTicketI
}

func (s *SendWeChatMessageAfterOrder) PlaceOrder(order Order) {
	s.Function.PlaceOrder(order)
	sendWeChatMessage()
}

func sendWeChatMessage() {
	fmt.Println("发送微信信息")
}
