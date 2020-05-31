package main

import "fmt"

/*
Let's talk about maps, which are a collection of key:value pairs! (dict in py, object in js,
 a hash in ruby).
Both the key and value are "Statically Typed"! ::: All of the keys in a map have to be of the
same exact type, and all of the values should be of the same type as well. (The key:value pair
do not have to be the same type)

*/

func main() {

	// [keys]values
	colors := map[string]string{
		"red":   "#00ffff",
		"blue":  "#898998",
		"white": "#000000",
	}
	fmt.Println(colors)
}
