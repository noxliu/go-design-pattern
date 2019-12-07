package decorator

type PlaceOrderI interface {
	PlaceOrder(order Order)
}

type Order struct {
	UserId    string
	ProductId string
}
