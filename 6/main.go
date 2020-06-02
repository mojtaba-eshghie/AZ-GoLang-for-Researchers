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

	teamMartialStatus := map[string]bool{
		"Mojtaba":  false,
		"Leyla":    false,
		"Mohammad": true,
	}

	fmt.Println(colors["red"])
	fmt.Println(colors)
	fmt.Println(teamMartialStatus)

	fmt.Println("================= Section 2 ===================")
	var newcolors map[string]string
	/*
		if we define a new variable in go, and we do not initialize it, go will initialize it with
		this variable's particular "zero value". Here, the zero value of a map is a just an empty map!
		Defining a map using the var notation is suitable for when we want to just create the map, and
		determine what to add to it later on.
	*/
	fmt.Println(newcolors)

	fmt.Println("=============== section 3 =====================")
	/*
		Alternative way to create maps (exactly the same outcome but another way to do it)!
	*/
	otherNewColors := make(map[string]string)

	otherNewColors["white"] = "#fffff"

	fmt.Println(otherNewColors["white"])
	// we do not have dot notation for maps (it is only available to structs in go)

	myArrayLikeMap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	fmt.Println(myArrayLikeMap)

	// deleting values from a map
	delete(myArrayLikeMap, 1)
	fmt.Println(myArrayLikeMap)

}
