package main

import "fmt"

func main() {
	x := 1
	y := 5

	if x > 0 {
		fmt.Println(x)
	} else {
		fmt.Println(y)
	}

	///for init; condition; post { } - a loop using the syntax borrowed from C;
	///for condition { } - a while loop, and;
	///for { } - an endless loop.

	upToten := 0
	for i := 0; i <= 10; i++ {
		count := upToten + i
		fmt.Printf("%v ", count)

	}
}
