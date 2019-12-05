package state

import (
	"fmt"
)

type Open struct {
}

/*
func (*Open) SetState(s State) {
	fmt.Println("")
}*/

func (c Open) Open() {
}

func (c Open) Close() {
	fmt.Println("关闭电梯门")
}

func (c Open) Run() {
}

func (c Open) Stop() {
}
