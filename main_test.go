package main

import "testing"

func TestSum(t *testing.T) {
	n := Number{
		Number1: 5,
		NUmber2: 6,
	}

	expected := 11

	res := n.sum()

	if res != expected {
		t.Errorf("Sum function doesn't work %d + %d isn't %d", n.Number1, n.NUmber2, res)
	}
}

func TestGetFullName(t *testing.T) {
	p := Person{
		FirstName: "A",
		LastName:  "B",
	}

	expected := "A B"

	res := p.GetFullName()

	if res != expected {
		t.Errorf("Get Full Name doest work properly First Name = %s, Last Name = %s, result Full Name = %s", p.FirstName, p.LastName, res)
	}
}
