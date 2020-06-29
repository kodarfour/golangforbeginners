package main

import "fmt"

var globe int // global variable

func evenOnly(i int) {
	if i%2 == 0 {
		fmt.Printf("%v is even\n", i)
	} else {
		return
	}

	globe = 5 //turns to local variable (within this scope)
	fmt.Println(globe)
}

func thenAddIt(y int, d func(dx int) int) int { //call back like so
	i := d(y)
	i += y
	return i
}

//Variadic Function
//Functions that take a variable number of parameters are known as variadic functions.
//To declare a function as variadic, do something like this:
func addThenSub(num ...int) { //this function alternates between adding and subtracting to find the total
	var total int
	for pos, i := range num {
		if pos%2 == 0 {
			total += num[i-1]
		} else if pos == 0 {
			total += num[i-1]
		} else {
			total -= num[i-1]
		}
	}
	fmt.Println(total)
}

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
	evenOnly(2) //printing only even numbers func
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//global variables and declaring variable as func
	globe = 3 //turns to local variable (within this scope)

	fmt.Println(globe)
	dabMachine := func() { // almost everything in go is a value so u can make functions a variable
		fmt.Println("*dabs*")
	}
	dabMachine()
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//callbacks (this section doubles the int then adds an int of the same int being doubles)
	//ex 3*2+3, 5*2+5
	doubleIt := func(x int) int { //declare function that does "something" (callbacks)
		i := x * 2
		return i
	}

	fmt.Println(thenAddIt(2, doubleIt)) //callback in main
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//defer notes
	for i := 0; i <= 10; i++ {
		defer fmt.Printf("%v ", i) //follows Last in First out format (this will always print last)
	}

	deferredWorld := func() {
		defer fmt.Println("World!") //waits til function is complete before printing world
		fmt.Println("Hello,")
	}

	deferredWorld()
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//Variadic Function
	addThenSub(1, 2, 3, 4) //adds 1 subs 2 adds 3 subs 4 = -2
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//Panic and Recover (with deferring)

	fmt.Println(addButAllergicTo8(3, 8))
	fmt.Println("This should print should there be no errors")
	defer fmt.Println("deferred call in main after printing addButAllergicTo8")

}
