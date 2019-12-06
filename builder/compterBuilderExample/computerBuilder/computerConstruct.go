package computerBuilder

import "strconv"

func (c ComputerBuilderImplement) Build() ComputerBuilder {
	c.ComputerItems.ComputerProduct = "台式电脑：" +
		c.ComputerItems.Cpu +
		c.ComputerItems.Memory +
		c.ComputerItems.HardDisk +
		strconv.Itoa(c.ComputerItems.ScreenSize)
	return c
}
