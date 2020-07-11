package main

import "fmt"

func main() {

	x := 8                       // declares x as 7
	y := &x                      // sets y to the memory value of x
	fmt.Printf("%v %v \n", x, y) // prints value of x and memory value of x
	*y = 7                       // points to the memory of y (which is x) and thus x equals 7
								 // known as dereferencing
	fmt.Println(x)

	// allocation notes
	//new(T) returns *T pointing to a zeroed T (basically creating empty memory map,slice or channel)
	//make(T) returns an initialized T (basically creating memory map,slice or channel with length and or capacity)
}
