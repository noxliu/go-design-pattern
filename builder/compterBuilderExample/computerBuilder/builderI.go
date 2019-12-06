package computerBuilder

type ComputerBuilder interface {
	SetCpu(cpu string) ComputerBuilder
	SetMemory(memory string) ComputerBuilder
	SetHardDisk(hardDisk string) ComputerBuilder
	SetScreenSize(screenSize int) ComputerBuilder
	GetComputerProduct() string //显示结果
}

type ComputerItems struct {
	Cpu             string
	Memory          string
	HardDisk        string
	ScreenSize      int
	ComputerProduct string
}
