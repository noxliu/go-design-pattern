package computerBuilder

type ComputerBuilderImplement struct {
	ComputerItems ComputerItems
}

func (c ComputerBuilderImplement) SetCpu(cpu string) ComputerBuilder {

	return nil
}

func (c ComputerBuilderImplement) SetMemory(memory string) ComputerBuilder {
	return nil
}

func (c ComputerBuilderImplement) SetHardDisk(hardDisk string) ComputerBuilder {
	return nil
}

func (c ComputerBuilderImplement) SetScreenSize(screenSize int) ComputerBuilder {
	return nil
}

func (c ComputerBuilderImplement) GetComputerProduct() string {
	return nil
}
