package main

import (
	. "go-design-pattern/command/commandExample/command"
)

func main() {
	//fmt.Println("....")
	command := Command{}
	command.ExecCommand()
}
