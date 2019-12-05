package state

import (
	"fmt"
)

type Close struct {
}

func (*Close) Open() {
	fmt.Println("打开电梯门")
}

func (*Close) Close() {
}

func (*Close) Run() {
	fmt.Println("启动电梯")
}

func (*Close) Stop() {
}
