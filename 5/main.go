package main

import "fmt"

/*
Let's talk about go Gotchas with pointers
*/

func main() {
	mySlice := []string{"Hi There", "how", "are", "you"}

	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	//replace the first element of the slice
	s[0] = "Bye!"

}

/*
Go also has arrays, but nobody tends to use them because they are not able
to grow or shrink over time.
A slice, in the other hand can grow. When creating a slice, we are actually creating two
separate DS's for us:
1. The the DS (i think it is like an array) that keeps length of the current items inside
	the slice, total capacity that the slice has (how many items can it contain?), and
	the ptr to the head of the real data items (this ptr is probably the first thing
	being put into this array)
2. The actual data array!

Therefore, when we call a function and pass a Slice into that function, we are actually
passing the value of the ptr to the actual array to it. Meaning that we are (!) passing by
reference! Not by value, and because of this, you see that if the contents of the slice
get changed inside the function, it will also get affected outside of the functions.
*/
