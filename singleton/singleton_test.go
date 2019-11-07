package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("Expected pointer to singleton after calling GetInstance(), not nil")
	}

	expectedCounter := counter1
	currentCount := counter1.AddOne()

	if currentCount != 1 {
		t.Errorf("After calling AddOne() for the first time its must be 1 not %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("Expected second instance have same instance with the first one")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling AddOne() for the second time its must be 2 not %d\n", currentCount)
	}
}
