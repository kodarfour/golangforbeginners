package main

import "fmt"

func varargs(num ...int) {
	for pos := range num {
		fmt.Println(num[pos])
	}

}

func main() {
	varargs(4, 3, 6, 1, 4, 5, 7)
}
