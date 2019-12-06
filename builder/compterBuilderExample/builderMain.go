package main

import (
	"fmt"
	. "go-design-pattern/builder/compterBuilderExample/computerBuilder"
)

func main() {
	builder := ComputerBuilderImplement{}
	ComputerConstruct := ComputerConstruct{}
	builder.SetCpu("Intel i7").SetMemory("16GB").SetHardDisk("1t ssd").SetScreenSize(21).SetConstruct(ComputerConstruct).Build()
	//ComputerConstruct.Build(&builder)
	fmt.Println(builder.GetComputerProduct())
}
