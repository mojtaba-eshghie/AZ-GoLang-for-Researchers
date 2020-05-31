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

func (r room) updateSize(newSize int16) {
	/*
		This is a wrong kind of update! Because it does not actually update the
		original object, rather it will update a copy of the original object
	*/
	r.size = newSize
	fmt.Printf("%+v \n", r)
}

func (r *room) correctlyUpdateSize(newSize int16) {
	r.size = newSize
	r.print()
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

	fmt.Println("\n--------------------Wrong:")
	// The last two updates should not be the same! This is because we missunderstood the idea of pointers in go.
	roomTwo.updateSize(13)
	roomTwo.print()

	fmt.Println("\n--------------------Correct:")
	//let's do this correctly using the concept of the pointers
	roomTwoPointer := &roomTwo
	roomTwoPointer.correctlyUpdateSize(17)
	roomTwoPointer.print()

	// A go pointer shortcut,
	fmt.Println("\n--------------------Pointer Shortcut:")
	/* whenever you have defined the receiver of a function as a receiver that would
	work on a pointer of a type, you could also use that function on a receiver of a
	regular variable of that type. This is called a pointer shortcut in GoLang. For
	instance, the following line is correct:
	*/
	roomTwo.correctlyUpdateSize(99)

}

/*

A very useful note:
	1. Whenever you see a * operator before a "type", then, we are looking for
	a pointer (probably to be passed to a function or smt like that). A good
	Example is line the usage of * in the definition of correctlyUpdateSize function

	2. Whenever you see a * operator before an actual pointer variable, it menas
	that we are dereferencing that pointer (getting the actual value inside the address
	that the pointer is pointing to)
*/
