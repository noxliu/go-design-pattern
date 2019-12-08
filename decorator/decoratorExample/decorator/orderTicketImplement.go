package decorator

import "fmt"

type OrderTicketImplement struct {
}

func (p *OrderTicketImplement) PlaceOrder(order Order) {
	fmt.Println("完成订单， 产品ID：" + order.ProductId + ", 用户ID：" + order.UserId)
}
