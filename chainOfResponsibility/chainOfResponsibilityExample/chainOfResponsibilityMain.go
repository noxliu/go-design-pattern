package main

import (
	"fmt"
	. "go-design-pattern/chainOfResponsibility/chainOfResponsibilityExample/chainOfResponsibility"
)

func main() {
	requestChain := RequestChain{}
	projectManager := ProjectManager{}
	requestChain.Manager = &projectManager

	fmt.Print("----")
}
