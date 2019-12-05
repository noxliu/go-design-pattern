package state

import "fmt"

type Stop struct {
}

func (*Stop) SetState(s State) {
	fmt.Println("")
}

func (c Stop) Open() {
	fmt.Println("打开电梯门")
}

func (c Stop) Close() {
}

func (c Stop) Run() {
}

func (c Stop) Stop() {

}
