package main

import "fmt"

func main() {
	fmt.Println("Example 1: for construct")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nExample 2: goto")
	count := 0

counter:
	if count < 11 {
		fmt.Printf("%d ", count)
		count += 1
		goto counter
	}

	fmt.Println("\nExample 3: array")
	var arr [11]int
	for i := 0; i <= 10; i++ {
		arr[i] = i
	}

	fmt.Println(arr)
}
