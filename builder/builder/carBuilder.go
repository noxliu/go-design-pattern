package builder

type Test struct {
	Name string
}

type CarBuilder struct {
	VehicleProduct VehicleProduct
}

func (c *CarBuilder) SetWheels(wheels int) BuildProcess {
	c.VehicleProduct.Wheels = wheels
	return c
}

func (c *CarBuilder) SetSeats(seats int) BuildProcess {
	c.VehicleProduct.Seats = seats
	return c
}

func (c *CarBuilder) SetStructure(structure string) BuildProcess {
	c.VehicleProduct.Structure = structure
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.VehicleProduct
}
