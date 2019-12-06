package computerBuilder

import "fmt"

type ComputerBuilder interface {
	SetCpu(cpu string) ComputerBuilder
	SetMemory(memory string) ComputerBuilder
	SetHardDisk(hardDisk string) ComputerBuilder
	SetScreenSize(screenSize int) ComputerBuilder
	GetComputerProduct() ComputerItems //显示结果
	SetConstruct(construct Construct) ComputerBuilder
	Build() ComputerItems
}

type ComputerItems struct {
	Cpu             string
	Memory          string
	HardDisk        string
	ScreenSize      int
	ComputerProduct string
	Construct       Construct
}

func (c ComputerItems) PrintComputerInfo() {
	fmt.Println(c.ComputerProduct)
}
