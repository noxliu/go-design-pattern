package decorator

import "fmt"

type SendWeChatMessageBeforeOrder struct {
	Function OrderTicketI
}

func (s *SendWeChatMessageBeforeOrder) PlaceOrder(order Order) {
	sendWeChatMessage()
	s.Function.PlaceOrder(order)
}

func sendWeChatMessage() {
	fmt.Println("发送微信信息")
}
