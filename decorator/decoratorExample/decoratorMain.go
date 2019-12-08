package main

import . "go-design-pattern/decorator/decoratorExample/decorator"

func main() {
	order := Order{
		UserId:    "10099202",
		ProductId: "09988779",
	}
	placeOrder := OrderTicketImplement{}
	sendMsgBeforeOrder := SendMsgBeforeOrder{}
	sendEmailAfterOrder := SendEmailAfterOrder{}

	//例子1 完成订单之前先发送短信，订单完成之后发送邮件
	//使用短信发送功能包装原有的订票功能
	sendMsgBeforeOrder.Function = &placeOrder

	//使用邮件发送功能包装上面的短信发送功能（短信发送功能已经包装了原有的订票功能），这样一来，就会实现先发送短信，然后订票，然后发送邮件
	sendEmailAfterOrder.Function = &sendMsgBeforeOrder
	sendEmailAfterOrder.PlaceOrder(order)

	//例子2 完成订单之前先发送短信，订单完成之后发送邮件和微信消息
}
