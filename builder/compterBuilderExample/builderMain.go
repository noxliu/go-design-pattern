package main

import (
	"fmt"
	. "go-design-pattern/builder/compterBuilderExample/computerBuilder"
)

func main() {
	builder := ComputerBuilderImplement{}
	test := builder.SetCpu("Intel i7").SetMemory("16GB").SetHardDisk("1t ssd").SetScreenSize(21)
	ComputerConstruct := ComputerConstruct{}
	fmt.Println("///")
	fmt.Println(test.ComputerItems.Cpu)
	ComputerConstruct.Build(&builder)
	fmt.Println(builder.GetComputerProduct())
}
