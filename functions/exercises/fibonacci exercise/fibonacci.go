package main

import "fmt"

func fibonacci(value int) []int {
	x := make([]int, value) 
	x[0] = 1 
	x[1] = 1       
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2] 
	}
	return x 
}

func main(){
	fmt.Println(fibonacci(5))

}