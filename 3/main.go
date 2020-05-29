package main

import "fmt"

type room struct {
	size     int16
	occupied bool
}

func (r room) print() {
	//print the fields in a struct in a civilized manner!
	fmt.Printf("%+v \n", r)
}

func main() {
	roomOne := room{12, true}
	roomOne.print()

	// Create, but assign no properties to a variable using the following approach
	var roomTwo room
	roomTwo.print()
}
