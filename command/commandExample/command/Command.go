package command

import "fmt"

type Command struct {
}

func (*Command) execCommand() {
	fmt.Println("exec command")
}
