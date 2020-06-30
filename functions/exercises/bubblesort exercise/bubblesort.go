package main

import "fmt"

func main(){
	arr := []int{3,4,9,32,2,58}
	fmt.Println(bubbleSort(arr))
}

func bubbleSort(unsorted []int) (sorted []int){
keepSwapping:

j:	for pos, _:= range unsorted{//goes through place holders
	if pos>=len(unsorted)-1{
		break j
	}	
	defer fmt.Println(pos)
		if pos>=len(unsorted)-1{
			panic("runtime error: out of range")
		}
		if unsorted[pos] > unsorted[pos+1]{
			swap(unsorted[pos],unsorted[pos+1])
		}	
	}
	goto keepSwapping

	sorted = unsorted
	return
}

func swap(x int, y int){
	var temp int
	temp = x
	x = y
	y = temp
}