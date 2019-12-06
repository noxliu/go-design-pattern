package example

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() { //构建
	f.builder.SetSeats(4).SetWheels(4).SetStructure("small")
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}
