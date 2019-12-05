package state

import (
	"fmt"
)

type Open struct {
}

func (*Open) SetState(s State) {
	fmt.Println("")
}

func (c Open) Open() {
	fmt.Println("电梯门打卡")
}

func (c Open) Close() {
}

func (c Open) Run() {
}

func (c Open) Stop() {
}
