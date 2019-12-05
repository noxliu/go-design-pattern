package main

import (
	"fmt"
	. "go-design-pattern/state/stateExample/state"
)

func main() {
	fmt.Println("")
	c := Context{}
	run := Run{}
	c.SetState(run)
	c.Run()
}
