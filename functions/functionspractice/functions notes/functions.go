package main

import "fmt"

func evenOnly(i int) {
	if i%2 == 0 {
		fmt.Printf("%v is even\n", i)
	} else {
		return
	}
}

func main() {
	evenOnly(2)
}
