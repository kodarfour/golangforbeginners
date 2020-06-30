package main

import "fmt"

func main() {
	arr := []int{4, 3, 1, 32, 2, 58}
	fmt.Println(bubbleSort(arr))
}

func bubbleSort(unsorted []int) (sorted []int) {
	var temp int
	var swapped bool
	swapped = false
keepSwapping:
	for pos := range unsorted { //goes through place holders
		defer fmt.Printf("%v ", pos)
		if pos >= 5 {
			break
		}
		if pos >= len(unsorted)-1 {
			panic("runtime error: out of range")
		}
		if unsorted[pos] > unsorted[pos+1] {

			temp = unsorted[pos]
			unsorted[pos] = unsorted[pos+1]
			unsorted[pos+1] = temp
			swapped = true
		}
	}
	if swapped != true {
		goto keepSwapping
	}
	sorted = unsorted
	return

}

func swap(x int, y int) {

}
