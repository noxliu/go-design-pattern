package main

import (
	"fmt"
	. "go-design-pattern/builder/compterBuilderExample/computerBuilder"
)

func main() {
	builder := ComputerBuilderImplement{}
	builder.SetCpu("Intel i7").SetMemory("16GB").SetHardDisk("1t ssd").SetScreenSize(21).GetComputerProduct()
	ComputerConstruct := ComputerConstruct{}
	fmt.Println("///")
	ComputerConstruct.Build(&builder)
	fmt.Println(builder.GetComputerProduct())
}
