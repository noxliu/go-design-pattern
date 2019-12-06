package computerBuilder

import "strconv"

type ComputerConstruct struct {
	ComputerBuilder ComputerBuilder
}

func (con ComputerConstruct) Build(c *ComputerBuilderImplement) ComputerBuilder {
	c.ComputerItems.ComputerProduct = "台式电脑：" +
		c.ComputerItems.Cpu +
		c.ComputerItems.Memory +
		c.ComputerItems.HardDisk +
		strconv.Itoa(c.ComputerItems.ScreenSize)
	return c
}
