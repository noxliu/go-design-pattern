package state

import (
	"fmt"
	"reflect"
)

type Run struct {
}

func (*Run) SetState(s State) {
	fmt.Println("")
}

func (c Run) Run() {
	fmt.Println(reflect.TypeOf(c))
	fmt.Println("启动电梯")
}

func (*Run) Stop() {
	fmt.Println("cannot stop")
}
