package main

import (
	"bigOrSmall"
	"fmt"
)

func main() {
	i := []int{4, 40, 79}
	fmt.Printf("Is this %d big??? Answer: %v\n", i[0], bigOrSmall.Big(i[0]))
	fmt.Printf("Is this %d small??? Answer: %v\n", i[1], bigOrSmall.Small(i[1]))
	fmt.Printf("Is this %d big??? Answer: %v\n", i[2], bigOrSmall.Big(i[2]))
}
