package state

import "fmt"

type Run struct {
}

/*
func (*Run) SetState(s State) {
	fmt.Println("")
}*/

func (c Run) Open() {
}

func (c Run) Close() {
}

func (c Run) Run() {
}

func (c Run) Stop() {
	fmt.Println("停止电梯")
}
