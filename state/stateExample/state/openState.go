package state

import "fmt"

type Run struct {
}

func SetState(s *State) {
	fmt.Println("")
}

func (*Run) Run() {
	fmt.Println("run")
}

func (*Run) Stop() {
	fmt.Println("cannot stop")
}
