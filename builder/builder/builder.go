package builder

type BuildProcess interface {
	SetWheels(wheels int) BuildProcess
	SetSeats(seats int) BuildProcess
	SetStructure(structure string) BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}
