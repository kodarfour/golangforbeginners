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

func doubleit(x int) { //declare function that does "something"
	fmt.Println("%v \n", x*2)
}

func callback(y int, d func(int)) { //call back like so
	d(y)
}

func main() {
	evenOnly(2)

	globe = 3 //turns to local variable (within this scope)

	fmt.Println(globe)
	dabMachine := func() { // almost everything in go is a value so u can make functions a variable
		fmt.Println("*dabs*")
	}
	dabMachine()

	for i := 0; i <= 10; i++ {
		defer fmt.Printf("%v ", i) //follows Last in First out format
	}

	deferredWorld := func() {
		defer fmt.Println("World!") //waits til function is complete before printing world
		fmt.Println("Hello,")
	}

	deferredWorld()

}
