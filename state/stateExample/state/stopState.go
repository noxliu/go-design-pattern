package state

import "fmt"

type Stop struct {
}

func (c Stop) Open() {
	fmt.Println("打开电梯门")
}

func (c Stop) Close() {
}

func (c Stop) Run() {
	fmt.Println("启动电梯")
}

func (c Stop) Stop() {

}
