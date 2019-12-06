package computerBuilder

import (
	"strconv"
)

type ComputerConstruct struct {
	ComputerBuilder ComputerBuilder
}

func (con ComputerConstruct) BuildComputer(ComputerItems ComputerItems) ComputerItems {
	ComputerItems.ComputerProduct = "台式电脑：" +
		ComputerItems.Cpu + " " +
		ComputerItems.Memory + " " +
		ComputerItems.HardDisk + " " +
		strconv.Itoa(ComputerItems.ScreenSize) + "寸显示屏"
	return ComputerItems
}
