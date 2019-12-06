package main

import (
	. "go-design-pattern/builder/compterBuilderExample/computerBuilder"
)

func main() {
	builder := ComputerBuilderImplement{}
	ComputerConstruct := ComputerConstruct{}
	builder.SetCpu("Intel i7").
		SetMemory("16GB").
		SetHardDisk("1TSSD").
		SetScreenSize(21).
		SetConstruct(ComputerConstruct).
		Build().PrintComputerInfo()
}
