package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       uint
}

type Number struct {
	Number1 int
	NUmber2 int
}

func main() {
	p := Person{
		FirstName: "Muhammad Fatih Abdus",
		LastName:  "Salam",
	}
	fmt.Println(p.GetFullName())
}

func (p *Person) GetFullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (n *Number) sum() int {
	return n.Number1 + n.NUmber2
}
