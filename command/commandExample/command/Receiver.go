package command

import "fmt"

type Receiver struct {
}

func (*Receiver) DoSomething() {
	fmt.Println("receiver do sth")
}
