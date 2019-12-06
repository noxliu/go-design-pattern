package computerBuilder

type ComputerBuilderImplement struct {
	ComputerItems ComputerItems
}

func (c ComputerBuilderImplement) SetCpu(cpu string) ComputerBuilder {
	c.ComputerItems.Cpu = cpu
	return c
}

func (c ComputerBuilderImplement) SetMemory(memory string) ComputerBuilder {
	c.ComputerItems.Memory = memory
	return c
}

func (c ComputerBuilderImplement) SetHardDisk(hardDisk string) ComputerBuilder {
	c.ComputerItems.HardDisk = hardDisk
	return c
}

func (c ComputerBuilderImplement) SetScreenSize(screenSize int) ComputerBuilder {
	c.ComputerItems.ScreenSize = screenSize
	return c
}

func (c ComputerBuilderImplement) GetComputerProduct() ComputerItems {
	return c.ComputerItems
}

func (c ComputerBuilderImplement) SetConstruct(construct Construct) ComputerBuilder {
	c.ComputerItems.Construct = construct
	return c
}

func (c ComputerBuilderImplement) Build() ComputerItems {
	return c.ComputerItems.Construct.BuildComputer(c.ComputerItems)
}
