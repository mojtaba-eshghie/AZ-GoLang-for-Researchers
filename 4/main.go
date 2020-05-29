package main

import "fmt"

type window struct {
	height int16
	width  int16
	bright bool
}

type room struct {
	size       int16
	occupied   bool
	roomWindow window
}

func (r room) print() {
	//print the fields in a struct in a civilized manner!
	fmt.Printf("%+v \n", r)
}

func main() {
	roomOne := room{
		size:     12,
		occupied: true,
		roomWindow: window{
			height: 180,
			width:  200,
			bright: true,
		},
	}
	roomOne.print()

	// Create, but assign no properties to a variable using the following approach
	var roomTwo room
	roomTwo.print()
}
