package main

import "fmt"

func mapFunc(x func(int) int, y []int) []int{// doubles items in a list
	for pos := range y{
		y[pos] = x(y[pos])
	}
	return y
}

func main(){
	doubleIt := func(i int) int{
		return i *2
	}

	arr := []int{1,2,3,4}

	fmt.Printf("%v ",mapFunc(doubleIt,arr))

}