package main

import (
	. "go-design-pattern/state/stateExample/state"
)

func main() {
	stop := Stop{}
	context := Context{}
	context.SetState(&stop)
	context.Open()

}
