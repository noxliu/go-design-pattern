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
	fmt.Println("cannot stop")
}

func (*Close) Run() {
}

func (*Close) Stop() {
}
