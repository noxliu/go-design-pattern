package main

import (
	. "go-design-pattern/chainOfResponsibility/chainOfResponsibilityExample/chainOfResponsibility"
)

func main() {
	projectManager := ProjectManager{}
	projectManagerChain := RequestChain{}
	projectManagerChain.Manager = &projectManager

	departmentManager := DepartmentManager{}
	departmentManagerChain := RequestChain{}
	departmentManagerChain.Manager = &departmentManager

	generalManager := GeneralManager{}
	generalManagerChain := RequestChain{}
	generalManagerChain.Manager = &generalManager

	departmentManagerChain.SetSuccessor(&generalManagerChain)
	projectManagerChain.SetSuccessor(&departmentManagerChain)

	projectManagerChain.HandleFeeRequest("rob", 100)

}
