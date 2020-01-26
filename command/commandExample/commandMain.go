package main

import (
	. "go-design-pattern/command/commandExample/command"
)

func main() {
	//fmt.Println("....")
	command := PlaceOrderCommand{}
	invoker := Invoker{}
	invoker.AddOrder(&command)
	invoker.ExecuteOrders()
}
