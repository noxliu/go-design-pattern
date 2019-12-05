package main

import (
	. "go-design-pattern/state/stateExample/state"
)

func main() {
	run := Run{}
	context := Context{}
	context.SetState(&run)
	context.Run()
}
