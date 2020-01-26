package main

import (
	. "go-design-pattern/command/commandExample/command"
)

func main() {
	onlyPlaceOrdercommand := PlaceOrderCommand{}
	placeOrderAndPostCommand := PlaceOrderAndPostCommand{}
	invoker := Invoker{}
	invoker.AddOrder(&onlyPlaceOrdercommand)
	invoker.AddOrder(&placeOrderAndPostCommand)
	invoker.ExecuteOrders()
}
