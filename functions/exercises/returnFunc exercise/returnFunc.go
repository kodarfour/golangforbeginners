package main

import "fmt"

func plusTwo() func(int) int {
	return func(x int) int {
		return x + 2
	}
}

/*General Code
func plusX(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
*/
func main() {
	p := plusTwo()

	fmt.Printf("%v\n", p(2))

}
