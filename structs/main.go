package main

import "fmt"

// typedefs
type person struct {
	firstName string
	lastName  string
}

type animal struct {
	height int8
	length int8
	weight int8
	name   string
	color  string
}

// functions
func (p person) print() {
	fmt.Println(p)
}
func (a animal) print() {
	fmt.Println("This animal (%s) is of size %v * %v", a.name, a.length, a.height)
}

func main() {
	p1 := person{"Mojtaba", "Eshghie"}
	p1.print()

	a := animal{100, 121, 50, "Mar", "Siyah"}
	a.print()
}
