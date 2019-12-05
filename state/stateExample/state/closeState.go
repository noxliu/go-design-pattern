package state

import (
	"fmt"
)

type Close struct {
}

func (*Close) SetState(s State) {
	fmt.Println("")
}

func (*Close) Open() {
}

func (*Close) Close() {
	fmt.Println("关闭电梯门")
}

func (*Close) Run() {
}

func (*Close) Stop() {
}
