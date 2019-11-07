package creational

type BuildProcess interface {
	SetWheel() BuildProcess
	SetSeat() BuildProcess
	SetStructure() BuildProcess
	Build() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

type VehicleProduct struct {
	Wheels    int
	Seat      int
	Structure string
}

var instance *ManufacturingDirector

func GetInstance() *ManufacturingDirector {
	if instance == nil {
		instance = new(ManufacturingDirector)
	}
	return instance
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetSeat().SetStructure().SetWheel()
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}
