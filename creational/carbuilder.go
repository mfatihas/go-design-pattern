package creational

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheel() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeat() BuildProcess {
	c.v.Seat = 4
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) Build() VehicleProduct {
	return c.v
}
