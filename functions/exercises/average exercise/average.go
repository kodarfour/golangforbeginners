package main

import "fmt"

func average(x []float64) (avg float64) { //calculates the average of a float64 slice
	var total float64
	switch len(x) {
	case 0: //if slice is empty return 0
		avg = 0
	default: //if not it adds up every instance in the slices in the range loop
		for _, j := range x {
			total += j
		}
		avg = float64(total) / float64(len(x)) //then divides sum of integers to the length
	}

	return

}

func main() {
	y := []float64{2, 3}
	fmt.Println(average(y))
}
