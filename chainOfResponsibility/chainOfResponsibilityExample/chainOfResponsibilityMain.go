package main

import (
	"./chainOfResponsibility"
)

func main() {
	projectManager := chainOfResponsibility.ProjectManager{}
	projectManagerChain := chainOfResponsibility.RequestChain{}
	projectManagerChain.Manager = &projectManager

	departmentManager := chainOfResponsibility.DepartmentManager{}
	departmentManagerChain := chainOfResponsibility.RequestChain{}
	departmentManagerChain.Manager = &departmentManager

	generalManager := chainOfResponsibility.GeneralManager{}
	generalManagerChain := chainOfResponsibility.RequestChain{}
	generalManagerChain.Manager = &generalManager

	departmentManagerChain.SetSuccessor(&generalManagerChain)
	projectManagerChain.SetSuccessor(&departmentManagerChain)

	projectManagerChain.HandleFeeRequest("张三", 100)
	projectManagerChain.HandleFeeRequest("张三", 600)
	projectManagerChain.HandleFeeRequest("李四", 499)
	projectManagerChain.HandleFeeRequest("王五", 1099)
	projectManagerChain.HandleFeeRequest("李四", 199)

}
