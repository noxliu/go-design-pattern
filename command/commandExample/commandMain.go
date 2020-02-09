package main

import (
	. "./command"
)

func main() {
	onlyPlaceOrdercommand := PlaceOrderCommand{}
	placeOrderAndPostCommand := PlaceOrderAndPostCommand{}
	invoker := Invoker{}
	invoker.AddOrder(&onlyPlaceOrdercommand)
	invoker.AddOrder(&placeOrderAndPostCommand)
	invoker.ExecuteOrders()
}
