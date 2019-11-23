package builder

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats(4).SetWheels(4).SetStructure("small")
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}
