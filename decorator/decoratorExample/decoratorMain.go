package main

import (
	. "./decorator"
	"fmt"
)

func main() {
	order := Order{
		UserId:    "10099202",
		ProductId: "09988779",
	}
	placeOrder := OrderTicketImplement{}
	sendMsgBeforeOrder := SendMsgBeforeOrder{}
	sendEmailAfterOrder := SendEmailAfterOrder{}
	sendWeChatMessageAfterOrder := SendWeChatMessageAfterOrder{}

	//例子1 完成订单之前先发送短信，订单完成之后发送邮件
	//使用短信发送功能包装原有的订票功能
	sendMsgBeforeOrder.Decorator = &placeOrder

	//使用邮件发送功能包装上面的短信发送功能(短信发送功能已经包装了原有的订票功能)，这样一来，就会实现先发送短信，然后订票，然后发送邮件
	sendEmailAfterOrder.Decorator = &sendMsgBeforeOrder
	sendEmailAfterOrder.PlaceOrder(order)

	fmt.Println()

	order2 := Order{
		UserId:    "10098547",
		ProductId: "09989948",
	}
	//例子2 完成订单之前先发送短信，订单完成之后发送邮件和微信消息
	//使用短信发送功能包装原有的订票功能
	sendMsgBeforeOrder.Decorator = &placeOrder

	//使用邮件发送功能包装上面的短信发送功能(短信发送功能已经包装了原有的订票功能)，这样一来，就会实现先发送短信，然后订票，然后发送邮件
	sendEmailAfterOrder.Decorator = &sendMsgBeforeOrder

	//使用微信消息功能包装上面的邮件发送功能(邮件发送功能已经在原有的短信包装了的订票功能上进行了包装)
	sendWeChatMessageAfterOrder.Decorator = &sendEmailAfterOrder
	sendWeChatMessageAfterOrder.PlaceOrder(order2)
}
