package main

import "fmt"

func main() {
	arr := []int{5, -1, 0, 12, 3, 5}
	fmt.Printf("unsorted %v\n", arr)
	bubblesort(arr)
	fmt.Printf("sorted %v\n", arr)
}

func bubblesort(n []int) {
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
}
