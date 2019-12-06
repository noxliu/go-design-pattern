package computerBuilder

type Construct interface {
	BuildComputer(ComputerItems ComputerItems) ComputerItems
}
