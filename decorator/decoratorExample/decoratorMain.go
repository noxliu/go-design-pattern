package main

import . "go-design-pattern/decorator/decoratorExample/decorator"

func main() {
	order := Order{
		UserId:    "10099202",
		ProductId: "09988779",
	}
	placeOrder := PlaceOrderImplement{}
	sendMsgBeforeOrder := SendMsgBeforeOrder{placeOrder}
	sendMsgBeforeOrde
	sendMsgBeforeOrder.PlaceOrder(order)
}
