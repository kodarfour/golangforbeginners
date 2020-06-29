package main

import "fmt"

//recover
func recoverFunc() { //
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

//Panic with deferred statement
func addButAllergicTo8(x int, y int) int {
	defer fmt.Println("deferred call in addButAllergicTo8")
	defer recoverFunc()
	if x == 8 {
		panic("runtime error: function used is allergic to 8") //will allow program to run but print out panic statement
	}

	if y == 8 {
		panic("runtime error: function is allergic to 8") //will allow program to run but print out panic statement
	}
	i := x + y
	return i
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	fmt.Println(addButAllergicTo8(3, 8))
	fmt.Println("This should print should there be no errors")
	defer fmt.Println("deferred call in main after printing addButAllergicTo8")
}
