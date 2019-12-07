package decorator

import "fmt"

type SendMsgBeforeOrder struct {
	Order Order
}

func (s *SendMsgBeforeOrder) SendMsgBeforePlaceOrder(placeOrder PlaceOrder) {
	sendMsg()
	placeOrder.PlaceOrder(s.Order)
}

func sendMsg() {
	fmt.Println("send message...")
}
