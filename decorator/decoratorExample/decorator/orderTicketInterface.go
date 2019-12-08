package decorator

type OrderTicketI interface {
	PlaceOrder(order Order)
}

type Order struct {
	UserId    string
	ProductId string
}
