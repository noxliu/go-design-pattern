package computerBuilder

import "strconv"

type ComputerConstruct struct {
	ComputerBuilder ComputerBuilder
}

//BuildComputer(ComputerItems ComputerItems) ComputerItems

func (con ComputerConstruct) BuildComputer(ComputerItems ComputerItems) ComputerBuilder {
	ComputerItems.ComputerProduct = "台式电脑：" +
		ComputerItems.Cpu +
		ComputerItems.Memory +
		ComputerItems.HardDisk +
		strconv.Itoa(ComputerItems.ScreenSize)
	return ComputerItems
}
