package main

import (
	"fmt"
)

func printMap(c map[string]string) {
	for color, code := range c {
		fmt.Println("The color ", color, " has code "+code)
	}
}

func main() {
	colors := map[string]string{
		"white": "#ffff",
		"black": "#0000",
		"blue":  "#8989",
	}
	printMap(colors)
}

/*
	maps vs structs in go
	1. In maps, all the keys together and all the values together should have
	the same type. But, in structs we are free to have fields of different types.
	2. Keys are indexed ==> what the benefit? ==> we can iterate over them using a
	for loop! The same is not true for structs, and iteration over key:value pairs is
	not possible
	3. Maps: reference types / structs: value types ==> meaning that when you
	pass a map, you are actually passing a pointer to the underlying data structure,
	whereas when we pass a struct to a function we make a copy of the entire struct,
	and that copy is passed to the function.

	Guideline:
		* Use a map when we want to have a collection of very closely related properties
		* With a map, we do not really need to know all the keys at compile time,
			and we could set them on the go when the program is running based on
			input, ... over time. And, we can delete key:value paris as well! However,
			with the structs, we have to very clearly define all of the property names
			and their types at compile time.
*/
