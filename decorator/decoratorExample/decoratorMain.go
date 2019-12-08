package main

import . "go-design-pattern/decorator/decoratorExample/decorator"

func main() {
	order := Order{
		UserId:    "10099202",
		ProductId: "09988779",
	}
	placeOrder := OrderTicketImplement{}
	sendMsgBeforeOrder := SendMsgBeforeOrder{&placeOrder}
	sendMsgBeforeOrder.Function = &placeOrder
	sendEmailAfterOrder := SendEmailAfterOrder{}
	sendEmailAfterOrder.Function = &sendMsgBeforeOrder
	sendEmailAfterOrder.PlaceOrder(order)
}
