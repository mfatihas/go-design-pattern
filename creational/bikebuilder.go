package creational

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheel() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeat() BuildProcess {
	b.v.Seat = 1
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) Build() VehicleProduct {
	return b.v
}
