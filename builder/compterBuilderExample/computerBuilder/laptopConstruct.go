package computerBuilder

import (
	"strconv"
)

type LaptopConstruct struct {
	ComputerBuilder ComputerBuilder
}

func (con LaptopConstruct) BuildComputer(ComputerItems ComputerItems) ComputerItems {
	ComputerItems.ComputerProduct = "笔记本电脑：" +
		ComputerItems.Cpu + " " +
		ComputerItems.Memory + " " +
		ComputerItems.HardDisk + " " +
		strconv.Itoa(ComputerItems.ScreenSize) + "寸触摸显示屏"
	return ComputerItems
}
