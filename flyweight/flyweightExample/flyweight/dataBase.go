package flyweight

var template1 = "感谢你的支持, 你购买的商品<<ProductName>>, 订单号<<OrderNo>>,已经发货, 快递单号为<<TrackNo>>"
var template2 = "客户你好, 你关注的商品<<ProductName>>, 正在促销打折! 数量不多!"
var template3 = "客户你好, 我们已经你对订单<<OrderNo>>的评价, 感谢!"

func QueryByType(templateNo int) string {
	if templateNo == 1 {
		return template1
	} else if templateNo == 2 {
		return template2
	} else if templateNo == 3 {
		return template3
	} else {
		return ""
	}
}
