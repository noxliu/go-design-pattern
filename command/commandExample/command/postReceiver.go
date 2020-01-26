package command

import "fmt"

type PostReceiver struct {
}

func (*PostReceiver) execPost() {
	fmt.Println("发货")
}
