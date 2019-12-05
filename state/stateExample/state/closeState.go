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
}

func (*Close) Run() {
	fmt.Println("启动电梯")
}

func (*Close) Stop() {
}
