package decorator

type PlaceOrder interface {
	PlaceOrder(order Order)
}

type Order struct {
	UserId    string
	ProductId string
}
