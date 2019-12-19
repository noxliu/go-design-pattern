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

	projectManagerChain.HandleFeeRequest("张三", 100)
	projectManagerChain.HandleFeeRequest("张三", 600)
	projectManagerChain.HandleFeeRequest("李四", 499)
	projectManagerChain.HandleFeeRequest("王五", 1099)
	projectManagerChain.HandleFeeRequest("李四", 199)

}
