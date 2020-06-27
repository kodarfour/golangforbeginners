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

func main() {
	evenOnly(2)

	globe = 3 //turns to local variable (within this scope)

	fmt.Println(globe)
	dabMachine := func() { // almost everything in go is a value so u can make functions a variable
		fmt.Println("*dabs*")
	}
	dabMachine()

}
