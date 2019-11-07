package creational

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := GetInstance()

	carBuilder := &CarBuilder{}

	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.Build()

	if car.Wheels != 4 {
		t.Errorf("Wheels of car must be 4, not %d\n", car.Wheels)
	}

	if car.Seat != 4 {
		t.Errorf("Seats of car must be 4 not %d\n", car.Seat)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure of the car must be 'Car' not %s\n", car.Structure)
	}

	bikeBuilder := &BikeBuilder{}

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	bike := bikeBuilder.Build()

	if bike.Wheels != 2 {
		t.Errorf("Wheels of bike must be 2, not %d\n", bike.Wheels)
	}

	if bike.Seat != 1 {
		t.Errorf("Seats of bike must be 1 not %d\n", bike.Seat)
	}

	if bike.Structure != "Bike" {
		t.Errorf("Structure of the Bike must be 'Bike' not %s\n", bike.Structure)
	}

}

func TestGetInstanceSingleton(t *testing.T) {
	manufactur1 := GetInstance()

	if manufactur1 == nil {
		t.Error("Expected pointer to singleton after calling GetInstance(), not nil")
	}

	expectedCounter := manufactur1

	manufactur2 := GetInstance()
	if manufactur2 != expectedCounter {
		t.Error("Expected second instance have same instance with the first one")
	}

}
